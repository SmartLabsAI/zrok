/** @module identity */
// Auto-generated, edits will be overwritten
import * as gateway from './gateway'

/**
 * @param {object} options Optional options
 * @param {module:types.accountRequest} [options.body] 
 * @return {Promise<module:types.accountResponse>} account created
 */
export function createAccount(options) {
  if (!options) options = {}
  const parameters = {
    body: {
      body: options.body
    }
  }
  return gateway.request(createAccountOperation, parameters)
}

/**
 * @param {object} options Optional options
 * @param {module:types.enableRequest} [options.body] 
 * @return {Promise<module:types.enableResponse>} environment enabled
 */
export function enable(options) {
  if (!options) options = {}
  const parameters = {
    body: {
      body: options.body
    }
  }
  return gateway.request(enableOperation, parameters)
}

/**
 * @param {object} options Optional options
 * @param {module:types.loginRequest} [options.body] 
 * @return {Promise<module:types.loginResponse>} login successful
 */
export function login(options) {
  if (!options) options = {}
  const parameters = {
    body: {
      body: options.body
    }
  }
  return gateway.request(loginOperation, parameters)
}

const createAccountOperation = {
  path: '/account',
  contentTypes: ['application/zrok.v1+json'],
  method: 'post'
}

const enableOperation = {
  path: '/enable',
  contentTypes: ['application/zrok.v1+json'],
  method: 'post',
  security: [
    {
      id: 'key'
    }
  ]
}

const loginOperation = {
  path: '/login',
  contentTypes: ['application/zrok.v1+json'],
  method: 'post'
}
