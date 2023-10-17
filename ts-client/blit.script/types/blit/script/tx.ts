/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { MsgRun, MsgRunResponse } from "./msgrun";
import { Params } from "./params";

export const protobufPackage = "blit.script";

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

export interface MsgCreateScript {
  creator: string;
  code: string;
  /** The list of MsgUrls the create will be granted access to via authz initially */
  msgTypePermissions: string[];
  grantee: string;
}

export interface MsgCreateScriptResponse {
}

export interface MsgUpdateScript {
  address: string;
  code: string;
  grantee: string;
}

export interface MsgUpdateScriptResponse {
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

function createBaseMsgCreateScript(): MsgCreateScript {
  return { creator: "", code: "", msgTypePermissions: [], grantee: "" };
}

export const MsgCreateScript = {
  encode(message: MsgCreateScript, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    for (const v of message.msgTypePermissions) {
      writer.uint32(26).string(v!);
    }
    if (message.grantee !== "") {
      writer.uint32(34).string(message.grantee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateScript {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateScript();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.code = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.msgTypePermissions.push(reader.string());
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

  fromJSON(object: any): MsgCreateScript {
    return {
      creator: isSet(object.creator) ? globalThis.String(object.creator) : "",
      code: isSet(object.code) ? globalThis.String(object.code) : "",
      msgTypePermissions: globalThis.Array.isArray(object?.msgTypePermissions)
        ? object.msgTypePermissions.map((e: any) => globalThis.String(e))
        : [],
      grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
    };
  },

  toJSON(message: MsgCreateScript): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.code !== "") {
      obj.code = message.code;
    }
    if (message.msgTypePermissions?.length) {
      obj.msgTypePermissions = message.msgTypePermissions;
    }
    if (message.grantee !== "") {
      obj.grantee = message.grantee;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateScript>, I>>(base?: I): MsgCreateScript {
    return MsgCreateScript.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateScript>, I>>(object: I): MsgCreateScript {
    const message = createBaseMsgCreateScript();
    message.creator = object.creator ?? "";
    message.code = object.code ?? "";
    message.msgTypePermissions = object.msgTypePermissions?.map((e) => e) || [];
    message.grantee = object.grantee ?? "";
    return message;
  },
};

function createBaseMsgCreateScriptResponse(): MsgCreateScriptResponse {
  return {};
}

export const MsgCreateScriptResponse = {
  encode(_: MsgCreateScriptResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateScriptResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateScriptResponse();
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

  fromJSON(_: any): MsgCreateScriptResponse {
    return {};
  },

  toJSON(_: MsgCreateScriptResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgCreateScriptResponse>, I>>(base?: I): MsgCreateScriptResponse {
    return MsgCreateScriptResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgCreateScriptResponse>, I>>(_: I): MsgCreateScriptResponse {
    const message = createBaseMsgCreateScriptResponse();
    return message;
  },
};

function createBaseMsgUpdateScript(): MsgUpdateScript {
  return { address: "", code: "", grantee: "" };
}

export const MsgUpdateScript = {
  encode(message: MsgUpdateScript, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.grantee !== "") {
      writer.uint32(26).string(message.grantee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateScript {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateScript();
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

          message.code = reader.string();
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

  fromJSON(object: any): MsgUpdateScript {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      code: isSet(object.code) ? globalThis.String(object.code) : "",
      grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
    };
  },

  toJSON(message: MsgUpdateScript): unknown {
    const obj: any = {};
    if (message.address !== "") {
      obj.address = message.address;
    }
    if (message.code !== "") {
      obj.code = message.code;
    }
    if (message.grantee !== "") {
      obj.grantee = message.grantee;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateScript>, I>>(base?: I): MsgUpdateScript {
    return MsgUpdateScript.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateScript>, I>>(object: I): MsgUpdateScript {
    const message = createBaseMsgUpdateScript();
    message.address = object.address ?? "";
    message.code = object.code ?? "";
    message.grantee = object.grantee ?? "";
    return message;
  },
};

function createBaseMsgUpdateScriptResponse(): MsgUpdateScriptResponse {
  return {};
}

export const MsgUpdateScriptResponse = {
  encode(_: MsgUpdateScriptResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateScriptResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateScriptResponse();
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

  fromJSON(_: any): MsgUpdateScriptResponse {
    return {};
  },

  toJSON(_: MsgUpdateScriptResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateScriptResponse>, I>>(base?: I): MsgUpdateScriptResponse {
    return MsgUpdateScriptResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateScriptResponse>, I>>(_: I): MsgUpdateScriptResponse {
    const message = createBaseMsgUpdateScriptResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** Since: cosmos-sdk 0.47 */
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
  CreateScript(request: MsgCreateScript): Promise<MsgCreateScriptResponse>;
  UpdateScript(request: MsgUpdateScript): Promise<MsgUpdateScriptResponse>;
  Run(request: MsgRun): Promise<MsgRunResponse>;
}

export const MsgServiceName = "blit.script.Msg";
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || MsgServiceName;
    this.rpc = rpc;
    this.UpdateParams = this.UpdateParams.bind(this);
    this.CreateScript = this.CreateScript.bind(this);
    this.UpdateScript = this.UpdateScript.bind(this);
    this.Run = this.Run.bind(this);
  }
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParams", data);
    return promise.then((data) => MsgUpdateParamsResponse.decode(_m0.Reader.create(data)));
  }

  CreateScript(request: MsgCreateScript): Promise<MsgCreateScriptResponse> {
    const data = MsgCreateScript.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateScript", data);
    return promise.then((data) => MsgCreateScriptResponse.decode(_m0.Reader.create(data)));
  }

  UpdateScript(request: MsgUpdateScript): Promise<MsgUpdateScriptResponse> {
    const data = MsgUpdateScript.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateScript", data);
    return promise.then((data) => MsgUpdateScriptResponse.decode(_m0.Reader.create(data)));
  }

  Run(request: MsgRun): Promise<MsgRunResponse> {
    const data = MsgRun.encode(request).finish();
    const promise = this.rpc.request(this.service, "Run", data);
    return promise.then((data) => MsgRunResponse.decode(_m0.Reader.create(data)));
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
