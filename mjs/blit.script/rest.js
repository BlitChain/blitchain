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
     * @name BlitScriptEval
     * @summary Runs the function and returns the result.
     * @request POST:/blit/script/eval/{script_address}
     */
    blitScriptEval = (scriptAddress, query, params = {}) => this.request({
        path: `/blit/script/eval/${scriptAddress}`,
        method: "POST",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name BlitScriptParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/blit/script/params
     */
    blitScriptParams = (params = {}) => this.request({
        path: `/blit/script/params`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name BlitScriptScriptAll
     * @request GET:/blit/script/script
     */
    blitScriptScriptAll = (query, params = {}) => this.request({
        path: `/blit/script/script`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name BlitScriptScript
     * @summary Queries a list of Script items.
     * @request GET:/blit/script/script/{address}
     */
    blitScriptScript = (address, params = {}) => this.request({
        path: `/blit/script/script/${address}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name BlitScriptWeb
     * @summary Queries the WSGI web application function of a script.
     * @request GET:/blit/script/web/{address}
     */
    blitScriptWeb = (address, query, params = {}) => this.request({
        path: `/blit/script/web/${address}`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map