/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "blit.script";

/** MsgRun runs a script at a specific address */
export interface MsgRun {
  callerAddress: string;
  scriptAddress: string;
  extraCode: string;
  functionName: string;
  kwargs: string;
  grantee: string;
}

export interface MsgRunResponse {
  response: string;
}

function createBaseMsgRun(): MsgRun {
  return { callerAddress: "", scriptAddress: "", extraCode: "", functionName: "", kwargs: "", grantee: "" };
}

export const MsgRun = {
  encode(message: MsgRun, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.callerAddress !== "") {
      writer.uint32(18).string(message.callerAddress);
    }
    if (message.scriptAddress !== "") {
      writer.uint32(26).string(message.scriptAddress);
    }
    if (message.extraCode !== "") {
      writer.uint32(34).string(message.extraCode);
    }
    if (message.functionName !== "") {
      writer.uint32(42).string(message.functionName);
    }
    if (message.kwargs !== "") {
      writer.uint32(50).string(message.kwargs);
    }
    if (message.grantee !== "") {
      writer.uint32(58).string(message.grantee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRun {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRun();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.callerAddress = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.scriptAddress = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.extraCode = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.functionName = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.kwargs = reader.string();
          continue;
        case 7:
          if (tag !== 58) {
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

  fromJSON(object: any): MsgRun {
    return {
      callerAddress: isSet(object.callerAddress) ? globalThis.String(object.callerAddress) : "",
      scriptAddress: isSet(object.scriptAddress) ? globalThis.String(object.scriptAddress) : "",
      extraCode: isSet(object.extraCode) ? globalThis.String(object.extraCode) : "",
      functionName: isSet(object.functionName) ? globalThis.String(object.functionName) : "",
      kwargs: isSet(object.kwargs) ? globalThis.String(object.kwargs) : "",
      grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
    };
  },

  toJSON(message: MsgRun): unknown {
    const obj: any = {};
    if (message.callerAddress !== "") {
      obj.callerAddress = message.callerAddress;
    }
    if (message.scriptAddress !== "") {
      obj.scriptAddress = message.scriptAddress;
    }
    if (message.extraCode !== "") {
      obj.extraCode = message.extraCode;
    }
    if (message.functionName !== "") {
      obj.functionName = message.functionName;
    }
    if (message.kwargs !== "") {
      obj.kwargs = message.kwargs;
    }
    if (message.grantee !== "") {
      obj.grantee = message.grantee;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgRun>, I>>(base?: I): MsgRun {
    return MsgRun.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgRun>, I>>(object: I): MsgRun {
    const message = createBaseMsgRun();
    message.callerAddress = object.callerAddress ?? "";
    message.scriptAddress = object.scriptAddress ?? "";
    message.extraCode = object.extraCode ?? "";
    message.functionName = object.functionName ?? "";
    message.kwargs = object.kwargs ?? "";
    message.grantee = object.grantee ?? "";
    return message;
  },
};

function createBaseMsgRunResponse(): MsgRunResponse {
  return { response: "" };
}

export const MsgRunResponse = {
  encode(message: MsgRunResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.response !== "") {
      writer.uint32(10).string(message.response);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRunResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRunResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.response = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgRunResponse {
    return { response: isSet(object.response) ? globalThis.String(object.response) : "" };
  },

  toJSON(message: MsgRunResponse): unknown {
    const obj: any = {};
    if (message.response !== "") {
      obj.response = message.response;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgRunResponse>, I>>(base?: I): MsgRunResponse {
    return MsgRunResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgRunResponse>, I>>(object: I): MsgRunResponse {
    const message = createBaseMsgRunResponse();
    message.response = object.response ?? "";
    return message;
  },
};

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
