// Typed models for the DietlyapiIntegration SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Barcode {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number
}

export interface BarcodeLoadMatch {
  id: string
}

export interface Food {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  count?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number
}

export interface FoodLoadMatch {
  id: number
}

export interface FoodListMatch {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  count?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number

  // Selects a custom action instead of the plain list:
  //   'category'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface Meta {
  foods_in_db?: number
  status?: string
}

export interface MetaLoadMatch {
  foods_in_db?: number
  status?: string
}

export interface Popular {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number
}

export interface PopularListMatch {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number
}

export interface Search {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number
}

export interface SearchListMatch {
  barcode?: string
  brand?: string
  calories_kcal?: number
  carbs_g?: number
  category?: string
  cholesterol_mg?: number
  confidence?: number
  fat_g?: number
  fiber_g?: number
  id?: number
  image_thumb_url?: string
  image_url?: string
  name?: string
  potassium_mg?: number
  protein_g?: number
  saturated_fat_g?: number
  serving_desc?: string
  serving_size_g?: number
  sodium_mg?: number
  source?: string
  static_url?: string
  sugar_g?: number
}

