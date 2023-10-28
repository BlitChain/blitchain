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
     * @name CosmosUpgradeV1Beta1AppliedPlan
     * @summary AppliedPlan queries a previously applied upgrade plan by its name.
     * @request GET:/cosmos/upgrade/v1beta1/applied_plan/{name}
     */
    cosmosUpgradeV1Beta1AppliedPlan = (name, params = {}) => this.request({
        path: `/cosmos/upgrade/v1beta1/applied_plan/${name}`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.46
     *
     * @tags Query
     * @name CosmosUpgradeV1Beta1Authority
     * @summary Returns the account with authority to conduct upgrades
     * @request GET:/cosmos/upgrade/v1beta1/authority
     */
    cosmosUpgradeV1Beta1Authority = (params = {}) => this.request({
        path: `/cosmos/upgrade/v1beta1/authority`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosUpgradeV1Beta1CurrentPlan
     * @summary CurrentPlan queries the current upgrade plan.
     * @request GET:/cosmos/upgrade/v1beta1/current_plan
     */
    cosmosUpgradeV1Beta1CurrentPlan = (params = {}) => this.request({
        path: `/cosmos/upgrade/v1beta1/current_plan`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.43
     *
     * @tags Query
     * @name CosmosUpgradeV1Beta1ModuleVersions
     * @summary ModuleVersions queries the list of module versions from state.
     * @request GET:/cosmos/upgrade/v1beta1/module_versions
     */
    cosmosUpgradeV1Beta1ModuleVersions = (query, params = {}) => this.request({
        path: `/cosmos/upgrade/v1beta1/module_versions`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosUpgradeV1Beta1UpgradedConsensusState
   * @summary UpgradedConsensusState queries the consensus state that will serve
  as a trusted kernel for the next version of this chain. It will only be
  stored at the last height of this chain.
  UpgradedConsensusState RPC not supported with legacy querier
  This rpc is deprecated now that IBC has its own replacement
  (https://github.com/cosmos/ibc-go/blob/2c880a22e9f9cc75f62b527ca94aa75ce1106001/proto/ibc/core/client/v1/query.proto#L54)
   * @request GET:/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}
   */
    cosmosUpgradeV1Beta1UpgradedConsensusState = (lastHeight, params = {}) => this.request({
        path: `/cosmos/upgrade/v1beta1/upgraded_consensus_state/${lastHeight}`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map