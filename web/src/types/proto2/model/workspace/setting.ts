// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: model/workspace/setting.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "workspace";

export enum SettingKey {
  WORKSPACE_SETTING_KEY_UNSPECIFIED = "WORKSPACE_SETTING_KEY_UNSPECIFIED",
  /** BASIC - BASIC is the key for basic settings. */
  BASIC = "BASIC",
  /** GENERAL - GENERAL is the key for general settings. */
  GENERAL = "GENERAL",
  /** STORAGE - STORAGE is the key for storage settings. */
  STORAGE = "STORAGE",
  /** MEMO_RELATED - MEMO_RELATED is the key for memo related settings. */
  MEMO_RELATED = "MEMO_RELATED",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function settingKeyFromJSON(object: any): SettingKey {
  switch (object) {
    case 0:
    case "WORKSPACE_SETTING_KEY_UNSPECIFIED":
      return SettingKey.WORKSPACE_SETTING_KEY_UNSPECIFIED;
    case 1:
    case "BASIC":
      return SettingKey.BASIC;
    case 2:
    case "GENERAL":
      return SettingKey.GENERAL;
    case 3:
    case "STORAGE":
      return SettingKey.STORAGE;
    case 4:
    case "MEMO_RELATED":
      return SettingKey.MEMO_RELATED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SettingKey.UNRECOGNIZED;
  }
}

export function settingKeyToNumber(object: SettingKey): number {
  switch (object) {
    case SettingKey.WORKSPACE_SETTING_KEY_UNSPECIFIED:
      return 0;
    case SettingKey.BASIC:
      return 1;
    case SettingKey.GENERAL:
      return 2;
    case SettingKey.STORAGE:
      return 3;
    case SettingKey.MEMO_RELATED:
      return 4;
    case SettingKey.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface Setting {
  /**
   * name is the name of the setting.
   * Format: settings/{setting}
   */
  name: string;
  generalSetting?: GeneralSetting | undefined;
  storageSetting?: StorageSetting | undefined;
  memoRelatedSetting?: MemoRelatedSetting | undefined;
}

export interface GeneralSetting {
  /** disallow_user_registration disallows user registration. */
  disallowUserRegistration: boolean;
  /** disallow_password_auth disallows password authentication. */
  disallowPasswordAuth: boolean;
  /** additional_script is the additional script. */
  additionalScript: string;
  /** additional_style is the additional style. */
  additionalStyle: string;
  /** custom_profile is the custom profile. */
  customProfile?:
    | CustomProfile
    | undefined;
  /**
   * week_start_day_offset is the week start day offset from Sunday.
   * 0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday
   * Default is Sunday.
   */
  weekStartDayOffset: number;
  /** disallow_change_username disallows changing username. */
  disallowChangeUsername: boolean;
  /** disallow_change_nickname disallows changing nickname. */
  disallowChangeNickname: boolean;
}

export interface CustomProfile {
  title: string;
  description: string;
  logoUrl: string;
  locale: string;
  appearance: string;
}

export interface StorageSetting {
  /** storage_type is the storage type. */
  storageType: StorageSetting_StorageType;
  /**
   * The template of file path.
   * e.g. assets/{timestamp}_{filename}
   */
  filepathTemplate: string;
  /** The max upload size in megabytes. */
  uploadSizeLimitMb: number;
  /** The S3 config. */
  s3Config?: StorageSetting_S3Config | undefined;
}

export enum StorageSetting_StorageType {
  STORAGE_TYPE_UNSPECIFIED = "STORAGE_TYPE_UNSPECIFIED",
  /** DATABASE - DATABASE is the database storage type. */
  DATABASE = "DATABASE",
  /** LOCAL - LOCAL is the local storage type. */
  LOCAL = "LOCAL",
  /** S3 - S3 is the S3 storage type. */
  S3 = "S3",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function storageSetting_StorageTypeFromJSON(object: any): StorageSetting_StorageType {
  switch (object) {
    case 0:
    case "STORAGE_TYPE_UNSPECIFIED":
      return StorageSetting_StorageType.STORAGE_TYPE_UNSPECIFIED;
    case 1:
    case "DATABASE":
      return StorageSetting_StorageType.DATABASE;
    case 2:
    case "LOCAL":
      return StorageSetting_StorageType.LOCAL;
    case 3:
    case "S3":
      return StorageSetting_StorageType.S3;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StorageSetting_StorageType.UNRECOGNIZED;
  }
}

export function storageSetting_StorageTypeToNumber(object: StorageSetting_StorageType): number {
  switch (object) {
    case StorageSetting_StorageType.STORAGE_TYPE_UNSPECIFIED:
      return 0;
    case StorageSetting_StorageType.DATABASE:
      return 1;
    case StorageSetting_StorageType.LOCAL:
      return 2;
    case StorageSetting_StorageType.S3:
      return 3;
    case StorageSetting_StorageType.UNRECOGNIZED:
    default:
      return -1;
  }
}

/** Reference: https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/ */
export interface StorageSetting_S3Config {
  accessKeyId: string;
  accessKeySecret: string;
  endpoint: string;
  region: string;
  bucket: string;
  usePathStyle: boolean;
}

export interface MemoRelatedSetting {
  /** disallow_public_visibility disallows set memo as public visibility. */
  disallowPublicVisibility: boolean;
  /** display_with_update_time orders and displays memo with update time. */
  displayWithUpdateTime: boolean;
  /** content_length_limit is the limit of content length. Unit is byte. */
  contentLengthLimit: number;
  /** enable_double_click_edit enables editing on double click. */
  enableDoubleClickEdit: boolean;
  /** enable_link_preview enables links preview. */
  enableLinkPreview: boolean;
  /** enable_comment enables comment. */
  enableComment: boolean;
  /** enable_location enables setting location for memo. */
  enableLocation: boolean;
  /** reactions is the list of reactions. */
  reactions: string[];
  /** disable_markdown_shortcuts disallow the registration of markdown shortcuts. */
  disableMarkdownShortcuts: boolean;
  /** enable_blur_nsfw_content enables blurring of content marked as not safe for work (NSFW). */
  enableBlurNsfwContent: boolean;
  /** nsfw_tags is the list of tags that mark content as NSFW for blurring. */
  nsfwTags: string[];
}

function createBaseSetting(): Setting {
  return { name: "", generalSetting: undefined, storageSetting: undefined, memoRelatedSetting: undefined };
}

export const Setting: MessageFns<Setting> = {
  encode(message: Setting, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.generalSetting !== undefined) {
      GeneralSetting.encode(message.generalSetting, writer.uint32(18).fork()).join();
    }
    if (message.storageSetting !== undefined) {
      StorageSetting.encode(message.storageSetting, writer.uint32(26).fork()).join();
    }
    if (message.memoRelatedSetting !== undefined) {
      MemoRelatedSetting.encode(message.memoRelatedSetting, writer.uint32(34).fork()).join();
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

          message.generalSetting = GeneralSetting.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.storageSetting = StorageSetting.decode(reader, reader.uint32());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.memoRelatedSetting = MemoRelatedSetting.decode(reader, reader.uint32());
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
    message.generalSetting = (object.generalSetting !== undefined && object.generalSetting !== null)
      ? GeneralSetting.fromPartial(object.generalSetting)
      : undefined;
    message.storageSetting = (object.storageSetting !== undefined && object.storageSetting !== null)
      ? StorageSetting.fromPartial(object.storageSetting)
      : undefined;
    message.memoRelatedSetting = (object.memoRelatedSetting !== undefined && object.memoRelatedSetting !== null)
      ? MemoRelatedSetting.fromPartial(object.memoRelatedSetting)
      : undefined;
    return message;
  },
};

function createBaseGeneralSetting(): GeneralSetting {
  return {
    disallowUserRegistration: false,
    disallowPasswordAuth: false,
    additionalScript: "",
    additionalStyle: "",
    customProfile: undefined,
    weekStartDayOffset: 0,
    disallowChangeUsername: false,
    disallowChangeNickname: false,
  };
}

export const GeneralSetting: MessageFns<GeneralSetting> = {
  encode(message: GeneralSetting, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.disallowUserRegistration !== false) {
      writer.uint32(8).bool(message.disallowUserRegistration);
    }
    if (message.disallowPasswordAuth !== false) {
      writer.uint32(16).bool(message.disallowPasswordAuth);
    }
    if (message.additionalScript !== "") {
      writer.uint32(26).string(message.additionalScript);
    }
    if (message.additionalStyle !== "") {
      writer.uint32(34).string(message.additionalStyle);
    }
    if (message.customProfile !== undefined) {
      CustomProfile.encode(message.customProfile, writer.uint32(42).fork()).join();
    }
    if (message.weekStartDayOffset !== 0) {
      writer.uint32(48).int32(message.weekStartDayOffset);
    }
    if (message.disallowChangeUsername !== false) {
      writer.uint32(56).bool(message.disallowChangeUsername);
    }
    if (message.disallowChangeNickname !== false) {
      writer.uint32(64).bool(message.disallowChangeNickname);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GeneralSetting {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGeneralSetting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.disallowUserRegistration = reader.bool();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.disallowPasswordAuth = reader.bool();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.additionalScript = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.additionalStyle = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.customProfile = CustomProfile.decode(reader, reader.uint32());
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.weekStartDayOffset = reader.int32();
          continue;
        }
        case 7: {
          if (tag !== 56) {
            break;
          }

          message.disallowChangeUsername = reader.bool();
          continue;
        }
        case 8: {
          if (tag !== 64) {
            break;
          }

          message.disallowChangeNickname = reader.bool();
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

  create(base?: DeepPartial<GeneralSetting>): GeneralSetting {
    return GeneralSetting.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GeneralSetting>): GeneralSetting {
    const message = createBaseGeneralSetting();
    message.disallowUserRegistration = object.disallowUserRegistration ?? false;
    message.disallowPasswordAuth = object.disallowPasswordAuth ?? false;
    message.additionalScript = object.additionalScript ?? "";
    message.additionalStyle = object.additionalStyle ?? "";
    message.customProfile = (object.customProfile !== undefined && object.customProfile !== null)
      ? CustomProfile.fromPartial(object.customProfile)
      : undefined;
    message.weekStartDayOffset = object.weekStartDayOffset ?? 0;
    message.disallowChangeUsername = object.disallowChangeUsername ?? false;
    message.disallowChangeNickname = object.disallowChangeNickname ?? false;
    return message;
  },
};

function createBaseCustomProfile(): CustomProfile {
  return { title: "", description: "", logoUrl: "", locale: "", appearance: "" };
}

export const CustomProfile: MessageFns<CustomProfile> = {
  encode(message: CustomProfile, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.logoUrl !== "") {
      writer.uint32(26).string(message.logoUrl);
    }
    if (message.locale !== "") {
      writer.uint32(34).string(message.locale);
    }
    if (message.appearance !== "") {
      writer.uint32(42).string(message.appearance);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CustomProfile {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCustomProfile();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
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

          message.logoUrl = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.locale = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.appearance = reader.string();
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

  create(base?: DeepPartial<CustomProfile>): CustomProfile {
    return CustomProfile.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CustomProfile>): CustomProfile {
    const message = createBaseCustomProfile();
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.logoUrl = object.logoUrl ?? "";
    message.locale = object.locale ?? "";
    message.appearance = object.appearance ?? "";
    return message;
  },
};

function createBaseStorageSetting(): StorageSetting {
  return {
    storageType: StorageSetting_StorageType.STORAGE_TYPE_UNSPECIFIED,
    filepathTemplate: "",
    uploadSizeLimitMb: 0,
    s3Config: undefined,
  };
}

export const StorageSetting: MessageFns<StorageSetting> = {
  encode(message: StorageSetting, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.storageType !== StorageSetting_StorageType.STORAGE_TYPE_UNSPECIFIED) {
      writer.uint32(8).int32(storageSetting_StorageTypeToNumber(message.storageType));
    }
    if (message.filepathTemplate !== "") {
      writer.uint32(18).string(message.filepathTemplate);
    }
    if (message.uploadSizeLimitMb !== 0) {
      writer.uint32(24).int64(message.uploadSizeLimitMb);
    }
    if (message.s3Config !== undefined) {
      StorageSetting_S3Config.encode(message.s3Config, writer.uint32(34).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): StorageSetting {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStorageSetting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.storageType = storageSetting_StorageTypeFromJSON(reader.int32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.filepathTemplate = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.uploadSizeLimitMb = longToNumber(reader.int64());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.s3Config = StorageSetting_S3Config.decode(reader, reader.uint32());
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

  create(base?: DeepPartial<StorageSetting>): StorageSetting {
    return StorageSetting.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<StorageSetting>): StorageSetting {
    const message = createBaseStorageSetting();
    message.storageType = object.storageType ?? StorageSetting_StorageType.STORAGE_TYPE_UNSPECIFIED;
    message.filepathTemplate = object.filepathTemplate ?? "";
    message.uploadSizeLimitMb = object.uploadSizeLimitMb ?? 0;
    message.s3Config = (object.s3Config !== undefined && object.s3Config !== null)
      ? StorageSetting_S3Config.fromPartial(object.s3Config)
      : undefined;
    return message;
  },
};

function createBaseStorageSetting_S3Config(): StorageSetting_S3Config {
  return { accessKeyId: "", accessKeySecret: "", endpoint: "", region: "", bucket: "", usePathStyle: false };
}

export const StorageSetting_S3Config: MessageFns<StorageSetting_S3Config> = {
  encode(message: StorageSetting_S3Config, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.accessKeyId !== "") {
      writer.uint32(10).string(message.accessKeyId);
    }
    if (message.accessKeySecret !== "") {
      writer.uint32(18).string(message.accessKeySecret);
    }
    if (message.endpoint !== "") {
      writer.uint32(26).string(message.endpoint);
    }
    if (message.region !== "") {
      writer.uint32(34).string(message.region);
    }
    if (message.bucket !== "") {
      writer.uint32(42).string(message.bucket);
    }
    if (message.usePathStyle !== false) {
      writer.uint32(48).bool(message.usePathStyle);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): StorageSetting_S3Config {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStorageSetting_S3Config();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.accessKeyId = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.accessKeySecret = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.endpoint = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.region = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.bucket = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.usePathStyle = reader.bool();
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

  create(base?: DeepPartial<StorageSetting_S3Config>): StorageSetting_S3Config {
    return StorageSetting_S3Config.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<StorageSetting_S3Config>): StorageSetting_S3Config {
    const message = createBaseStorageSetting_S3Config();
    message.accessKeyId = object.accessKeyId ?? "";
    message.accessKeySecret = object.accessKeySecret ?? "";
    message.endpoint = object.endpoint ?? "";
    message.region = object.region ?? "";
    message.bucket = object.bucket ?? "";
    message.usePathStyle = object.usePathStyle ?? false;
    return message;
  },
};

function createBaseMemoRelatedSetting(): MemoRelatedSetting {
  return {
    disallowPublicVisibility: false,
    displayWithUpdateTime: false,
    contentLengthLimit: 0,
    enableDoubleClickEdit: false,
    enableLinkPreview: false,
    enableComment: false,
    enableLocation: false,
    reactions: [],
    disableMarkdownShortcuts: false,
    enableBlurNsfwContent: false,
    nsfwTags: [],
  };
}

export const MemoRelatedSetting: MessageFns<MemoRelatedSetting> = {
  encode(message: MemoRelatedSetting, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.disallowPublicVisibility !== false) {
      writer.uint32(8).bool(message.disallowPublicVisibility);
    }
    if (message.displayWithUpdateTime !== false) {
      writer.uint32(16).bool(message.displayWithUpdateTime);
    }
    if (message.contentLengthLimit !== 0) {
      writer.uint32(24).int32(message.contentLengthLimit);
    }
    if (message.enableDoubleClickEdit !== false) {
      writer.uint32(40).bool(message.enableDoubleClickEdit);
    }
    if (message.enableLinkPreview !== false) {
      writer.uint32(48).bool(message.enableLinkPreview);
    }
    if (message.enableComment !== false) {
      writer.uint32(56).bool(message.enableComment);
    }
    if (message.enableLocation !== false) {
      writer.uint32(64).bool(message.enableLocation);
    }
    for (const v of message.reactions) {
      writer.uint32(82).string(v!);
    }
    if (message.disableMarkdownShortcuts !== false) {
      writer.uint32(88).bool(message.disableMarkdownShortcuts);
    }
    if (message.enableBlurNsfwContent !== false) {
      writer.uint32(96).bool(message.enableBlurNsfwContent);
    }
    for (const v of message.nsfwTags) {
      writer.uint32(106).string(v!);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): MemoRelatedSetting {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMemoRelatedSetting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.disallowPublicVisibility = reader.bool();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.displayWithUpdateTime = reader.bool();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.contentLengthLimit = reader.int32();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.enableDoubleClickEdit = reader.bool();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.enableLinkPreview = reader.bool();
          continue;
        }
        case 7: {
          if (tag !== 56) {
            break;
          }

          message.enableComment = reader.bool();
          continue;
        }
        case 8: {
          if (tag !== 64) {
            break;
          }

          message.enableLocation = reader.bool();
          continue;
        }
        case 10: {
          if (tag !== 82) {
            break;
          }

          message.reactions.push(reader.string());
          continue;
        }
        case 11: {
          if (tag !== 88) {
            break;
          }

          message.disableMarkdownShortcuts = reader.bool();
          continue;
        }
        case 12: {
          if (tag !== 96) {
            break;
          }

          message.enableBlurNsfwContent = reader.bool();
          continue;
        }
        case 13: {
          if (tag !== 106) {
            break;
          }

          message.nsfwTags.push(reader.string());
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

  create(base?: DeepPartial<MemoRelatedSetting>): MemoRelatedSetting {
    return MemoRelatedSetting.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MemoRelatedSetting>): MemoRelatedSetting {
    const message = createBaseMemoRelatedSetting();
    message.disallowPublicVisibility = object.disallowPublicVisibility ?? false;
    message.displayWithUpdateTime = object.displayWithUpdateTime ?? false;
    message.contentLengthLimit = object.contentLengthLimit ?? 0;
    message.enableDoubleClickEdit = object.enableDoubleClickEdit ?? false;
    message.enableLinkPreview = object.enableLinkPreview ?? false;
    message.enableComment = object.enableComment ?? false;
    message.enableLocation = object.enableLocation ?? false;
    message.reactions = object.reactions?.map((e) => e) || [];
    message.disableMarkdownShortcuts = object.disableMarkdownShortcuts ?? false;
    message.enableBlurNsfwContent = object.enableBlurNsfwContent ?? false;
    message.nsfwTags = object.nsfwTags?.map((e) => e) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

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
