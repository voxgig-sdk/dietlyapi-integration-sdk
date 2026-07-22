# frozen_string_literal: true

# Typed models for the DietlyapiIntegration SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Barcode entity data model.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
Barcode = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

# Request payload for Barcode#load.
#
# @!attribute [rw] id
#   @return [String]
BarcodeLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Food entity data model.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
Food = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :count,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

# Request payload for Food#load.
#
# @!attribute [rw] id
#   @return [Integer]
FoodLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Food#list.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
FoodListMatch = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :count,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

# Meta entity data model.
#
# @!attribute [rw] foods_in_db
#   @return [Integer, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
Meta = Struct.new(
  :foods_in_db,
  :status,
  keyword_init: true
)

# Request payload for Meta#load.
#
# @!attribute [rw] foods_in_db
#   @return [Integer, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
MetaLoadMatch = Struct.new(
  :foods_in_db,
  :status,
  keyword_init: true
)

# Popular entity data model.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
Popular = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

# Request payload for Popular#list.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
PopularListMatch = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
Search = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

# Request payload for Search#list.
#
# @!attribute [rw] barcode
#   @return [String, nil]
#
# @!attribute [rw] brand
#   @return [String, nil]
#
# @!attribute [rw] calories_kcal
#   @return [Float, nil]
#
# @!attribute [rw] carbs_g
#   @return [Float, nil]
#
# @!attribute [rw] category
#   @return [String, nil]
#
# @!attribute [rw] cholesterol_mg
#   @return [Float, nil]
#
# @!attribute [rw] confidence
#   @return [Float, nil]
#
# @!attribute [rw] fat_g
#   @return [Float, nil]
#
# @!attribute [rw] fiber_g
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [Integer, nil]
#
# @!attribute [rw] image_thumb_url
#   @return [String, nil]
#
# @!attribute [rw] image_url
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] potassium_mg
#   @return [Float, nil]
#
# @!attribute [rw] protein_g
#   @return [Float, nil]
#
# @!attribute [rw] saturated_fat_g
#   @return [Float, nil]
#
# @!attribute [rw] serving_desc
#   @return [String, nil]
#
# @!attribute [rw] serving_size_g
#   @return [Float, nil]
#
# @!attribute [rw] sodium_mg
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] static_url
#   @return [String, nil]
#
# @!attribute [rw] sugar_g
#   @return [Float, nil]
SearchListMatch = Struct.new(
  :barcode,
  :brand,
  :calories_kcal,
  :carbs_g,
  :category,
  :cholesterol_mg,
  :confidence,
  :fat_g,
  :fiber_g,
  :id,
  :image_thumb_url,
  :image_url,
  :name,
  :potassium_mg,
  :protein_g,
  :saturated_fat_g,
  :serving_desc,
  :serving_size_g,
  :sodium_mg,
  :source,
  :static_url,
  :sugar_g,
  keyword_init: true
)

