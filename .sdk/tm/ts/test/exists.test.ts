
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { DietlyapiIntegrationSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await DietlyapiIntegrationSDK.test()
    equal(null !== testsdk, true)
  })

})
