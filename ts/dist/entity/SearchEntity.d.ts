import { DietlyapiIntegrationEntityBase } from '../DietlyapiIntegrationEntityBase';
import type { DietlyapiIntegrationSDK } from '../DietlyapiIntegrationSDK';
import type { Control } from '../types';
import type { Search, SearchListMatch } from '../DietlyapiIntegrationTypes';
declare class SearchEntity extends DietlyapiIntegrationEntityBase<Search> {
    constructor(client: DietlyapiIntegrationSDK, entopts: any);
    make(this: SearchEntity): SearchEntity;
    list(this: any, reqmatch?: SearchListMatch, ctrl?: Control): Promise<SearchEntity[]>;
}
export { SearchEntity };
