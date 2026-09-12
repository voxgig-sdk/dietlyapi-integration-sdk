import { DietlyapiIntegrationEntityBase } from '../DietlyapiIntegrationEntityBase';
import type { DietlyapiIntegrationSDK } from '../DietlyapiIntegrationSDK';
import type { Control } from '../types';
import type { Barcode, BarcodeLoadMatch } from '../DietlyapiIntegrationTypes';
declare class BarcodeEntity extends DietlyapiIntegrationEntityBase<Barcode> {
    constructor(client: DietlyapiIntegrationSDK, entopts: any);
    make(this: BarcodeEntity): BarcodeEntity;
    load(this: any, reqmatch?: BarcodeLoadMatch, ctrl?: Control): Promise<BarcodeEntity>;
}
export { BarcodeEntity };
