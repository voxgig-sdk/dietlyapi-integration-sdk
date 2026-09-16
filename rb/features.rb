# DietlyapiIntegration SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module DietlyapiIntegrationFeatures
  def self.make_feature(name)
    case name
    when "base"
      DietlyapiIntegrationBaseFeature.new
    when "ratelimit"
      DietlyapiIntegrationRatelimitFeature.new
    when "retry"
      DietlyapiIntegrationRetryFeature.new
    when "test"
      DietlyapiIntegrationTestFeature.new
    when "timeout"
      DietlyapiIntegrationTimeoutFeature.new
    else
      DietlyapiIntegrationBaseFeature.new
    end
  end
end
