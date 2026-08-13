<?php
declare(strict_types=1);

// Barcode entity test

require_once __DIR__ . '/../dietlyapiintegration_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class BarcodeEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = DietlyapiIntegrationSDK::test(null, null);
        $ent = $testsdk->Barcode(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = barcode_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "barcode." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $barcode_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.barcode")));
        $barcode_ref01_data = null;
        if (count($barcode_ref01_data_raw) > 0) {
            $barcode_ref01_data = Helpers::to_map($barcode_ref01_data_raw[0][1]);
        }

        // LOAD
        $barcode_ref01_ent = $client->Barcode(null);
        $barcode_ref01_match_dt0 = [
            "id" => $barcode_ref01_data["id"],
        ];
        $barcode_ref01_data_dt0_loaded = $barcode_ref01_ent->load($barcode_ref01_match_dt0, null);
        $barcode_ref01_data_dt0_load_result = Helpers::to_map(is_object($barcode_ref01_data_dt0_loaded) && method_exists($barcode_ref01_data_dt0_loaded, 'data_get') ? $barcode_ref01_data_dt0_loaded->data_get() : $barcode_ref01_data_dt0_loaded);
        $this->assertNotNull($barcode_ref01_data_dt0_load_result);
        $this->assertEquals($barcode_ref01_data_dt0_load_result["id"], $barcode_ref01_data["id"]);

    }
}

function barcode_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/barcode/BarcodeTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = DietlyapiIntegrationSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["barcode01", "barcode02", "barcode03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID" => $idmap,
        "DIETLYAPI_INTEGRATION_TEST_LIVE" => "FALSE",
        "DIETLYAPI_INTEGRATION_TEST_EXPLAIN" => "FALSE",
        "DIETLYAPI_INTEGRATION_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["DIETLYAPI_INTEGRATION_TEST_BARCODE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["DIETLYAPI_INTEGRATION_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["DIETLYAPI_INTEGRATION_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new DietlyapiIntegrationSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["DIETLYAPI_INTEGRATION_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["DIETLYAPI_INTEGRATION_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
