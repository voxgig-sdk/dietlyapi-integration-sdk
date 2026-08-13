# DietlyapiIntegration SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

DietlyapiIntegrationUtility.registrar = ->(u) {
  u.clean = DietlyapiIntegrationUtilities::Clean
  u.done = DietlyapiIntegrationUtilities::Done
  u.make_error = DietlyapiIntegrationUtilities::MakeError
  u.feature_add = DietlyapiIntegrationUtilities::FeatureAdd
  u.feature_hook = DietlyapiIntegrationUtilities::FeatureHook
  u.feature_init = DietlyapiIntegrationUtilities::FeatureInit
  u.fetcher = DietlyapiIntegrationUtilities::Fetcher
  u.make_fetch_def = DietlyapiIntegrationUtilities::MakeFetchDef
  u.make_context = DietlyapiIntegrationUtilities::MakeContext
  u.make_options = DietlyapiIntegrationUtilities::MakeOptions
  u.make_request = DietlyapiIntegrationUtilities::MakeRequest
  u.make_response = DietlyapiIntegrationUtilities::MakeResponse
  u.make_result = DietlyapiIntegrationUtilities::MakeResult
  u.make_point = DietlyapiIntegrationUtilities::MakePoint
  u.make_spec = DietlyapiIntegrationUtilities::MakeSpec
  u.make_url = DietlyapiIntegrationUtilities::MakeUrl
  u.param = DietlyapiIntegrationUtilities::Param
  u.prepare_auth = DietlyapiIntegrationUtilities::PrepareAuth
  u.prepare_body = DietlyapiIntegrationUtilities::PrepareBody
  u.prepare_headers = DietlyapiIntegrationUtilities::PrepareHeaders
  u.prepare_method = DietlyapiIntegrationUtilities::PrepareMethod
  u.prepare_params = DietlyapiIntegrationUtilities::PrepareParams
  u.prepare_path = DietlyapiIntegrationUtilities::PreparePath
  u.prepare_query = DietlyapiIntegrationUtilities::PrepareQuery
  u.graphql_body = DietlyapiIntegrationUtilities::GraphqlBody
  u.graphql_errors = DietlyapiIntegrationUtilities::GraphqlErrors
  u.result_basic = DietlyapiIntegrationUtilities::ResultBasic
  u.result_body = DietlyapiIntegrationUtilities::ResultBody
  u.result_headers = DietlyapiIntegrationUtilities::ResultHeaders
  u.transform_request = DietlyapiIntegrationUtilities::TransformRequest
  u.transform_response = DietlyapiIntegrationUtilities::TransformResponse
}
