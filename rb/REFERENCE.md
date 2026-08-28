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
| `calories_kcal` | `Float` | No | Energy in kcal per 100 g |
| `carbs_g` | `Float` | No | Carbohydrate in grams per 100 g |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No | Cholesterol in mg per 100 g |
| `confidence` | `Float` | No | Data confidence score, 0–1 |
| `fat_g` | `Float` | No | Total fat in grams per 100 g |
| `fiber_g` | `Float` | No | Fibre in grams per 100 g |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No | Potassium in mg per 100 g |
| `protein_g` | `Float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `Float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `String` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `Float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `Float` | No | Sodium in mg per 100 g |
| `source` | `String` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `String` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `Float` | No | Sugars in grams per 100 g |

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
| `calories_kcal` | `Float` | No | Energy in kcal per 100 g |
| `carbs_g` | `Float` | No | Carbohydrate in grams per 100 g |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No | Cholesterol in mg per 100 g |
| `confidence` | `Float` | No | Data confidence score, 0–1 |
| `count` | `Integer` | No |  |
| `fat_g` | `Float` | No | Total fat in grams per 100 g |
| `fiber_g` | `Float` | No | Fibre in grams per 100 g |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No | Potassium in mg per 100 g |
| `protein_g` | `Float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `Float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `String` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `Float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `Float` | No | Sodium in mg per 100 g |
| `source` | `String` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `String` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `Float` | No | Sugars in grams per 100 g |

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
| `calories_kcal` | `Float` | No | Energy in kcal per 100 g |
| `carbs_g` | `Float` | No | Carbohydrate in grams per 100 g |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No | Cholesterol in mg per 100 g |
| `confidence` | `Float` | No | Data confidence score, 0–1 |
| `fat_g` | `Float` | No | Total fat in grams per 100 g |
| `fiber_g` | `Float` | No | Fibre in grams per 100 g |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No | Potassium in mg per 100 g |
| `protein_g` | `Float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `Float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `String` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `Float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `Float` | No | Sodium in mg per 100 g |
| `source` | `String` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `String` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `Float` | No | Sugars in grams per 100 g |

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
| `calories_kcal` | `Float` | No | Energy in kcal per 100 g |
| `carbs_g` | `Float` | No | Carbohydrate in grams per 100 g |
| `category` | `String` | No |  |
| `cholesterol_mg` | `Float` | No | Cholesterol in mg per 100 g |
| `confidence` | `Float` | No | Data confidence score, 0–1 |
| `fat_g` | `Float` | No | Total fat in grams per 100 g |
| `fiber_g` | `Float` | No | Fibre in grams per 100 g |
| `id` | `Integer` | No |  |
| `image_thumb_url` | `String` | No |  |
| `image_url` | `String` | No |  |
| `name` | `String` | No |  |
| `potassium_mg` | `Float` | No | Potassium in mg per 100 g |
| `protein_g` | `Float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `Float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `String` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `Float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `Float` | No | Sodium in mg per 100 g |
| `source` | `String` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `String` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `Float` | No | Sugars in grams per 100 g |

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


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

