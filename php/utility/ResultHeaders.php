<?php
declare(strict_types=1);

// DietlyapiIntegration SDK utility: result_headers

class DietlyapiIntegrationResultHeaders
{
    public static function call(DietlyapiIntegrationContext $ctx): ?DietlyapiIntegrationResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
