// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: api/v1/user.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { FieldMask } from "../../google/protobuf/field_mask";
import { Timestamp } from "../../google/protobuf/timestamp";
import { AccessToken } from "../../model/user/access_token";
import { Setting } from "../../model/user/setting";
import { Stats } from "../../model/user/stats";
import { User } from "../../model/user/user";

export const protobufPackage = "api.v1";

export interface CreateAccessTokenRequest {
  /** The name of the user. */
  name: string;
  description: string;
  expiresAt?: Date | undefined;
}

export interface ListAllUserStatsRequest {
}

export interface ListAllUserStatsResponse {
  userStats: Stats[];
}

export interface GetUserStatsRequest {
  /** The name of the user. */
  name: string;
}

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

function createBaseCreateAccessTokenRequest(): CreateAccessTokenRequest {
  return { name: "", description: "", expiresAt: undefined };
}

export const CreateAccessTokenRequest: MessageFns<CreateAccessTokenRequest> = {
  encode(message: CreateAccessTokenRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.expiresAt !== undefined) {
      Timestamp.encode(toTimestamp(message.expiresAt), writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateAccessTokenRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateAccessTokenRequest();
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

          message.description = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.expiresAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
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

  create(base?: DeepPartial<CreateAccessTokenRequest>): CreateAccessTokenRequest {
    return CreateAccessTokenRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CreateAccessTokenRequest>): CreateAccessTokenRequest {
    const message = createBaseCreateAccessTokenRequest();
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.expiresAt = object.expiresAt ?? undefined;
    return message;
  },
};

function createBaseListAllUserStatsRequest(): ListAllUserStatsRequest {
  return {};
}

export const ListAllUserStatsRequest: MessageFns<ListAllUserStatsRequest> = {
  encode(_: ListAllUserStatsRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListAllUserStatsRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAllUserStatsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<ListAllUserStatsRequest>): ListAllUserStatsRequest {
    return ListAllUserStatsRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<ListAllUserStatsRequest>): ListAllUserStatsRequest {
    const message = createBaseListAllUserStatsRequest();
    return message;
  },
};

function createBaseListAllUserStatsResponse(): ListAllUserStatsResponse {
  return { userStats: [] };
}

export const ListAllUserStatsResponse: MessageFns<ListAllUserStatsResponse> = {
  encode(message: ListAllUserStatsResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.userStats) {
      Stats.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListAllUserStatsResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAllUserStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.userStats.push(Stats.decode(reader, reader.uint32()));
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

  create(base?: DeepPartial<ListAllUserStatsResponse>): ListAllUserStatsResponse {
    return ListAllUserStatsResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ListAllUserStatsResponse>): ListAllUserStatsResponse {
    const message = createBaseListAllUserStatsResponse();
    message.userStats = object.userStats?.map((e) => Stats.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetUserStatsRequest(): GetUserStatsRequest {
  return { name: "" };
}

export const GetUserStatsRequest: MessageFns<GetUserStatsRequest> = {
  encode(message: GetUserStatsRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetUserStatsRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserStatsRequest();
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

  create(base?: DeepPartial<GetUserStatsRequest>): GetUserStatsRequest {
    return GetUserStatsRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetUserStatsRequest>): GetUserStatsRequest {
    const message = createBaseGetUserStatsRequest();
    message.name = object.name ?? "";
    return message;
  },
};

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
    /** ListAllUserStats returns all user stats. */
    listAllUserStats: {
      name: "ListAllUserStats",
      requestType: ListAllUserStatsRequest,
      requestStream: false,
      responseType: ListAllUserStatsResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              23,
              34,
              21,
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
              115,
              47,
              45,
              47,
              115,
              116,
              97,
              116,
              115,
            ]),
          ],
        },
      },
    },
    /** GetUserStats returns the stats of a user. */
    getUserStats: {
      name: "GetUserStats",
      requestType: GetUserStatsRequest,
      requestStream: false,
      responseType: Stats,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              30,
              18,
              28,
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
              116,
              97,
              116,
              115,
            ]),
          ],
        },
      },
    },
    /** CreateUserAccessToken creates a new access token for a user. */
    createAccessToken: {
      name: "CreateAccessToken",
      requestType: CreateAccessTokenRequest,
      requestStream: false,
      responseType: AccessToken,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              41,
              58,
              1,
              42,
              34,
              36,
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
              97,
              99,
              99,
              101,
              115,
              115,
              95,
              116,
              111,
              107,
              101,
              110,
              115,
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

function toTimestamp(date: Date): Timestamp {
  const seconds = Math.trunc(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
