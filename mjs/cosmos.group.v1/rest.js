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
* ProposalExecutorResult defines types of proposal executor results.

 - PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED: An empty value is not allowed.
 - PROPOSAL_EXECUTOR_RESULT_NOT_RUN: We have not yet run the executor.
 - PROPOSAL_EXECUTOR_RESULT_SUCCESS: The executor was successful and proposed action updated state.
 - PROPOSAL_EXECUTOR_RESULT_FAILURE: The executor returned an error and proposed action didn't update state.
*/
export var CosmosGroupV1ProposalExecutorResult;
(function (CosmosGroupV1ProposalExecutorResult) {
    CosmosGroupV1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED"] = "PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED";
    CosmosGroupV1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_NOT_RUN"] = "PROPOSAL_EXECUTOR_RESULT_NOT_RUN";
    CosmosGroupV1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_SUCCESS"] = "PROPOSAL_EXECUTOR_RESULT_SUCCESS";
    CosmosGroupV1ProposalExecutorResult["PROPOSAL_EXECUTOR_RESULT_FAILURE"] = "PROPOSAL_EXECUTOR_RESULT_FAILURE";
})(CosmosGroupV1ProposalExecutorResult || (CosmosGroupV1ProposalExecutorResult = {}));
/**
* ProposalStatus defines proposal statuses.

 - PROPOSAL_STATUS_UNSPECIFIED: An empty value is invalid and not allowed.
 - PROPOSAL_STATUS_SUBMITTED: Initial status of a proposal when submitted.
 - PROPOSAL_STATUS_ACCEPTED: Final status of a proposal when the final tally is done and the outcome
passes the group policy's decision policy.
 - PROPOSAL_STATUS_REJECTED: Final status of a proposal when the final tally is done and the outcome
is rejected by the group policy's decision policy.
 - PROPOSAL_STATUS_ABORTED: Final status of a proposal when the group policy is modified before the
final tally.
 - PROPOSAL_STATUS_WITHDRAWN: A proposal can be withdrawn before the voting start time by the owner.
When this happens the final status is Withdrawn.
*/
export var CosmosGroupV1ProposalStatus;
(function (CosmosGroupV1ProposalStatus) {
    CosmosGroupV1ProposalStatus["PROPOSAL_STATUS_UNSPECIFIED"] = "PROPOSAL_STATUS_UNSPECIFIED";
    CosmosGroupV1ProposalStatus["PROPOSAL_STATUS_SUBMITTED"] = "PROPOSAL_STATUS_SUBMITTED";
    CosmosGroupV1ProposalStatus["PROPOSAL_STATUS_ACCEPTED"] = "PROPOSAL_STATUS_ACCEPTED";
    CosmosGroupV1ProposalStatus["PROPOSAL_STATUS_REJECTED"] = "PROPOSAL_STATUS_REJECTED";
    CosmosGroupV1ProposalStatus["PROPOSAL_STATUS_ABORTED"] = "PROPOSAL_STATUS_ABORTED";
    CosmosGroupV1ProposalStatus["PROPOSAL_STATUS_WITHDRAWN"] = "PROPOSAL_STATUS_WITHDRAWN";
})(CosmosGroupV1ProposalStatus || (CosmosGroupV1ProposalStatus = {}));
/**
* VoteOption enumerates the valid vote options for a given proposal.

 - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines an unspecified vote option which will
return an error.
 - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option.
 - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option.
 - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option.
 - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
*/
export var CosmosGroupV1VoteOption;
(function (CosmosGroupV1VoteOption) {
    CosmosGroupV1VoteOption["VOTE_OPTION_UNSPECIFIED"] = "VOTE_OPTION_UNSPECIFIED";
    CosmosGroupV1VoteOption["VOTE_OPTION_YES"] = "VOTE_OPTION_YES";
    CosmosGroupV1VoteOption["VOTE_OPTION_ABSTAIN"] = "VOTE_OPTION_ABSTAIN";
    CosmosGroupV1VoteOption["VOTE_OPTION_NO"] = "VOTE_OPTION_NO";
    CosmosGroupV1VoteOption["VOTE_OPTION_NO_WITH_VETO"] = "VOTE_OPTION_NO_WITH_VETO";
})(CosmosGroupV1VoteOption || (CosmosGroupV1VoteOption = {}));
/**
* Exec defines modes of execution of a proposal on creation or on new vote.

 - EXEC_UNSPECIFIED: An empty value means that there should be a separate
MsgExec request for the proposal to execute.
 - EXEC_TRY: Try to execute the proposal immediately.
If the proposal is not allowed per the DecisionPolicy,
the proposal will still be open and could
be executed at a later point.
*/
export var CosmosGroupV1Exec;
(function (CosmosGroupV1Exec) {
    CosmosGroupV1Exec["EXEC_UNSPECIFIED"] = "EXEC_UNSPECIFIED";
    CosmosGroupV1Exec["EXEC_TRY"] = "EXEC_TRY";
})(CosmosGroupV1Exec || (CosmosGroupV1Exec = {}));
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
     * @name CosmosGroupV1GroupInfo
     * @summary GroupInfo queries group info based on group id.
     * @request GET:/cosmos/group/v1/group_info/{group_id}
     */
    cosmosGroupV1GroupInfo = (groupId, params = {}) => this.request({
        path: `/cosmos/group/v1/group_info/${groupId}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1GroupMembers
     * @summary GroupMembers queries members of a group by group id.
     * @request GET:/cosmos/group/v1/group_members/{group_id}
     */
    cosmosGroupV1GroupMembers = (groupId, query, params = {}) => this.request({
        path: `/cosmos/group/v1/group_members/${groupId}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1GroupPoliciesByAdmin
     * @summary GroupPoliciesByAdmin queries group policies by admin address.
     * @request GET:/cosmos/group/v1/group_policies_by_admin/{admin}
     */
    cosmosGroupV1GroupPoliciesByAdmin = (admin, query, params = {}) => this.request({
        path: `/cosmos/group/v1/group_policies_by_admin/${admin}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1GroupPoliciesByGroup
     * @summary GroupPoliciesByGroup queries group policies by group id.
     * @request GET:/cosmos/group/v1/group_policies_by_group/{group_id}
     */
    cosmosGroupV1GroupPoliciesByGroup = (groupId, query, params = {}) => this.request({
        path: `/cosmos/group/v1/group_policies_by_group/${groupId}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1GroupPolicyInfo
     * @summary GroupPolicyInfo queries group policy info based on account address of group policy.
     * @request GET:/cosmos/group/v1/group_policy_info/{address}
     */
    cosmosGroupV1GroupPolicyInfo = (address, params = {}) => this.request({
        path: `/cosmos/group/v1/group_policy_info/${address}`,
        method: "GET",
        ...params,
    });
    /**
     * @description Since: cosmos-sdk 0.47.1
     *
     * @tags Query
     * @name CosmosGroupV1Groups
     * @summary Groups queries all groups in state.
     * @request GET:/cosmos/group/v1/groups
     */
    cosmosGroupV1Groups = (query, params = {}) => this.request({
        path: `/cosmos/group/v1/groups`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1GroupsByAdmin
     * @summary GroupsByAdmin queries groups by admin address.
     * @request GET:/cosmos/group/v1/groups_by_admin/{admin}
     */
    cosmosGroupV1GroupsByAdmin = (admin, query, params = {}) => this.request({
        path: `/cosmos/group/v1/groups_by_admin/${admin}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1GroupsByMember
     * @summary GroupsByMember queries groups by member address.
     * @request GET:/cosmos/group/v1/groups_by_member/{address}
     */
    cosmosGroupV1GroupsByMember = (address, query, params = {}) => this.request({
        path: `/cosmos/group/v1/groups_by_member/${address}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1Proposal
     * @summary Proposal queries a proposal based on proposal id.
     * @request GET:/cosmos/group/v1/proposal/{proposal_id}
     */
    cosmosGroupV1Proposal = (proposalId, params = {}) => this.request({
        path: `/cosmos/group/v1/proposal/${proposalId}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosGroupV1TallyResult
   * @summary TallyResult returns the tally result of a proposal. If the proposal is
  still in voting period, then this query computes the current tally state,
  which might not be final. On the other hand, if the proposal is final,
  then it simply returns the `final_tally_result` state stored in the
  proposal itself.
   * @request GET:/cosmos/group/v1/proposals/{proposal_id}/tally
   */
    cosmosGroupV1TallyResult = (proposalId, params = {}) => this.request({
        path: `/cosmos/group/v1/proposals/${proposalId}/tally`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1ProposalsByGroupPolicy
     * @summary ProposalsByGroupPolicy queries proposals based on account address of group policy.
     * @request GET:/cosmos/group/v1/proposals_by_group_policy/{address}
     */
    cosmosGroupV1ProposalsByGroupPolicy = (address, query, params = {}) => this.request({
        path: `/cosmos/group/v1/proposals_by_group_policy/${address}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1VoteByProposalVoter
     * @summary VoteByProposalVoter queries a vote by proposal id and voter.
     * @request GET:/cosmos/group/v1/vote_by_proposal_voter/{proposal_id}/{voter}
     */
    cosmosGroupV1VoteByProposalVoter = (proposalId, voter, params = {}) => this.request({
        path: `/cosmos/group/v1/vote_by_proposal_voter/${proposalId}/${voter}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1VotesByProposal
     * @summary VotesByProposal queries a vote by proposal id.
     * @request GET:/cosmos/group/v1/votes_by_proposal/{proposal_id}
     */
    cosmosGroupV1VotesByProposal = (proposalId, query, params = {}) => this.request({
        path: `/cosmos/group/v1/votes_by_proposal/${proposalId}`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosGroupV1VotesByVoter
     * @summary VotesByVoter queries a vote by voter.
     * @request GET:/cosmos/group/v1/votes_by_voter/{voter}
     */
    cosmosGroupV1VotesByVoter = (voter, query, params = {}) => this.request({
        path: `/cosmos/group/v1/votes_by_voter/${voter}`,
        method: "GET",
        query: query,
        ...params,
    });
}
//# sourceMappingURL=rest.js.map