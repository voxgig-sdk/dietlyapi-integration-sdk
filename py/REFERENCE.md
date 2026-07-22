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
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `str` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `str` | No |  |
| `static_url` | `str` | No |  |
| `sugar_g` | `float` | No |  |

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
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `count` | `int` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `str` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `str` | No |  |
| `static_url` | `str` | No |  |
| `sugar_g` | `float` | No |  |

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
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `str` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `str` | No |  |
| `static_url` | `str` | No |  |
| `sugar_g` | `float` | No |  |

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
| `calories_kcal` | `float` | No |  |
| `carbs_g` | `float` | No |  |
| `category` | `str` | No |  |
| `cholesterol_mg` | `float` | No |  |
| `confidence` | `float` | No |  |
| `fat_g` | `float` | No |  |
| `fiber_g` | `float` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `str` | No |  |
| `image_url` | `str` | No |  |
| `name` | `str` | No |  |
| `potassium_mg` | `float` | No |  |
| `protein_g` | `float` | No |  |
| `saturated_fat_g` | `float` | No |  |
| `serving_desc` | `str` | No |  |
| `serving_size_g` | `float` | No |  |
| `sodium_mg` | `float` | No |  |
| `source` | `str` | No |  |
| `static_url` | `str` | No |  |
| `sugar_g` | `float` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Search().list()
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
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = DietlyapiIntegrationSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

