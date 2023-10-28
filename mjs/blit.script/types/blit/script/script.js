/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
export const protobufPackage = "blit.script";
function createBaseScript() {
    return { address: "", code: "", version: 0 };
}
export const Script = {
    encode(message, writer = _m0.Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.code !== "") {
            writer.uint32(18).string(message.code);
        }
        if (message.version !== 0) {
            writer.uint32(24).int64(message.version);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseScript();
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
                    if (tag !== 24) {
                        break;
                    }
                    message.version = longToNumber(reader.int64());
                    continue;
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skipType(tag & 7);
        }
        return message;
    },
    fromJSON(object) {
        return {
            address: isSet(object.address) ? globalThis.String(object.address) : "",
            code: isSet(object.code) ? globalThis.String(object.code) : "",
            version: isSet(object.version) ? globalThis.Number(object.version) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.address !== "") {
            obj.address = message.address;
        }
        if (message.code !== "") {
            obj.code = message.code;
        }
        if (message.version !== 0) {
            obj.version = Math.round(message.version);
        }
        return obj;
    },
    create(base) {
        return Script.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseScript();
        message.address = object.address ?? "";
        message.code = object.code ?? "";
        message.version = object.version ?? 0;
        return message;
    },
};
function longToNumber(long) {
    if (long.gt(globalThis.Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (_m0.util.Long !== Long) {
    _m0.util.Long = Long;
    _m0.configure();
}
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=script.js.map