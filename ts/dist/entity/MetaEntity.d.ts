import { DietlyapiIntegrationEntityBase } from '../DietlyapiIntegrationEntityBase';
import type { DietlyapiIntegrationSDK } from '../DietlyapiIntegrationSDK';
import type { Control } from '../types';
import type { Meta, MetaLoadMatch } from '../DietlyapiIntegrationTypes';
declare class MetaEntity extends DietlyapiIntegrationEntityBase<Meta> {
    constructor(client: DietlyapiIntegrationSDK, entopts: any);
    make(this: MetaEntity): MetaEntity;
    load(this: any, reqmatch?: MetaLoadMatch, ctrl?: Control): Promise<MetaEntity>;
}
export { MetaEntity };
