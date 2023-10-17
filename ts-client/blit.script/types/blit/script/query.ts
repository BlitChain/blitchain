/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { MsgRun, MsgRunResponse } from "./msgrun";
import { Params } from "./params";
import { Script } from "./script";

export const protobufPackage = "blit.script";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetScriptRequest {
  address: string;
}

export interface QueryGetScriptResponse {
  script: Script | undefined;
}

export interface QueryAllScriptRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllScriptResponse {
  script: Script[];
  pagination: PageResponse | undefined;
}

export interface QueryWebRequest {
  address: string;
  httprequest: string;
}

export interface QueryWebResponse {
  httpresponse: string;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(base?: I): QueryParamsRequest {
    return QueryParamsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(base?: I): QueryParamsResponse {
    return QueryParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetScriptRequest(): QueryGetScriptRequest {
  return { address: "" };
}

export const QueryGetScriptRequest = {
  encode(message: QueryGetScriptRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetScriptRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetScriptRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.address = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetScriptRequest {
    return { address: isSet(object.address) ? globalThis.String(object.address) : "" };
  },

  toJSON(message: QueryGetScriptRequest): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetScriptRequest>, I>>(base?: I): QueryGetScriptRequest {
    return QueryGetScriptRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetScriptRequest>, I>>(object: I): QueryGetScriptRequest {
    const message = createBaseQueryGetScriptRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetScriptResponse(): QueryGetScriptResponse {
  return { script: undefined };
}

export const QueryGetScriptResponse = {
  encode(message: QueryGetScriptResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.script !== undefined) {
      Script.encode(message.script, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetScriptResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetScriptResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.script = Script.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetScriptResponse {
    return { script: isSet(object.script) ? Script.fromJSON(object.script) : undefined };
  },

  toJSON(message: QueryGetScriptResponse): unknown {
    const obj: any = {};
    if (message.script !== undefined) {
      obj.script = Script.toJSON(message.script);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetScriptResponse>, I>>(base?: I): QueryGetScriptResponse {
    return QueryGetScriptResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetScriptResponse>, I>>(object: I): QueryGetScriptResponse {
    const message = createBaseQueryGetScriptResponse();
    message.script = (object.script !== undefined && object.script !== null)
      ? Script.fromPartial(object.script)
      : undefined;
    return message;
  },
};

function createBaseQueryAllScriptRequest(): QueryAllScriptRequest {
  return { pagination: undefined };
}

export const QueryAllScriptRequest = {
  encode(message: QueryAllScriptRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllScriptRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllScriptRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllScriptRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllScriptRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllScriptRequest>, I>>(base?: I): QueryAllScriptRequest {
    return QueryAllScriptRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllScriptRequest>, I>>(object: I): QueryAllScriptRequest {
    const message = createBaseQueryAllScriptRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllScriptResponse(): QueryAllScriptResponse {
  return { script: [], pagination: undefined };
}

export const QueryAllScriptResponse = {
  encode(message: QueryAllScriptResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.script) {
      Script.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllScriptResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllScriptResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.script.push(Script.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllScriptResponse {
    return {
      script: globalThis.Array.isArray(object?.script) ? object.script.map((e: any) => Script.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllScriptResponse): unknown {
    const obj: any = {};
    if (message.script?.length) {
      obj.script = message.script.map((e) => Script.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllScriptResponse>, I>>(base?: I): QueryAllScriptResponse {
    return QueryAllScriptResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllScriptResponse>, I>>(object: I): QueryAllScriptResponse {
    const message = createBaseQueryAllScriptResponse();
    message.script = object.script?.map((e) => Script.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryWebRequest(): QueryWebRequest {
  return { address: "", httprequest: "" };
}

export const QueryWebRequest = {
  encode(message: QueryWebRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.httprequest !== "") {
      writer.uint32(18).string(message.httprequest);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryWebRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryWebRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.address = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.httprequest = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryWebRequest {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      httprequest: isSet(object.httprequest) ? globalThis.String(object.httprequest) : "",
    };
  },

  toJSON(message: QueryWebRequest): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.httprequest !== "") {
      obj.httprequest = message.httprequest;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryWebRequest>, I>>(base?: I): QueryWebRequest {
    return QueryWebRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryWebRequest>, I>>(object: I): QueryWebRequest {
    const message = createBaseQueryWebRequest();
    message.address = object.address ?? "";
    message.httprequest = object.httprequest ?? "";
    return message;
  },
};

function createBaseQueryWebResponse(): QueryWebResponse {
  return { httpresponse: "" };
}

export const QueryWebResponse = {
  encode(message: QueryWebResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.httpresponse !== "") {
      writer.uint32(10).string(message.httpresponse);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryWebResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryWebResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.httpresponse = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryWebResponse {
    return { httpresponse: isSet(object.httpresponse) ? globalThis.String(object.httpresponse) : "" };
  },

  toJSON(message: QueryWebResponse): unknown {
    const obj: any = {};
    if (message.httpresponse !== "") {
      obj.httpresponse = message.httpresponse;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryWebResponse>, I>>(base?: I): QueryWebResponse {
    return QueryWebResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryWebResponse>, I>>(object: I): QueryWebResponse {
    const message = createBaseQueryWebResponse();
    message.httpresponse = object.httpresponse ?? "";
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Script items. */
  Script(request: QueryGetScriptRequest): Promise<QueryGetScriptResponse>;
  ScriptAll(request: QueryAllScriptRequest): Promise<QueryAllScriptResponse>;
  /** Runs the function and returns the result. */
  Eval(request: MsgRun): Promise<MsgRunResponse>;
  /** Queries the WSGI web application function of a script. */
  Web(request: QueryWebRequest): Promise<QueryWebResponse>;
}

export const QueryServiceName = "blit.script.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Script = this.Script.bind(this);
    this.ScriptAll = this.ScriptAll.bind(this);
    this.Eval = this.Eval.bind(this);
    this.Web = this.Web.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(_m0.Reader.create(data)));
  }

  Script(request: QueryGetScriptRequest): Promise<QueryGetScriptResponse> {
    const data = QueryGetScriptRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Script", data);
    return promise.then((data) => QueryGetScriptResponse.decode(_m0.Reader.create(data)));
  }

  ScriptAll(request: QueryAllScriptRequest): Promise<QueryAllScriptResponse> {
    const data = QueryAllScriptRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ScriptAll", data);
    return promise.then((data) => QueryAllScriptResponse.decode(_m0.Reader.create(data)));
  }

  Eval(request: MsgRun): Promise<MsgRunResponse> {
    const data = MsgRun.encode(request).finish();
    const promise = this.rpc.request(this.service, "Eval", data);
    return promise.then((data) => MsgRunResponse.decode(_m0.Reader.create(data)));
  }

  Web(request: QueryWebRequest): Promise<QueryWebResponse> {
    const data = QueryWebRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Web", data);
    return promise.then((data) => QueryWebResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
