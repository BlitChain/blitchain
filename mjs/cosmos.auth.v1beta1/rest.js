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
     * @description Since: cosmos-sdk 0.47
     *
     * @tags Query
     * @name CosmosAuthV1Beta1AccountInfo
     * @summary AccountInfo queries account info which is common to all account types.
     * @request GET:/cosmos/auth/v1beta1/account_info/{address}
     */
    cosmosAuthV1Beta1AccountInfo = (address, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/account_info/${address}`,
        method: "GET",
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set. Since: cosmos-sdk 0.43
     *
     * @tags Query
     * @name CosmosAuthV1Beta1Accounts
     * @summary Accounts returns all the existing accounts.
     * @request GET:/cosmos/auth/v1beta1/accounts
     */
    cosmosAuthV1Beta1Accounts = (query, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/accounts`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosAuthV1Beta1Account
     * @summary Account returns account details based on address.
     * @request GET:/cosmos/auth/v1beta1/accounts/{address}
     */
    cosmosAuthV1Beta1Account = (address, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/accounts/${address}`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46.2
     *
     * @tags Query
     * @name CosmosAuthV1Beta1AccountAddressById
     * @summary AccountAddressByID returns account address based on account number.
     * @request GET:/cosmos/auth/v1beta1/address_by_id/{id}
     */
    cosmosAuthV1Beta1AccountAddressByID = (id, query, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/address_by_id/${id}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46
     *
     * @tags Query
     * @name CosmosAuthV1Beta1Bech32Prefix
     * @summary Bech32Prefix queries bech32Prefix
     * @request GET:/cosmos/auth/v1beta1/bech32
     */
    cosmosAuthV1Beta1Bech32Prefix = (params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/bech32`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46
     *
     * @tags Query
     * @name CosmosAuthV1Beta1AddressBytesToString
     * @summary AddressBytesToString converts Account Address bytes to string
     * @request GET:/cosmos/auth/v1beta1/bech32/{address_bytes}
     */
    cosmosAuthV1Beta1AddressBytesToString = (addressBytes, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/bech32/${addressBytes}`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46
     *
     * @tags Query
     * @name CosmosAuthV1Beta1AddressStringToBytes
     * @summary AddressStringToBytes converts Address string to bytes
     * @request GET:/cosmos/auth/v1beta1/bech32/{address_string}
     */
    cosmosAuthV1Beta1AddressStringToBytes = (addressString, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/bech32/${addressString}`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46
     *
     * @tags Query
     * @name CosmosAuthV1Beta1ModuleAccounts
     * @summary ModuleAccounts returns all the existing module accounts.
     * @request GET:/cosmos/auth/v1beta1/module_accounts
     */
    cosmosAuthV1Beta1ModuleAccounts = (params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/module_accounts`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosAuthV1Beta1ModuleAccountByName
     * @summary ModuleAccountByName returns the module account info by module name
     * @request GET:/cosmos/auth/v1beta1/module_accounts/{name}
     */
    cosmosAuthV1Beta1ModuleAccountByName = (name, params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/module_accounts/${name}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosAuthV1Beta1Params
     * @summary Params queries all parameters.
     * @request GET:/cosmos/auth/v1beta1/params
     */
    cosmosAuthV1Beta1Params = (params = {}) => this.request({
        path: `/cosmos/auth/v1beta1/params`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map