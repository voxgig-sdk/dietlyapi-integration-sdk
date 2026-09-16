# DietlyapiIntegration SDK feature factory

from dietlyapiintegration_sdk.feature.base_feature import DietlyapiIntegrationBaseFeature
from dietlyapiintegration_sdk.feature.ratelimit_feature import DietlyapiIntegrationRatelimitFeature
from dietlyapiintegration_sdk.feature.retry_feature import DietlyapiIntegrationRetryFeature
from dietlyapiintegration_sdk.feature.test_feature import DietlyapiIntegrationTestFeature
from dietlyapiintegration_sdk.feature.timeout_feature import DietlyapiIntegrationTimeoutFeature


_FEATURES = {
    "base": lambda: DietlyapiIntegrationBaseFeature(),
    "ratelimit": lambda: DietlyapiIntegrationRatelimitFeature(),
    "retry": lambda: DietlyapiIntegrationRetryFeature(),
    "test": lambda: DietlyapiIntegrationTestFeature(),
    "timeout": lambda: DietlyapiIntegrationTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
