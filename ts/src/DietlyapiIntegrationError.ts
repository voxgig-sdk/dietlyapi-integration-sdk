
import { Context } from './Context'


class DietlyapiIntegrationError extends Error {

  isDietlyapiIntegrationError = true

  sdk = 'DietlyapiIntegration'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  DietlyapiIntegrationError
}

