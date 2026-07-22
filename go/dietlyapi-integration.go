package voxgigdietlyapiintegrationsdk

import (
	"github.com/voxgig-sdk/dietlyapi-integration-sdk/go/core"
	"github.com/voxgig-sdk/dietlyapi-integration-sdk/go/entity"
	"github.com/voxgig-sdk/dietlyapi-integration-sdk/go/feature"
	_ "github.com/voxgig-sdk/dietlyapi-integration-sdk/go/utility"
)

// Type aliases preserve external API.
type DietlyapiIntegrationSDK = core.DietlyapiIntegrationSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type DietlyapiIntegrationEntity = core.DietlyapiIntegrationEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type DietlyapiIntegrationError = core.DietlyapiIntegrationError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewBarcodeEntityFunc = func(client *core.DietlyapiIntegrationSDK, entopts map[string]any) core.DietlyapiIntegrationEntity {
		return entity.NewBarcodeEntity(client, entopts)
	}
	core.NewFoodEntityFunc = func(client *core.DietlyapiIntegrationSDK, entopts map[string]any) core.DietlyapiIntegrationEntity {
		return entity.NewFoodEntity(client, entopts)
	}
	core.NewMetaEntityFunc = func(client *core.DietlyapiIntegrationSDK, entopts map[string]any) core.DietlyapiIntegrationEntity {
		return entity.NewMetaEntity(client, entopts)
	}
	core.NewPopularEntityFunc = func(client *core.DietlyapiIntegrationSDK, entopts map[string]any) core.DietlyapiIntegrationEntity {
		return entity.NewPopularEntity(client, entopts)
	}
	core.NewSearchEntityFunc = func(client *core.DietlyapiIntegrationSDK, entopts map[string]any) core.DietlyapiIntegrationEntity {
		return entity.NewSearchEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewDietlyapiIntegrationSDK = core.NewDietlyapiIntegrationSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewDietlyapiIntegrationSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *DietlyapiIntegrationSDK  { return NewDietlyapiIntegrationSDK(nil) }
func Test() *DietlyapiIntegrationSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
