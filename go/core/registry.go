package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewBarcodeEntityFunc func(client *DietlyapiIntegrationSDK, entopts map[string]any) DietlyapiIntegrationEntity

var NewFoodEntityFunc func(client *DietlyapiIntegrationSDK, entopts map[string]any) DietlyapiIntegrationEntity

var NewMetaEntityFunc func(client *DietlyapiIntegrationSDK, entopts map[string]any) DietlyapiIntegrationEntity

var NewPopularEntityFunc func(client *DietlyapiIntegrationSDK, entopts map[string]any) DietlyapiIntegrationEntity

var NewSearchEntityFunc func(client *DietlyapiIntegrationSDK, entopts map[string]any) DietlyapiIntegrationEntity

