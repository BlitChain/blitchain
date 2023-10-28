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
/**
* - TYPE_UNSPECIFIED: Default zero value enumeration
 - TYPE_EXECUTE_TX: Execute a transaction on an interchain accounts host chain
*/
export var IbcApplicationsInterchainAccountsV1Type;
(function (IbcApplicationsInterchainAccountsV1Type) {
    IbcApplicationsInterchainAccountsV1Type["TYPE_UNSPECIFIED"] = "TYPE_UNSPECIFIED";
    IbcApplicationsInterchainAccountsV1Type["TYPE_EXECUTE_TX"] = "TYPE_EXECUTE_TX";
})(IbcApplicationsInterchainAccountsV1Type || (IbcApplicationsInterchainAccountsV1Type = {}));
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
     * No description
     *
     * @tags Query
     * @name IbcApplicationsInterchainAccountsControllerV1InterchainAccount
     * @summary InterchainAccount returns the interchain account address for a given owner address on a given connection
     * @request GET:/ibc/apps/interchain_accounts/controller/v1/owners/{owner}/connections/{connection_id}
     */
    ibcApplicationsInterchainAccountsControllerV1InterchainAccount = (owner, connectionId, params = {}) => this.request({
        path: `/ibc/apps/interchain_accounts/controller/v1/owners/${owner}/connections/${connectionId}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcApplicationsInterchainAccountsControllerV1Params
     * @summary Params queries all parameters of the ICA controller submodule.
     * @request GET:/ibc/apps/interchain_accounts/controller/v1/params
     */
    ibcApplicationsInterchainAccountsControllerV1Params = (params = {}) => this.request({
        path: `/ibc/apps/interchain_accounts/controller/v1/params`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map