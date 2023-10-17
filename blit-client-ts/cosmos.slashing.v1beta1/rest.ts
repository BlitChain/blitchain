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

export interface GoogleProtobufAny {
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
 * Params represents the parameters used for by the slashing module.
 */
export interface CosmosSlashingV1Beta1Params {
  /** @format int64 */
  signed_blocks_window?: string;

  /** @format byte */
  min_signed_per_window?: string;
  downtime_jail_duration?: string;

  /** @format byte */
  slash_fraction_double_sign?: string;

  /** @format byte */
  slash_fraction_downtime?: string;
}

export interface CosmosSlashingV1Beta1QueryParamsResponse {
  /** Params represents the parameters used for by the slashing module. */
  params?: {
    signed_blocks_window?: string;
    min_signed_per_window?: string;
    downtime_jail_duration?: string;
    slash_fraction_double_sign?: string;
    slash_fraction_downtime?: string;
  };
}

export interface CosmosSlashingV1Beta1QuerySigningInfoResponse {
  /**
   * val_signing_info is the signing info of requested val cons address
   * ValidatorSigningInfo defines a validator's signing info for monitoring their
   * liveness activity.
   */
  val_signing_info?: {
    address?: string;
    start_height?: string;
    index_offset?: string;
    jailed_until?: string;
    tombstoned?: boolean;
    missed_blocks_counter?: string;
  };
}

export interface CosmosSlashingV1Beta1QuerySigningInfosResponse {
  /** info is the signing info of all validators */
  info?: {
    address?: string;
    start_height?: string;
    index_offset?: string;
    jailed_until?: string;
    tombstoned?: boolean;
    missed_blocks_counter?: string;
  }[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: { next_key?: string; total?: string };
}

/**
* ValidatorSigningInfo defines a validator's signing info for monitoring their
liveness activity.
*/
export interface CosmosSlashingV1Beta1ValidatorSigningInfo {
  address?: string;

  /**
   * Height at which validator was first a candidate OR was unjailed
   * @format int64
   */
  start_height?: string;

  /**
   * Index which is incremented each time the validator was a bonded
   * in a block and may have signed a precommit or not. This in conjunction with the
   * `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
   * @format int64
   */
  index_offset?: string;

  /**
   * Timestamp until which the validator is jailed due to liveness downtime.
   * @format date-time
   */
  jailed_until?: string;

  /**
   * Whether or not a validator has been tombstoned (killed out of validator set). It is set
   * once the validator commits an equivocation or for any other configured misbehiavor.
   */
  tombstoned?: boolean;

  /**
   * A counter kept to avoid unnecessary array reads.
   * Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
   * @format int64
   */
  missed_blocks_counter?: string;
}

export type CosmosSlashingV1Beta1MsgUnjailResponse = object;

/**
* MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: cosmos-sdk 0.47
*/
export type CosmosSlashingV1Beta1MsgUpdateParamsResponse = object;

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
   * @name CosmosSlashingV1Beta1Params
   * @summary Params queries the parameters of slashing module
   * @request GET:/cosmos/slashing/v1beta1/params
   */
  cosmosSlashingV1Beta1Params = (params: RequestParams = {}) =>
    this.request<
      {
        params?: {
          signed_blocks_window?: string;
          min_signed_per_window?: string;
          downtime_jail_duration?: string;
          slash_fraction_double_sign?: string;
          slash_fraction_downtime?: string;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/slashing/v1beta1/params`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosSlashingV1Beta1SigningInfos
   * @summary SigningInfos queries signing info of all validators
   * @request GET:/cosmos/slashing/v1beta1/signing_infos
   */
  cosmosSlashingV1Beta1SigningInfos = (
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
        info?: {
          address?: string;
          start_height?: string;
          index_offset?: string;
          jailed_until?: string;
          tombstoned?: boolean;
          missed_blocks_counter?: string;
        }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/slashing/v1beta1/signing_infos`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosSlashingV1Beta1SigningInfo
   * @summary SigningInfo queries the signing info of given cons address
   * @request GET:/cosmos/slashing/v1beta1/signing_infos/{cons_address}
   */
  cosmosSlashingV1Beta1SigningInfo = (consAddress: string, params: RequestParams = {}) =>
    this.request<
      {
        val_signing_info?: {
          address?: string;
          start_height?: string;
          index_offset?: string;
          jailed_until?: string;
          tombstoned?: boolean;
          missed_blocks_counter?: string;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/slashing/v1beta1/signing_infos/${consAddress}`,
      method: "GET",
      ...params,
    });
}
