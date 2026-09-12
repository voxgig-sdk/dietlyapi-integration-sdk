import { BarcodeEntity } from './entity/BarcodeEntity';
import { FoodEntity } from './entity/FoodEntity';
import { MetaEntity } from './entity/MetaEntity';
import { PopularEntity } from './entity/PopularEntity';
import { SearchEntity } from './entity/SearchEntity';
export type * from './DietlyapiIntegrationTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { DietlyapiIntegrationEntityBase } from './DietlyapiIntegrationEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class DietlyapiIntegrationSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Barcode(entopts?: Record<string, any>): BarcodeEntity;
    Food(entopts?: Record<string, any>): FoodEntity;
    Meta(entopts?: Record<string, any>): MetaEntity;
    Popular(entopts?: Record<string, any>): PopularEntity;
    Search(entopts?: Record<string, any>): SearchEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): DietlyapiIntegrationSDK;
    tester(testopts?: any, sdkopts?: any): DietlyapiIntegrationSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof DietlyapiIntegrationSDK;
export { stdutil, config, BaseFeature, DietlyapiIntegrationEntityBase, DietlyapiIntegrationSDK, SDK, };
