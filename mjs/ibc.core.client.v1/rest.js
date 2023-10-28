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
     * @name IbcCoreClientV1ClientStates
     * @summary ClientStates queries all the IBC light clients of a chain.
     * @request GET:/ibc/core/client/v1/client_states
     */
    ibcCoreClientV1ClientStates = (query, params = {}) => this.request({
        path: `/ibc/core/client/v1/client_states`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreClientV1ClientState
     * @summary ClientState queries an IBC light client.
     * @request GET:/ibc/core/client/v1/client_states/{client_id}
     */
    ibcCoreClientV1ClientState = (clientId, params = {}) => this.request({
        path: `/ibc/core/client/v1/client_states/${clientId}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreClientV1ClientStatus
     * @summary Status queries the status of an IBC client.
     * @request GET:/ibc/core/client/v1/client_status/{client_id}
     */
    ibcCoreClientV1ClientStatus = (clientId, params = {}) => this.request({
        path: `/ibc/core/client/v1/client_status/${clientId}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreClientV1ConsensusStates
   * @summary ConsensusStates queries all the consensus state associated with a given
  client.
   * @request GET:/ibc/core/client/v1/consensus_states/{client_id}
   */
    ibcCoreClientV1ConsensusStates = (clientId, query, params = {}) => this.request({
        path: `/ibc/core/client/v1/consensus_states/${clientId}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreClientV1ConsensusStateHeights
     * @summary ConsensusStateHeights queries the height of every consensus states associated with a given client.
     * @request GET:/ibc/core/client/v1/consensus_states/{client_id}/heights
     */
    ibcCoreClientV1ConsensusStateHeights = (clientId, query, params = {}) => this.request({
        path: `/ibc/core/client/v1/consensus_states/${clientId}/heights`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreClientV1ConsensusState
   * @summary ConsensusState queries a consensus state associated with a client state at
  a given height.
   * @request GET:/ibc/core/client/v1/consensus_states/{client_id}/revision/{revision_number}/height/{revision_height}
   */
    ibcCoreClientV1ConsensusState = (clientId, revisionNumber, revisionHeight, query, params = {}) => this.request({
        path: `/ibc/core/client/v1/consensus_states/${clientId}/revision/${revisionNumber}/height/${revisionHeight}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreClientV1ClientParams
     * @summary ClientParams queries all parameters of the ibc client submodule.
     * @request GET:/ibc/core/client/v1/params
     */
    ibcCoreClientV1ClientParams = (params = {}) => this.request({
        path: `/ibc/core/client/v1/params`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreClientV1UpgradedClientState
     * @summary UpgradedClientState queries an Upgraded IBC light client.
     * @request GET:/ibc/core/client/v1/upgraded_client_states
     */
    ibcCoreClientV1UpgradedClientState = (params = {}) => this.request({
        path: `/ibc/core/client/v1/upgraded_client_states`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreClientV1UpgradedConsensusState
     * @summary UpgradedConsensusState queries an Upgraded IBC consensus state.
     * @request GET:/ibc/core/client/v1/upgraded_consensus_states
     */
    ibcCoreClientV1UpgradedConsensusState = (params = {}) => this.request({
        path: `/ibc/core/client/v1/upgraded_consensus_states`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map