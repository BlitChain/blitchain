/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */
import axios from "axios";
export var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType || (ContentType = {}));
export class HttpClient {
    instance;
    securityData = null;
    securityWorker;
    secure;
    format;
    constructor({ securityWorker, secure, format, ...axiosConfig } = {}) {
        this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
        this.secure = secure;
        this.format = format;
        this.securityWorker = securityWorker;
    }
    setSecurityData = (data) => {
        this.securityData = data;
    };
    mergeRequestParams(params1, params2) {
        return {
            ...this.instance.defaults,
            ...params1,
            ...(params2 || {}),
            headers: {
                ...(this.instance.defaults.headers || {}),
                ...(params1.headers || {}),
                ...((params2 && params2.headers) || {}),
            },
        };
    }
    createFormData(input) {
        return Object.keys(input || {}).reduce((formData, key) => {
            const property = input[key];
            formData.append(key, property instanceof Blob
                ? property
                : typeof property === "object" && property !== null
                    ? JSON.stringify(property)
                    : `${property}`);
            return formData;
        }, new FormData());
    }
    request = async ({ secure, path, type, query, format, body, ...params }) => {
        const secureParams = ((typeof secure === "boolean" ? secure : this.secure) &&
            this.securityWorker &&
            (await this.securityWorker(this.securityData))) ||
            {};
        const requestParams = this.mergeRequestParams(params, secureParams);
        const responseFormat = (format && this.format) || void 0;
        if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
            requestParams.headers.common = { Accept: "*/*" };
            requestParams.headers.post = {};
            requestParams.headers.put = {};
            body = this.createFormData(body);
        }
        return this.instance.request({
            ...requestParams,
            headers: {
                ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
                ...(requestParams.headers || {}),
            },
            params: query,
            responseType: responseFormat,
            data: body,
            url: path,
        });
    };
}
/**
 * @title HTTP API Console
 */
export class Api extends HttpClient {
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosBankV1Beta1AllBalances
     * @summary AllBalances queries the balance of all coins for a single account.
     * @request GET:/cosmos/bank/v1beta1/balances/{address}
     */
    cosmosBankV1Beta1AllBalances = (address, query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/balances/${address}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosBankV1Beta1Balance
     * @summary Balance queries the balance of a single coin for a single account.
     * @request GET:/cosmos/bank/v1beta1/balances/{address}/by_denom
     */
    cosmosBankV1Beta1Balance = (address, query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/balances/${address}/by_denom`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set. Since: cosmos-sdk 0.46
   *
   * @tags Query
   * @name CosmosBankV1Beta1DenomOwners
   * @summary DenomOwners queries for all account addresses that own a particular token
  denomination.
   * @request GET:/cosmos/bank/v1beta1/denom_owners/{denom}
   */
    cosmosBankV1Beta1DenomOwners = (denom, query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/denom_owners/${denom}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosBankV1Beta1DenomsMetadata
   * @summary DenomsMetadata queries the client metadata for all registered coin
  denominations.
   * @request GET:/cosmos/bank/v1beta1/denoms_metadata
   */
    cosmosBankV1Beta1DenomsMetadata = (query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/denoms_metadata`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosBankV1Beta1DenomMetadata
     * @summary DenomsMetadata queries the client metadata of a given coin denomination.
     * @request GET:/cosmos/bank/v1beta1/denoms_metadata/{denom}
     */
    cosmosBankV1Beta1DenomMetadata = (denom, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/denoms_metadata/${denom}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosBankV1Beta1Params
     * @summary Params queries the parameters of x/bank module.
     * @request GET:/cosmos/bank/v1beta1/params
     */
    cosmosBankV1Beta1Params = (params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/params`,
        method: "GET",
        ...params,
    });
    /**
     * @description This query only returns denominations that have specific SendEnabled settings. Any denomination that does not have a specific setting will use the default params.default_send_enabled, and will not be returned by this query. Since: cosmos-sdk 0.47
     *
     * @tags Query
     * @name CosmosBankV1Beta1SendEnabled
     * @summary SendEnabled queries for SendEnabled entries.
     * @request GET:/cosmos/bank/v1beta1/send_enabled
     */
    cosmosBankV1Beta1SendEnabled = (query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/send_enabled`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set. Since: cosmos-sdk 0.46
   *
   * @tags Query
   * @name CosmosBankV1Beta1SpendableBalances
   * @summary SpendableBalances queries the spendable balance of all coins for a single
  account.
   * @request GET:/cosmos/bank/v1beta1/spendable_balances/{address}
   */
    cosmosBankV1Beta1SpendableBalances = (address, query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/spendable_balances/${address}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set. Since: cosmos-sdk 0.47
   *
   * @tags Query
   * @name CosmosBankV1Beta1SpendableBalanceByDenom
   * @summary SpendableBalanceByDenom queries the spendable balance of a single denom for
  a single account.
   * @request GET:/cosmos/bank/v1beta1/spendable_balances/{address}/by_denom
   */
    cosmosBankV1Beta1SpendableBalanceByDenom = (address, query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/spendable_balances/${address}/by_denom`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosBankV1Beta1TotalSupply
     * @summary TotalSupply queries the total supply of all coins.
     * @request GET:/cosmos/bank/v1beta1/supply
     */
    cosmosBankV1Beta1TotalSupply = (query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/supply`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosBankV1Beta1SupplyOf
     * @summary SupplyOf queries the supply of a single coin.
     * @request GET:/cosmos/bank/v1beta1/supply/by_denom
     */
    cosmosBankV1Beta1SupplyOf = (query, params = {}) => this.request({
        path: `/cosmos/bank/v1beta1/supply/by_denom`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map