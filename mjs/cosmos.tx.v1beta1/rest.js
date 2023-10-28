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
* SignMode represents a signing mode with its own security guarantees.

This enum should be considered a registry of all known sign modes
in the Cosmos ecosystem. Apps are not expected to support all known
sign modes. Apps that would like to support custom  sign modes are
encouraged to open a small PR against this file to add a new case
to this SignMode enum describing their sign mode so that different
apps have a consistent version of this enum.

 - SIGN_MODE_UNSPECIFIED: SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be
rejected.
 - SIGN_MODE_DIRECT: SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is
verified with raw bytes from Tx.
 - SIGN_MODE_TEXTUAL: SIGN_MODE_TEXTUAL is a future signing mode that will verify some
human-readable textual representation on top of the binary representation
from SIGN_MODE_DIRECT. It is currently not supported.
 - SIGN_MODE_DIRECT_AUX: SIGN_MODE_DIRECT_AUX specifies a signing mode which uses
SignDocDirectAux. As opposed to SIGN_MODE_DIRECT, this sign mode does not
require signers signing over other signers' `signer_info`. It also allows
for adding Tips in transactions.

Since: cosmos-sdk 0.46
 - SIGN_MODE_LEGACY_AMINO_JSON: SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses
Amino JSON and will be removed in the future.
 - SIGN_MODE_EIP_191: SIGN_MODE_EIP_191 specifies the sign mode for EIP 191 signing on the Cosmos
SDK. Ref: https://eips.ethereum.org/EIPS/eip-191

Currently, SIGN_MODE_EIP_191 is registered as a SignMode enum variant,
but is not implemented on the SDK by default. To enable EIP-191, you need
to pass a custom `TxConfig` that has an implementation of
`SignModeHandler` for EIP-191. The SDK may decide to fully support
EIP-191 in the future.

Since: cosmos-sdk 0.45.2
*/
export var CosmosTxSigningV1Beta1SignMode;
(function (CosmosTxSigningV1Beta1SignMode) {
    CosmosTxSigningV1Beta1SignMode["SIGN_MODE_UNSPECIFIED"] = "SIGN_MODE_UNSPECIFIED";
    CosmosTxSigningV1Beta1SignMode["SIGN_MODE_DIRECT"] = "SIGN_MODE_DIRECT";
    CosmosTxSigningV1Beta1SignMode["SIGN_MODE_TEXTUAL"] = "SIGN_MODE_TEXTUAL";
    CosmosTxSigningV1Beta1SignMode["SIGN_MODE_DIRECT_AUX"] = "SIGN_MODE_DIRECT_AUX";
    CosmosTxSigningV1Beta1SignMode["SIGN_MODE_LEGACY_AMINO_JSON"] = "SIGN_MODE_LEGACY_AMINO_JSON";
    CosmosTxSigningV1Beta1SignMode["SIGNMODEEIP191"] = "SIGN_MODE_EIP_191";
})(CosmosTxSigningV1Beta1SignMode || (CosmosTxSigningV1Beta1SignMode = {}));
/**
* BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.

 - BROADCAST_MODE_UNSPECIFIED: zero-value for mode ordering
 - BROADCAST_MODE_BLOCK: DEPRECATED: use BROADCAST_MODE_SYNC instead,
BROADCAST_MODE_BLOCK is not supported by the SDK from v0.47.x onwards.
 - BROADCAST_MODE_SYNC: BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for
a CheckTx execution response only.
 - BROADCAST_MODE_ASYNC: BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns
immediately.
*/
export var CosmosTxV1Beta1BroadcastMode;
(function (CosmosTxV1Beta1BroadcastMode) {
    CosmosTxV1Beta1BroadcastMode["BROADCAST_MODE_UNSPECIFIED"] = "BROADCAST_MODE_UNSPECIFIED";
    CosmosTxV1Beta1BroadcastMode["BROADCAST_MODE_BLOCK"] = "BROADCAST_MODE_BLOCK";
    CosmosTxV1Beta1BroadcastMode["BROADCAST_MODE_SYNC"] = "BROADCAST_MODE_SYNC";
    CosmosTxV1Beta1BroadcastMode["BROADCAST_MODE_ASYNC"] = "BROADCAST_MODE_ASYNC";
})(CosmosTxV1Beta1BroadcastMode || (CosmosTxV1Beta1BroadcastMode = {}));
/**
* - ORDER_BY_UNSPECIFIED: ORDER_BY_UNSPECIFIED specifies an unknown sorting order. OrderBy defaults to ASC in this case.
 - ORDER_BY_ASC: ORDER_BY_ASC defines ascending order
 - ORDER_BY_DESC: ORDER_BY_DESC defines descending order
*/
export var CosmosTxV1Beta1OrderBy;
(function (CosmosTxV1Beta1OrderBy) {
    CosmosTxV1Beta1OrderBy["ORDER_BY_UNSPECIFIED"] = "ORDER_BY_UNSPECIFIED";
    CosmosTxV1Beta1OrderBy["ORDER_BY_ASC"] = "ORDER_BY_ASC";
    CosmosTxV1Beta1OrderBy["ORDER_BY_DESC"] = "ORDER_BY_DESC";
})(CosmosTxV1Beta1OrderBy || (CosmosTxV1Beta1OrderBy = {}));
export var TendermintTypesBlockIDFlag;
(function (TendermintTypesBlockIDFlag) {
    TendermintTypesBlockIDFlag["BLOCK_ID_FLAG_UNKNOWN"] = "BLOCK_ID_FLAG_UNKNOWN";
    TendermintTypesBlockIDFlag["BLOCK_ID_FLAG_ABSENT"] = "BLOCK_ID_FLAG_ABSENT";
    TendermintTypesBlockIDFlag["BLOCK_ID_FLAG_COMMIT"] = "BLOCK_ID_FLAG_COMMIT";
    TendermintTypesBlockIDFlag["BLOCK_ID_FLAG_NIL"] = "BLOCK_ID_FLAG_NIL";
})(TendermintTypesBlockIDFlag || (TendermintTypesBlockIDFlag = {}));
/**
* SignedMsgType is a type of signed message in the consensus.

 - SIGNED_MSG_TYPE_PREVOTE: Votes
 - SIGNED_MSG_TYPE_PROPOSAL: Proposals
*/
export var TendermintTypesSignedMsgType;
(function (TendermintTypesSignedMsgType) {
    TendermintTypesSignedMsgType["SIGNED_MSG_TYPE_UNKNOWN"] = "SIGNED_MSG_TYPE_UNKNOWN";
    TendermintTypesSignedMsgType["SIGNED_MSG_TYPE_PREVOTE"] = "SIGNED_MSG_TYPE_PREVOTE";
    TendermintTypesSignedMsgType["SIGNED_MSG_TYPE_PRECOMMIT"] = "SIGNED_MSG_TYPE_PRECOMMIT";
    TendermintTypesSignedMsgType["SIGNED_MSG_TYPE_PROPOSAL"] = "SIGNED_MSG_TYPE_PROPOSAL";
})(TendermintTypesSignedMsgType || (TendermintTypesSignedMsgType = {}));
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
     * @tags Service
     * @name CosmosTxV1Beta1TxDecode
     * @summary TxDecode decodes the transaction.
     * @request POST:/cosmos/tx/v1beta1/decode
     */
    cosmosTxV1Beta1TxDecode = (body, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/decode`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.47
     *
     * @tags Service
     * @name CosmosTxV1Beta1TxDecodeAmino
     * @summary TxDecodeAmino decodes an Amino transaction from encoded bytes to JSON.
     * @request POST:/cosmos/tx/v1beta1/decode/amino
     */
    cosmosTxV1Beta1TxDecodeAmino = (body, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/decode/amino`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.47
     *
     * @tags Service
     * @name CosmosTxV1Beta1TxEncode
     * @summary TxEncode encodes the transaction.
     * @request POST:/cosmos/tx/v1beta1/encode
     */
    cosmosTxV1Beta1TxEncode = (body, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/encode`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.47
     *
     * @tags Service
     * @name CosmosTxV1Beta1TxEncodeAmino
     * @summary TxEncodeAmino encodes an Amino transaction from JSON to encoded bytes.
     * @request POST:/cosmos/tx/v1beta1/encode/amino
     */
    cosmosTxV1Beta1TxEncodeAmino = (body, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/encode/amino`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        ...params,
    });
    /**
     * No description
     *
     * @tags Service
     * @name CosmosTxV1Beta1Simulate
     * @summary Simulate simulates executing a transaction for estimating gas usage.
     * @request POST:/cosmos/tx/v1beta1/simulate
     */
    cosmosTxV1Beta1Simulate = (body, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/simulate`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        ...params,
    });
    /**
     * No description
     *
     * @tags Service
     * @name CosmosTxV1Beta1GetTxsEvent
     * @summary GetTxsEvent fetches txs by event.
     * @request GET:/cosmos/tx/v1beta1/txs
     */
    cosmosTxV1Beta1GetTxsEvent = (query, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/txs`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Service
     * @name CosmosTxV1Beta1BroadcastTx
     * @summary BroadcastTx broadcast transaction.
     * @request POST:/cosmos/tx/v1beta1/txs
     */
    cosmosTxV1Beta1BroadcastTx = (body, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/txs`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.45.2
     *
     * @tags Service
     * @name CosmosTxV1Beta1GetBlockWithTxs
     * @summary GetBlockWithTxs fetches a block with decoded txs.
     * @request GET:/cosmos/tx/v1beta1/txs/block/{height}
     */
    cosmosTxV1Beta1GetBlockWithTxs = (height, query, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/txs/block/${height}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Service
     * @name CosmosTxV1Beta1GetTx
     * @summary GetTx fetches a tx by hash.
     * @request GET:/cosmos/tx/v1beta1/txs/{hash}
     */
    cosmosTxV1Beta1GetTx = (hash, params = {}) => this.request({
        path: `/cosmos/tx/v1beta1/txs/${hash}`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map