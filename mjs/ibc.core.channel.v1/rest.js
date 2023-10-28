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
* - ORDER_NONE_UNSPECIFIED: zero-value for channel ordering
 - ORDER_UNORDERED: packets can be delivered in any order, which may differ from the order in
which they were sent.
 - ORDER_ORDERED: packets are delivered exactly in the order which they were sent
*/
export var IbcCoreChannelV1Order;
(function (IbcCoreChannelV1Order) {
    IbcCoreChannelV1Order["ORDER_NONE_UNSPECIFIED"] = "ORDER_NONE_UNSPECIFIED";
    IbcCoreChannelV1Order["ORDER_UNORDERED"] = "ORDER_UNORDERED";
    IbcCoreChannelV1Order["ORDER_ORDERED"] = "ORDER_ORDERED";
})(IbcCoreChannelV1Order || (IbcCoreChannelV1Order = {}));
/**
* State defines if a channel is in one of the following states:
CLOSED, INIT, TRYOPEN, OPEN or UNINITIALIZED.

 - STATE_UNINITIALIZED_UNSPECIFIED: Default State
 - STATE_INIT: A channel has just started the opening handshake.
 - STATE_TRYOPEN: A channel has acknowledged the handshake step on the counterparty chain.
 - STATE_OPEN: A channel has completed the handshake. Open channels are
ready to send and receive packets.
 - STATE_CLOSED: A channel has been closed and can no longer be used to send or receive
packets.
*/
export var IbcCoreChannelV1State;
(function (IbcCoreChannelV1State) {
    IbcCoreChannelV1State["STATE_UNINITIALIZED_UNSPECIFIED"] = "STATE_UNINITIALIZED_UNSPECIFIED";
    IbcCoreChannelV1State["STATE_INIT"] = "STATE_INIT";
    IbcCoreChannelV1State["STATE_TRYOPEN"] = "STATE_TRYOPEN";
    IbcCoreChannelV1State["STATE_OPEN"] = "STATE_OPEN";
    IbcCoreChannelV1State["STATE_CLOSED"] = "STATE_CLOSED";
})(IbcCoreChannelV1State || (IbcCoreChannelV1State = {}));
/**
* - RESPONSE_RESULT_TYPE_UNSPECIFIED: Default zero value enumeration
 - RESPONSE_RESULT_TYPE_NOOP: The message did not call the IBC application callbacks (because, for example, the packet had already been relayed)
 - RESPONSE_RESULT_TYPE_SUCCESS: The message was executed successfully
*/
export var IbcCoreChannelV1ResponseResultType;
(function (IbcCoreChannelV1ResponseResultType) {
    IbcCoreChannelV1ResponseResultType["RESPONSE_RESULT_TYPE_UNSPECIFIED"] = "RESPONSE_RESULT_TYPE_UNSPECIFIED";
    IbcCoreChannelV1ResponseResultType["RESPONSE_RESULT_TYPE_NOOP"] = "RESPONSE_RESULT_TYPE_NOOP";
    IbcCoreChannelV1ResponseResultType["RESPONSE_RESULT_TYPE_SUCCESS"] = "RESPONSE_RESULT_TYPE_SUCCESS";
})(IbcCoreChannelV1ResponseResultType || (IbcCoreChannelV1ResponseResultType = {}));
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
     * @name IbcCoreChannelV1Channels
     * @summary Channels queries all the IBC channels of a chain.
     * @request GET:/ibc/core/channel/v1/channels
     */
    ibcCoreChannelV1Channels = (query, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreChannelV1Channel
     * @summary Channel queries an IBC Channel.
     * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}
     */
    ibcCoreChannelV1Channel = (channelId, portId, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1ChannelClientState
   * @summary ChannelClientState queries for the client state for the channel associated
  with the provided channel identifiers.
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/client_state
   */
    ibcCoreChannelV1ChannelClientState = (channelId, portId, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/client_state`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1ChannelConsensusState
   * @summary ChannelConsensusState queries for the consensus state for the channel
  associated with the provided channel identifiers.
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/consensus_state/revision/{revision_number}/height/{revision_height}
   */
    ibcCoreChannelV1ChannelConsensusState = (channelId, portId, revisionNumber, revisionHeight, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/consensus_state/revision/${revisionNumber}/height/${revisionHeight}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreChannelV1NextSequenceReceive
     * @summary NextSequenceReceive returns the next receive sequence for a given channel.
     * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/next_sequence
     */
    ibcCoreChannelV1NextSequenceReceive = (channelId, portId, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/next_sequence`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1PacketAcknowledgements
   * @summary PacketAcknowledgements returns all the packet acknowledgements associated
  with a channel.
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_acknowledgements
   */
    ibcCoreChannelV1PacketAcknowledgements = (channelId, portId, query, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_acknowledgements`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreChannelV1PacketAcknowledgement
     * @summary PacketAcknowledgement queries a stored packet acknowledgement hash.
     * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_acks/{sequence}
     */
    ibcCoreChannelV1PacketAcknowledgement = (channelId, portId, sequence, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_acks/${sequence}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1PacketCommitments
   * @summary PacketCommitments returns all the packet commitments hashes associated
  with a channel.
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments
   */
    ibcCoreChannelV1PacketCommitments = (channelId, portId, query, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1UnreceivedAcks
   * @summary UnreceivedAcks returns all the unreceived IBC acknowledgements associated
  with a channel and sequences.
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_ack_sequences}/unreceived_acks
   */
    ibcCoreChannelV1UnreceivedAcks = (channelId, portId, packetAckSequences, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments/${packetAckSequences}/unreceived_acks`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1UnreceivedPackets
   * @summary UnreceivedPackets returns all the unreceived IBC packets associated with a
  channel and sequences.
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{packet_commitment_sequences}/unreceived_packets
   */
    ibcCoreChannelV1UnreceivedPackets = (channelId, portId, packetCommitmentSequences, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments/${packetCommitmentSequences}/unreceived_packets`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name IbcCoreChannelV1PacketCommitment
     * @summary PacketCommitment queries a stored packet commitment hash.
     * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_commitments/{sequence}
     */
    ibcCoreChannelV1PacketCommitment = (channelId, portId, sequence, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_commitments/${sequence}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1PacketReceipt
   * @summary PacketReceipt queries if a given packet sequence has been received on the
  queried chain
   * @request GET:/ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/packet_receipts/{sequence}
   */
    ibcCoreChannelV1PacketReceipt = (channelId, portId, sequence, params = {}) => this.request({
        path: `/ibc/core/channel/v1/channels/${channelId}/ports/${portId}/packet_receipts/${sequence}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name IbcCoreChannelV1ConnectionChannels
   * @summary ConnectionChannels queries all the channels associated with a connection
  end.
   * @request GET:/ibc/core/channel/v1/connections/{connection}/channels
   */
    ibcCoreChannelV1ConnectionChannels = (connection, query, params = {}) => this.request({
        path: `/ibc/core/channel/v1/connections/${connection}/channels`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map