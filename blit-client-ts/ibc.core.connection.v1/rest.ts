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
* Normally the RevisionHeight is incremented at each height while keeping
RevisionNumber the same. However some consensus algorithms may choose to
reset the height in certain conditions e.g. hard forks, state-machine
breaking changes In these cases, the RevisionNumber is incremented so that
height continues to be monitonically increasing even as the RevisionHeight
gets reset
*/
export interface IbcCoreClientV1Height {
  /**
   * the revision that the client is currently on
   * @format uint64
   */
  revision_number?: string;

  /**
   * the height within the given revision
   * @format uint64
   */
  revision_height?: string;
}

/**
* IdentifiedClientState defines a client state with an additional client
identifier field.
*/
export interface IbcCoreClientV1IdentifiedClientState {
  /** client identifier */
  client_id?: string;

  /**
   * client state
   * `Any` contains an arbitrary serialized protocol buffer message along with a
   * URL that describes the type of the serialized message.
   *
   * Protobuf library provides support to pack/unpack Any values in the form
   * of utility functions or additional generated methods of the Any type.
   * Example 1: Pack and unpack a message in C++.
   *     Foo foo = ...;
   *     Any any;
   *     any.PackFrom(foo);
   *     ...
   *     if (any.UnpackTo(&foo)) {
   *       ...
   *     }
   * Example 2: Pack and unpack a message in Java.
   *     Any any = Any.pack(foo);
   *     if (any.is(Foo.class)) {
   *       foo = any.unpack(Foo.class);
   *     // or ...
   *     if (any.isSameTypeAs(Foo.getDefaultInstance())) {
   *       foo = any.unpack(Foo.getDefaultInstance());
   * Example 3: Pack and unpack a message in Python.
   *     foo = Foo(...)
   *     any = Any()
   *     any.Pack(foo)
   *     if any.Is(Foo.DESCRIPTOR):
   *       any.Unpack(foo)
   * Example 4: Pack and unpack a message in Go
   *      foo := &pb.Foo{...}
   *      any, err := anypb.New(foo)
   *      if err != nil {
   *        ...
   *      }
   *      ...
   *      foo := &pb.Foo{}
   *      if err := any.UnmarshalTo(foo); err != nil {
   * The pack methods provided by protobuf library will by default use
   * 'type.googleapis.com/full.type.name' as the type URL and the unpack
   * methods only use the fully qualified type name after the last '/'
   * in the type URL, for example "foo.bar.com/x/y.z" will yield type
   * name "y.z".
   * JSON
   * The JSON representation of an `Any` value uses the regular
   * representation of the deserialized, embedded message, with an
   * additional field `@type` which contains the type URL. Example:
   *     package google.profile;
   *     message Person {
   *       string first_name = 1;
   *       string last_name = 2;
   *     {
   *       "@type": "type.googleapis.com/google.profile.Person",
   *       "firstName": <string>,
   *       "lastName": <string>
   * If the embedded message type is well-known and has a custom JSON
   * representation, that representation will be embedded adding a field
   * `value` which holds the custom JSON in addition to the `@type`
   * field. Example (for message [google.protobuf.Duration][]):
   *       "@type": "type.googleapis.com/google.protobuf.Duration",
   *       "value": "1.212s"
   */
  client_state?: { "@type"?: string };
}

export interface IbcCoreCommitmentV1MerklePrefix {
  /** @format byte */
  key_prefix?: string;
}

/**
* ConnectionEnd defines a stateful object on a chain connected to another
separate one.
NOTE: there must only be 2 defined ConnectionEnds to establish
a connection between two chains.
*/
export interface IbcCoreConnectionV1ConnectionEnd {
  /** client associated with this connection. */
  client_id?: string;

  /**
   * IBC version which can be utilised to determine encodings or protocols for
   * channels or packets utilising this connection.
   */
  versions?: { identifier?: string; features?: string[] }[];

  /** current state of the connection end. */
  state?: "STATE_UNINITIALIZED_UNSPECIFIED" | "STATE_INIT" | "STATE_TRYOPEN" | "STATE_OPEN";

  /** counterparty chain associated with this connection. */
  counterparty?: { client_id?: string; connection_id?: string; prefix?: { key_prefix?: string } };

  /**
   * delay period that must pass before a consensus state can be used for
   * packet-verification NOTE: delay period logic is only implemented by some
   * clients.
   * @format uint64
   */
  delay_period?: string;
}

/**
 * Counterparty defines the counterparty chain associated with a connection end.
 */
export interface IbcCoreConnectionV1Counterparty {
  /**
   * identifies the client on the counterparty chain associated with a given
   * connection.
   */
  client_id?: string;

  /**
   * identifies the connection end on the counterparty chain associated with a
   * given connection.
   */
  connection_id?: string;

  /**
   * MerklePrefix is merkle path prefixed to the key.
   * The constructed key from the Path and the key will be append(Path.KeyPath,
   * append(Path.KeyPrefix, key...))
   * commitment merkle prefix of the counterparty chain.
   */
  prefix?: { key_prefix?: string };
}

/**
* IdentifiedConnection defines a connection with additional connection
identifier field.
*/
export interface IbcCoreConnectionV1IdentifiedConnection {
  /** connection identifier. */
  id?: string;

  /** client associated with this connection. */
  client_id?: string;

  /**
   * IBC version which can be utilised to determine encodings or protocols for
   * channels or packets utilising this connection
   */
  versions?: { identifier?: string; features?: string[] }[];

  /** current state of the connection end. */
  state?: "STATE_UNINITIALIZED_UNSPECIFIED" | "STATE_INIT" | "STATE_TRYOPEN" | "STATE_OPEN";

  /** counterparty chain associated with this connection. */
  counterparty?: { client_id?: string; connection_id?: string; prefix?: { key_prefix?: string } };

  /**
   * delay period associated with this connection.
   * @format uint64
   */
  delay_period?: string;
}

/**
 * Params defines the set of Connection parameters.
 */
export interface IbcCoreConnectionV1Params {
  /**
   * maximum expected time per block (in nanoseconds), used to enforce block delay. This parameter should reflect the
   * largest amount of time that the chain might reasonably take to produce the next block under normal operating
   * conditions. A safe choice is 3-5x the expected time per block.
   * @format uint64
   */
  max_expected_time_per_block?: string;
}

export interface IbcCoreConnectionV1QueryClientConnectionsResponse {
  /** slice of all the connection paths associated with a client. */
  connection_paths?: string[];

  /**
   * merkle proof of existence
   * @format byte
   */
  proof?: string;

  /**
   * height at which the proof was generated
   * Normally the RevisionHeight is incremented at each height while keeping
   * RevisionNumber the same. However some consensus algorithms may choose to
   * reset the height in certain conditions e.g. hard forks, state-machine
   * breaking changes In these cases, the RevisionNumber is incremented so that
   * height continues to be monitonically increasing even as the RevisionHeight
   * gets reset
   */
  proof_height?: { revision_number?: string; revision_height?: string };
}

export interface IbcCoreConnectionV1QueryConnectionClientStateResponse {
  /**
   * client state associated with the channel
   * IdentifiedClientState defines a client state with an additional client
   * identifier field.
   */
  identified_client_state?: { client_id?: string; client_state?: { "@type"?: string } };

  /**
   * merkle proof of existence
   * @format byte
   */
  proof?: string;

  /**
   * height at which the proof was retrieved
   * Normally the RevisionHeight is incremented at each height while keeping
   * RevisionNumber the same. However some consensus algorithms may choose to
   * reset the height in certain conditions e.g. hard forks, state-machine
   * breaking changes In these cases, the RevisionNumber is incremented so that
   * height continues to be monitonically increasing even as the RevisionHeight
   * gets reset
   */
  proof_height?: { revision_number?: string; revision_height?: string };
}

export interface IbcCoreConnectionV1QueryConnectionConsensusStateResponse {
  /**
   * consensus state associated with the channel
   * `Any` contains an arbitrary serialized protocol buffer message along with a
   * URL that describes the type of the serialized message.
   *
   * Protobuf library provides support to pack/unpack Any values in the form
   * of utility functions or additional generated methods of the Any type.
   * Example 1: Pack and unpack a message in C++.
   *     Foo foo = ...;
   *     Any any;
   *     any.PackFrom(foo);
   *     ...
   *     if (any.UnpackTo(&foo)) {
   *       ...
   *     }
   * Example 2: Pack and unpack a message in Java.
   *     Any any = Any.pack(foo);
   *     if (any.is(Foo.class)) {
   *       foo = any.unpack(Foo.class);
   *     // or ...
   *     if (any.isSameTypeAs(Foo.getDefaultInstance())) {
   *       foo = any.unpack(Foo.getDefaultInstance());
   * Example 3: Pack and unpack a message in Python.
   *     foo = Foo(...)
   *     any = Any()
   *     any.Pack(foo)
   *     if any.Is(Foo.DESCRIPTOR):
   *       any.Unpack(foo)
   * Example 4: Pack and unpack a message in Go
   *      foo := &pb.Foo{...}
   *      any, err := anypb.New(foo)
   *      if err != nil {
   *        ...
   *      }
   *      ...
   *      foo := &pb.Foo{}
   *      if err := any.UnmarshalTo(foo); err != nil {
   * The pack methods provided by protobuf library will by default use
   * 'type.googleapis.com/full.type.name' as the type URL and the unpack
   * methods only use the fully qualified type name after the last '/'
   * in the type URL, for example "foo.bar.com/x/y.z" will yield type
   * name "y.z".
   * JSON
   * The JSON representation of an `Any` value uses the regular
   * representation of the deserialized, embedded message, with an
   * additional field `@type` which contains the type URL. Example:
   *     package google.profile;
   *     message Person {
   *       string first_name = 1;
   *       string last_name = 2;
   *     {
   *       "@type": "type.googleapis.com/google.profile.Person",
   *       "firstName": <string>,
   *       "lastName": <string>
   * If the embedded message type is well-known and has a custom JSON
   * representation, that representation will be embedded adding a field
   * `value` which holds the custom JSON in addition to the `@type`
   * field. Example (for message [google.protobuf.Duration][]):
   *       "@type": "type.googleapis.com/google.protobuf.Duration",
   *       "value": "1.212s"
   */
  consensus_state?: { "@type"?: string };

  /** client ID associated with the consensus state */
  client_id?: string;

  /**
   * merkle proof of existence
   * @format byte
   */
  proof?: string;

  /**
   * height at which the proof was retrieved
   * Normally the RevisionHeight is incremented at each height while keeping
   * RevisionNumber the same. However some consensus algorithms may choose to
   * reset the height in certain conditions e.g. hard forks, state-machine
   * breaking changes In these cases, the RevisionNumber is incremented so that
   * height continues to be monitonically increasing even as the RevisionHeight
   * gets reset
   */
  proof_height?: { revision_number?: string; revision_height?: string };
}

/**
 * QueryConnectionParamsResponse is the response type for the Query/ConnectionParams RPC method.
 */
export interface IbcCoreConnectionV1QueryConnectionParamsResponse {
  /** params defines the parameters of the module. */
  params?: { max_expected_time_per_block?: string };
}

/**
* QueryConnectionResponse is the response type for the Query/Connection RPC
method. Besides the connection end, it includes a proof and the height from
which the proof was retrieved.
*/
export interface IbcCoreConnectionV1QueryConnectionResponse {
  /**
   * connection associated with the request identifier
   * ConnectionEnd defines a stateful object on a chain connected to another
   * separate one.
   * NOTE: there must only be 2 defined ConnectionEnds to establish
   * a connection between two chains.
   */
  connection?: {
    client_id?: string;
    versions?: { identifier?: string; features?: string[] }[];
    state?: "STATE_UNINITIALIZED_UNSPECIFIED" | "STATE_INIT" | "STATE_TRYOPEN" | "STATE_OPEN";
    counterparty?: { client_id?: string; connection_id?: string; prefix?: { key_prefix?: string } };
    delay_period?: string;
  };

  /**
   * merkle proof of existence
   * @format byte
   */
  proof?: string;

  /**
   * height at which the proof was retrieved
   * Normally the RevisionHeight is incremented at each height while keeping
   * RevisionNumber the same. However some consensus algorithms may choose to
   * reset the height in certain conditions e.g. hard forks, state-machine
   * breaking changes In these cases, the RevisionNumber is incremented so that
   * height continues to be monitonically increasing even as the RevisionHeight
   * gets reset
   */
  proof_height?: { revision_number?: string; revision_height?: string };
}

/**
* QueryConnectionsResponse is the response type for the Query/Connections RPC
method.
*/
export interface IbcCoreConnectionV1QueryConnectionsResponse {
  /** list of stored connections of the chain. */
  connections?: {
    id?: string;
    client_id?: string;
    versions?: { identifier?: string; features?: string[] }[];
    state?: "STATE_UNINITIALIZED_UNSPECIFIED" | "STATE_INIT" | "STATE_TRYOPEN" | "STATE_OPEN";
    counterparty?: { client_id?: string; connection_id?: string; prefix?: { key_prefix?: string } };
    delay_period?: string;
  }[];

  /**
   * pagination response
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: { next_key?: string; total?: string };

  /**
   * query block height
   * Normally the RevisionHeight is incremented at each height while keeping
   * RevisionNumber the same. However some consensus algorithms may choose to
   * reset the height in certain conditions e.g. hard forks, state-machine
   * breaking changes In these cases, the RevisionNumber is incremented so that
   * height continues to be monitonically increasing even as the RevisionHeight
   * gets reset
   */
  height?: { revision_number?: string; revision_height?: string };
}

/**
* State defines if a connection is in one of the following states:
INIT, TRYOPEN, OPEN or UNINITIALIZED.

 - STATE_UNINITIALIZED_UNSPECIFIED: Default State
 - STATE_INIT: A connection end has just started the opening handshake.
 - STATE_TRYOPEN: A connection end has acknowledged the handshake step on the counterparty
chain.
 - STATE_OPEN: A connection end has completed the handshake.
*/
export enum IbcCoreConnectionV1State {
  STATE_UNINITIALIZED_UNSPECIFIED = "STATE_UNINITIALIZED_UNSPECIFIED",
  STATE_INIT = "STATE_INIT",
  STATE_TRYOPEN = "STATE_TRYOPEN",
  STATE_OPEN = "STATE_OPEN",
}

/**
* Version defines the versioning scheme used to negotiate the IBC verison in
the connection handshake.
*/
export interface IbcCoreConnectionV1Version {
  /** unique version identifier */
  identifier?: string;

  /** list of features compatible with the specified identifier */
  features?: string[];
}

/**
 * MsgConnectionOpenAckResponse defines the Msg/ConnectionOpenAck response type.
 */
export type IbcCoreConnectionV1MsgConnectionOpenAckResponse = object;

/**
* MsgConnectionOpenConfirmResponse defines the Msg/ConnectionOpenConfirm
response type.
*/
export type IbcCoreConnectionV1MsgConnectionOpenConfirmResponse = object;

/**
* MsgConnectionOpenInitResponse defines the Msg/ConnectionOpenInit response
type.
*/
export type IbcCoreConnectionV1MsgConnectionOpenInitResponse = object;

/**
 * MsgConnectionOpenTryResponse defines the Msg/ConnectionOpenTry response type.
 */
export type IbcCoreConnectionV1MsgConnectionOpenTryResponse = object;

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
 * @name IbcCoreConnectionV1ClientConnections
 * @summary ClientConnections queries the connection paths associated with a client
state.
 * @request GET:/ibc/core/connection/v1/client_connections/{client_id}
 */
  ibcCoreConnectionV1ClientConnections = (clientId: string, params: RequestParams = {}) =>
    this.request<
      {
        connection_paths?: string[];
        proof?: string;
        proof_height?: { revision_number?: string; revision_height?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/ibc/core/connection/v1/client_connections/${clientId}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name IbcCoreConnectionV1Connections
   * @summary Connections queries all the IBC connections of a chain.
   * @request GET:/ibc/core/connection/v1/connections
   */
  ibcCoreConnectionV1Connections = (
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
        connections?: {
          id?: string;
          client_id?: string;
          versions?: { identifier?: string; features?: string[] }[];
          state?: "STATE_UNINITIALIZED_UNSPECIFIED" | "STATE_INIT" | "STATE_TRYOPEN" | "STATE_OPEN";
          counterparty?: { client_id?: string; connection_id?: string; prefix?: { key_prefix?: string } };
          delay_period?: string;
        }[];
        pagination?: { next_key?: string; total?: string };
        height?: { revision_number?: string; revision_height?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/ibc/core/connection/v1/connections`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name IbcCoreConnectionV1Connection
   * @summary Connection queries an IBC connection end.
   * @request GET:/ibc/core/connection/v1/connections/{connection_id}
   */
  ibcCoreConnectionV1Connection = (connectionId: string, params: RequestParams = {}) =>
    this.request<
      {
        connection?: {
          client_id?: string;
          versions?: { identifier?: string; features?: string[] }[];
          state?: "STATE_UNINITIALIZED_UNSPECIFIED" | "STATE_INIT" | "STATE_TRYOPEN" | "STATE_OPEN";
          counterparty?: { client_id?: string; connection_id?: string; prefix?: { key_prefix?: string } };
          delay_period?: string;
        };
        proof?: string;
        proof_height?: { revision_number?: string; revision_height?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/ibc/core/connection/v1/connections/${connectionId}`,
      method: "GET",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name IbcCoreConnectionV1ConnectionClientState
 * @summary ConnectionClientState queries the client state associated with the
connection.
 * @request GET:/ibc/core/connection/v1/connections/{connection_id}/client_state
 */
  ibcCoreConnectionV1ConnectionClientState = (connectionId: string, params: RequestParams = {}) =>
    this.request<
      {
        identified_client_state?: { client_id?: string; client_state?: { "@type"?: string } };
        proof?: string;
        proof_height?: { revision_number?: string; revision_height?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/ibc/core/connection/v1/connections/${connectionId}/client_state`,
      method: "GET",
      ...params,
    });

  /**
 * No description
 * 
 * @tags Query
 * @name IbcCoreConnectionV1ConnectionConsensusState
 * @summary ConnectionConsensusState queries the consensus state associated with the
connection.
 * @request GET:/ibc/core/connection/v1/connections/{connection_id}/consensus_state/revision/{revision_number}/height/{revision_height}
 */
  ibcCoreConnectionV1ConnectionConsensusState = (
    connectionId: string,
    revisionNumber: string,
    revisionHeight: string,
    params: RequestParams = {},
  ) =>
    this.request<
      {
        consensus_state?: { "@type"?: string };
        client_id?: string;
        proof?: string;
        proof_height?: { revision_number?: string; revision_height?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/ibc/core/connection/v1/connections/${connectionId}/consensus_state/revision/${revisionNumber}/height/${revisionHeight}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name IbcCoreConnectionV1ConnectionParams
   * @summary ConnectionParams queries all parameters of the ibc connection submodule.
   * @request GET:/ibc/core/connection/v1/params
   */
  ibcCoreConnectionV1ConnectionParams = (params: RequestParams = {}) =>
    this.request<
      { params?: { max_expected_time_per_block?: string } },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/ibc/core/connection/v1/params`,
      method: "GET",
      ...params,
    });
}
