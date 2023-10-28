/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
export const protobufPackage = "blit.script";
function createBaseMsgRun() {
    return { callerAddress: "", scriptAddress: "", extraCode: "", functionName: "", kwargs: "", grantee: "" };
}
export const MsgRun = {
    encode(message, writer = _m0.Writer.create()) {
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
    decode(input, length) {
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
    fromJSON(object) {
        return {
            callerAddress: isSet(object.callerAddress) ? globalThis.String(object.callerAddress) : "",
            scriptAddress: isSet(object.scriptAddress) ? globalThis.String(object.scriptAddress) : "",
            extraCode: isSet(object.extraCode) ? globalThis.String(object.extraCode) : "",
            functionName: isSet(object.functionName) ? globalThis.String(object.functionName) : "",
            kwargs: isSet(object.kwargs) ? globalThis.String(object.kwargs) : "",
            grantee: isSet(object.grantee) ? globalThis.String(object.grantee) : "",
        };
    },
    toJSON(message) {
        const obj = {};
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
    create(base) {
        return MsgRun.fromPartial(base ?? {});
    },
    fromPartial(object) {
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
function createBaseMsgRunResponse() {
    return { response: "" };
}
export const MsgRunResponse = {
    encode(message, writer = _m0.Writer.create()) {
        if (message.response !== "") {
            writer.uint32(10).string(message.response);
        }
        return writer;
    },
    decode(input, length) {
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
    fromJSON(object) {
        return { response: isSet(object.response) ? globalThis.String(object.response) : "" };
    },
    toJSON(message) {
        const obj = {};
        if (message.response !== "") {
            obj.response = message.response;
        }
        return obj;
    },
    create(base) {
        return MsgRunResponse.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseMsgRunResponse();
        message.response = object.response ?? "";
        return message;
    },
};
function isSet(value) {
    return value !== null && value !== undefined;
}
//# sourceMappingURL=msgrun.js.map