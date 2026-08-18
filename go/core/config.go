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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"type": "`$NUMBER`",
					},
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
								"parts": []any{
									"barcode",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"code": "id",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "count",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"type": "`$NUMBER`",
					},
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
								"parts": []any{
									"foods",
									"categories",
								},
								"select": map[string]any{
									"$action": "category",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
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
								"parts": []any{
									"food",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"food_id": "id",
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
								"parts": []any{
									"health",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"type": "`$NUMBER`",
					},
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
								"parts": []any{
									"foods",
									"popular",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "carbs_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cholesterol_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "confidence",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fiber_g",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "protein_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "saturated_fat_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "serving_desc",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "serving_size_g",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sodium_mg",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "static_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sugar_g",
						"type": "`$NUMBER`",
					},
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
								"parts": []any{
									"search",
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
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
