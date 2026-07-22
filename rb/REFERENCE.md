# DietlyapiIntegration Ruby SDK Reference

Complete API reference for the DietlyapiIntegration Ruby SDK.


## DietlyapiIntegrationSDK

### Constructor

```ruby
require_relative 'DietlyapiIntegration_sdk'

client = DietlyapiIntegrationSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DietlyapiIntegrationSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = DietlyapiIntegrationSDK.test
```


### Instance Methods

#### `Barcode(data = nil)`

Create a new `Barcode` entity instance. Pass `nil` for no initial data.

#### `Food(data = nil)`

Create a new `Food` entity instance. Pass `nil` for no initial data.

#### `Meta(data = nil)`

Create a new `Meta` entity instance. Pass `nil` for no initial data.

#### `Popular(data = nil)`

Create a new `Popular` entity instance. Pass `nil` for no initial data.

#### `Search(data = nil)`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## BarcodeEntity

```ruby
barcode = client.Barcode
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `String` | No |  |
| `brand` | `String` | No |  |
| `calories_kcal` | `Float` | No |  |
| `carbs_g` | `Float` | No |  |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No |  |
| `confidence` | `Float` | No |  |
| `fat_g` | `Float` | No |  |
| `fiber_g` | `Float` | No |  |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No |  |
| `protein_g` | `Float` | No |  |
| `saturated_fat_g` | `Float` | No |  |
| `serving_desc` | `String` | No |  |
| `serving_size_g` | `Float` | No |  |
| `sodium_mg` | `Float` | No |  |
| `source` | `String` | No |  |
| `static_url` | `String` | No |  |
| `sugar_g` | `Float` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Barcode.load({ "id" => "barcode_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BarcodeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FoodEntity

```ruby
food = client.Food
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `String` | No |  |
| `brand` | `String` | No |  |
| `calories_kcal` | `Float` | No |  |
| `carbs_g` | `Float` | No |  |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No |  |
| `confidence` | `Float` | No |  |
| `count` | `Integer` | No |  |
| `fat_g` | `Float` | No |  |
| `fiber_g` | `Float` | No |  |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No |  |
| `protein_g` | `Float` | No |  |
| `saturated_fat_g` | `Float` | No |  |
| `serving_desc` | `String` | No |  |
| `serving_size_g` | `Float` | No |  |
| `sodium_mg` | `Float` | No |  |
| `source` | `String` | No |  |
| `static_url` | `String` | No |  |
| `sugar_g` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Food.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Food.load({ "id" => 1 })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FoodEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MetaEntity

```ruby
meta = client.Meta
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `foods_in_db` | `Integer` | No |  |
| `status` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Meta.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MetaEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PopularEntity

```ruby
popular = client.Popular
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `String` | No |  |
| `brand` | `String` | No |  |
| `calories_kcal` | `Float` | No |  |
| `carbs_g` | `Float` | No |  |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No |  |
| `confidence` | `Float` | No |  |
| `fat_g` | `Float` | No |  |
| `fiber_g` | `Float` | No |  |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No |  |
| `protein_g` | `Float` | No |  |
| `saturated_fat_g` | `Float` | No |  |
| `serving_desc` | `String` | No |  |
| `serving_size_g` | `Float` | No |  |
| `sodium_mg` | `Float` | No |  |
| `source` | `String` | No |  |
| `static_url` | `String` | No |  |
| `sugar_g` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Popular.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PopularEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SearchEntity

```ruby
search = client.Search
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `String` | No |  |
| `brand` | `String` | No |  |
| `calories_kcal` | `Float` | No |  |
| `carbs_g` | `Float` | No |  |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No |  |
| `confidence` | `Float` | No |  |
| `fat_g` | `Float` | No |  |
| `fiber_g` | `Float` | No |  |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No |  |
| `protein_g` | `Float` | No |  |
| `saturated_fat_g` | `Float` | No |  |
| `serving_desc` | `String` | No |  |
| `serving_size_g` | `Float` | No |  |
| `sodium_mg` | `Float` | No |  |
| `source` | `String` | No |  |
| `static_url` | `String` | No |  |
| `sugar_g` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Search.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = DietlyapiIntegrationSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

