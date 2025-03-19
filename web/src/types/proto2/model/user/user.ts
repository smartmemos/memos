// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: model/user/user.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "user";

export enum State {
  STATE_UNSPECIFIED = "STATE_UNSPECIFIED",
  NORMAL = "NORMAL",
  ARCHIVED = "ARCHIVED",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function stateFromJSON(object: any): State {
  switch (object) {
    case 0:
    case "STATE_UNSPECIFIED":
      return State.STATE_UNSPECIFIED;
    case 1:
    case "NORMAL":
      return State.NORMAL;
    case 2:
    case "ARCHIVED":
      return State.ARCHIVED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return State.UNRECOGNIZED;
  }
}

export function stateToNumber(object: State): number {
  switch (object) {
    case State.STATE_UNSPECIFIED:
      return 0;
    case State.NORMAL:
      return 1;
    case State.ARCHIVED:
      return 2;
    case State.UNRECOGNIZED:
    default:
      return -1;
  }
}

export enum Direction {
  DIRECTION_UNSPECIFIED = "DIRECTION_UNSPECIFIED",
  ASC = "ASC",
  DESC = "DESC",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function directionFromJSON(object: any): Direction {
  switch (object) {
    case 0:
    case "DIRECTION_UNSPECIFIED":
      return Direction.DIRECTION_UNSPECIFIED;
    case 1:
    case "ASC":
      return Direction.ASC;
    case 2:
    case "DESC":
      return Direction.DESC;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Direction.UNRECOGNIZED;
  }
}

export function directionToNumber(object: Direction): number {
  switch (object) {
    case Direction.DIRECTION_UNSPECIFIED:
      return 0;
    case Direction.ASC:
      return 1;
    case Direction.DESC:
      return 2;
    case Direction.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface User {
  id: number;
  name: string;
  username: string;
  role: User_Role;
  nickname: string;
  email: string;
  avatarUrl: string;
  description: string;
  password: string;
  state: State;
  createAt?: Date | undefined;
  updateAt?: Date | undefined;
}

export enum User_Role {
  ROLE_UNSPECIFIED = "ROLE_UNSPECIFIED",
  HOST = "HOST",
  ADMIN = "ADMIN",
  USER = "USER",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function user_RoleFromJSON(object: any): User_Role {
  switch (object) {
    case 0:
    case "ROLE_UNSPECIFIED":
      return User_Role.ROLE_UNSPECIFIED;
    case 1:
    case "HOST":
      return User_Role.HOST;
    case 2:
    case "ADMIN":
      return User_Role.ADMIN;
    case 3:
    case "USER":
      return User_Role.USER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return User_Role.UNRECOGNIZED;
  }
}

export function user_RoleToNumber(object: User_Role): number {
  switch (object) {
    case User_Role.ROLE_UNSPECIFIED:
      return 0;
    case User_Role.HOST:
      return 1;
    case User_Role.ADMIN:
      return 2;
    case User_Role.USER:
      return 3;
    case User_Role.UNRECOGNIZED:
    default:
      return -1;
  }
}

/** Used internally for obfuscating the page token. */
export interface PageToken {
  limit: number;
  offset: number;
}

function createBaseUser(): User {
  return {
    id: 0,
    name: "",
    username: "",
    role: User_Role.ROLE_UNSPECIFIED,
    nickname: "",
    email: "",
    avatarUrl: "",
    description: "",
    password: "",
    state: State.STATE_UNSPECIFIED,
    createAt: undefined,
    updateAt: undefined,
  };
}

export const User: MessageFns<User> = {
  encode(message: User, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.username !== "") {
      writer.uint32(26).string(message.username);
    }
    if (message.role !== User_Role.ROLE_UNSPECIFIED) {
      writer.uint32(32).int32(user_RoleToNumber(message.role));
    }
    if (message.nickname !== "") {
      writer.uint32(42).string(message.nickname);
    }
    if (message.email !== "") {
      writer.uint32(50).string(message.email);
    }
    if (message.avatarUrl !== "") {
      writer.uint32(58).string(message.avatarUrl);
    }
    if (message.description !== "") {
      writer.uint32(66).string(message.description);
    }
    if (message.password !== "") {
      writer.uint32(74).string(message.password);
    }
    if (message.state !== State.STATE_UNSPECIFIED) {
      writer.uint32(80).int32(stateToNumber(message.state));
    }
    if (message.createAt !== undefined) {
      Timestamp.encode(toTimestamp(message.createAt), writer.uint32(90).fork()).join();
    }
    if (message.updateAt !== undefined) {
      Timestamp.encode(toTimestamp(message.updateAt), writer.uint32(98).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): User {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.id = longToNumber(reader.int64());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.username = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.role = user_RoleFromJSON(reader.int32());
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.nickname = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.avatarUrl = reader.string();
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.description = reader.string();
          continue;
        }
        case 9: {
          if (tag !== 74) {
            break;
          }

          message.password = reader.string();
          continue;
        }
        case 10: {
          if (tag !== 80) {
            break;
          }

          message.state = stateFromJSON(reader.int32());
          continue;
        }
        case 11: {
          if (tag !== 90) {
            break;
          }

          message.createAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        }
        case 12: {
          if (tag !== 98) {
            break;
          }

          message.updateAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
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

  create(base?: DeepPartial<User>): User {
    return User.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<User>): User {
    const message = createBaseUser();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.username = object.username ?? "";
    message.role = object.role ?? User_Role.ROLE_UNSPECIFIED;
    message.nickname = object.nickname ?? "";
    message.email = object.email ?? "";
    message.avatarUrl = object.avatarUrl ?? "";
    message.description = object.description ?? "";
    message.password = object.password ?? "";
    message.state = object.state ?? State.STATE_UNSPECIFIED;
    message.createAt = object.createAt ?? undefined;
    message.updateAt = object.updateAt ?? undefined;
    return message;
  },
};

function createBasePageToken(): PageToken {
  return { limit: 0, offset: 0 };
}

export const PageToken: MessageFns<PageToken> = {
  encode(message: PageToken, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.limit !== 0) {
      writer.uint32(8).int32(message.limit);
    }
    if (message.offset !== 0) {
      writer.uint32(16).int32(message.offset);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): PageToken {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePageToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.limit = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.offset = reader.int32();
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

  create(base?: DeepPartial<PageToken>): PageToken {
    return PageToken.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<PageToken>): PageToken {
    const message = createBasePageToken();
    message.limit = object.limit ?? 0;
    message.offset = object.offset ?? 0;
    return message;
  },
};

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

function longToNumber(int64: { toString(): string }): number {
  const num = globalThis.Number(int64.toString());
  if (num > globalThis.Number.MAX_SAFE_INTEGER) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  if (num < globalThis.Number.MIN_SAFE_INTEGER) {
    throw new globalThis.Error("Value is smaller than Number.MIN_SAFE_INTEGER");
  }
  return num;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
