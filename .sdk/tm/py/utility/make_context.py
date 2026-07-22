# DietlyapiIntegration SDK utility: make_context

from core.context import DietlyapiIntegrationContext


def make_context_util(ctxmap, basectx):
    return DietlyapiIntegrationContext(ctxmap, basectx)
