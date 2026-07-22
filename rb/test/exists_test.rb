# DietlyapiIntegration SDK exists test

require "minitest/autorun"
require_relative "../DietlyapiIntegration_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = DietlyapiIntegrationSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
