# DietlyapiIntegration Ruby SDK



The Ruby SDK for the DietlyapiIntegration API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Barcode` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dietlyapi-integration-sdk/releases](https://github.com/voxgig-sdk/dietlyapi-integration-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "DietlyapiIntegration_sdk"

client = DietlyapiIntegrationSDK.new({
  "apikey" => ENV["DIETLYAPI_INTEGRATION_APIKEY"],
})
```

### 3. Load a barcode

```ruby
begin
  # load returns the bare Barcode record (raises on error).
  barcode = client.Barcode.load({ "id" => "example_id" })
  puts barcode
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  barcode = client.Barcode.load({ "id" => "example_id" })
rescue => err
  warn "load failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = DietlyapiIntegrationSDK.test({
  "entity" => { "barcode" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the bare mock record (raises on error).
barcode = client.Barcode.load({ "id" => "test01" })
puts barcode
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = DietlyapiIntegrationSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
DIETLYAPI_INTEGRATION_TEST_LIVE=TRUE
DIETLYAPI_INTEGRATION_APIKEY=<your-key>
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### DietlyapiIntegrationSDK

```ruby
require_relative "DietlyapiIntegration_sdk"
client = DietlyapiIntegrationSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = DietlyapiIntegrationSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### DietlyapiIntegrationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Barcode` | `(data) -> BarcodeEntity` | Create a Barcode entity instance. |
| `Food` | `(data) -> FoodEntity` | Create a Food entity instance. |
| `Meta` | `(data) -> MetaEntity` | Create a Meta entity instance. |
| `Popular` | `(data) -> PopularEntity` | Create a Popular entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `DietlyapiIntegrationError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Barcode

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` |  |
| `carbs_g` |  |
| `category` |  |
| `cholesterol_mg` |  |
| `confidence` |  |
| `fat_g` |  |
| `fiber_g` |  |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` |  |
| `protein_g` |  |
| `saturated_fat_g` |  |
| `serving_desc` |  |
| `serving_size_g` |  |
| `sodium_mg` |  |
| `source` |  |
| `static_url` |  |
| `sugar_g` |  |

Operations: Load.

API path: `/barcode/{code}`

#### Food

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` |  |
| `carbs_g` |  |
| `category` |  |
| `cholesterol_mg` |  |
| `confidence` |  |
| `count` |  |
| `fat_g` |  |
| `fiber_g` |  |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` |  |
| `protein_g` |  |
| `saturated_fat_g` |  |
| `serving_desc` |  |
| `serving_size_g` |  |
| `sodium_mg` |  |
| `source` |  |
| `static_url` |  |
| `sugar_g` |  |

Operations: List, Load.

API path: `/foods/categories`

#### Meta

| Field | Description |
| --- | --- |
| `foods_in_db` |  |
| `status` |  |

Operations: Load.

API path: `/health`

#### Popular

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` |  |
| `carbs_g` |  |
| `category` |  |
| `cholesterol_mg` |  |
| `confidence` |  |
| `fat_g` |  |
| `fiber_g` |  |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` |  |
| `protein_g` |  |
| `saturated_fat_g` |  |
| `serving_desc` |  |
| `serving_size_g` |  |
| `sodium_mg` |  |
| `source` |  |
| `static_url` |  |
| `sugar_g` |  |

Operations: List.

API path: `/foods/popular`

#### Search

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` |  |
| `carbs_g` |  |
| `category` |  |
| `cholesterol_mg` |  |
| `confidence` |  |
| `fat_g` |  |
| `fiber_g` |  |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` |  |
| `protein_g` |  |
| `saturated_fat_g` |  |
| `serving_desc` |  |
| `serving_size_g` |  |
| `sodium_mg` |  |
| `source` |  |
| `static_url` |  |
| `sugar_g` |  |

Operations: List.

API path: `/search`



## Entities


### Barcode

Create an instance: `barcode = client.Barcode`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `String` |  |
| `brand` | `String` |  |
| `calories_kcal` | `Float` |  |
| `carbs_g` | `Float` |  |
| `category` | `String` |  |
| `cholesterol_mg` | `Float` |  |
| `confidence` | `Float` |  |
| `fat_g` | `Float` |  |
| `fiber_g` | `Float` |  |
| `id` | `Integer` |  |
| `image_thumb_url` | `String` |  |
| `image_url` | `String` |  |
| `name` | `String` |  |
| `potassium_mg` | `Float` |  |
| `protein_g` | `Float` |  |
| `saturated_fat_g` | `Float` |  |
| `serving_desc` | `String` |  |
| `serving_size_g` | `Float` |  |
| `sodium_mg` | `Float` |  |
| `source` | `String` |  |
| `static_url` | `String` |  |
| `sugar_g` | `Float` |  |

#### Example: Load

```ruby
# load returns the bare Barcode record (raises on error).
barcode = client.Barcode.load({ "id" => "barcode_id" })
```


### Food

Create an instance: `food = client.Food`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `String` |  |
| `brand` | `String` |  |
| `calories_kcal` | `Float` |  |
| `carbs_g` | `Float` |  |
| `category` | `String` |  |
| `cholesterol_mg` | `Float` |  |
| `confidence` | `Float` |  |
| `count` | `Integer` |  |
| `fat_g` | `Float` |  |
| `fiber_g` | `Float` |  |
| `id` | `Integer` |  |
| `image_thumb_url` | `String` |  |
| `image_url` | `String` |  |
| `name` | `String` |  |
| `potassium_mg` | `Float` |  |
| `protein_g` | `Float` |  |
| `saturated_fat_g` | `Float` |  |
| `serving_desc` | `String` |  |
| `serving_size_g` | `Float` |  |
| `sodium_mg` | `Float` |  |
| `source` | `String` |  |
| `static_url` | `String` |  |
| `sugar_g` | `Float` |  |

#### Example: Load

```ruby
# load returns the bare Food record (raises on error).
food = client.Food.load({ "id" => 1 })
```

#### Example: List

```ruby
# list returns an Array of Food records (raises on error).
foods = client.Food.list
```


### Meta

Create an instance: `meta = client.Meta`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `foods_in_db` | `Integer` |  |
| `status` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Meta record (raises on error).
meta = client.Meta.load()
```


### Popular

Create an instance: `popular = client.Popular`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `String` |  |
| `brand` | `String` |  |
| `calories_kcal` | `Float` |  |
| `carbs_g` | `Float` |  |
| `category` | `String` |  |
| `cholesterol_mg` | `Float` |  |
| `confidence` | `Float` |  |
| `fat_g` | `Float` |  |
| `fiber_g` | `Float` |  |
| `id` | `Integer` |  |
| `image_thumb_url` | `String` |  |
| `image_url` | `String` |  |
| `name` | `String` |  |
| `potassium_mg` | `Float` |  |
| `protein_g` | `Float` |  |
| `saturated_fat_g` | `Float` |  |
| `serving_desc` | `String` |  |
| `serving_size_g` | `Float` |  |
| `sodium_mg` | `Float` |  |
| `source` | `String` |  |
| `static_url` | `String` |  |
| `sugar_g` | `Float` |  |

#### Example: List

```ruby
# list returns an Array of Popular records (raises on error).
populars = client.Popular.list
```


### Search

Create an instance: `search = client.Search`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `String` |  |
| `brand` | `String` |  |
| `calories_kcal` | `Float` |  |
| `carbs_g` | `Float` |  |
| `category` | `String` |  |
| `cholesterol_mg` | `Float` |  |
| `confidence` | `Float` |  |
| `fat_g` | `Float` |  |
| `fiber_g` | `Float` |  |
| `id` | `Integer` |  |
| `image_thumb_url` | `String` |  |
| `image_url` | `String` |  |
| `name` | `String` |  |
| `potassium_mg` | `Float` |  |
| `protein_g` | `Float` |  |
| `saturated_fat_g` | `Float` |  |
| `serving_desc` | `String` |  |
| `serving_size_g` | `Float` |  |
| `sodium_mg` | `Float` |  |
| `source` | `String` |  |
| `static_url` | `String` |  |
| `sugar_g` | `Float` |  |

#### Example: List

```ruby
# list returns an Array of Search records (raises on error).
searchs = client.Search.list
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── DietlyapiIntegration_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`DietlyapiIntegration_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
barcode = client.Barcode
barcode.load({ "id" => "example_id" })

# barcode.data_get now returns the barcode data from the last load
# barcode.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
