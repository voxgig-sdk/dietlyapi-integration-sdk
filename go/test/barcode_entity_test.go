package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/dietlyapi-integration-sdk/go"
	"github.com/voxgig-sdk/dietlyapi-integration-sdk/go/core"

	vs "github.com/voxgig-sdk/dietlyapi-integration-sdk/go/utility/struct"
)

func TestBarcodeEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Barcode(nil)
		if ent == nil {
			t.Fatal("expected non-nil BarcodeEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := barcodeBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "barcode." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set DIETLYAPIINTEGRATION_TEST_BARCODE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		barcodeRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.barcode", setup.data)))
		var barcodeRef01Data map[string]any
		if len(barcodeRef01DataRaw) > 0 {
			barcodeRef01Data = core.ToMapAny(barcodeRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = barcodeRef01Data

		// LOAD
		barcodeRef01Ent := client.Barcode(nil)
		barcodeRef01MatchDt0 := map[string]any{
			"id": barcodeRef01Data["id"],
		}
		barcodeRef01DataDt0Loaded, err := barcodeRef01Ent.Load(barcodeRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		barcodeRef01DataDt0LoadResult := core.ToMapAny(barcodeRef01DataDt0Loaded)
		if barcodeRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if barcodeRef01DataDt0LoadResult["id"] != barcodeRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func barcodeBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "barcode", "BarcodeTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read barcode test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse barcode test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"barcode01", "barcode02", "barcode03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("DIETLYAPIINTEGRATION_TEST_BARCODE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"DIETLYAPIINTEGRATION_TEST_BARCODE_ENTID": idmap,
		"DIETLYAPIINTEGRATION_TEST_LIVE":      "FALSE",
		"DIETLYAPIINTEGRATION_TEST_EXPLAIN":   "FALSE",
		"DIETLYAPIINTEGRATION_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["DIETLYAPIINTEGRATION_TEST_BARCODE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["DIETLYAPIINTEGRATION_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["DIETLYAPIINTEGRATION_APIKEY"],
			},
			extra,
		})
		client = sdk.NewDietlyapiIntegrationSDK(core.ToMapAny(mergedOpts))
	}

	live := env["DIETLYAPIINTEGRATION_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["DIETLYAPIINTEGRATION_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
