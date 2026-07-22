# DietlyapiIntegration SDK utility: make_context
require_relative '../core/context'
module DietlyapiIntegrationUtilities
  MakeContext = ->(ctxmap, basectx) {
    DietlyapiIntegrationContext.new(ctxmap, basectx)
  }
end
