-- DietlyapiIntegration SDK error

local DietlyapiIntegrationError = {}
DietlyapiIntegrationError.__index = DietlyapiIntegrationError


function DietlyapiIntegrationError.new(code, msg, ctx)
  local self = setmetatable({}, DietlyapiIntegrationError)
  self.is_sdk_error = true
  self.sdk = "DietlyapiIntegration"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function DietlyapiIntegrationError:error()
  return self.msg
end


function DietlyapiIntegrationError:__tostring()
  return self.msg
end


return DietlyapiIntegrationError
