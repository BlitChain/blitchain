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
* ParamChange defines an individual parameter change, for use in
ParameterChangeProposal.
*/
export interface CosmosParamsV1Beta1ParamChange {
  subspace?: string;
  key?: string;
  value?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CosmosParamsV1Beta1QueryParamsResponse {
  /** param defines the queried parameter. */
  param?: { subspace?: string; key?: string; value?: string };
}

/**
* QuerySubspacesResponse defines the response types for querying for all
registered subspaces and all keys for a subspace.

Since: cosmos-sdk 0.46
*/
export interface CosmosParamsV1Beta1QuerySubspacesResponse {
  subspaces?: { subspace?: string; keys?: string[] }[];
}

/**
* Subspace defines a parameter subspace name and all the keys that exist for
the subspace.

Since: cosmos-sdk 0.46
*/
export interface CosmosParamsV1Beta1Subspace {
  subspace?: string;
  keys?: string[];
}

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
 * @name CosmosParamsV1Beta1Params
 * @summary Params queries a specific parameter of a module, given its subspace and
key.
 * @request GET:/cosmos/params/v1beta1/params
 */
  cosmosParamsV1Beta1Params = (query?: { subspace?: string; key?: string }, params: RequestParams = {}) =>
    this.request<
      { param?: { subspace?: string; key?: string; value?: string } },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/params/v1beta1/params`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * @description Since: cosmos-sdk 0.46
   *
   * @tags Query
   * @name CosmosParamsV1Beta1Subspaces
   * @summary Subspaces queries for all registered subspaces and all keys for a subspace.
   * @request GET:/cosmos/params/v1beta1/subspaces
   */
  cosmosParamsV1Beta1Subspaces = (params: RequestParams = {}) =>
    this.request<
      { subspaces?: { subspace?: string; keys?: string[] }[] },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/cosmos/params/v1beta1/subspaces`,
      method: "GET",
      ...params,
    });
}
