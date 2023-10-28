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
     * No description
     *
     * @tags Query
     * @name CosmosFeegrantV1Beta1Allowance
     * @summary Allowance returns fee granted to the grantee by the granter.
     * @request GET:/cosmos/feegrant/v1beta1/allowance/{granter}/{grantee}
     */
    cosmosFeegrantV1Beta1Allowance = (granter, grantee, params = {}) => this.request({
        path: `/cosmos/feegrant/v1beta1/allowance/${granter}/${grantee}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosFeegrantV1Beta1Allowances
     * @summary Allowances returns all the grants for address.
     * @request GET:/cosmos/feegrant/v1beta1/allowances/{grantee}
     */
    cosmosFeegrantV1Beta1Allowances = (grantee, query, params = {}) => this.request({
        path: `/cosmos/feegrant/v1beta1/allowances/${grantee}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46
     *
     * @tags Query
     * @name CosmosFeegrantV1Beta1AllowancesByGranter
     * @summary AllowancesByGranter returns all the grants given by an address
     * @request GET:/cosmos/feegrant/v1beta1/issued/{granter}
     */
    cosmosFeegrantV1Beta1AllowancesByGranter = (granter, query, params = {}) => this.request({
        path: `/cosmos/feegrant/v1beta1/issued/${granter}`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map