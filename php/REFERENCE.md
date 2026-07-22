# DietlyapiIntegration PHP SDK Reference

Complete API reference for the DietlyapiIntegration PHP SDK.


## DietlyapiIntegrationSDK

### Constructor

```php
require_once __DIR__ . '/dietlyapiintegration_sdk.php';

$client = new DietlyapiIntegrationSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DietlyapiIntegrationSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = DietlyapiIntegrationSDK::test();
```


### Instance Methods

#### `Barcode($data = null)`

Create a new `BarcodeEntity` instance. Pass `null` for no initial data.

#### `Food($data = null)`

Create a new `FoodEntity` instance. Pass `null` for no initial data.

#### `Meta($data = null)`

Create a new `MetaEntity` instance. Pass `null` for no initial data.

#### `Popular($data = null)`

Create a new `PopularEntity` instance. Pass `null` for no initial data.

#### `Search($data = null)`

Create a new `SearchEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): DietlyapiIntegrationUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## BarcodeEntity

```php
$barcode = $client->Barcode();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Barcode()->load(["id" => "barcode_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BarcodeEntity`

Create a new `BarcodeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FoodEntity

```php
$food = $client->Food();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `count` | `int` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Food()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Food()->load(["id" => 1]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FoodEntity`

Create a new `FoodEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MetaEntity

```php
$meta = $client->Meta();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `foods_in_db` | `int` | No |  |
| `status` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Meta()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MetaEntity`

Create a new `MetaEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PopularEntity

```php
$popular = $client->Popular();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Popular()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PopularEntity`

Create a new `PopularEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SearchEntity

```php
$search = $client->Search();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Search()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SearchEntity`

Create a new `SearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new DietlyapiIntegrationSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

