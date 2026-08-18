# DietlyapiIntegration SDK configuration


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "DietlyapiIntegration",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://api.getdietly.com",
            "auth": {
                "prefix": "Bearer",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "barcode": {},
                "food": {},
                "meta": {},
                "popular": {},
                "search": {},
            },
        },
        "entity": {
      "barcode": {
        "fields": [
          {
            "name": "barcode",
            "type": "`$STRING`",
          },
          {
            "name": "brand",
            "type": "`$STRING`",
          },
          {
            "name": "calories_kcal",
            "type": "`$NUMBER`",
          },
          {
            "name": "carbs_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "category",
            "type": "`$STRING`",
          },
          {
            "name": "cholesterol_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "confidence",
            "type": "`$NUMBER`",
          },
          {
            "name": "fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "fiber_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "type": "`$INTEGER`",
          },
          {
            "name": "image_thumb_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "potassium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "protein_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "saturated_fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "serving_desc",
            "type": "`$STRING`",
          },
          {
            "name": "serving_size_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "sodium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "source",
            "type": "`$STRING`",
          },
          {
            "name": "static_url",
            "type": "`$STRING`",
          },
          {
            "name": "sugar_g",
            "type": "`$NUMBER`",
          },
        ],
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
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/barcode/{code}",
                "parts": [
                  "barcode",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "code": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "food": {
        "fields": [
          {
            "name": "barcode",
            "type": "`$STRING`",
          },
          {
            "name": "brand",
            "type": "`$STRING`",
          },
          {
            "name": "calories_kcal",
            "type": "`$NUMBER`",
          },
          {
            "name": "carbs_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "category",
            "type": "`$STRING`",
          },
          {
            "name": "cholesterol_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "confidence",
            "type": "`$NUMBER`",
          },
          {
            "name": "count",
            "type": "`$INTEGER`",
          },
          {
            "name": "fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "fiber_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "type": "`$INTEGER`",
          },
          {
            "name": "image_thumb_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "potassium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "protein_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "saturated_fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "serving_desc",
            "type": "`$STRING`",
          },
          {
            "name": "serving_size_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "sodium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "source",
            "type": "`$STRING`",
          },
          {
            "name": "static_url",
            "type": "`$STRING`",
          },
          {
            "name": "sugar_g",
            "type": "`$NUMBER`",
          },
        ],
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
                "parts": [
                  "foods",
                  "categories",
                ],
                "select": {
                  "$action": "category",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
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
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/food/{food_id}",
                "parts": [
                  "food",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "food_id": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "meta": {
        "fields": [
          {
            "name": "foods_in_db",
            "type": "`$INTEGER`",
          },
          {
            "name": "status",
            "type": "`$STRING`",
          },
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
                "parts": [
                  "health",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "popular": {
        "fields": [
          {
            "name": "barcode",
            "type": "`$STRING`",
          },
          {
            "name": "brand",
            "type": "`$STRING`",
          },
          {
            "name": "calories_kcal",
            "type": "`$NUMBER`",
          },
          {
            "name": "carbs_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "category",
            "type": "`$STRING`",
          },
          {
            "name": "cholesterol_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "confidence",
            "type": "`$NUMBER`",
          },
          {
            "name": "fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "fiber_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "type": "`$INTEGER`",
          },
          {
            "name": "image_thumb_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "potassium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "protein_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "saturated_fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "serving_desc",
            "type": "`$STRING`",
          },
          {
            "name": "serving_size_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "sodium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "source",
            "type": "`$STRING`",
          },
          {
            "name": "static_url",
            "type": "`$STRING`",
          },
          {
            "name": "sugar_g",
            "type": "`$NUMBER`",
          },
        ],
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
                      "type": "`$STRING`",
                    },
                    {
                      "example": True,
                      "kind": "query",
                      "name": "has_image",
                      "orig": "has_image",
                      "type": "`$BOOLEAN`",
                    },
                    {
                      "example": 100,
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                    },
                    {
                      "example": 0,
                      "kind": "query",
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/foods/popular",
                "parts": [
                  "foods",
                  "popular",
                ],
                "select": {
                  "exist": [
                    "category",
                    "has_image",
                    "limit",
                    "offset",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "search": {
        "fields": [
          {
            "name": "barcode",
            "type": "`$STRING`",
          },
          {
            "name": "brand",
            "type": "`$STRING`",
          },
          {
            "name": "calories_kcal",
            "type": "`$NUMBER`",
          },
          {
            "name": "carbs_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "category",
            "type": "`$STRING`",
          },
          {
            "name": "cholesterol_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "confidence",
            "type": "`$NUMBER`",
          },
          {
            "name": "fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "fiber_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "id",
            "type": "`$INTEGER`",
          },
          {
            "name": "image_thumb_url",
            "type": "`$STRING`",
          },
          {
            "name": "image_url",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "potassium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "protein_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "saturated_fat_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "serving_desc",
            "type": "`$STRING`",
          },
          {
            "name": "serving_size_g",
            "type": "`$NUMBER`",
          },
          {
            "name": "sodium_mg",
            "type": "`$NUMBER`",
          },
          {
            "name": "source",
            "type": "`$STRING`",
          },
          {
            "name": "static_url",
            "type": "`$STRING`",
          },
          {
            "name": "sugar_g",
            "type": "`$NUMBER`",
          },
        ],
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
                      "type": "`$INTEGER`",
                    },
                    {
                      "example": "greek yogurt",
                      "kind": "query",
                      "name": "q",
                      "orig": "q",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "source",
                      "orig": "source",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/search",
                "parts": [
                  "search",
                ],
                "select": {
                  "exist": [
                    "limit",
                    "q",
                    "source",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
