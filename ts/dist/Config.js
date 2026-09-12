"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FEATURE_PLUGINS = exports.config = void 0;
const TestFeature_1 = require("./feature/test/TestFeature");
const FEATURE_CLASS = {
    test: TestFeature_1.TestFeature,
};
// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS = {};
exports.FEATURE_PLUGINS = FEATURE_PLUGINS;
class Config {
    makeFeature(fn) {
        const fc = FEATURE_CLASS[fn];
        const fi = new fc();
        // TODO: errors etc
        return fi;
    }
    // False for a feature added at runtime via options.extend (station's
    // adopt path) - the constructor uses this to skip makeFeature for names
    // no generated class backs.
    hasFeature(fn) {
        return null != FEATURE_CLASS[fn];
    }
    main = {
        name: 'DietlyapiIntegration',
        slug: "dietlyapi-integration",
        version: "0.0.1",
        target: "ts",
    };
    feature = {
        test: {
            "options": {
                "active": false
            },
            "transport": "base"
        },
    };
    options = {
        base: "https://api.getdietly.com",
        auth: {
            prefix: 'Bearer',
        },
        headers: {
            "content-type": "application/json"
        },
        entity: {
            barcode: {},
            food: {},
            meta: {},
            popular: {},
            search: {},
        }
    };
    entity = {
        "barcode": {
            "fields": [
                {
                    "name": "barcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "brand",
                    "type": "`$STRING`"
                },
                {
                    "name": "calories_kcal",
                    "short": "Energy in kcal per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "carbs_g",
                    "short": "Carbohydrate in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "category",
                    "type": "`$STRING`"
                },
                {
                    "name": "cholesterol_mg",
                    "short": "Cholesterol in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "confidence",
                    "short": "Data confidence score, 0–1",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fat_g",
                    "short": "Total fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fiber_g",
                    "short": "Fibre in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "id",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "image_thumb_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "image_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "potassium_mg",
                    "short": "Potassium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "protein_g",
                    "short": "Protein in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "saturated_fat_g",
                    "short": "Saturated fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "serving_desc",
                    "short": "Human-readable label for one serving, e.g.",
                    "type": "`$STRING`"
                },
                {
                    "name": "serving_size_g",
                    "short": "Grams in one manufacturer serving, where the source declares one.",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "sodium_mg",
                    "short": "Sodium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "source",
                    "short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
                    "type": "`$STRING`"
                },
                {
                    "name": "static_url",
                    "short": "Path of the human-readable page on www.getdietly.com",
                    "type": "`$STRING`"
                },
                {
                    "name": "sugar_g",
                    "short": "Sugars in grams per 100 g",
                    "type": "`$NUMBER`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "barcode",
            "op": {
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "args": {
                                "params": [
                                    {
                                        "example": "0855088005245",
                                        "kind": "param",
                                        "name": "id",
                                        "orig": "code",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/barcode/{code}",
                            "rename": {
                                "param": {
                                    "code": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "barcode"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "barcode",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "food": {
            "fields": [
                {
                    "name": "barcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "brand",
                    "type": "`$STRING`"
                },
                {
                    "name": "calories_kcal",
                    "short": "Energy in kcal per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "carbs_g",
                    "short": "Carbohydrate in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "category",
                    "type": "`$STRING`"
                },
                {
                    "name": "cholesterol_mg",
                    "short": "Cholesterol in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "confidence",
                    "short": "Data confidence score, 0–1",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "count",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "fat_g",
                    "short": "Total fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fiber_g",
                    "short": "Fibre in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "id",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "image_thumb_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "image_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "potassium_mg",
                    "short": "Potassium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "protein_g",
                    "short": "Protein in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "saturated_fat_g",
                    "short": "Saturated fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "serving_desc",
                    "short": "Human-readable label for one serving, e.g.",
                    "type": "`$STRING`"
                },
                {
                    "name": "serving_size_g",
                    "short": "Grams in one manufacturer serving, where the source declares one.",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "sodium_mg",
                    "short": "Sodium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "source",
                    "short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
                    "type": "`$STRING`"
                },
                {
                    "name": "static_url",
                    "short": "Path of the human-readable page on www.getdietly.com",
                    "type": "`$STRING`"
                },
                {
                    "name": "sugar_g",
                    "short": "Sugars in grams per 100 g",
                    "type": "`$NUMBER`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "food",
            "op": {
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "args": {},
                            "kind": "http",
                            "method": "GET",
                            "orig": "/foods/categories",
                            "segments": [
                                {
                                    "lit": "foods"
                                },
                                {
                                    "lit": "categories"
                                }
                            ],
                            "select": {
                                "$action": "category"
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "foods",
                                "categories"
                            ]
                        }
                    ]
                },
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "args": {
                                "params": [
                                    {
                                        "example": 1068319,
                                        "kind": "param",
                                        "name": "id",
                                        "orig": "food_id",
                                        "reqd": true,
                                        "type": "`$INTEGER`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/food/{food_id}",
                            "rename": {
                                "param": {
                                    "food_id": "id"
                                }
                            },
                            "segments": [
                                {
                                    "lit": "food"
                                },
                                {
                                    "var": "id"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "id"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "food",
                                "{id}"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "meta": {
            "fields": [
                {
                    "name": "foods_in_db",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "status",
                    "type": "`$STRING`"
                }
            ],
            "name": "meta",
            "op": {
                "load": {
                    "input": "data",
                    "name": "load",
                    "points": [
                        {
                            "args": {},
                            "kind": "http",
                            "method": "GET",
                            "orig": "/health",
                            "segments": [
                                {
                                    "lit": "health"
                                }
                            ],
                            "select": {},
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "health"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "popular": {
            "fields": [
                {
                    "name": "barcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "brand",
                    "type": "`$STRING`"
                },
                {
                    "name": "calories_kcal",
                    "short": "Energy in kcal per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "carbs_g",
                    "short": "Carbohydrate in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "category",
                    "type": "`$STRING`"
                },
                {
                    "name": "cholesterol_mg",
                    "short": "Cholesterol in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "confidence",
                    "short": "Data confidence score, 0–1",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fat_g",
                    "short": "Total fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fiber_g",
                    "short": "Fibre in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "id",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "image_thumb_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "image_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "potassium_mg",
                    "short": "Potassium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "protein_g",
                    "short": "Protein in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "saturated_fat_g",
                    "short": "Saturated fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "serving_desc",
                    "short": "Human-readable label for one serving, e.g.",
                    "type": "`$STRING`"
                },
                {
                    "name": "serving_size_g",
                    "short": "Grams in one manufacturer serving, where the source declares one.",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "sodium_mg",
                    "short": "Sodium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "source",
                    "short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
                    "type": "`$STRING`"
                },
                {
                    "name": "static_url",
                    "short": "Path of the human-readable page on www.getdietly.com",
                    "type": "`$STRING`"
                },
                {
                    "name": "sugar_g",
                    "short": "Sugars in grams per 100 g",
                    "type": "`$NUMBER`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "popular",
            "op": {
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "kind": "query",
                                        "name": "category",
                                        "orig": "category",
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "example": true,
                                        "kind": "query",
                                        "name": "has_image",
                                        "orig": "has_image",
                                        "type": "`$BOOLEAN`"
                                    },
                                    {
                                        "example": 100,
                                        "kind": "query",
                                        "name": "limit",
                                        "orig": "limit",
                                        "type": "`$INTEGER`"
                                    },
                                    {
                                        "example": 0,
                                        "kind": "query",
                                        "name": "offset",
                                        "orig": "offset",
                                        "type": "`$INTEGER`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/foods/popular",
                            "segments": [
                                {
                                    "lit": "foods"
                                },
                                {
                                    "lit": "popular"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "category",
                                    "has_image",
                                    "limit",
                                    "offset"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "foods",
                                "popular"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        },
        "search": {
            "fields": [
                {
                    "name": "barcode",
                    "type": "`$STRING`"
                },
                {
                    "name": "brand",
                    "type": "`$STRING`"
                },
                {
                    "name": "calories_kcal",
                    "short": "Energy in kcal per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "carbs_g",
                    "short": "Carbohydrate in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "category",
                    "type": "`$STRING`"
                },
                {
                    "name": "cholesterol_mg",
                    "short": "Cholesterol in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "confidence",
                    "short": "Data confidence score, 0–1",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fat_g",
                    "short": "Total fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "fiber_g",
                    "short": "Fibre in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "id",
                    "type": "`$INTEGER`"
                },
                {
                    "name": "image_thumb_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "image_url",
                    "type": "`$STRING`"
                },
                {
                    "name": "name",
                    "type": "`$STRING`"
                },
                {
                    "name": "potassium_mg",
                    "short": "Potassium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "protein_g",
                    "short": "Protein in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "saturated_fat_g",
                    "short": "Saturated fat in grams per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "serving_desc",
                    "short": "Human-readable label for one serving, e.g.",
                    "type": "`$STRING`"
                },
                {
                    "name": "serving_size_g",
                    "short": "Grams in one manufacturer serving, where the source declares one.",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "sodium_mg",
                    "short": "Sodium in mg per 100 g",
                    "type": "`$NUMBER`"
                },
                {
                    "name": "source",
                    "short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
                    "type": "`$STRING`"
                },
                {
                    "name": "static_url",
                    "short": "Path of the human-readable page on www.getdietly.com",
                    "type": "`$STRING`"
                },
                {
                    "name": "sugar_g",
                    "short": "Sugars in grams per 100 g",
                    "type": "`$NUMBER`"
                }
            ],
            "id": {
                "field": "id",
                "name": "id"
            },
            "name": "search",
            "op": {
                "list": {
                    "input": "data",
                    "name": "list",
                    "points": [
                        {
                            "args": {
                                "query": [
                                    {
                                        "example": 5,
                                        "kind": "query",
                                        "name": "limit",
                                        "orig": "limit",
                                        "type": "`$INTEGER`"
                                    },
                                    {
                                        "example": "greek yogurt",
                                        "kind": "query",
                                        "name": "q",
                                        "orig": "q",
                                        "reqd": true,
                                        "type": "`$STRING`"
                                    },
                                    {
                                        "kind": "query",
                                        "name": "source",
                                        "orig": "source",
                                        "type": "`$STRING`"
                                    }
                                ]
                            },
                            "kind": "http",
                            "method": "GET",
                            "orig": "/search",
                            "segments": [
                                {
                                    "lit": "search"
                                }
                            ],
                            "select": {
                                "exist": [
                                    "limit",
                                    "q",
                                    "source"
                                ]
                            },
                            "transform": {
                                "req": "`reqdata`",
                                "res": "`body`"
                            },
                            "parts": [
                                "search"
                            ]
                        }
                    ]
                }
            },
            "relations": {
                "ancestors": []
            }
        }
    };
}
const config = new Config();
exports.config = config;
//# sourceMappingURL=Config.js.map