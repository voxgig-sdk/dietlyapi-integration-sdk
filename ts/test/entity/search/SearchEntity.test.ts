

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { DietlyapiIntegrationSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('SearchEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when DIETLYAPI_INTEGRATION_TEST_LIVE=TRUE.
  afterEach(liveDelay('DIETLYAPI_INTEGRATION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = DietlyapiIntegrationSDK.test()
    const ent = testsdk.Search()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.DIETLYAPI_INTEGRATION_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'search.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"barcode","req":false,"type":"`$STRING`","index$":0},{"active":true,"name":"brand","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"calories_kcal","req":false,"short":"Energy in kcal per 100 g","type":"`$NUMBER`","index$":2},{"active":true,"name":"carbs_g","req":false,"short":"Carbohydrate in grams per 100 g","type":"`$NUMBER`","index$":3},{"active":true,"name":"category","req":false,"type":"`$STRING`","index$":4},{"active":true,"name":"cholesterol_mg","req":false,"short":"Cholesterol in mg per 100 g","type":"`$NUMBER`","index$":5},{"active":true,"name":"confidence","req":false,"short":"Data confidence score, 0–1","type":"`$NUMBER`","index$":6},{"active":true,"name":"fat_g","req":false,"short":"Total fat in grams per 100 g","type":"`$NUMBER`","index$":7},{"active":true,"name":"fiber_g","req":false,"short":"Fibre in grams per 100 g","type":"`$NUMBER`","index$":8},{"active":true,"name":"id","req":false,"type":"`$INTEGER`","index$":9},{"active":true,"name":"image_thumb_url","req":false,"type":"`$STRING`","index$":10},{"active":true,"name":"image_url","req":false,"type":"`$STRING`","index$":11},{"active":true,"name":"name","req":false,"type":"`$STRING`","index$":12},{"active":true,"name":"potassium_mg","req":false,"short":"Potassium in mg per 100 g","type":"`$NUMBER`","index$":13},{"active":true,"name":"protein_g","req":false,"short":"Protein in grams per 100 g","type":"`$NUMBER`","index$":14},{"active":true,"name":"saturated_fat_g","req":false,"short":"Saturated fat in grams per 100 g","type":"`$NUMBER`","index$":15},{"active":true,"name":"serving_desc","req":false,"short":"Human-readable label for one serving, e.g.","type":"`$STRING`","index$":16},{"active":true,"name":"serving_size_g","req":false,"short":"Grams in one manufacturer serving, where the source declares one.","type":"`$NUMBER`","index$":17},{"active":true,"name":"sodium_mg","req":false,"short":"Sodium in mg per 100 g","type":"`$NUMBER`","index$":18},{"active":true,"name":"source","req":false,"short":"Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community","type":"`$STRING`","index$":19},{"active":true,"name":"static_url","req":false,"short":"Path of the human-readable page on www.getdietly.com","type":"`$STRING`","index$":20},{"active":true,"name":"sugar_g","req":false,"short":"Sugars in grams per 100 g","type":"`$NUMBER`","index$":21}],"id":{"field":"id","name":"id"},"name":"search","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"query":[{"active":true,"example":5,"kind":"query","name":"limit","orig":"limit","reqd":false,"type":"`$INTEGER`","index$":0},{"active":true,"example":"greek yogurt","kind":"query","name":"q","orig":"q","reqd":true,"type":"`$STRING`","index$":1},{"active":true,"kind":"query","name":"source","orig":"source","reqd":false,"type":"`$STRING`","index$":2}]},"contract":{"id":"GET /search","json":"{\"operationId\":\"searchFoods\",\"parameters\":[{\"description\":\"Search term (min 2 characters)\",\"example\":\"greek yogurt\",\"in\":\"query\",\"name\":\"q\",\"required\":true,\"schema\":{\"minLength\":2,\"type\":\"string\"}},{\"in\":\"query\",\"name\":\"limit\",\"schema\":{\"default\":5,\"maximum\":50,\"minimum\":1,\"type\":\"integer\"}},{\"description\":\"Optional source filter: off, usda or claude. Omit this parameter to search all indexed sources.\",\"in\":\"query\",\"name\":\"source\",\"schema\":{\"enum\":[\"off\",\"usda\",\"claude\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"example\":[{\"barcode\":\"0855088005245\",\"brand\":null,\"calories_kcal\":105.9,\"carbs_g\":12.4,\"category\":\"Dairies\",\"cholesterol_mg\":18,\"confidence\":0.9,\"fat_g\":3.5,\"fiber_g\":0,\"id\":1068319,\"image_thumb_url\":\"https://api.getdietly.com/img?u=...\",\"image_url\":\"https://api.getdietly.com/img?u=...\",\"name\":\"Greek yogurt\",\"potassium_mg\":112,\"protein_g\":7.6,\"saturated_fat_g\":2.4,\"serving_desc\":\"3/4 cup (170 g)\",\"serving_size_g\":170,\"sodium_mg\":44.1,\"source\":\"off\",\"static_url\":\"food/19/greek-yogurt-2.html\",\"sugar_g\":10}],\"schema\":{\"items\":{\"description\":\"A single food record. Every nutrient value is expressed PER 100 GRAMS of the food, not per serving. To report a portion, multiply by grams and divide by 100 — e.g. an olive oil returning calories_kcal 831 is 831 kcal per 100 g, so a 13 g tablespoon is 108 kcal. Use serving_size_g to translate a serving into grams, but never treat the nutrient values as a per-serving figure.\",\"example\":{\"barcode\":\"0855088005245\",\"brand\":null,\"calories_kcal\":105.9,\"carbs_g\":12.4,\"category\":\"Dairies\",\"cholesterol_mg\":18,\"confidence\":0.9,\"fat_g\":3.5,\"fiber_g\":0,\"id\":1068319,\"image_thumb_url\":\"https://api.getdietly.com/img?u=...\",\"image_url\":\"https://api.getdietly.com/img?u=...\",\"name\":\"Greek yogurt\",\"potassium_mg\":112,\"protein_g\":7.6,\"saturated_fat_g\":2.4,\"serving_desc\":\"3/4 cup (170 g)\",\"serving_size_g\":170,\"sodium_mg\":44.1,\"source\":\"off\",\"static_url\":\"food/19/greek-yogurt-2.html\",\"sugar_g\":10},\"properties\":{\"barcode\":{\"nullable\":true,\"type\":\"string\"},\"brand\":{\"nullable\":true,\"type\":\"string\"},\"calories_kcal\":{\"description\":\"Energy in kcal per 100 g\",\"nullable\":true,\"type\":\"number\"},\"carbs_g\":{\"description\":\"Carbohydrate in grams per 100 g\",\"nullable\":true,\"type\":\"number\"},\"category\":{\"nullable\":true,\"type\":\"string\"},\"cholesterol_mg\":{\"description\":\"Cholesterol in mg per 100 g\",\"nullable\":true,\"type\":\"number\"},\"confidence\":{\"description\":\"Data confidence score, 0–1\",\"type\":\"number\"},\"fat_g\":{\"description\":\"Total fat in grams per 100 g\",\"nullable\":true,\"type\":\"number\"},\"fiber_g\":{\"description\":\"Fibre in grams per 100 g\",\"nullable\":true,\"type\":\"number\"},\"id\":{\"type\":\"integer\"},\"image_thumb_url\":{\"nullable\":true,\"type\":\"string\"},\"image_url\":{\"nullable\":true,\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"potassium_mg\":{\"description\":\"Potassium in mg per 100 g\",\"nullable\":true,\"type\":\"number\"},\"protein_g\":{\"description\":\"Protein in grams per 100 g\",\"nullable\":true,\"type\":\"number\"},\"saturated_fat_g\":{\"description\":\"Saturated fat in grams per 100 g\",\"nullable\":true,\"type\":\"number\"},\"serving_desc\":{\"description\":\"Human-readable label for one serving, e.g. '3/4 cup (170 g)'.\",\"nullable\":true,\"type\":\"string\"},\"serving_size_g\":{\"description\":\"Grams in one manufacturer serving, where the source declares one. Descriptive only: the nutrient fields below are not scaled to it. Null on roughly a third of records.\",\"nullable\":true,\"type\":\"number\"},\"sodium_mg\":{\"description\":\"Sodium in mg per 100 g\",\"nullable\":true,\"type\":\"number\"},\"source\":{\"description\":\"Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community\",\"type\":\"string\"},\"static_url\":{\"description\":\"Path of the human-readable page on www.getdietly.com\",\"nullable\":true,\"type\":\"string\"},\"sugar_g\":{\"description\":\"Sugars in grams per 100 g\",\"nullable\":true,\"type\":\"number\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Ranked list of matching foods\"},\"429\":{\"description\":\"Rate limit exceeded. Check the Retry-After and X-RateLimit-* headers.\",\"headers\":{\"Retry-After\":{\"description\":\"Seconds to wait before retrying\",\"schema\":{\"type\":\"integer\"}},\"X-RateLimit-Limit\":{\"schema\":{\"type\":\"integer\"}},\"X-RateLimit-Remaining\":{\"schema\":{\"type\":\"integer\"}},\"X-RateLimit-Reset\":{\"schema\":{\"type\":\"integer\"}}}}},\"security\":[{},{\"bearerAuth\":[]}],\"securitySchemes\":{\"bearerAuth\":{\"description\":\"Optional for the endpoints in this spec. Get a free key instantly at https://www.getdietly.com/account (no card required).\",\"scheme\":\"bearer\",\"type\":\"http\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/search","segments":[{"lit":"search"}],"select":{"exist":["limit","q","source"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"search","name__orig":"search","Name":"Search","name_":"search","name-":"search","NAME":"SEARCH","index$":4}, {"active":true,"entity":"search","key$":"BasicSearchFlow","kind":"basic","name":"BasicSearchFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"search_ref01"}}],"index$":0}]}, 'Search')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let search_ref01_data = Object.values(setup.data.existing.search)[0] as any

    // LIST
    const search_ref01_ent = client.Search()
    const search_ref01_match: any = {}

    const search_ref01_list = (await search_ref01_ent.list(search_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/search/SearchTestData.json')

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
    ['search01','search02','search03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'DIETLYAPI_INTEGRATION_TEST_SEARCH_ENTID': idmap,
    'DIETLYAPI_INTEGRATION_TEST_LIVE': 'FALSE',
    'DIETLYAPI_INTEGRATION_TEST_EXPLAIN': 'FALSE',
    'DIETLYAPI_INTEGRATION_APIKEY': '',
  })

  idmap = env['DIETLYAPI_INTEGRATION_TEST_SEARCH_ENTID']

  const live = 'TRUE' === env.DIETLYAPI_INTEGRATION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['DIETLYAPI_INTEGRATION_TEST_SEARCH_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new DietlyapiIntegrationSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.DIETLYAPI_INTEGRATION_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
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
    transport,
    now: Date.now(),
  }

  return setup
}
  
