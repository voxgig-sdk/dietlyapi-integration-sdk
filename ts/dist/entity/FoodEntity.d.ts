import { DietlyapiIntegrationEntityBase } from '../DietlyapiIntegrationEntityBase';
import type { DietlyapiIntegrationSDK } from '../DietlyapiIntegrationSDK';
import type { Control } from '../types';
import type { Food, FoodLoadMatch, FoodListMatch } from '../DietlyapiIntegrationTypes';
declare class FoodEntity extends DietlyapiIntegrationEntityBase<Food> {
    constructor(client: DietlyapiIntegrationSDK, entopts: any);
    make(this: FoodEntity): FoodEntity;
    load(this: any, reqmatch?: FoodLoadMatch, ctrl?: Control): Promise<FoodEntity>;
    list(this: any, reqmatch?: FoodListMatch, ctrl?: Control): Promise<FoodEntity[]>;
}
export { FoodEntity };
