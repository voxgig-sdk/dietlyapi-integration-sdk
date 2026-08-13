# DietlyapiIntegration TypeScript SDK Reference

Complete API reference for the DietlyapiIntegration TypeScript SDK.


## DietlyapiIntegrationSDK

### Constructor

```ts
new DietlyapiIntegrationSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DietlyapiIntegrationSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = DietlyapiIntegrationSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `DietlyapiIntegrationSDK` instance in test mode.


### Instance Methods

#### `Barcode(data?: object)`

Create a new `Barcode` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BarcodeEntity` instance.

#### `Food(data?: object)`

Create a new `Food` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FoodEntity` instance.

#### `Meta(data?: object)`

Create a new `Meta` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MetaEntity` instance.

#### `Popular(data?: object)`

Create a new `Popular` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PopularEntity` instance.

#### `Search(data?: object)`

Create a new `Search` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SearchEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `DietlyapiIntegrationSDK.test()`.

**Returns:** `DietlyapiIntegrationSDK` instance in test mode.


---

## BarcodeEntity

```ts
const barcode = client.Barcode()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No |  |
| `carbs_g` | `number` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No |  |
| `confidence` | `number` | No |  |
| `fat_g` | `number` | No |  |
| `fiber_g` | `number` | No |  |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No |  |
| `protein_g` | `number` | No |  |
| `saturated_fat_g` | `number` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `number` | No |  |
| `sodium_mg` | `number` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Barcode().load({ id: 'barcode_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BarcodeEntity` instance with the same client and
options.

#### `client()`

Return the parent `DietlyapiIntegrationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FoodEntity

```ts
const food = client.Food()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No |  |
| `carbs_g` | `number` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No |  |
| `confidence` | `number` | No |  |
| `count` | `number` | No |  |
| `fat_g` | `number` | No |  |
| `fiber_g` | `number` | No |  |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No |  |
| `protein_g` | `number` | No |  |
| `saturated_fat_g` | `number` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `number` | No |  |
| `sodium_mg` | `number` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `number` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `category` | `/foods/categories` | `client.Food().list({ $action: 'category', ... })` |

An action returns that action's OWN response, which is not necessarily a
Food record — check the API definition for its shape.

```ts
const result = await client.Food().list({
  $action: 'category',
  /* ...the action's own arguments */
})
```

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Food().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Food().load({ id: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FoodEntity` instance with the same client and
options.

#### `client()`

Return the parent `DietlyapiIntegrationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MetaEntity

```ts
const meta = client.Meta()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `foods_in_db` | `number` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Meta().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MetaEntity` instance with the same client and
options.

#### `client()`

Return the parent `DietlyapiIntegrationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PopularEntity

```ts
const popular = client.Popular()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No |  |
| `carbs_g` | `number` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No |  |
| `confidence` | `number` | No |  |
| `fat_g` | `number` | No |  |
| `fiber_g` | `number` | No |  |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No |  |
| `protein_g` | `number` | No |  |
| `saturated_fat_g` | `number` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `number` | No |  |
| `sodium_mg` | `number` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Popular().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PopularEntity` instance with the same client and
options.

#### `client()`

Return the parent `DietlyapiIntegrationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SearchEntity

```ts
const search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No |  |
| `carbs_g` | `number` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No |  |
| `confidence` | `number` | No |  |
| `fat_g` | `number` | No |  |
| `fiber_g` | `number` | No |  |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No |  |
| `protein_g` | `number` | No |  |
| `saturated_fat_g` | `number` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `number` | No |  |
| `sodium_mg` | `number` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Search().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `DietlyapiIntegrationSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new DietlyapiIntegrationSDK({
  feature: {
    test: { active: true },
  }
})
```

