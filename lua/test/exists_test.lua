-- DietlyapiIntegration SDK exists test

local sdk = require("dietlyapi-integration_sdk")

describe("DietlyapiIntegrationSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
