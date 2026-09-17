package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "DietlyapiIntegration",
			"slug": "dietlyapi-integration",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api.getdietly.com",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"barcode": map[string]any{},
				"food": map[string]any{},
				"meta": map[string]any{},
				"popular": map[string]any{},
				"search": map[string]any{},
			},
		},
		"entity": map[string]any{
			"barcode": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "barcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brand",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "calories_kcal",
						"short": "Energy in kcal per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"short": "Carbohydrate in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"short": "Cholesterol in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"short": "Data confidence score, 0–1",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"short": "Total fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
						"short": "Fibre in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "image_thumb_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "potassium_mg",
						"short": "Potassium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"short": "Protein in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"short": "Saturated fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"short": "Human-readable label for one serving, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"short": "Grams in one manufacturer serving, where the source declares one.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"short": "Sodium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"short": "Path of the human-readable page on www.getdietly.com",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"short": "Sugars in grams per 100 g",
						"type": "`$NUMBER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "barcode",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "0855088005245",
											"kind": "param",
											"name": "id",
											"orig": "code",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/barcode/{code}",
								"rename": map[string]any{
									"param": map[string]any{
										"code": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "barcode",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"barcode",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"food": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "barcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brand",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "calories_kcal",
						"short": "Energy in kcal per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"short": "Carbohydrate in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"short": "Cholesterol in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"short": "Data confidence score, 0–1",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"short": "Total fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
						"short": "Fibre in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "image_thumb_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "potassium_mg",
						"short": "Potassium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"short": "Protein in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"short": "Saturated fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"short": "Human-readable label for one serving, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"short": "Grams in one manufacturer serving, where the source declares one.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"short": "Sodium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"short": "Path of the human-readable page on www.getdietly.com",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"short": "Sugars in grams per 100 g",
						"type": "`$NUMBER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "food",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/foods/categories",
								"segments": []any{
									map[string]any{
										"lit": "foods",
									},
									map[string]any{
										"lit": "categories",
									},
								},
								"select": map[string]any{
									"$action": "category",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"foods",
									"categories",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 1068319,
											"kind": "param",
											"name": "id",
											"orig": "food_id",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/food/{food_id}",
								"rename": map[string]any{
									"param": map[string]any{
										"food_id": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "food",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"food",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"meta": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "foods_in_db",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
				},
				"name": "meta",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/health",
								"segments": []any{
									map[string]any{
										"lit": "health",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"health",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"popular": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "barcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brand",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "calories_kcal",
						"short": "Energy in kcal per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"short": "Carbohydrate in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"short": "Cholesterol in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"short": "Data confidence score, 0–1",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"short": "Total fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
						"short": "Fibre in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "image_thumb_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "potassium_mg",
						"short": "Potassium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"short": "Protein in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"short": "Saturated fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"short": "Human-readable label for one serving, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"short": "Grams in one manufacturer serving, where the source declares one.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"short": "Sodium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"short": "Path of the human-readable page on www.getdietly.com",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"short": "Sugars in grams per 100 g",
						"type": "`$NUMBER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "popular",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "category",
											"orig": "category",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": true,
											"kind": "query",
											"name": "has_image",
											"orig": "has_image",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/foods/popular",
								"segments": []any{
									map[string]any{
										"lit": "foods",
									},
									map[string]any{
										"lit": "popular",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"category",
										"has_image",
										"limit",
										"offset",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"foods",
									"popular",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "barcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "brand",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "calories_kcal",
						"short": "Energy in kcal per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"short": "Carbohydrate in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"short": "Cholesterol in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"short": "Data confidence score, 0–1",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"short": "Total fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
						"short": "Fibre in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "image_thumb_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "image_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "potassium_mg",
						"short": "Potassium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"short": "Protein in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"short": "Saturated fat in grams per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"short": "Human-readable label for one serving, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"short": "Grams in one manufacturer serving, where the source declares one.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"short": "Sodium in mg per 100 g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"short": "Data source: off (Open Food Facts), usda (USDA FoodData Central), claude (labeled AI estimate) or community",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"short": "Path of the human-readable page on www.getdietly.com",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"short": "Sugars in grams per 100 g",
						"type": "`$NUMBER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": 5,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "greek yogurt",
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "source",
											"orig": "source",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/search",
								"segments": []any{
									map[string]any{
										"lit": "search",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"q",
										"source",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"search",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
