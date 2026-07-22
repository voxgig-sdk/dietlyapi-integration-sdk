<?php
declare(strict_types=1);

// DietlyapiIntegration SDK exists test

require_once __DIR__ . '/../dietlyapiintegration_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = DietlyapiIntegrationSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
