
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { DietlyapiIntegrationSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('BarcodeEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DIETLYAPI_INTEGRATION_TEST_LIVE=TRUE.
  afterEach(liveDelay('DIETLYAPI_INTEGRATION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DietlyapiIntegrationSDK.test()
    const ent = testsdk.Barcode()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DIETLYAPI_INTEGRATION_TEST_LIVE
    for (const op of ['load']) {
      if (maybeSkipControl(t, 'entityOp', 'barcode.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let barcode_ref01_data = Object.values(setup.data.existing.barcode)[0] as any

    // LOAD
    const barcode_ref01_ent = client.Barcode()
    const barcode_ref01_match_dt0: any = {}
    barcode_ref01_match_dt0.id = barcode_ref01_data.id
    const barcode_ref01_data_dt0 = (await barcode_ref01_ent.load(barcode_ref01_match_dt0)).data()
    assert(barcode_ref01_data_dt0.id === barcode_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/barcode/BarcodeTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = DietlyapiIntegrationSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['barcode01','barcode02','barcode03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID': idmap,
    'DIETLYAPI_INTEGRATION_TEST_LIVE': 'FALSE',
    'DIETLYAPI_INTEGRATION_TEST_EXPLAIN': 'FALSE',
    'DIETLYAPI_INTEGRATION_APIKEY': 'NONE',
  })

  idmap = env['DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID']

  const live = 'TRUE' === env.DIETLYAPI_INTEGRATION_TEST_LIVE

  if (live) {
    client = new DietlyapiIntegrationSDK(merge([
      {
        apikey: env.DIETLYAPI_INTEGRATION_APIKEY,
      },
      extra
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.DIETLYAPI_INTEGRATION_TEST_EXPLAIN,
    live,
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
