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
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }
    // or ...
    if (any.isSameTypeAs(Foo.getDefaultInstance())) {
      foo = any.unpack(Foo.getDefaultInstance());
    }

Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".

JSON

The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface GoogleProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

export interface GoogleRpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: { "@type"?: string }[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface CosmosBaseQueryV1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface CosmosBaseQueryV1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface CosmosBaseV1Beta1Coin {
  denom?: string;
  amount?: string;
}

/**
* Deposit defines an amount deposited by an account address to an active
proposal.
*/
export interface CosmosGovV1Deposit {
  /**
   * proposal_id defines the unique id of the proposal.
   * @format uint64
   */
  proposal_id?: string;

  /** depositor defines the deposit addresses from the proposals. */
  depositor?: string;

  /** amount to be deposited by depositor. */
  amount?: { denom?: string; amount?: string }[];
}

/**
 * DepositParams defines the params for deposits on governance proposals.
 */
export interface CosmosGovV1DepositParams {
  /** Minimum deposit for a proposal to enter voting period. */
  min_deposit?: { denom?: string; amount?: string }[];

  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   * months.
   */
  max_deposit_period?: string;
}

/**
* Params defines the parameters for the x/gov module.

Since: cosmos-sdk 0.47
*/
export interface CosmosGovV1Params {
  /** Minimum deposit for a proposal to enter voting period. */
  min_deposit?: { denom?: string; amount?: string }[];

  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   * months.
   */
  max_deposit_period?: string;

  /** Duration of the voting period. */
  voting_period?: string;

  /**
   * Minimum percentage of total stake needed to vote for a result to be
   *  considered valid.
   */
  quorum?: string;

  /** Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. */
  threshold?: string;

  /**
   * Minimum value of Veto votes to Total votes ratio for proposal to be
   *  vetoed. Default value: 1/3.
   */
  veto_threshold?: string;

  /** The ratio representing the proportion of the deposit value that must be paid at proposal submission. */
  min_initial_deposit_ratio?: string;

  /** burn deposits if a proposal does not meet quorum */
  burn_vote_quorum?: boolean;

  /** burn deposits if the proposal does not enter voting period */
  burn_proposal_deposit_prevote?: boolean;

  /** burn deposits if quorum with vote type no_veto is met */
  burn_vote_veto?: boolean;
}

/**
 * Proposal defines the core field members of a governance proposal.
 */
export interface CosmosGovV1Proposal {
  /**
   * id defines the unique id of the proposal.
   * @format uint64
   */
  id?: string;

  /** messages are the arbitrary messages to be executed if the proposal passes. */
  messages?: { "@type"?: string }[];

  /** status defines the proposal status. */
  status?:
    | "PROPOSAL_STATUS_UNSPECIFIED"
    | "PROPOSAL_STATUS_DEPOSIT_PERIOD"
    | "PROPOSAL_STATUS_VOTING_PERIOD"
    | "PROPOSAL_STATUS_PASSED"
    | "PROPOSAL_STATUS_REJECTED"
    | "PROPOSAL_STATUS_FAILED";

  /**
   * final_tally_result is the final tally result of the proposal. When
   * querying a proposal via gRPC, this field is not populated until the
   * proposal's voting period has ended.
   */
  final_tally_result?: { yes_count?: string; abstain_count?: string; no_count?: string; no_with_veto_count?: string };

  /**
   * submit_time is the time of proposal submission.
   * @format date-time
   */
  submit_time?: string;

  /**
   * deposit_end_time is the end time for deposition.
   * @format date-time
   */
  deposit_end_time?: string;

  /** total_deposit is the total deposit on the proposal. */
  total_deposit?: { denom?: string; amount?: string }[];

  /**
   * voting_start_time is the starting time to vote on a proposal.
   * @format date-time
   */
  voting_start_time?: string;

  /**
   * voting_end_time is the end time of voting on a proposal.
   * @format date-time
   */
  voting_end_time?: string;

  /** metadata is any arbitrary metadata attached to the proposal. */
  metadata?: string;

  /**
   * title is the title of the proposal
   * Since: cosmos-sdk 0.47
   */
  title?: string;

  /**
   * summary is a short summary of the proposal
   * Since: cosmos-sdk 0.47
   */
  summary?: string;

  /**
   * Proposer is the address of the proposal sumbitter
   * Since: cosmos-sdk 0.47
   */
  proposer?: string;
}

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
export enum CosmosGovV1ProposalStatus {
  PROPOSAL_STATUS_UNSPECIFIED = "PROPOSAL_STATUS_UNSPECIFIED",
  PROPOSAL_STATUS_DEPOSIT_PERIOD = "PROPOSAL_STATUS_DEPOSIT_PERIOD",
  PROPOSAL_STATUS_VOTING_PERIOD = "PROPOSAL_STATUS_VOTING_PERIOD",
  PROPOSAL_STATUS_PASSED = "PROPOSAL_STATUS_PASSED",
  PROPOSAL_STATUS_REJECTED = "PROPOSAL_STATUS_REJECTED",
  PROPOSAL_STATUS_FAILED = "PROPOSAL_STATUS_FAILED",
}

/**
 * QueryDepositResponse is the response type for the Query/Deposit RPC method.
 */
export interface CosmosGovV1QueryDepositResponse {
  /** deposit defines the requested deposit. */
  deposit?: { proposal_id?: string; depositor?: string; amount?: { denom?: string; amount?: string }[] };
}

/**
 * QueryDepositsResponse is the response type for the Query/Deposits RPC method.
 */
export interface CosmosGovV1QueryDepositsResponse {
  /** deposits defines the requested deposits. */
  deposits?: { proposal_id?: string; depositor?: string; amount?: { denom?: string; amount?: string }[] }[];

  /** pagination defines the pagination in the response. */
  pagination?: { next_key?: string; total?: string };
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 */
export interface CosmosGovV1QueryParamsResponse {
  /**
   * Deprecated: Prefer to use `params` instead.
   * voting_params defines the parameters related to voting.
   */
  voting_params?: { voting_period?: string };

  /**
   * Deprecated: Prefer to use `params` instead.
   * deposit_params defines the parameters related to deposit.
   */
  deposit_params?: { min_deposit?: { denom?: string; amount?: string }[]; max_deposit_period?: string };

  /**
   * Deprecated: Prefer to use `params` instead.
   * tally_params defines the parameters related to tally.
   */
  tally_params?: { quorum?: string; threshold?: string; veto_threshold?: string };

  /**
   * params defines all the paramaters of x/gov module.
   *
   * Since: cosmos-sdk 0.47
   */
  params?: {
    min_deposit?: { denom?: string; amount?: string }[];
    max_deposit_period?: string;
    voting_period?: string;
    quorum?: string;
    threshold?: string;
    veto_threshold?: string;
    min_initial_deposit_ratio?: string;
    burn_vote_quorum?: boolean;
    burn_proposal_deposit_prevote?: boolean;
    burn_vote_veto?: boolean;
  };
}

/**
 * QueryProposalResponse is the response type for the Query/Proposal RPC method.
 */
export interface CosmosGovV1QueryProposalResponse {
  /** proposal is the requested governance proposal. */
  proposal?: {
    id?: string;
    messages?: { "@type"?: string }[];
    status?:
      | "PROPOSAL_STATUS_UNSPECIFIED"
      | "PROPOSAL_STATUS_DEPOSIT_PERIOD"
      | "PROPOSAL_STATUS_VOTING_PERIOD"
      | "PROPOSAL_STATUS_PASSED"
      | "PROPOSAL_STATUS_REJECTED"
      | "PROPOSAL_STATUS_FAILED";
    final_tally_result?: { yes_count?: string; abstain_count?: string; no_count?: string; no_with_veto_count?: string };
    submit_time?: string;
    deposit_end_time?: string;
    total_deposit?: { denom?: string; amount?: string }[];
    voting_start_time?: string;
    voting_end_time?: string;
    metadata?: string;
    title?: string;
    summary?: string;
    proposer?: string;
  };
}

/**
* QueryProposalsResponse is the response type for the Query/Proposals RPC
method.
*/
export interface CosmosGovV1QueryProposalsResponse {
  /** proposals defines all the requested governance proposals. */
  proposals?: {
    id?: string;
    messages?: { "@type"?: string }[];
    status?:
      | "PROPOSAL_STATUS_UNSPECIFIED"
      | "PROPOSAL_STATUS_DEPOSIT_PERIOD"
      | "PROPOSAL_STATUS_VOTING_PERIOD"
      | "PROPOSAL_STATUS_PASSED"
      | "PROPOSAL_STATUS_REJECTED"
      | "PROPOSAL_STATUS_FAILED";
    final_tally_result?: { yes_count?: string; abstain_count?: string; no_count?: string; no_with_veto_count?: string };
    submit_time?: string;
    deposit_end_time?: string;
    total_deposit?: { denom?: string; amount?: string }[];
    voting_start_time?: string;
    voting_end_time?: string;
    metadata?: string;
    title?: string;
    summary?: string;
    proposer?: string;
  }[];

  /** pagination defines the pagination in the response. */
  pagination?: { next_key?: string; total?: string };
}

/**
 * QueryTallyResultResponse is the response type for the Query/Tally RPC method.
 */
export interface CosmosGovV1QueryTallyResultResponse {
  /** tally defines the requested tally. */
  tally?: { yes_count?: string; abstain_count?: string; no_count?: string; no_with_veto_count?: string };
}

/**
 * QueryVoteResponse is the response type for the Query/Vote RPC method.
 */
export interface CosmosGovV1QueryVoteResponse {
  /** vote defines the queried vote. */
  vote?: {
    proposal_id?: string;
    voter?: string;
    options?: {
      option?:
        | "VOTE_OPTION_UNSPECIFIED"
        | "VOTE_OPTION_YES"
        | "VOTE_OPTION_ABSTAIN"
        | "VOTE_OPTION_NO"
        | "VOTE_OPTION_NO_WITH_VETO";
      weight?: string;
    }[];
    metadata?: string;
  };
}

/**
 * QueryVotesResponse is the response type for the Query/Votes RPC method.
 */
export interface CosmosGovV1QueryVotesResponse {
  /** votes defines the queried votes. */
  votes?: {
    proposal_id?: string;
    voter?: string;
    options?: {
      option?:
        | "VOTE_OPTION_UNSPECIFIED"
        | "VOTE_OPTION_YES"
        | "VOTE_OPTION_ABSTAIN"
        | "VOTE_OPTION_NO"
        | "VOTE_OPTION_NO_WITH_VETO";
      weight?: string;
    }[];
    metadata?: string;
  }[];

  /** pagination defines the pagination in the response. */
  pagination?: { next_key?: string; total?: string };
}

/**
 * TallyParams defines the params for tallying votes on governance proposals.
 */
export interface CosmosGovV1TallyParams {
  /**
   * Minimum percentage of total stake needed to vote for a result to be
   * considered valid.
   */
  quorum?: string;

  /** Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. */
  threshold?: string;

  /**
   * Minimum value of Veto votes to Total votes ratio for proposal to be
   * vetoed. Default value: 1/3.
   */
  veto_threshold?: string;
}

/**
 * TallyResult defines a standard tally for a governance proposal.
 */
export interface CosmosGovV1TallyResult {
  /** yes_count is the number of yes votes on a proposal. */
  yes_count?: string;

  /** abstain_count is the number of abstain votes on a proposal. */
  abstain_count?: string;

  /** no_count is the number of no votes on a proposal. */
  no_count?: string;

  /** no_with_veto_count is the number of no with veto votes on a proposal. */
  no_with_veto_count?: string;
}

/**
* Vote defines a vote on a governance proposal.
A Vote consists of a proposal ID, the voter, and the vote option.
*/
export interface CosmosGovV1Vote {
  /**
   * proposal_id defines the unique id of the proposal.
   * @format uint64
   */
  proposal_id?: string;

  /** voter is the voter address of the proposal. */
  voter?: string;

  /** options is the weighted vote options. */
  options?: {
    option?:
      | "VOTE_OPTION_UNSPECIFIED"
      | "VOTE_OPTION_YES"
      | "VOTE_OPTION_ABSTAIN"
      | "VOTE_OPTION_NO"
      | "VOTE_OPTION_NO_WITH_VETO";
    weight?: string;
  }[];

  /** metadata is any  arbitrary metadata to attached to the vote. */
  metadata?: string;
}

/**
* VoteOption enumerates the valid vote options for a given governance proposal.

 - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
 - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option.
 - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option.
 - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option.
 - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
*/
export enum CosmosGovV1VoteOption {
  VOTE_OPTION_UNSPECIFIED = "VOTE_OPTION_UNSPECIFIED",
  VOTE_OPTION_YES = "VOTE_OPTION_YES",
  VOTE_OPTION_ABSTAIN = "VOTE_OPTION_ABSTAIN",
  VOTE_OPTION_NO = "VOTE_OPTION_NO",
  VOTE_OPTION_NO_WITH_VETO = "VOTE_OPTION_NO_WITH_VETO",
}

/**
 * VotingParams defines the params for voting on governance proposals.
 */
export interface CosmosGovV1VotingParams {
  /** Duration of the voting period. */
  voting_period?: string;
}

/**
 * WeightedVoteOption defines a unit of vote for vote split.
 */
export interface CosmosGovV1WeightedVoteOption {
  /** option defines the valid vote options, it must not contain duplicate vote options. */
  option?:
    | "VOTE_OPTION_UNSPECIFIED"
    | "VOTE_OPTION_YES"
    | "VOTE_OPTION_ABSTAIN"
    | "VOTE_OPTION_NO"
    | "VOTE_OPTION_NO_WITH_VETO";

  /** weight is the vote weight associated with the vote option. */
  weight?: string;
}

/**
 * MsgDepositResponse defines the Msg/Deposit response type.
 */
export type CosmosGovV1MsgDepositResponse = object;

/**
 * MsgExecLegacyContentResponse defines the Msg/ExecLegacyContent response type.
 */
export type CosmosGovV1MsgExecLegacyContentResponse = object;

/**
 * MsgSubmitProposalResponse defines the Msg/SubmitProposal response type.
 */
export interface CosmosGovV1MsgSubmitProposalResponse {
  /**
   * proposal_id defines the unique id of the proposal.
   * @format uint64
   */
  proposal_id?: string;
}

/**
* MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: cosmos-sdk 0.47
*/
export type CosmosGovV1MsgUpdateParamsResponse = object;

/**
 * MsgVoteResponse defines the Msg/Vote response type.
 */
export type CosmosGovV1MsgVoteResponse = object;

/**
 * MsgVoteWeightedResponse defines the Msg/VoteWeighted response type.
 */
export type CosmosGovV1MsgVoteWeightedResponse = object;

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
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

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
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
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Params
   * @summary Params queries all parameters of the gov module.
   * @request GET:/cosmos/gov/v1/params/{params_type}
   */
  cosmosGovV1Params = (paramsType: string, params: RequestParams = {}) =>
    this.request<
      {
        voting_params?: { voting_period?: string };
        deposit_params?: { min_deposit?: { denom?: string; amount?: string }[]; max_deposit_period?: string };
        tally_params?: { quorum?: string; threshold?: string; veto_threshold?: string };
        params?: {
          min_deposit?: { denom?: string; amount?: string }[];
          max_deposit_period?: string;
          voting_period?: string;
          quorum?: string;
          threshold?: string;
          veto_threshold?: string;
          min_initial_deposit_ratio?: string;
          burn_vote_quorum?: boolean;
          burn_proposal_deposit_prevote?: boolean;
          burn_vote_veto?: boolean;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/params/${paramsType}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Proposals
   * @summary Proposals queries all proposals based on given status.
   * @request GET:/cosmos/gov/v1/proposals
   */
  cosmosGovV1Proposals = (
    query?: {
      proposal_status?:
        | "PROPOSAL_STATUS_UNSPECIFIED"
        | "PROPOSAL_STATUS_DEPOSIT_PERIOD"
        | "PROPOSAL_STATUS_VOTING_PERIOD"
        | "PROPOSAL_STATUS_PASSED"
        | "PROPOSAL_STATUS_REJECTED"
        | "PROPOSAL_STATUS_FAILED";
      voter?: string;
      depositor?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        proposals?: {
          id?: string;
          messages?: { "@type"?: string }[];
          status?:
            | "PROPOSAL_STATUS_UNSPECIFIED"
            | "PROPOSAL_STATUS_DEPOSIT_PERIOD"
            | "PROPOSAL_STATUS_VOTING_PERIOD"
            | "PROPOSAL_STATUS_PASSED"
            | "PROPOSAL_STATUS_REJECTED"
            | "PROPOSAL_STATUS_FAILED";
          final_tally_result?: {
            yes_count?: string;
            abstain_count?: string;
            no_count?: string;
            no_with_veto_count?: string;
          };
          submit_time?: string;
          deposit_end_time?: string;
          total_deposit?: { denom?: string; amount?: string }[];
          voting_start_time?: string;
          voting_end_time?: string;
          metadata?: string;
          title?: string;
          summary?: string;
          proposer?: string;
        }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Proposal
   * @summary Proposal queries proposal details based on ProposalID.
   * @request GET:/cosmos/gov/v1/proposals/{proposal_id}
   */
  cosmosGovV1Proposal = (proposalId: string, params: RequestParams = {}) =>
    this.request<
      {
        proposal?: {
          id?: string;
          messages?: { "@type"?: string }[];
          status?:
            | "PROPOSAL_STATUS_UNSPECIFIED"
            | "PROPOSAL_STATUS_DEPOSIT_PERIOD"
            | "PROPOSAL_STATUS_VOTING_PERIOD"
            | "PROPOSAL_STATUS_PASSED"
            | "PROPOSAL_STATUS_REJECTED"
            | "PROPOSAL_STATUS_FAILED";
          final_tally_result?: {
            yes_count?: string;
            abstain_count?: string;
            no_count?: string;
            no_with_veto_count?: string;
          };
          submit_time?: string;
          deposit_end_time?: string;
          total_deposit?: { denom?: string; amount?: string }[];
          voting_start_time?: string;
          voting_end_time?: string;
          metadata?: string;
          title?: string;
          summary?: string;
          proposer?: string;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals/${proposalId}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Deposits
   * @summary Deposits queries all deposits of a single proposal.
   * @request GET:/cosmos/gov/v1/proposals/{proposal_id}/deposits
   */
  cosmosGovV1Deposits = (
    proposalId: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        deposits?: { proposal_id?: string; depositor?: string; amount?: { denom?: string; amount?: string }[] }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals/${proposalId}/deposits`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Deposit
   * @summary Deposit queries single deposit information based proposalID, depositAddr.
   * @request GET:/cosmos/gov/v1/proposals/{proposal_id}/deposits/{depositor}
   */
  cosmosGovV1Deposit = (proposalId: string, depositor: string, params: RequestParams = {}) =>
    this.request<
      { deposit?: { proposal_id?: string; depositor?: string; amount?: { denom?: string; amount?: string }[] } },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals/${proposalId}/deposits/${depositor}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1TallyResult
   * @summary TallyResult queries the tally of a proposal vote.
   * @request GET:/cosmos/gov/v1/proposals/{proposal_id}/tally
   */
  cosmosGovV1TallyResult = (proposalId: string, params: RequestParams = {}) =>
    this.request<
      { tally?: { yes_count?: string; abstain_count?: string; no_count?: string; no_with_veto_count?: string } },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals/${proposalId}/tally`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Votes
   * @summary Votes queries votes of a given proposal.
   * @request GET:/cosmos/gov/v1/proposals/{proposal_id}/votes
   */
  cosmosGovV1Votes = (
    proposalId: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        votes?: {
          proposal_id?: string;
          voter?: string;
          options?: {
            option?:
              | "VOTE_OPTION_UNSPECIFIED"
              | "VOTE_OPTION_YES"
              | "VOTE_OPTION_ABSTAIN"
              | "VOTE_OPTION_NO"
              | "VOTE_OPTION_NO_WITH_VETO";
            weight?: string;
          }[];
          metadata?: string;
        }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals/${proposalId}/votes`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosGovV1Vote
   * @summary Vote queries voted information based on proposalID, voterAddr.
   * @request GET:/cosmos/gov/v1/proposals/{proposal_id}/votes/{voter}
   */
  cosmosGovV1Vote = (proposalId: string, voter: string, params: RequestParams = {}) =>
    this.request<
      {
        vote?: {
          proposal_id?: string;
          voter?: string;
          options?: {
            option?:
              | "VOTE_OPTION_UNSPECIFIED"
              | "VOTE_OPTION_YES"
              | "VOTE_OPTION_ABSTAIN"
              | "VOTE_OPTION_NO"
              | "VOTE_OPTION_NO_WITH_VETO";
            weight?: string;
          }[];
          metadata?: string;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/gov/v1/proposals/${proposalId}/votes/${voter}`,
      method: "GET",
      ...params,
    });
}
