/**
 * Ferret Server API
 * API of Ferret Server
 *
 * OpenAPI spec version: 1.0.0-rc.2
 * Contact: mont.ferret@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * The properties that are included when fetching a list of Data.
 */
export interface DataOutput {
    id: string;
    rev: string;
    createdAt: string;
    updatedAt?: string;
    jobId: string;
    scriptId: string;
    scriptRev?: string;
}
