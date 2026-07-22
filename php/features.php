<?php
declare(strict_types=1);

// DietlyapiIntegration SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class DietlyapiIntegrationFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new DietlyapiIntegrationBaseFeature();
            case "test":
                return new DietlyapiIntegrationTestFeature();
            default:
                return new DietlyapiIntegrationBaseFeature();
        }
    }
}
