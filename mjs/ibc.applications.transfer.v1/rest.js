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
     * @name IbcApplicationsTransferV1EscrowAddress
     * @summary EscrowAddress returns the escrow address for a particular port and channel id.
     * @request GET:/ibc/apps/transfer/v1/channels/{channel_id}/ports/{port_id}/escrow_address
     */
    ibcApplicationsTransferV1EscrowAddress = (channelId, portId, params = {}) => this.request({
        path: `/ibc/apps/transfer/v1/channels/${channelId}/ports/${portId}/escrow_address`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcApplicationsTransferV1DenomHash
     * @summary DenomHash queries a denomination hash information.
     * @request GET:/ibc/apps/transfer/v1/denom_hashes/{trace}
     */
    ibcApplicationsTransferV1DenomHash = (trace, params = {}) => this.request({
        path: `/ibc/apps/transfer/v1/denom_hashes/${trace}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcApplicationsTransferV1DenomTraces
     * @summary DenomTraces queries all denomination traces.
     * @request GET:/ibc/apps/transfer/v1/denom_traces
     */
    ibcApplicationsTransferV1DenomTraces = (query, params = {}) => this.request({
        path: `/ibc/apps/transfer/v1/denom_traces`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcApplicationsTransferV1DenomTrace
     * @summary DenomTrace queries a denomination trace information.
     * @request GET:/ibc/apps/transfer/v1/denom_traces/{hash}
     */
    ibcApplicationsTransferV1DenomTrace = (hash, params = {}) => this.request({
        path: `/ibc/apps/transfer/v1/denom_traces/${hash}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcApplicationsTransferV1TotalEscrowForDenom
     * @summary TotalEscrowForDenom returns the total amount of tokens in escrow based on the denom.
     * @request GET:/ibc/apps/transfer/v1/denoms/{denom}/total_escrow
     */
    ibcApplicationsTransferV1TotalEscrowForDenom = (denom, params = {}) => this.request({
        path: `/ibc/apps/transfer/v1/denoms/${denom}/total_escrow`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcApplicationsTransferV1Params
     * @summary Params queries all parameters of the ibc-transfer module.
     * @request GET:/ibc/apps/transfer/v1/params
     */
    ibcApplicationsTransferV1Params = (params = {}) => this.request({
        path: `/ibc/apps/transfer/v1/params`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map