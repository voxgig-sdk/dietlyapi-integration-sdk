# DietlyapiIntegration Golang SDK Reference

Complete API reference for the DietlyapiIntegration Golang SDK.


## DietlyapiIntegrationSDK

### Constructor

```go
func NewDietlyapiIntegrationSDK(options map[string]any) *DietlyapiIntegrationSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *DietlyapiIntegrationSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *DietlyapiIntegrationSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Barcode(data map[string]any) DietlyapiIntegrationEntity`

Create a new `Barcode` entity instance. Pass `nil` for no initial data.

#### `Food(data map[string]any) DietlyapiIntegrationEntity`

Create a new `Food` entity instance. Pass `nil` for no initial data.

#### `Meta(data map[string]any) DietlyapiIntegrationEntity`

Create a new `Meta` entity instance. Pass `nil` for no initial data.

#### `Popular(data map[string]any) DietlyapiIntegrationEntity`

Create a new `Popular` entity instance. Pass `nil` for no initial data.

#### `Search(data map[string]any) DietlyapiIntegrationEntity`

Create a new `Search` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## BarcodeEntity

```go
barcode := client.Barcode(nil)
fmt.Println(barcode.GetName()) // "barcode"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float64` | No |  |
| `carbs_g` | `float64` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float64` | No |  |
| `confidence` | `float64` | No |  |
| `fat_g` | `float64` | No |  |
| `fiber_g` | `float64` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float64` | No |  |
| `protein_g` | `float64` | No |  |
| `saturated_fat_g` | `float64` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float64` | No |  |
| `sodium_mg` | `float64` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float64` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Barcode(nil).Load(map[string]any{"id": "barcode_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BarcodeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FoodEntity

```go
food := client.Food(nil)
fmt.Println(food.GetName()) // "food"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float64` | No |  |
| `carbs_g` | `float64` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float64` | No |  |
| `confidence` | `float64` | No |  |
| `count` | `int` | No |  |
| `fat_g` | `float64` | No |  |
| `fiber_g` | `float64` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float64` | No |  |
| `protein_g` | `float64` | No |  |
| `saturated_fat_g` | `float64` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float64` | No |  |
| `sodium_mg` | `float64` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Food(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Food(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FoodEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MetaEntity

```go
meta := client.Meta(nil)
fmt.Println(meta.GetName()) // "meta"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `foods_in_db` | `int` | No |  |
| `status` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Meta(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MetaEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PopularEntity

```go
popular := client.Popular(nil)
fmt.Println(popular.GetName()) // "popular"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float64` | No |  |
| `carbs_g` | `float64` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float64` | No |  |
| `confidence` | `float64` | No |  |
| `fat_g` | `float64` | No |  |
| `fiber_g` | `float64` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float64` | No |  |
| `protein_g` | `float64` | No |  |
| `saturated_fat_g` | `float64` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float64` | No |  |
| `sodium_mg` | `float64` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Popular(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PopularEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SearchEntity

```go
search := client.Search(nil)
fmt.Println(search.GetName()) // "search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `barcode` | `string` | No |  |
| `brand` | `string` | No |  |
| `calories_kcal` | `float64` | No |  |
| `carbs_g` | `float64` | No |  |
| `category` | `string` | No |  |
| `cholesterol_mg` | `float64` | No |  |
| `confidence` | `float64` | No |  |
| `fat_g` | `float64` | No |  |
| `fiber_g` | `float64` | No |  |
| `id` | `int` | No |  |
| `image_thumb_url` | `string` | No |  |
| `image_url` | `string` | No |  |
| `name` | `string` | No |  |
| `potassium_mg` | `float64` | No |  |
| `protein_g` | `float64` | No |  |
| `saturated_fat_g` | `float64` | No |  |
| `serving_desc` | `string` | No |  |
| `serving_size_g` | `float64` | No |  |
| `sodium_mg` | `float64` | No |  |
| `source` | `string` | No |  |
| `static_url` | `string` | No |  |
| `sugar_g` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Search(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewDietlyapiIntegrationSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

