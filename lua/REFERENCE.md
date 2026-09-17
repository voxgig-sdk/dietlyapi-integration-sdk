# DietlyapiIntegration Lua SDK Reference

Complete API reference for the DietlyapiIntegration Lua SDK.


## DietlyapiIntegrationSDK

### Constructor

```lua
local sdk = require("dietlyapi-integration_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Barcode(data)`

Create a new `Barcode` entity instance. Pass `nil` for no initial data.

#### `Food(data)`

Create a new `Food` entity instance. Pass `nil` for no initial data.

#### `Meta(data)`

Create a new `Meta` entity instance. Pass `nil` for no initial data.

#### `Popular(data)`

Create a new `Popular` entity instance. Pass `nil` for no initial data.

#### `Search(data)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## BarcodeEntity

```lua
local barcode = client:Barcode(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No | Energy in kcal per 100 g |
| `carbs_g` | `number` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No | Cholesterol in mg per 100 g |
| `confidence` | `number` | No | Data confidence score, 0–1 |
| `fat_g` | `number` | No | Total fat in grams per 100 g |
| `fiber_g` | `number` | No | Fibre in grams per 100 g |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No | Potassium in mg per 100 g |
| `protein_g` | `number` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `number` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `number` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `number` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `number` | No | Sugars in grams per 100 g |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Barcode():load({ id = "barcode_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BarcodeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FoodEntity

```lua
local food = client:Food(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No | Energy in kcal per 100 g |
| `carbs_g` | `number` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No | Cholesterol in mg per 100 g |
| `confidence` | `number` | No | Data confidence score, 0–1 |
| `fat_g` | `number` | No | Total fat in grams per 100 g |
| `fiber_g` | `number` | No | Fibre in grams per 100 g |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No | Potassium in mg per 100 g |
| `protein_g` | `number` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `number` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `number` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `number` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `number` | No | Sugars in grams per 100 g |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Food():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Food():load({ id = 1 })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FoodEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MetaEntity

```lua
local meta = client:Meta(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `foods_in_db` | `number` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Meta():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetaEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PopularEntity

```lua
local popular = client:Popular(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No | Energy in kcal per 100 g |
| `carbs_g` | `number` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No | Cholesterol in mg per 100 g |
| `confidence` | `number` | No | Data confidence score, 0–1 |
| `fat_g` | `number` | No | Total fat in grams per 100 g |
| `fiber_g` | `number` | No | Fibre in grams per 100 g |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No | Potassium in mg per 100 g |
| `protein_g` | `number` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `number` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `number` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `number` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `number` | No | Sugars in grams per 100 g |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Popular():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PopularEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SearchEntity

```lua
local search = client:Search(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `number` | No | Energy in kcal per 100 g |
| `carbs_g` | `number` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `number` | No | Cholesterol in mg per 100 g |
| `confidence` | `number` | No | Data confidence score, 0–1 |
| `fat_g` | `number` | No | Total fat in grams per 100 g |
| `fiber_g` | `number` | No | Fibre in grams per 100 g |
| `id` | `number` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `number` | No | Potassium in mg per 100 g |
| `protein_g` | `number` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `number` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `number` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `number` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `number` | No | Sugars in grams per 100 g |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Search():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    ratelimit = { active = true },
    retry = { active = true },
    test = { active = true },
    timeout = { active = true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Client-side rate limiting via a token bucket.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Automatic retry of transient failures with exponential backoff.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Per-request timeout with transport abort.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

