// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: model/user/setting.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "user";

export interface Setting {
  /** The name of the user. */
  name: string;
  /** The preferred locale of the user. */
  locale: string;
  /** The preferred appearance of the user. */
  appearance: string;
  /** The default visibility of the memo. */
  memoVisibility: string;
}

function createBaseSetting(): Setting {
  return { name: "", locale: "", appearance: "", memoVisibility: "" };
}

export const Setting: MessageFns<Setting> = {
  encode(message: Setting, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.locale !== "") {
      writer.uint32(18).string(message.locale);
    }
    if (message.appearance !== "") {
      writer.uint32(26).string(message.appearance);
    }
    if (message.memoVisibility !== "") {
      writer.uint32(34).string(message.memoVisibility);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Setting {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.locale = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.appearance = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.memoVisibility = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<Setting>): Setting {
    return Setting.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<Setting>): Setting {
    const message = createBaseSetting();
    message.name = object.name ?? "";
    message.locale = object.locale ?? "";
    message.appearance = object.appearance ?? "";
    message.memoVisibility = object.memoVisibility ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
