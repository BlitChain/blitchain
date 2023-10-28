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
* BondStatus is the status of a validator.

 - BOND_STATUS_UNSPECIFIED: UNSPECIFIED defines an invalid validator status.
 - BOND_STATUS_UNBONDED: UNBONDED defines a validator that is not bonded.
 - BOND_STATUS_UNBONDING: UNBONDING defines a validator that is unbonding.
 - BOND_STATUS_BONDED: BONDED defines a validator that is bonded.
*/
export var CosmosStakingV1Beta1BondStatus;
(function (CosmosStakingV1Beta1BondStatus) {
    CosmosStakingV1Beta1BondStatus["BOND_STATUS_UNSPECIFIED"] = "BOND_STATUS_UNSPECIFIED";
    CosmosStakingV1Beta1BondStatus["BOND_STATUS_UNBONDED"] = "BOND_STATUS_UNBONDED";
    CosmosStakingV1Beta1BondStatus["BOND_STATUS_UNBONDING"] = "BOND_STATUS_UNBONDING";
    CosmosStakingV1Beta1BondStatus["BOND_STATUS_BONDED"] = "BOND_STATUS_BONDED";
})(CosmosStakingV1Beta1BondStatus || (CosmosStakingV1Beta1BondStatus = {}));
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
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosStakingV1Beta1DelegatorDelegations
     * @summary DelegatorDelegations queries all delegations of a given delegator address.
     * @request GET:/cosmos/staking/v1beta1/delegations/{delegator_addr}
     */
    cosmosStakingV1Beta1DelegatorDelegations = (delegatorAddr, query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/delegations/${delegatorAddr}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosStakingV1Beta1Redelegations
     * @summary Redelegations queries redelegations of given address.
     * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations
     */
    cosmosStakingV1Beta1Redelegations = (delegatorAddr, query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/redelegations`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
   *
   * @tags Query
   * @name CosmosStakingV1Beta1DelegatorUnbondingDelegations
   * @summary DelegatorUnbondingDelegations queries all unbonding delegations of a given
  delegator address.
   * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations
   */
    cosmosStakingV1Beta1DelegatorUnbondingDelegations = (delegatorAddr, query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/unbonding_delegations`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
   *
   * @tags Query
   * @name CosmosStakingV1Beta1DelegatorValidators
   * @summary DelegatorValidators queries all validators info for given delegator
  address.
   * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators
   */
    cosmosStakingV1Beta1DelegatorValidators = (delegatorAddr, query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/validators`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosStakingV1Beta1DelegatorValidator
   * @summary DelegatorValidator queries validator info for given delegator validator
  pair.
   * @request GET:/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}
   */
    cosmosStakingV1Beta1DelegatorValidator = (delegatorAddr, validatorAddr, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/delegators/${delegatorAddr}/validators/${validatorAddr}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosStakingV1Beta1HistoricalInfo
     * @summary HistoricalInfo queries the historical info for given height.
     * @request GET:/cosmos/staking/v1beta1/historical_info/{height}
     */
    cosmosStakingV1Beta1HistoricalInfo = (height, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/historical_info/${height}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosStakingV1Beta1Params
     * @summary Parameters queries the staking parameters.
     * @request GET:/cosmos/staking/v1beta1/params
     */
    cosmosStakingV1Beta1Params = (params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/params`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosStakingV1Beta1Pool
     * @summary Pool queries the pool info.
     * @request GET:/cosmos/staking/v1beta1/pool
     */
    cosmosStakingV1Beta1Pool = (params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/pool`,
        method: "GET",
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosStakingV1Beta1Validators
     * @summary Validators queries all validators that match the given status.
     * @request GET:/cosmos/staking/v1beta1/validators
     */
    cosmosStakingV1Beta1Validators = (query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/validators`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosStakingV1Beta1Validator
     * @summary Validator queries validator info for given validator address.
     * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}
     */
    cosmosStakingV1Beta1Validator = (validatorAddr, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/validators/${validatorAddr}`,
        method: "GET",
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosStakingV1Beta1ValidatorDelegations
     * @summary ValidatorDelegations queries delegate info for given validator.
     * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/delegations
     */
    cosmosStakingV1Beta1ValidatorDelegations = (validatorAddr, query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/delegations`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosStakingV1Beta1Delegation
     * @summary Delegation queries delegate info for given validator delegator pair.
     * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}
     */
    cosmosStakingV1Beta1Delegation = (validatorAddr, delegatorAddr, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/delegations/${delegatorAddr}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosStakingV1Beta1UnbondingDelegation
   * @summary UnbondingDelegation queries unbonding info for given validator delegator
  pair.
   * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation
   */
    cosmosStakingV1Beta1UnbondingDelegation = (validatorAddr, delegatorAddr, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/delegations/${delegatorAddr}/unbonding_delegation`,
        method: "GET",
        ...params,
    });
    /**
     * @description When called from another module, this query might consume a high amount of gas if the pagination field is incorrectly set.
     *
     * @tags Query
     * @name CosmosStakingV1Beta1ValidatorUnbondingDelegations
     * @summary ValidatorUnbondingDelegations queries unbonding delegations of a validator.
     * @request GET:/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations
     */
    cosmosStakingV1Beta1ValidatorUnbondingDelegations = (validatorAddr, query, params = {}) => this.request({
        path: `/cosmos/staking/v1beta1/validators/${validatorAddr}/unbonding_delegations`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map