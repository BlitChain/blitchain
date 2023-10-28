/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
export const protobufPackage = "blit.storage";
function createBaseParams() {
    return { gasPerChar: "" };
}
export const Params = {
    encode(message, writer = _m0.Writer.create()) {
        if (message.gasPerChar !== "") {
            writer.uint32(10).string(message.gasPerChar);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseParams();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if (tag !== 10) {
                        break;
                    }
                    message.gasPerChar = reader.string();
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
        return { gasPerChar: isSet(object.gasPerChar) ? globalThis.String(object.gasPerChar) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.gasPerChar !== "") {
            obj.gasPerChar = message.gasPerChar;
        }
        return obj;
    },
    create(base) {
        return Params.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseParams();
        message.gasPerChar = object.gasPerChar ?? "";
        return message;
    },
};
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=params.js.map