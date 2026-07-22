<?php
declare(strict_types=1);

// Typed models for the DietlyapiIntegration SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Barcode entity data model. */
class Barcode
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

/** Request payload for Barcode#load. */
class BarcodeLoadMatch
{
    public string $id;
}

/** Food entity data model. */
class Food
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?int $count = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

/** Request payload for Food#load. */
class FoodLoadMatch
{
    public int $id;
}

/** Request payload for Food#list. */
class FoodListMatch
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?int $count = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

/** Meta entity data model. */
class Meta
{
    public ?int $foods_in_db = null;
    public ?string $status = null;
}

/** Request payload for Meta#load. */
class MetaLoadMatch
{
    public ?int $foods_in_db = null;
    public ?string $status = null;
}

/** Popular entity data model. */
class Popular
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

/** Request payload for Popular#list. */
class PopularListMatch
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

/** Search entity data model. */
class Search
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

/** Request payload for Search#list. */
class SearchListMatch
{
    public ?string $barcode = null;
    public ?string $brand = null;
    public ?float $calories_kcal = null;
    public ?float $carbs_g = null;
    public ?string $category = null;
    public ?float $cholesterol_mg = null;
    public ?float $confidence = null;
    public ?float $fat_g = null;
    public ?float $fiber_g = null;
    public ?int $id = null;
    public ?string $image_thumb_url = null;
    public ?string $image_url = null;
    public ?string $name = null;
    public ?float $potassium_mg = null;
    public ?float $protein_g = null;
    public ?float $saturated_fat_g = null;
    public ?string $serving_desc = null;
    public ?float $serving_size_g = null;
    public ?float $sodium_mg = null;
    public ?string $source = null;
    public ?string $static_url = null;
    public ?float $sugar_g = null;
}

