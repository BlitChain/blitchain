/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Params } from "./params";

export const protobufPackage = "blit.storage";

/**
 * MsgUpdateParams is the Msg/UpdateParams request type.
 *
 * Since: cosmos-sdk 0.47
 */
export interface MsgUpdateParams {
  /** authority is the address that controls the module (defaults to x/gov unless overwritten). */
  authority: string;
  /** NOTE: All parameters must be supplied. */
  params: Params | undefined;
}

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 *
 * Since: cosmos-sdk 0.47
 */
export interface MsgUpdateParamsResponse {
}

export interface MsgCreateStorage {
  address: string;
  index: string;
  data: string;
  grantee: string;
}

export interface MsgCreateStorageResponse {
}

export interface MsgUpdateStorage {
  address: string;
  index: string;
  data: string;
  grantee: string;
}

export interface MsgUpdateStorageResponse {
}

export interface MsgDeleteStorage {
  address: string;
  index: string;
  grantee: string;
}

export interface MsgDeleteStorageResponse {
}

function createBaseMsgUpdateParams(): MsgUpdateParams {
  return { authority: "", params: undefined };
}

export const MsgUpdateParams = {
  encode(message: MsgUpdateParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParams {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.authority = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
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

  fromJSON(object: any): MsgUpdateParams {
    return {
      authority: isSet(object.authority) ? globalThis.String(object.authority) : "",
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
    };
  },

  toJSON(message: MsgUpdateParams): unknown {
    const obj: any = {};
    if (message.authority !== "") {
      obj.authority = message.authority;
    }
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(base?: I): MsgUpdateParams {
    return MsgUpdateParams.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(object: I): MsgUpdateParams {
    const message = createBaseMsgUpdateParams();
    message.authority = object.authority ?? "";
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateParamsResponse(): MsgUpdateParamsResponse {
  return {};
}

export const MsgUpdateParamsResponse = {
  encode(_: MsgUpdateParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParamsResponse();
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

  fromJSON(_: any): MsgUpdateParamsResponse {
    return {};
  },

  toJSON(_: MsgUpdateParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(base?: I): MsgUpdateParamsResponse {
    return MsgUpdateParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(_: I): MsgUpdateParamsResponse {
    const message = createBaseMsgUpdateParamsResponse();
    return message;
  },
};

function createBaseMsgCreateStorage(): MsgCreateStorage {
  return { address: "", index: "", data: "", grantee: "" };
}

export const MsgCreateStorage = {
  encode(message: MsgCreateStorage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.data !== "") {
      writer.uint32(26).string(message.data);
    }
    if (message.grantee !== "") {
      writer.uint32(34).string(message.grantee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateStorage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateStorage();
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
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.grantee = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCreateStorage {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      index: isSet(object.index) ? globalThis.String(object.index) : "",
      data: isSet(object.data) ? globalThis.String(object.data) : "",
      grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
    };
  },

  toJSON(message: MsgCreateStorage): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.index !== "") {
      obj.index = message.index;
    }
    if (message.data !== "") {
      obj.data = message.data;
    }
    if (message.grantee !== "") {
      obj.grantee = message.grantee;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateStorage>, I>>(base?: I): MsgCreateStorage {
    return MsgCreateStorage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateStorage>, I>>(object: I): MsgCreateStorage {
    const message = createBaseMsgCreateStorage();
    message.address = object.address ?? "";
    message.index = object.index ?? "";
    message.data = object.data ?? "";
    message.grantee = object.grantee ?? "";
    return message;
  },
};

function createBaseMsgCreateStorageResponse(): MsgCreateStorageResponse {
  return {};
}

export const MsgCreateStorageResponse = {
  encode(_: MsgCreateStorageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateStorageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateStorageResponse();
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

  fromJSON(_: any): MsgCreateStorageResponse {
    return {};
  },

  toJSON(_: MsgCreateStorageResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateStorageResponse>, I>>(base?: I): MsgCreateStorageResponse {
    return MsgCreateStorageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateStorageResponse>, I>>(_: I): MsgCreateStorageResponse {
    const message = createBaseMsgCreateStorageResponse();
    return message;
  },
};

function createBaseMsgUpdateStorage(): MsgUpdateStorage {
  return { address: "", index: "", data: "", grantee: "" };
}

export const MsgUpdateStorage = {
  encode(message: MsgUpdateStorage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.data !== "") {
      writer.uint32(26).string(message.data);
    }
    if (message.grantee !== "") {
      writer.uint32(34).string(message.grantee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateStorage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateStorage();
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
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.grantee = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateStorage {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      index: isSet(object.index) ? globalThis.String(object.index) : "",
      data: isSet(object.data) ? globalThis.String(object.data) : "",
      grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
    };
  },

  toJSON(message: MsgUpdateStorage): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.index !== "") {
      obj.index = message.index;
    }
    if (message.data !== "") {
      obj.data = message.data;
    }
    if (message.grantee !== "") {
      obj.grantee = message.grantee;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateStorage>, I>>(base?: I): MsgUpdateStorage {
    return MsgUpdateStorage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateStorage>, I>>(object: I): MsgUpdateStorage {
    const message = createBaseMsgUpdateStorage();
    message.address = object.address ?? "";
    message.index = object.index ?? "";
    message.data = object.data ?? "";
    message.grantee = object.grantee ?? "";
    return message;
  },
};

function createBaseMsgUpdateStorageResponse(): MsgUpdateStorageResponse {
  return {};
}

export const MsgUpdateStorageResponse = {
  encode(_: MsgUpdateStorageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateStorageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateStorageResponse();
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

  fromJSON(_: any): MsgUpdateStorageResponse {
    return {};
  },

  toJSON(_: MsgUpdateStorageResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateStorageResponse>, I>>(base?: I): MsgUpdateStorageResponse {
    return MsgUpdateStorageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateStorageResponse>, I>>(_: I): MsgUpdateStorageResponse {
    const message = createBaseMsgUpdateStorageResponse();
    return message;
  },
};

function createBaseMsgDeleteStorage(): MsgDeleteStorage {
  return { address: "", index: "", grantee: "" };
}

export const MsgDeleteStorage = {
  encode(message: MsgDeleteStorage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.grantee !== "") {
      writer.uint32(26).string(message.grantee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteStorage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteStorage();
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
        case 3:
          if (tag !== 26) {
            break;
          }

          message.grantee = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteStorage {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      index: isSet(object.index) ? globalThis.String(object.index) : "",
      grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
    };
  },

  toJSON(message: MsgDeleteStorage): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.index !== "") {
      obj.index = message.index;
    }
    if (message.grantee !== "") {
      obj.grantee = message.grantee;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteStorage>, I>>(base?: I): MsgDeleteStorage {
    return MsgDeleteStorage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteStorage>, I>>(object: I): MsgDeleteStorage {
    const message = createBaseMsgDeleteStorage();
    message.address = object.address ?? "";
    message.index = object.index ?? "";
    message.grantee = object.grantee ?? "";
    return message;
  },
};

function createBaseMsgDeleteStorageResponse(): MsgDeleteStorageResponse {
  return {};
}

export const MsgDeleteStorageResponse = {
  encode(_: MsgDeleteStorageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteStorageResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteStorageResponse();
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

  fromJSON(_: any): MsgDeleteStorageResponse {
    return {};
  },

  toJSON(_: MsgDeleteStorageResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgDeleteStorageResponse>, I>>(base?: I): MsgDeleteStorageResponse {
    return MsgDeleteStorageResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgDeleteStorageResponse>, I>>(_: I): MsgDeleteStorageResponse {
    const message = createBaseMsgDeleteStorageResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** Since: cosmos-sdk 0.47 */
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
  CreateStorage(request: MsgCreateStorage): Promise<MsgCreateStorageResponse>;
  UpdateStorage(request: MsgUpdateStorage): Promise<MsgUpdateStorageResponse>;
  DeleteStorage(request: MsgDeleteStorage): Promise<MsgDeleteStorageResponse>;
}

export const MsgServiceName = "blit.storage.Msg";
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || MsgServiceName;
    this.rpc = rpc;
    this.UpdateParams = this.UpdateParams.bind(this);
    this.CreateStorage = this.CreateStorage.bind(this);
    this.UpdateStorage = this.UpdateStorage.bind(this);
    this.DeleteStorage = this.DeleteStorage.bind(this);
  }
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParams", data);
    return promise.then((data) => MsgUpdateParamsResponse.decode(_m0.Reader.create(data)));
  }

  CreateStorage(request: MsgCreateStorage): Promise<MsgCreateStorageResponse> {
    const data = MsgCreateStorage.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateStorage", data);
    return promise.then((data) => MsgCreateStorageResponse.decode(_m0.Reader.create(data)));
  }

  UpdateStorage(request: MsgUpdateStorage): Promise<MsgUpdateStorageResponse> {
    const data = MsgUpdateStorage.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateStorage", data);
    return promise.then((data) => MsgUpdateStorageResponse.decode(_m0.Reader.create(data)));
  }

  DeleteStorage(request: MsgDeleteStorage): Promise<MsgDeleteStorageResponse> {
    const data = MsgDeleteStorage.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteStorage", data);
    return promise.then((data) => MsgDeleteStorageResponse.decode(_m0.Reader.create(data)));
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
