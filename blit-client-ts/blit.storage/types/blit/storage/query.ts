/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { Storage } from "./storage";

export const protobufPackage = "blit.storage";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetStorageRequest {
  address: string;
  index: string;
}

export interface QueryGetStorageResponse {
  storage: Storage | undefined;
}

export interface QueryAllStorageRequest {
  filterAddress: string;
  filterIndexPrefix: string;
  pagination: PageRequest | undefined;
}

export interface QueryAllStorageResponse {
  storage: Storage[];
  pagination: PageResponse | undefined;
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

function createBaseQueryGetStorageRequest(): QueryGetStorageRequest {
  return { address: "", index: "" };
}

export const QueryGetStorageRequest = {
  encode(message: QueryGetStorageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStorageRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStorageRequest();
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

          message.index = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetStorageRequest {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      index: isSet(object.index) ? globalThis.String(object.index) : "",
    };
  },

  toJSON(message: QueryGetStorageRequest): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.index !== "") {
      obj.index = message.index;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetStorageRequest>, I>>(base?: I): QueryGetStorageRequest {
    return QueryGetStorageRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetStorageRequest>, I>>(object: I): QueryGetStorageRequest {
    const message = createBaseQueryGetStorageRequest();
    message.address = object.address ?? "";
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetStorageResponse(): QueryGetStorageResponse {
  return { storage: undefined };
}

export const QueryGetStorageResponse = {
  encode(message: QueryGetStorageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storage !== undefined) {
      Storage.encode(message.storage, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStorageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStorageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.storage = Storage.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetStorageResponse {
    return { storage: isSet(object.storage) ? Storage.fromJSON(object.storage) : undefined };
  },

  toJSON(message: QueryGetStorageResponse): unknown {
    const obj: any = {};
    if (message.storage !== undefined) {
      obj.storage = Storage.toJSON(message.storage);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetStorageResponse>, I>>(base?: I): QueryGetStorageResponse {
    return QueryGetStorageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetStorageResponse>, I>>(object: I): QueryGetStorageResponse {
    const message = createBaseQueryGetStorageResponse();
    message.storage = (object.storage !== undefined && object.storage !== null)
      ? Storage.fromPartial(object.storage)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStorageRequest(): QueryAllStorageRequest {
  return { filterAddress: "", filterIndexPrefix: "", pagination: undefined };
}

export const QueryAllStorageRequest = {
  encode(message: QueryAllStorageRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.filterAddress !== "") {
      writer.uint32(10).string(message.filterAddress);
    }
    if (message.filterIndexPrefix !== "") {
      writer.uint32(18).string(message.filterIndexPrefix);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStorageRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStorageRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.filterAddress = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.filterIndexPrefix = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
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

  fromJSON(object: any): QueryAllStorageRequest {
    return {
      filterAddress: isSet(object.filterAddress) ? globalThis.String(object.filterAddress) : "",
      filterIndexPrefix: isSet(object.filterIndexPrefix) ? globalThis.String(object.filterIndexPrefix) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStorageRequest): unknown {
    const obj: any = {};
    if (message.filterAddress !== "") {
      obj.filterAddress = message.filterAddress;
    }
    if (message.filterIndexPrefix !== "") {
      obj.filterIndexPrefix = message.filterIndexPrefix;
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllStorageRequest>, I>>(base?: I): QueryAllStorageRequest {
    return QueryAllStorageRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllStorageRequest>, I>>(object: I): QueryAllStorageRequest {
    const message = createBaseQueryAllStorageRequest();
    message.filterAddress = object.filterAddress ?? "";
    message.filterIndexPrefix = object.filterIndexPrefix ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStorageResponse(): QueryAllStorageResponse {
  return { storage: [], pagination: undefined };
}

export const QueryAllStorageResponse = {
  encode(message: QueryAllStorageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.storage) {
      Storage.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStorageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStorageResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.storage.push(Storage.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllStorageResponse {
    return {
      storage: globalThis.Array.isArray(object?.storage) ? object.storage.map((e: any) => Storage.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStorageResponse): unknown {
    const obj: any = {};
    if (message.storage?.length) {
      obj.storage = message.storage.map((e) => Storage.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllStorageResponse>, I>>(base?: I): QueryAllStorageResponse {
    return QueryAllStorageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllStorageResponse>, I>>(object: I): QueryAllStorageResponse {
    const message = createBaseQueryAllStorageResponse();
    message.storage = object.storage?.map((e) => Storage.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Storage items. */
  Storage(request: QueryGetStorageRequest): Promise<QueryGetStorageResponse>;
  StorageAll(request: QueryAllStorageRequest): Promise<QueryAllStorageResponse>;
}

export const QueryServiceName = "blit.storage.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Storage = this.Storage.bind(this);
    this.StorageAll = this.StorageAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(_m0.Reader.create(data)));
  }

  Storage(request: QueryGetStorageRequest): Promise<QueryGetStorageResponse> {
    const data = QueryGetStorageRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Storage", data);
    return promise.then((data) => QueryGetStorageResponse.decode(_m0.Reader.create(data)));
  }

  StorageAll(request: QueryAllStorageRequest): Promise<QueryAllStorageResponse> {
    const data = QueryAllStorageRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "StorageAll", data);
    return promise.then((data) => QueryAllStorageResponse.decode(_m0.Reader.create(data)));
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
