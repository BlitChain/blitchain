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
 * Params defines the parameters for the x/mint module.
 */
export interface CosmosMintV1Beta1Params {
  /** type of coin to mint */
  mint_denom?: string;

  /** maximum annual change in inflation rate */
  inflation_rate_change?: string;

  /** maximum inflation rate */
  inflation_max?: string;

  /** minimum inflation rate */
  inflation_min?: string;

  /** goal of percent bonded atoms */
  goal_bonded?: string;

  /**
   * expected blocks per year
   * @format uint64
   */
  blocks_per_year?: string;
}

/**
* QueryAnnualProvisionsResponse is the response type for the
Query/AnnualProvisions RPC method.
*/
export interface CosmosMintV1Beta1QueryAnnualProvisionsResponse {
  /**
   * annual_provisions is the current minting annual provisions value.
   * @format byte
   */
  annual_provisions?: string;
}

/**
* QueryInflationResponse is the response type for the Query/Inflation RPC
method.
*/
export interface CosmosMintV1Beta1QueryInflationResponse {
  /**
   * inflation is the current minting inflation value.
   * @format byte
   */
  inflation?: string;
}

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 */
export interface CosmosMintV1Beta1QueryParamsResponse {
  /** params defines the parameters of the module. */
  params?: {
    mint_denom?: string;
    inflation_rate_change?: string;
    inflation_max?: string;
    inflation_min?: string;
    goal_bonded?: string;
    blocks_per_year?: string;
  };
}

/**
* MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: cosmos-sdk 0.47
*/
export type CosmosMintV1Beta1MsgUpdateParamsResponse = object;

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
   * @name CosmosMintV1Beta1AnnualProvisions
   * @summary AnnualProvisions current minting annual provisions value.
   * @request GET:/cosmos/mint/v1beta1/annual_provisions
   */
  cosmosMintV1Beta1AnnualProvisions = (params: RequestParams = {}) =>
    this.request<{ annual_provisions?: string }, { code?: number; message?: string; details?: { "@type"?: string }[] }>(
      {
        path: `/cosmos/mint/v1beta1/annual_provisions`,
        method: "GET",
        ...params,
      },
    );

  /**
   * No description
   *
   * @tags Query
   * @name CosmosMintV1Beta1Inflation
   * @summary Inflation returns the current minting inflation value.
   * @request GET:/cosmos/mint/v1beta1/inflation
   */
  cosmosMintV1Beta1Inflation = (params: RequestParams = {}) =>
    this.request<{ inflation?: string }, { code?: number; message?: string; details?: { "@type"?: string }[] }>({
      path: `/cosmos/mint/v1beta1/inflation`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name CosmosMintV1Beta1Params
   * @summary Params returns the total set of minting parameters.
   * @request GET:/cosmos/mint/v1beta1/params
   */
  cosmosMintV1Beta1Params = (params: RequestParams = {}) =>
    this.request<
      {
        params?: {
          mint_denom?: string;
          inflation_rate_change?: string;
          inflation_max?: string;
          inflation_min?: string;
          goal_bonded?: string;
          blocks_per_year?: string;
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/mint/v1beta1/params`,
      method: "GET",
      ...params,
    });
}
