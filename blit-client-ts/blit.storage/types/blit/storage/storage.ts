/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "blit.storage";

export interface Storage {
  address: string;
  index: string;
  data: string;
}

function createBaseStorage(): Storage {
  return { address: "", index: "", data: "" };
}

export const Storage = {
  encode(message: Storage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.data !== "") {
      writer.uint32(26).string(message.data);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Storage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStorage();
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
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Storage {
    return {
      address: isSet(object.address) ? globalThis.String(object.address) : "",
      index: isSet(object.index) ? globalThis.String(object.index) : "",
      data: isSet(object.data) ? globalThis.String(object.data) : "",
    };
  },

  toJSON(message: Storage): unknown {
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
    return obj;
  },

  create<I extends Exact<DeepPartial<Storage>, I>>(base?: I): Storage {
    return Storage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Storage>, I>>(object: I): Storage {
    const message = createBaseStorage();
    message.address = object.address ?? "";
    message.index = object.index ?? "";
    message.data = object.data ?? "";
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
