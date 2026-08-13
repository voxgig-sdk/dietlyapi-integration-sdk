# Typed models for the DietlyapiIntegration SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Barcode(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float


class BarcodeLoadMatch(TypedDict):
    id: str


class Food(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    count: int
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float


class FoodLoadMatch(TypedDict):
    id: int


class FoodListMatch(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    count: int
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float


class Meta(TypedDict, total=False):
    foods_in_db: int
    status: str


class MetaLoadMatch(TypedDict, total=False):
    foods_in_db: int
    status: str


class Popular(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float


class PopularListMatch(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float


class Search(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float


class SearchListMatch(TypedDict, total=False):
    barcode: str
    brand: str
    calories_kcal: float
    carbs_g: float
    category: str
    cholesterol_mg: float
    confidence: float
    fat_g: float
    fiber_g: float
    id: int
    image_thumb_url: str
    image_url: str
    name: str
    potassium_mg: float
    protein_g: float
    saturated_fat_g: float
    serving_desc: str
    serving_size_g: float
    sodium_mg: float
    source: str
    static_url: str
    sugar_g: float
