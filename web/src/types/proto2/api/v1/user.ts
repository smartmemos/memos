// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: api/v1/user.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { FieldMask } from "../../google/protobuf/field_mask";
import { Setting } from "../../model/user/setting";
import { User } from "../../model/user/user";

export const protobufPackage = "api.v1";

export interface CreateUserRequest {
  name: string;
  username: string;
  email: string;
  nickname: string;
  avatarUrl: string;
  description: string;
  password: string;
}

export interface GetUserSettingRequest {
  /** The name of the user. */
  name: string;
}

export interface UpdateUserSettingRequest {
  setting?: Setting | undefined;
  updateMask?: string[] | undefined;
}

function createBaseCreateUserRequest(): CreateUserRequest {
  return { name: "", username: "", email: "", nickname: "", avatarUrl: "", description: "", password: "" };
}

export const CreateUserRequest: MessageFns<CreateUserRequest> = {
  encode(message: CreateUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.username !== "") {
      writer.uint32(18).string(message.username);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    if (message.nickname !== "") {
      writer.uint32(34).string(message.nickname);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(42).string(message.avatarUrl);
    }
    if (message.description !== "") {
      writer.uint32(50).string(message.description);
    }
    if (message.password !== "") {
      writer.uint32(58).string(message.password);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateUserRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
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

          message.username = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.nickname = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.avatarUrl = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.description = reader.string();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.password = reader.string();
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

  create(base?: DeepPartial<CreateUserRequest>): CreateUserRequest {
    return CreateUserRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CreateUserRequest>): CreateUserRequest {
    const message = createBaseCreateUserRequest();
    message.name = object.name ?? "";
    message.username = object.username ?? "";
    message.email = object.email ?? "";
    message.nickname = object.nickname ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    message.description = object.description ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseGetUserSettingRequest(): GetUserSettingRequest {
  return { name: "" };
}

export const GetUserSettingRequest: MessageFns<GetUserSettingRequest> = {
  encode(message: GetUserSettingRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetUserSettingRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserSettingRequest();
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
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<GetUserSettingRequest>): GetUserSettingRequest {
    return GetUserSettingRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetUserSettingRequest>): GetUserSettingRequest {
    const message = createBaseGetUserSettingRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseUpdateUserSettingRequest(): UpdateUserSettingRequest {
  return { setting: undefined, updateMask: undefined };
}

export const UpdateUserSettingRequest: MessageFns<UpdateUserSettingRequest> = {
  encode(message: UpdateUserSettingRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.setting !== undefined) {
      Setting.encode(message.setting, writer.uint32(10).fork()).join();
    }
    if (message.updateMask !== undefined) {
      FieldMask.encode(FieldMask.wrap(message.updateMask), writer.uint32(18).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UpdateUserSettingRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserSettingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.setting = Setting.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.updateMask = FieldMask.unwrap(FieldMask.decode(reader, reader.uint32()));
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

  create(base?: DeepPartial<UpdateUserSettingRequest>): UpdateUserSettingRequest {
    return UpdateUserSettingRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UpdateUserSettingRequest>): UpdateUserSettingRequest {
    const message = createBaseUpdateUserSettingRequest();
    message.setting = (object.setting !== undefined && object.setting !== null)
      ? Setting.fromPartial(object.setting)
      : undefined;
    message.updateMask = object.updateMask ?? undefined;
    return message;
  },
};

export type UserServiceDefinition = typeof UserServiceDefinition;
export const UserServiceDefinition = {
  name: "UserService",
  fullName: "api.v1.UserService",
  methods: {
    /** CreateUser creates a new user. */
    createUser: {
      name: "CreateUser",
      requestType: CreateUserRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [
            new Uint8Array([
              22,
              110,
              97,
              109,
              101,
              44,
              117,
              115,
              101,
              114,
              110,
              97,
              109,
              101,
              44,
              112,
              97,
              115,
              115,
              119,
              111,
              114,
              100,
            ]),
          ],
          578365826: [
            new Uint8Array([
              20,
              34,
              18,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              117,
              115,
              101,
              114,
              47,
              117,
              115,
              101,
              114,
              115,
            ]),
          ],
        },
      },
    },
    /** GetUserSetting gets the setting of a user. */
    getUserSetting: {
      name: "GetUserSetting",
      requestType: GetUserSettingRequest,
      requestStream: false,
      responseType: Setting,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              32,
              18,
              30,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              117,
              115,
              101,
              114,
              115,
              47,
              42,
              125,
              47,
              115,
              101,
              116,
              116,
              105,
              110,
              103,
            ]),
          ],
        },
      },
    },
    /** UpdateUserSetting updates the setting of a user. */
    updateUserSetting: {
      name: "UpdateUserSetting",
      requestType: UpdateUserSettingRequest,
      requestStream: false,
      responseType: Setting,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [
            new Uint8Array([
              19,
              115,
              101,
              116,
              116,
              105,
              110,
              103,
              44,
              117,
              112,
              100,
              97,
              116,
              101,
              95,
              109,
              97,
              115,
              107,
            ]),
          ],
          578365826: [
            new Uint8Array([
              49,
              58,
              7,
              115,
              101,
              116,
              116,
              105,
              110,
              103,
              50,
              38,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              123,
              115,
              101,
              116,
              116,
              105,
              110,
              103,
              46,
              110,
              97,
              109,
              101,
              61,
              117,
              115,
              101,
              114,
              115,
              47,
              42,
              47,
              115,
              101,
              116,
              116,
              105,
              110,
              103,
              125,
            ]),
          ],
        },
      },
    },
  },
} as const;

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
