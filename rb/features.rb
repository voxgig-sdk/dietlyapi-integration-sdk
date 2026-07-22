# DietlyapiIntegration SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module DietlyapiIntegrationFeatures
  def self.make_feature(name)
    case name
    when "base"
      DietlyapiIntegrationBaseFeature.new
    when "test"
      DietlyapiIntegrationTestFeature.new
    else
      DietlyapiIntegrationBaseFeature.new
    end
  end
end
