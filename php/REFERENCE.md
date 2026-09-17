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
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

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
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

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
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

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
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `string` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `string` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `string` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

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
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```php
$client = new DietlyapiIntegrationSDK([
  "feature" => [
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
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

