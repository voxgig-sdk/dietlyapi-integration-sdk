// Typed models for the DietlyapiIntegration SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Barcode is the typed data model for the barcode entity.
type Barcode struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// BarcodeLoadMatch is the typed request payload for Barcode.LoadTyped.
type BarcodeLoadMatch struct {
	Id string `json:"id"`
}

// Food is the typed data model for the food entity.
type Food struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	Count *int `json:"count,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// FoodLoadMatch is the typed request payload for Food.LoadTyped.
type FoodLoadMatch struct {
	Id int `json:"id"`
}

// FoodListMatch is the typed request payload for Food.ListTyped.
type FoodListMatch struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	Count *int `json:"count,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// Meta is the typed data model for the meta entity.
type Meta struct {
	FoodsInDb *int `json:"foods_in_db,omitempty"`
	Status *string `json:"status,omitempty"`
}

// MetaLoadMatch is the typed request payload for Meta.LoadTyped.
type MetaLoadMatch struct {
	FoodsInDb *int `json:"foods_in_db,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Popular is the typed data model for the popular entity.
type Popular struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// PopularListMatch is the typed request payload for Popular.ListTyped.
type PopularListMatch struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// Search is the typed data model for the search entity.
type Search struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	Barcode *string `json:"barcode,omitempty"`
	Brand *string `json:"brand,omitempty"`
	CaloriesKcal *float64 `json:"calories_kcal,omitempty"`
	CarbsG *float64 `json:"carbs_g,omitempty"`
	Category *string `json:"category,omitempty"`
	CholesterolMg *float64 `json:"cholesterol_mg,omitempty"`
	Confidence *float64 `json:"confidence,omitempty"`
	FatG *float64 `json:"fat_g,omitempty"`
	FiberG *float64 `json:"fiber_g,omitempty"`
	Id *int `json:"id,omitempty"`
	ImageThumbUrl *string `json:"image_thumb_url,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Name *string `json:"name,omitempty"`
	PotassiumMg *float64 `json:"potassium_mg,omitempty"`
	ProteinG *float64 `json:"protein_g,omitempty"`
	SaturatedFatG *float64 `json:"saturated_fat_g,omitempty"`
	ServingDesc *string `json:"serving_desc,omitempty"`
	ServingSizeG *float64 `json:"serving_size_g,omitempty"`
	SodiumMg *float64 `json:"sodium_mg,omitempty"`
	Source *string `json:"source,omitempty"`
	StaticUrl *string `json:"static_url,omitempty"`
	SugarG *float64 `json:"sugar_g,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
