# DietlyapiIntegration SDK feature factory

from dietlyapiintegration_sdk.feature.base_feature import DietlyapiIntegrationBaseFeature
from dietlyapiintegration_sdk.feature.test_feature import DietlyapiIntegrationTestFeature


def _make_feature(name):
    features = {
        "base": lambda: DietlyapiIntegrationBaseFeature(),
        "test": lambda: DietlyapiIntegrationTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
