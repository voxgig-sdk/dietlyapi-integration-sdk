import { DietlyapiIntegrationEntityBase } from '../DietlyapiIntegrationEntityBase';
import type { DietlyapiIntegrationSDK } from '../DietlyapiIntegrationSDK';
import type { Control } from '../types';
import type { Popular, PopularListMatch } from '../DietlyapiIntegrationTypes';
declare class PopularEntity extends DietlyapiIntegrationEntityBase<Popular> {
    constructor(client: DietlyapiIntegrationSDK, entopts: any);
    make(this: PopularEntity): PopularEntity;
    list(this: any, reqmatch?: PopularListMatch, ctrl?: Control): Promise<PopularEntity[]>;
}
export { PopularEntity };
