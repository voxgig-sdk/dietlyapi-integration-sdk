<?php
declare(strict_types=1);

// DietlyapiIntegration SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class DietlyapiIntegrationMakeContext
{
    public static function call(array $ctxmap, ?DietlyapiIntegrationContext $basectx): DietlyapiIntegrationContext
    {
        return new DietlyapiIntegrationContext($ctxmap, $basectx);
    }
}
