/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
export const protobufPackage = "google.protobuf";
function createBaseDuration() {
    return { seconds: 0, nanos: 0 };
}
export const Duration = {
    encode(message, writer = _m0.Writer.create()) {
        if (message.seconds !== 0) {
            writer.uint32(8).int64(message.seconds);
        }
        if (message.nanos !== 0) {
            writer.uint32(16).int32(message.nanos);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDuration();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if (tag !== 8) {
                        break;
                    }
                    message.seconds = longToNumber(reader.int64());
                    continue;
                case 2:
                    if (tag !== 16) {
                        break;
                    }
                    message.nanos = reader.int32();
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
            seconds: isSet(object.seconds) ? globalThis.Number(object.seconds) : 0,
            nanos: isSet(object.nanos) ? globalThis.Number(object.nanos) : 0,
        };
    },
    toJSON(message) {
        const obj = {};
        if (message.seconds !== 0) {
            obj.seconds = Math.round(message.seconds);
        }
        if (message.nanos !== 0) {
            obj.nanos = Math.round(message.nanos);
        }
        return obj;
    },
    create(base) {
        return Duration.fromPartial(base ?? {});
    },
    fromPartial(object) {
        const message = createBaseDuration();
        message.seconds = object.seconds ?? 0;
        message.nanos = object.nanos ?? 0;
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
//# sourceMappingURL=duration.js.map