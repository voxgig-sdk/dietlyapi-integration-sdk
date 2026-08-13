# DietlyapiIntegration SDK utility: make_context

from projectname_sdk.core.context import DietlyapiIntegrationContext


def make_context_util(ctxmap, basectx):
    return DietlyapiIntegrationContext(ctxmap, basectx)
