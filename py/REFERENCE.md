# DietlyapiIntegration Python SDK Reference

Complete API reference for the DietlyapiIntegration Python SDK.


## DietlyapiIntegrationSDK

### Constructor

```python
from dietlyapiintegration_sdk import DietlyapiIntegrationSDK

client = DietlyapiIntegrationSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DietlyapiIntegrationSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = DietlyapiIntegrationSDK.test()
```


### Instance Methods

#### `Barcode(data=None)`

Create a new `BarcodeEntity` instance. Pass `None` for no initial data.

#### `Food(data=None)`

Create a new `FoodEntity` instance. Pass `None` for no initial data.

#### `Meta(data=None)`

Create a new `MetaEntity` instance. Pass `None` for no initial data.

#### `Popular(data=None)`

Create a new `PopularEntity` instance. Pass `None` for no initial data.

#### `Search(data=None)`

Create a new `SearchEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## BarcodeEntity

```python
barcode = client.Barcode()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `str` | No |  |
| `brand` | `str` | No |  |
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `str` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `str` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `str` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Barcode().load({"id": "barcode_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BarcodeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FoodEntity

```python
food = client.Food()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `str` | No |  |
| `brand` | `str` | No |  |
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `count` | `int` | No |  |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `str` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `str` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `str` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Food().list()
for food in results:
    print(food)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Food().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FoodEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MetaEntity

```python
meta = client.Meta()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `foods_in_db` | `int` | No |  |
| `status` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Meta().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetaEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PopularEntity

```python
popular = client.Popular()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `str` | No |  |
| `brand` | `str` | No |  |
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `str` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `str` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `str` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Popular().list()
for popular in results:
    print(popular)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PopularEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SearchEntity

```python
search = client.Search()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `str` | No |  |
| `brand` | `str` | No |  |
| `calories_kcal` | `float` | No | Energy in kcal per 100 g |
| `carbs_g` | `float` | No | Carbohydrate in grams per 100 g |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No | Cholesterol in mg per 100 g |
| `confidence` | `float` | No | Data confidence score, 0–1 |
| `fat_g` | `float` | No | Total fat in grams per 100 g |
| `fiber_g` | `float` | No | Fibre in grams per 100 g |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No | Potassium in mg per 100 g |
| `protein_g` | `float` | No | Protein in grams per 100 g |
| `saturated_fat_g` | `float` | No | Saturated fat in grams per 100 g |
| `serving_desc` | `str` | No | Human-readable label for one serving, e.g. |
| `serving_size_g` | `float` | No | Grams in one manufacturer serving, where the source declares one. |
| `sodium_mg` | `float` | No | Sodium in mg per 100 g |
| `source` | `str` | No | Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community |
| `static_url` | `str` | No | Path of the human-readable page on www.getdietly.com |
| `sugar_g` | `float` | No | Sugars in grams per 100 g |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list({"q": "example"})
for search in results:
    print(search)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SearchEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = DietlyapiIntegrationSDK({
    "feature": {
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

