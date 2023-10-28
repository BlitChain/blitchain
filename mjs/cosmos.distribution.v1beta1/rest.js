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
     * @name CosmosDistributionV1Beta1CommunityPool
     * @summary CommunityPool queries the community pool coins.
     * @request GET:/cosmos/distribution/v1beta1/community_pool
     */
    cosmosDistributionV1Beta1CommunityPool = (params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/community_pool`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosDistributionV1Beta1DelegationTotalRewards
   * @summary DelegationTotalRewards queries the total rewards accrued by a each
  validator.
   * @request GET:/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards
   */
    cosmosDistributionV1Beta1DelegationTotalRewards = (delegatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/delegators/${delegatorAddress}/rewards`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1DelegationRewards
     * @summary DelegationRewards queries the total rewards accrued by a delegation.
     * @request GET:/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards/{validator_address}
     */
    cosmosDistributionV1Beta1DelegationRewards = (delegatorAddress, validatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/delegators/${delegatorAddress}/rewards/${validatorAddress}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1DelegatorValidators
     * @summary DelegatorValidators queries the validators of a delegator.
     * @request GET:/cosmos/distribution/v1beta1/delegators/{delegator_address}/validators
     */
    cosmosDistributionV1Beta1DelegatorValidators = (delegatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/delegators/${delegatorAddress}/validators`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1DelegatorWithdrawAddress
     * @summary DelegatorWithdrawAddress queries withdraw address of a delegator.
     * @request GET:/cosmos/distribution/v1beta1/delegators/{delegator_address}/withdraw_address
     */
    cosmosDistributionV1Beta1DelegatorWithdrawAddress = (delegatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/delegators/${delegatorAddress}/withdraw_address`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1Params
     * @summary Params queries params of the distribution module.
     * @request GET:/cosmos/distribution/v1beta1/params
     */
    cosmosDistributionV1Beta1Params = (params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/params`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1ValidatorDistributionInfo
     * @summary ValidatorDistributionInfo queries validator commission and self-delegation rewards for validator
     * @request GET:/cosmos/distribution/v1beta1/validators/{validator_address}
     */
    cosmosDistributionV1Beta1ValidatorDistributionInfo = (validatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/validators/${validatorAddress}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1ValidatorCommission
     * @summary ValidatorCommission queries accumulated commission for a validator.
     * @request GET:/cosmos/distribution/v1beta1/validators/{validator_address}/commission
     */
    cosmosDistributionV1Beta1ValidatorCommission = (validatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/validators/${validatorAddress}/commission`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1ValidatorOutstandingRewards
     * @summary ValidatorOutstandingRewards queries rewards of a validator address.
     * @request GET:/cosmos/distribution/v1beta1/validators/{validator_address}/outstanding_rewards
     */
    cosmosDistributionV1Beta1ValidatorOutstandingRewards = (validatorAddress, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/validators/${validatorAddress}/outstanding_rewards`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosDistributionV1Beta1ValidatorSlashes
     * @summary ValidatorSlashes queries slash events of a validator.
     * @request GET:/cosmos/distribution/v1beta1/validators/{validator_address}/slashes
     */
    cosmosDistributionV1Beta1ValidatorSlashes = (validatorAddress, query, params = {}) => this.request({
        path: `/cosmos/distribution/v1beta1/validators/${validatorAddress}/slashes`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map