# DietlyapiIntegration SDK exists test

import pytest
from dietlyapiintegration_sdk import DietlyapiIntegrationSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = DietlyapiIntegrationSDK.test(None, None)
        assert testsdk is not None
