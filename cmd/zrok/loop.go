package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/openziti-test-kitchen/zrok/model"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok"
	"github.com/openziti-test-kitchen/zrok/rest_client_zrok/tunnel"
	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
	"github.com/openziti-test-kitchen/zrok/util"
	"github.com/openziti-test-kitchen/zrok/zrokdir"
	"github.com/openziti/sdk-golang/ziti"
	"github.com/openziti/sdk-golang/ziti/config"
	"github.com/openziti/sdk-golang/ziti/edge"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	testCmd.AddCommand(newLoopCmd().cmd)
}

type loopCmd struct {
	cmd            *cobra.Command
	loopers        int
	iterations     int
	statusEvery    int
	dwellSeconds   int
	timeoutSeconds int
	minPayload     int
	maxPayload     int
	minPacingMs    int
	maxPacingMs    int
}

func newLoopCmd() *loopCmd {
	cmd := &cobra.Command{
		Use:   "loop",
		Short: "Start a loop agent",
		Args:  cobra.ExactArgs(0),
	}
	r := &loopCmd{cmd: cmd}
	cmd.Run = r.run
	cmd.Flags().IntVarP(&r.loopers, "loopers", "l", 1, "Number of current loopers to start")
	cmd.Flags().IntVarP(&r.iterations, "iterations", "i", 1, "Number of iterations per looper")
	cmd.Flags().IntVarP(&r.statusEvery, "status-every", "E", 100, "Show status every # iterations")
	cmd.Flags().IntVarP(&r.dwellSeconds, "dwell-seconds", "D", 1, "Dwell # seconds before starting iterations")
	cmd.Flags().IntVarP(&r.timeoutSeconds, "timeout-seconds", "T", 30, "Time out after # seconds when sending http requests")
	cmd.Flags().IntVar(&r.minPayload, "min-payload", 64, "Minimum payload size in bytes")
	cmd.Flags().IntVar(&r.maxPayload, "max-payload", 10240, "Maximum payload size in bytes")
	cmd.Flags().IntVar(&r.minPacingMs, "min-pacing-ms", 0, "Minimum pacing in milliseconds")
	cmd.Flags().IntVar(&r.maxPacingMs, "max-pacing-ms", 0, "Maximum pacing in milliseconds")
	return r
}

func (r *loopCmd) run(_ *cobra.Command, _ []string) {
	var loopers []*looper
	for i := 0; i < r.loopers; i++ {
		l := newLooper(i, r)
		loopers = append(loopers, l)
		go l.run()
	}
	for _, l := range loopers {
		<-l.done
	}
	totalMismatches := 0
	totalXfer := int64(0)
	for _, l := range loopers {
		deltaSeconds := l.stopTime.Sub(l.startTime).Seconds()
		xfer := int64(float64(l.bytes) / deltaSeconds)
		totalXfer += xfer
		totalMismatches += l.mismatches
		xferSec := util.BytesToSize(xfer)
		logrus.Infof("looper #%d: %d mismatches, %s/sec", l.id, l.mismatches, xferSec)
	}
	totalXferSec := util.BytesToSize(totalXfer)
	logrus.Infof("total: %d mismatches, %s/sec", totalMismatches, totalXferSec)
}

type looper struct {
	id            int
	cmd           *loopCmd
	env           *zrokdir.Environment
	done          chan struct{}
	listener      edge.Listener
	zif           string
	zrok          *rest_client_zrok.Zrok
	service       string
	proxyEndpoint string
	auth          runtime.ClientAuthInfoWriter
	mismatches    int
	bytes         int64
	startTime     time.Time
	stopTime      time.Time
}

func newLooper(id int, cmd *loopCmd) *looper {
	return &looper{
		id:   id,
		cmd:  cmd,
		done: make(chan struct{}),
	}
}

func (l *looper) run() {
	defer close(l.done)
	defer logrus.Infof("stopping #%d", l.id)

	l.startup()
	logrus.Infof("looper #%d, service: %v, frontend: %v", l.id, l.service, l.proxyEndpoint)
	go l.serviceListener()
	l.dwell()
	l.iterate()
	logrus.Infof("looper #%d: complete", l.id)
	l.shutdown()
}

func (l *looper) serviceListener() {
	zcfg, err := config.NewFromFile(l.zif)
	if err != nil {
		logrus.Errorf("error opening ziti config '%v': %v", l.zif, err)
		return
	}
	opts := ziti.ListenOptions{
		ConnectTimeout: 5 * time.Minute,
		MaxConnections: 10,
	}
	if l.listener, err = ziti.NewContextWithConfig(zcfg).ListenWithOptions(l.service, &opts); err == nil {
		if err := http.Serve(l.listener, l); err != nil {
			logrus.Errorf("looper #%d, error serving: %v", l.id, err)
		}
	} else {
		logrus.Errorf("looper #%d, error listening: %v", l.id, err)
	}
}

func (l *looper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	io.Copy(buf, r.Body)
	w.Write(buf.Bytes())
}

func (l *looper) startup() {
	logrus.Infof("starting #%d", l.id)

	var err error
	l.env, err = zrokdir.LoadEnvironment()
	if err != nil {
		panic(err)
	}
	l.zif, err = zrokdir.ZitiIdentityFile("backend")
	if err != nil {
		panic(err)
	}
	l.zrok, err = zrokdir.ZrokClient(l.env.ApiEndpoint)
	if err != nil {
		panic(err)
	}
	l.auth = httptransport.APIKeyAuth("x-token", "header", l.env.ZrokToken)
	tunnelReq := tunnel.NewTunnelParams()
	tunnelReq.Body = &rest_model_zrok.TunnelRequest{
		ZitiIdentityID: l.env.ZitiIdentityId,
		Endpoint:       fmt.Sprintf("looper#%d", l.id),
		AuthScheme:     string(model.None),
	}
	tunnelResp, err := l.zrok.Tunnel.Tunnel(tunnelReq, l.auth)
	if err != nil {
		panic(err)
	}
	l.service = tunnelResp.Payload.Service
	l.proxyEndpoint = tunnelResp.Payload.ProxyEndpoint
}

func (l *looper) dwell() {
	time.Sleep(time.Duration(l.cmd.dwellSeconds) * time.Second)
}

func (l *looper) iterate() {
	l.startTime = time.Now()
	defer func() { l.stopTime = time.Now() }()

	for i := 0; i < l.cmd.iterations; i++ {
		if i > 0 && i%l.cmd.statusEvery == 0 {
			logrus.Infof("looper #%d: iteration #%d", l.id, i)
		}
		sz := l.cmd.maxPayload
		if l.cmd.maxPayload-l.cmd.minPayload > 0 {
			sz = rand.Intn(l.cmd.maxPayload-l.cmd.minPayload) + l.cmd.minPayload
		}
		outpayload := make([]byte, sz)
		outbase64 := base64.StdEncoding.EncodeToString(outpayload)
		rand.Read(outpayload)
		if req, err := http.NewRequest("POST", l.proxyEndpoint, bytes.NewBufferString(outbase64)); err == nil {
			client := &http.Client{Timeout: time.Second * time.Duration(l.cmd.timeoutSeconds)}
			if resp, err := client.Do(req); err == nil {
				inpayload := new(bytes.Buffer)
				io.Copy(inpayload, resp.Body)
				inbase64 := inpayload.String()
				if inbase64 != outbase64 {
					logrus.Errorf("looper #%d payload mismatch!", l.id)
					l.mismatches++
				} else {
					l.bytes += int64(len(outbase64))
					logrus.Debugf("looper #%d payload match", l.id)
				}
			} else {
				logrus.Errorf("looper #%d error: %v", l.id, err)
			}
		} else {
			logrus.Errorf("looper #%d error creating request: %v", l.id, err)
		}
		pacingMs := l.cmd.maxPayload
		if l.cmd.maxPacingMs-l.cmd.minPacingMs > 0 {
			pacingMs = rand.Intn(l.cmd.maxPacingMs-l.cmd.minPacingMs) + l.cmd.minPacingMs
			time.Sleep(time.Duration(pacingMs) * time.Millisecond)
		}
	}
}

func (l *looper) shutdown() {
	if l.listener != nil {
		if err := l.listener.Close(); err != nil {
			logrus.Errorf("looper #%d error closing listener: %v", l.id, err)
		}
	}

	untunnelReq := tunnel.NewUntunnelParams()
	untunnelReq.Body = &rest_model_zrok.UntunnelRequest{
		ZitiIdentityID: l.env.ZitiIdentityId,
		Service:        l.service,
	}
	if _, err := l.zrok.Tunnel.Untunnel(untunnelReq, l.auth); err != nil {
		logrus.Errorf("error shutting down looper #%d: %v", l.id, err)
	}
}
