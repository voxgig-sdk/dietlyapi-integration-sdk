# DietlyapiIntegration PHP SDK



The PHP SDK for the DietlyapiIntegration API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Barcode()` — with named operations (`list`/`load`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dietlyapi-integration-sdk/releases](https://github.com/voxgig-sdk/dietlyapi-integration-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'dietlyapiintegration_sdk.php';

$client = new DietlyapiIntegrationSDK([
    "apikey" => getenv("DIETLYAPI_INTEGRATION_APIKEY"),
]);
```

### 3. Load a barcode

```php
try {
    // load() returns the ENTITY — call data_get() for the Barcode record (throws on error).
    $barcode = $client->Barcode()->load(["id" => "example_id"]);
    print_r($barcode);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $searchs = $client->Search()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = DietlyapiIntegrationSDK::test();

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$search = $client->Search()->list();
print_r($search);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new DietlyapiIntegrationSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
DIETLYAPI_INTEGRATION_TEST_LIVE=TRUE
DIETLYAPI_INTEGRATION_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### DietlyapiIntegrationSDK

```php
require_once 'dietlyapiintegration_sdk.php';
$client = new DietlyapiIntegrationSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = DietlyapiIntegrationSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### DietlyapiIntegrationSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Barcode` | `($data): BarcodeEntity` | Create a Barcode entity instance. |
| `Food` | `($data): FoodEntity` | Create a Food entity instance. |
| `Meta` | `($data): MetaEntity` | Create a Meta entity instance. |
| `Popular` | `($data): PopularEntity` | Create a Popular entity instance. |
| `Search` | `($data): SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

### Entities

#### Barcode

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` | Energy in kcal per 100 g |
| `carbs_g` | Carbohydrate in grams per 100 g |
| `category` |  |
| `cholesterol_mg` | Cholesterol in mg per 100 g |
| `confidence` | Data confidence score, 0–1 |
| `fat_g` | Total fat in grams per 100 g |
| `fiber_g` | Fibre in grams per 100 g |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` | Potassium in mg per 100 g |
| `protein_g` | Protein in grams per 100 g |
| `saturated_fat_g` | Saturated fat in grams per 100 g |
| `serving_desc` | Human-readable label for one serving, e.g. |
| `serving_size_g` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | Sodium in mg per 100 g |
| `source` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | Sugars in grams per 100 g |

Operations: Load.

API path: `/barcode/{code}`

#### Food

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` | Energy in kcal per 100 g |
| `carbs_g` | Carbohydrate in grams per 100 g |
| `category` |  |
| `cholesterol_mg` | Cholesterol in mg per 100 g |
| `confidence` | Data confidence score, 0–1 |
| `count` |  |
| `fat_g` | Total fat in grams per 100 g |
| `fiber_g` | Fibre in grams per 100 g |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` | Potassium in mg per 100 g |
| `protein_g` | Protein in grams per 100 g |
| `saturated_fat_g` | Saturated fat in grams per 100 g |
| `serving_desc` | Human-readable label for one serving, e.g. |
| `serving_size_g` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | Sodium in mg per 100 g |
| `source` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | Sugars in grams per 100 g |

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
| `calories_kcal` | Energy in kcal per 100 g |
| `carbs_g` | Carbohydrate in grams per 100 g |
| `category` |  |
| `cholesterol_mg` | Cholesterol in mg per 100 g |
| `confidence` | Data confidence score, 0–1 |
| `fat_g` | Total fat in grams per 100 g |
| `fiber_g` | Fibre in grams per 100 g |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` | Potassium in mg per 100 g |
| `protein_g` | Protein in grams per 100 g |
| `saturated_fat_g` | Saturated fat in grams per 100 g |
| `serving_desc` | Human-readable label for one serving, e.g. |
| `serving_size_g` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | Sodium in mg per 100 g |
| `source` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | Sugars in grams per 100 g |

Operations: List.

API path: `/foods/popular`

#### Search

| Field | Description |
| --- | --- |
| `barcode` |  |
| `brand` |  |
| `calories_kcal` | Energy in kcal per 100 g |
| `carbs_g` | Carbohydrate in grams per 100 g |
| `category` |  |
| `cholesterol_mg` | Cholesterol in mg per 100 g |
| `confidence` | Data confidence score, 0–1 |
| `fat_g` | Total fat in grams per 100 g |
| `fiber_g` | Fibre in grams per 100 g |
| `id` |  |
| `image_thumb_url` |  |
| `image_url` |  |
| `name` |  |
| `potassium_mg` | Potassium in mg per 100 g |
| `protein_g` | Protein in grams per 100 g |
| `saturated_fat_g` | Saturated fat in grams per 100 g |
| `serving_desc` | Human-readable label for one serving, e.g. |
| `serving_size_g` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | Sodium in mg per 100 g |
| `source` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | Sugars in grams per 100 g |

Operations: List.

API path: `/search`



## Entities


### Barcode

Create an instance: `$barcode = $client->Barcode();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `string` |  |
| `brand` | `string` |  |
| `calories_kcal` | `float` | Energy in kcal per 100 g |
| `carbs_g` | `float` | Carbohydrate in grams per 100 g |
| `category` | `string` |  |
| `cholesterol_mg` | `float` | Cholesterol in mg per 100 g |
| `confidence` | `float` | Data confidence score, 0–1 |
| `fat_g` | `float` | Total fat in grams per 100 g |
| `fiber_g` | `float` | Fibre in grams per 100 g |
| `id` | `int` |  |
| `image_thumb_url` | `string` |  |
| `image_url` | `string` |  |
| `name` | `string` |  |
| `potassium_mg` | `float` | Potassium in mg per 100 g |
| `protein_g` | `float` | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | Sodium in mg per 100 g |
| `source` | `string` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | Sugars in grams per 100 g |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Barcode record (throws on error).
$barcode = $client->Barcode()->load(["id" => "barcode_id"]);
```


### Food

Create an instance: `$food = $client->Food();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `string` |  |
| `brand` | `string` |  |
| `calories_kcal` | `float` | Energy in kcal per 100 g |
| `carbs_g` | `float` | Carbohydrate in grams per 100 g |
| `category` | `string` |  |
| `cholesterol_mg` | `float` | Cholesterol in mg per 100 g |
| `confidence` | `float` | Data confidence score, 0–1 |
| `count` | `int` |  |
| `fat_g` | `float` | Total fat in grams per 100 g |
| `fiber_g` | `float` | Fibre in grams per 100 g |
| `id` | `int` |  |
| `image_thumb_url` | `string` |  |
| `image_url` | `string` |  |
| `name` | `string` |  |
| `potassium_mg` | `float` | Potassium in mg per 100 g |
| `protein_g` | `float` | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | Sodium in mg per 100 g |
| `source` | `string` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | Sugars in grams per 100 g |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Food record (throws on error).
$food = $client->Food()->load(["id" => 1]);
```

#### Example: List

```php
// list() returns an array of Food records (throws on error).
$foods = $client->Food()->list();
```


### Meta

Create an instance: `$meta = $client->Meta();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `foods_in_db` | `int` |  |
| `status` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Meta record (throws on error).
$meta = $client->Meta()->load();
```


### Popular

Create an instance: `$popular = $client->Popular();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `string` |  |
| `brand` | `string` |  |
| `calories_kcal` | `float` | Energy in kcal per 100 g |
| `carbs_g` | `float` | Carbohydrate in grams per 100 g |
| `category` | `string` |  |
| `cholesterol_mg` | `float` | Cholesterol in mg per 100 g |
| `confidence` | `float` | Data confidence score, 0–1 |
| `fat_g` | `float` | Total fat in grams per 100 g |
| `fiber_g` | `float` | Fibre in grams per 100 g |
| `id` | `int` |  |
| `image_thumb_url` | `string` |  |
| `image_url` | `string` |  |
| `name` | `string` |  |
| `potassium_mg` | `float` | Potassium in mg per 100 g |
| `protein_g` | `float` | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | Sodium in mg per 100 g |
| `source` | `string` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | Sugars in grams per 100 g |

#### Example: List

```php
// list() returns an array of Popular records (throws on error).
$populars = $client->Popular()->list();
```


### Search

Create an instance: `$search = $client->Search();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `barcode` | `string` |  |
| `brand` | `string` |  |
| `calories_kcal` | `float` | Energy in kcal per 100 g |
| `carbs_g` | `float` | Carbohydrate in grams per 100 g |
| `category` | `string` |  |
| `cholesterol_mg` | `float` | Cholesterol in mg per 100 g |
| `confidence` | `float` | Data confidence score, 0–1 |
| `fat_g` | `float` | Total fat in grams per 100 g |
| `fiber_g` | `float` | Fibre in grams per 100 g |
| `id` | `int` |  |
| `image_thumb_url` | `string` |  |
| `image_url` | `string` |  |
| `name` | `string` |  |
| `potassium_mg` | `float` | Potassium in mg per 100 g |
| `protein_g` | `float` | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | Sodium in mg per 100 g |
| `source` | `string` | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | Sugars in grams per 100 g |

#### Example: List

```php
// list() returns an array of Search records (throws on error).
$searchs = $client->Search()->list();
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── dietlyapiintegration_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`dietlyapiintegration_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$search = $client->Search();
$search->list();

// $search->data_get() now returns the search data from the last list
// $search->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
