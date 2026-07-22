<?php
declare(strict_types=1);

// DietlyapiIntegration SDK utility: result_body

class DietlyapiIntegrationResultBody
{
    public static function call(DietlyapiIntegrationContext $ctx): ?DietlyapiIntegrationResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
