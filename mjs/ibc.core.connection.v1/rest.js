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
* State defines if a connection is in one of the following states:
INIT, TRYOPEN, OPEN or UNINITIALIZED.

 - STATE_UNINITIALIZED_UNSPECIFIED: Default State
 - STATE_INIT: A connection end has just started the opening handshake.
 - STATE_TRYOPEN: A connection end has acknowledged the handshake step on the counterparty
chain.
 - STATE_OPEN: A connection end has completed the handshake.
*/
export var IbcCoreConnectionV1State;
(function (IbcCoreConnectionV1State) {
    IbcCoreConnectionV1State["STATE_UNINITIALIZED_UNSPECIFIED"] = "STATE_UNINITIALIZED_UNSPECIFIED";
    IbcCoreConnectionV1State["STATE_INIT"] = "STATE_INIT";
    IbcCoreConnectionV1State["STATE_TRYOPEN"] = "STATE_TRYOPEN";
    IbcCoreConnectionV1State["STATE_OPEN"] = "STATE_OPEN";
})(IbcCoreConnectionV1State || (IbcCoreConnectionV1State = {}));
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
   * @name IbcCoreConnectionV1ClientConnections
   * @summary ClientConnections queries the connection paths associated with a client
  state.
   * @request GET:/ibc/core/connection/v1/client_connections/{client_id}
   */
    ibcCoreConnectionV1ClientConnections = (clientId, params = {}) => this.request({
        path: `/ibc/core/connection/v1/client_connections/${clientId}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreConnectionV1Connections
     * @summary Connections queries all the IBC connections of a chain.
     * @request GET:/ibc/core/connection/v1/connections
     */
    ibcCoreConnectionV1Connections = (query, params = {}) => this.request({
        path: `/ibc/core/connection/v1/connections`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreConnectionV1Connection
     * @summary Connection queries an IBC connection end.
     * @request GET:/ibc/core/connection/v1/connections/{connection_id}
     */
    ibcCoreConnectionV1Connection = (connectionId, params = {}) => this.request({
        path: `/ibc/core/connection/v1/connections/${connectionId}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreConnectionV1ConnectionClientState
   * @summary ConnectionClientState queries the client state associated with the
  connection.
   * @request GET:/ibc/core/connection/v1/connections/{connection_id}/client_state
   */
    ibcCoreConnectionV1ConnectionClientState = (connectionId, params = {}) => this.request({
        path: `/ibc/core/connection/v1/connections/${connectionId}/client_state`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreConnectionV1ConnectionConsensusState
   * @summary ConnectionConsensusState queries the consensus state associated with the
  connection.
   * @request GET:/ibc/core/connection/v1/connections/{connection_id}/consensus_state/revision/{revision_number}/height/{revision_height}
   */
    ibcCoreConnectionV1ConnectionConsensusState = (connectionId, revisionNumber, revisionHeight, params = {}) => this.request({
        path: `/ibc/core/connection/v1/connections/${connectionId}/consensus_state/revision/${revisionNumber}/height/${revisionHeight}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreConnectionV1ConnectionParams
     * @summary ConnectionParams queries all parameters of the ibc connection submodule.
     * @request GET:/ibc/core/connection/v1/params
     */
    ibcCoreConnectionV1ConnectionParams = (params = {}) => this.request({
        path: `/ibc/core/connection/v1/params`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map