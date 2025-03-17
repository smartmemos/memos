// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: store/idp.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "memos.store";

export interface IdentityProvider {
  id: number;
  name: string;
  type: IdentityProvider_Type;
  identifierFilter: string;
  config?: IdentityProviderConfig | undefined;
}

export enum IdentityProvider_Type {
  TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
  OAUTH2 = "OAUTH2",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function identityProvider_TypeFromJSON(object: any): IdentityProvider_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return IdentityProvider_Type.TYPE_UNSPECIFIED;
    case 1:
    case "OAUTH2":
      return IdentityProvider_Type.OAUTH2;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IdentityProvider_Type.UNRECOGNIZED;
  }
}

export function identityProvider_TypeToNumber(object: IdentityProvider_Type): number {
  switch (object) {
    case IdentityProvider_Type.TYPE_UNSPECIFIED:
      return 0;
    case IdentityProvider_Type.OAUTH2:
      return 1;
    case IdentityProvider_Type.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface IdentityProviderConfig {
  oauth2Config?: OAuth2Config | undefined;
}

export interface FieldMapping {
  identifier: string;
  displayName: string;
  email: string;
}

export interface OAuth2Config {
  clientId: string;
  clientSecret: string;
  authUrl: string;
  tokenUrl: string;
  userInfoUrl: string;
  scopes: string[];
  fieldMapping?: FieldMapping | undefined;
}

function createBaseIdentityProvider(): IdentityProvider {
  return { id: 0, name: "", type: IdentityProvider_Type.TYPE_UNSPECIFIED, identifierFilter: "", config: undefined };
}

export const IdentityProvider: MessageFns<IdentityProvider> = {
  encode(message: IdentityProvider, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.type !== IdentityProvider_Type.TYPE_UNSPECIFIED) {
      writer.uint32(24).int32(identityProvider_TypeToNumber(message.type));
    }
    if (message.identifierFilter !== "") {
      writer.uint32(34).string(message.identifierFilter);
    }
    if (message.config !== undefined) {
      IdentityProviderConfig.encode(message.config, writer.uint32(42).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): IdentityProvider {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProvider();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
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
          if (tag !== 24) {
            break;
          }

          message.type = identityProvider_TypeFromJSON(reader.int32());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.identifierFilter = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.config = IdentityProviderConfig.decode(reader, reader.uint32());
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

  create(base?: DeepPartial<IdentityProvider>): IdentityProvider {
    return IdentityProvider.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IdentityProvider>): IdentityProvider {
    const message = createBaseIdentityProvider();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.type = object.type ?? IdentityProvider_Type.TYPE_UNSPECIFIED;
    message.identifierFilter = object.identifierFilter ?? "";
    message.config = (object.config !== undefined && object.config !== null)
      ? IdentityProviderConfig.fromPartial(object.config)
      : undefined;
    return message;
  },
};

function createBaseIdentityProviderConfig(): IdentityProviderConfig {
  return { oauth2Config: undefined };
}

export const IdentityProviderConfig: MessageFns<IdentityProviderConfig> = {
  encode(message: IdentityProviderConfig, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oauth2Config !== undefined) {
      OAuth2Config.encode(message.oauth2Config, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): IdentityProviderConfig {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.oauth2Config = OAuth2Config.decode(reader, reader.uint32());
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

  create(base?: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    return IdentityProviderConfig.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    const message = createBaseIdentityProviderConfig();
    message.oauth2Config = (object.oauth2Config !== undefined && object.oauth2Config !== null)
      ? OAuth2Config.fromPartial(object.oauth2Config)
      : undefined;
    return message;
  },
};

function createBaseFieldMapping(): FieldMapping {
  return { identifier: "", displayName: "", email: "" };
}

export const FieldMapping: MessageFns<FieldMapping> = {
  encode(message: FieldMapping, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.identifier !== "") {
      writer.uint32(10).string(message.identifier);
    }
    if (message.displayName !== "") {
      writer.uint32(18).string(message.displayName);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): FieldMapping {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFieldMapping();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.identifier = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.displayName = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.email = reader.string();
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

  create(base?: DeepPartial<FieldMapping>): FieldMapping {
    return FieldMapping.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<FieldMapping>): FieldMapping {
    const message = createBaseFieldMapping();
    message.identifier = object.identifier ?? "";
    message.displayName = object.displayName ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseOAuth2Config(): OAuth2Config {
  return {
    clientId: "",
    clientSecret: "",
    authUrl: "",
    tokenUrl: "",
    userInfoUrl: "",
    scopes: [],
    fieldMapping: undefined,
  };
}

export const OAuth2Config: MessageFns<OAuth2Config> = {
  encode(message: OAuth2Config, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.clientId !== "") {
      writer.uint32(10).string(message.clientId);
    }
    if (message.clientSecret !== "") {
      writer.uint32(18).string(message.clientSecret);
    }
    if (message.authUrl !== "") {
      writer.uint32(26).string(message.authUrl);
    }
    if (message.tokenUrl !== "") {
      writer.uint32(34).string(message.tokenUrl);
    }
    if (message.userInfoUrl !== "") {
      writer.uint32(42).string(message.userInfoUrl);
    }
    for (const v of message.scopes) {
      writer.uint32(50).string(v!);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(58).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): OAuth2Config {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuth2Config();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.clientId = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.clientSecret = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.authUrl = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.tokenUrl = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.userInfoUrl = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.scopes.push(reader.string());
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
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

  create(base?: DeepPartial<OAuth2Config>): OAuth2Config {
    return OAuth2Config.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<OAuth2Config>): OAuth2Config {
    const message = createBaseOAuth2Config();
    message.clientId = object.clientId ?? "";
    message.clientSecret = object.clientSecret ?? "";
    message.authUrl = object.authUrl ?? "";
    message.tokenUrl = object.tokenUrl ?? "";
    message.userInfoUrl = object.userInfoUrl ?? "";
    message.scopes = object.scopes?.map((e) => e) || [];
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
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
