// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: api/v1/auth.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Empty } from "../../google/protobuf/empty";
import { User } from "../../model/system/user";

export const protobufPackage = "api.v1";

export interface SignInRequest {
  /** The username to sign in with. */
  username: string;
  /** The password to sign in with. */
  password: string;
  /** Whether the session should never expire. */
  neverExpire: boolean;
}

export interface SignUpRequest {
  /** The username to sign up with. */
  username: string;
  /** The password to sign up with. */
  password: string;
}

export interface SignOutRequest {
}

export interface GetAuthStatusRequest {
}

export interface GetAuthStatusResponse {
  user?: User | undefined;
}

function createBaseSignInRequest(): SignInRequest {
  return { username: "", password: "", neverExpire: false };
}

export const SignInRequest: MessageFns<SignInRequest> = {
  encode(message: SignInRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    if (message.neverExpire !== false) {
      writer.uint32(24).bool(message.neverExpire);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignInRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.password = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.neverExpire = reader.bool();
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

  create(base?: DeepPartial<SignInRequest>): SignInRequest {
    return SignInRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SignInRequest>): SignInRequest {
    const message = createBaseSignInRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    message.neverExpire = object.neverExpire ?? false;
    return message;
  },
};

function createBaseSignUpRequest(): SignUpRequest {
  return { username: "", password: "" };
}

export const SignUpRequest: MessageFns<SignUpRequest> = {
  encode(message: SignUpRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignUpRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignUpRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
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

  create(base?: DeepPartial<SignUpRequest>): SignUpRequest {
    return SignUpRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SignUpRequest>): SignUpRequest {
    const message = createBaseSignUpRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseSignOutRequest(): SignOutRequest {
  return {};
}

export const SignOutRequest: MessageFns<SignOutRequest> = {
  encode(_: SignOutRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignOutRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignOutRequest();
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

  create(base?: DeepPartial<SignOutRequest>): SignOutRequest {
    return SignOutRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<SignOutRequest>): SignOutRequest {
    const message = createBaseSignOutRequest();
    return message;
  },
};

function createBaseGetAuthStatusRequest(): GetAuthStatusRequest {
  return {};
}

export const GetAuthStatusRequest: MessageFns<GetAuthStatusRequest> = {
  encode(_: GetAuthStatusRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetAuthStatusRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAuthStatusRequest();
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

  create(base?: DeepPartial<GetAuthStatusRequest>): GetAuthStatusRequest {
    return GetAuthStatusRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<GetAuthStatusRequest>): GetAuthStatusRequest {
    const message = createBaseGetAuthStatusRequest();
    return message;
  },
};

function createBaseGetAuthStatusResponse(): GetAuthStatusResponse {
  return { user: undefined };
}

export const GetAuthStatusResponse: MessageFns<GetAuthStatusResponse> = {
  encode(message: GetAuthStatusResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetAuthStatusResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAuthStatusResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
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

  create(base?: DeepPartial<GetAuthStatusResponse>): GetAuthStatusResponse {
    return GetAuthStatusResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetAuthStatusResponse>): GetAuthStatusResponse {
    const message = createBaseGetAuthStatusResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

export type AuthServiceDefinition = typeof AuthServiceDefinition;
export const AuthServiceDefinition = {
  name: "AuthService",
  fullName: "api.v1.AuthService",
  methods: {
    /** GetAuthStatus returns the current auth status of the user. */
    getAuthStatus: {
      name: "GetAuthStatus",
      requestType: GetAuthStatusRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              116,
              97,
              116,
              117,
              115,
            ]),
          ],
        },
      },
    },
    /** SignIn signs in the user with the given username and password. */
    signIn: {
      name: "SignIn",
      requestType: SignInRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              105,
              110,
            ]),
          ],
        },
      },
    },
    /** SignUp signs up the user with the given username and password. */
    signUp: {
      name: "SignUp",
      requestType: SignUpRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              117,
              112,
            ]),
          ],
        },
      },
    },
    /** SignOut signs out the user. */
    signOut: {
      name: "SignOut",
      requestType: SignOutRequest,
      requestStream: false,
      responseType: Empty,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              22,
              34,
              20,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              111,
              117,
              116,
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
