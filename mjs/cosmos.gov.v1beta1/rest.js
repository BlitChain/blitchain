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
* ProposalStatus enumerates the valid statuses of a proposal.

 - PROPOSAL_STATUS_UNSPECIFIED: PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status.
 - PROPOSAL_STATUS_DEPOSIT_PERIOD: PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
period.
 - PROPOSAL_STATUS_VOTING_PERIOD: PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
period.
 - PROPOSAL_STATUS_PASSED: PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
passed.
 - PROPOSAL_STATUS_REJECTED: PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
been rejected.
 - PROPOSAL_STATUS_FAILED: PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
failed.
*/
export var CosmosGovV1Beta1ProposalStatus;
(function (CosmosGovV1Beta1ProposalStatus) {
    CosmosGovV1Beta1ProposalStatus["PROPOSAL_STATUS_UNSPECIFIED"] = "PROPOSAL_STATUS_UNSPECIFIED";
    CosmosGovV1Beta1ProposalStatus["PROPOSAL_STATUS_DEPOSIT_PERIOD"] = "PROPOSAL_STATUS_DEPOSIT_PERIOD";
    CosmosGovV1Beta1ProposalStatus["PROPOSAL_STATUS_VOTING_PERIOD"] = "PROPOSAL_STATUS_VOTING_PERIOD";
    CosmosGovV1Beta1ProposalStatus["PROPOSAL_STATUS_PASSED"] = "PROPOSAL_STATUS_PASSED";
    CosmosGovV1Beta1ProposalStatus["PROPOSAL_STATUS_REJECTED"] = "PROPOSAL_STATUS_REJECTED";
    CosmosGovV1Beta1ProposalStatus["PROPOSAL_STATUS_FAILED"] = "PROPOSAL_STATUS_FAILED";
})(CosmosGovV1Beta1ProposalStatus || (CosmosGovV1Beta1ProposalStatus = {}));
/**
* VoteOption enumerates the valid vote options for a given governance proposal.

 - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
 - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option.
 - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option.
 - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option.
 - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
*/
export var CosmosGovV1Beta1VoteOption;
(function (CosmosGovV1Beta1VoteOption) {
    CosmosGovV1Beta1VoteOption["VOTE_OPTION_UNSPECIFIED"] = "VOTE_OPTION_UNSPECIFIED";
    CosmosGovV1Beta1VoteOption["VOTE_OPTION_YES"] = "VOTE_OPTION_YES";
    CosmosGovV1Beta1VoteOption["VOTE_OPTION_ABSTAIN"] = "VOTE_OPTION_ABSTAIN";
    CosmosGovV1Beta1VoteOption["VOTE_OPTION_NO"] = "VOTE_OPTION_NO";
    CosmosGovV1Beta1VoteOption["VOTE_OPTION_NO_WITH_VETO"] = "VOTE_OPTION_NO_WITH_VETO";
})(CosmosGovV1Beta1VoteOption || (CosmosGovV1Beta1VoteOption = {}));
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
     * @name CosmosGovV1Beta1Params
     * @summary Params queries all parameters of the gov module.
     * @request GET:/cosmos/gov/v1beta1/params/{params_type}
     */
    cosmosGovV1Beta1Params = (paramsType, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/params/${paramsType}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1Proposals
     * @summary Proposals queries all proposals based on given status.
     * @request GET:/cosmos/gov/v1beta1/proposals
     */
    cosmosGovV1Beta1Proposals = (query, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1Proposal
     * @summary Proposal queries proposal details based on ProposalID.
     * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}
     */
    cosmosGovV1Beta1Proposal = (proposalId, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals/${proposalId}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1Deposits
     * @summary Deposits queries all deposits of a single proposal.
     * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits
     */
    cosmosGovV1Beta1Deposits = (proposalId, query, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals/${proposalId}/deposits`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1Deposit
     * @summary Deposit queries single deposit information based proposalID, depositAddr.
     * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}
     */
    cosmosGovV1Beta1Deposit = (proposalId, depositor, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals/${proposalId}/deposits/${depositor}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1TallyResult
     * @summary TallyResult queries the tally of a proposal vote.
     * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/tally
     */
    cosmosGovV1Beta1TallyResult = (proposalId, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals/${proposalId}/tally`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1Votes
     * @summary Votes queries votes of a given proposal.
     * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/votes
     */
    cosmosGovV1Beta1Votes = (proposalId, query, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals/${proposalId}/votes`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGovV1Beta1Vote
     * @summary Vote queries voted information based on proposalID, voterAddr.
     * @request GET:/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}
     */
    cosmosGovV1Beta1Vote = (proposalId, voter, params = {}) => this.request({
        path: `/cosmos/gov/v1beta1/proposals/${proposalId}/votes/${voter}`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map