<?php
declare(strict_types=1);

// DietlyapiIntegration SDK utility: prepare_body

class DietlyapiIntegrationPrepareBody
{
    public static function call(DietlyapiIntegrationContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
