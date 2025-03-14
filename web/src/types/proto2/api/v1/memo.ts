// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: api/v1/memo.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Memo } from "../../model/memo/memo";

export const protobufPackage = "api.v1";

export interface CreateMemoRequest {
  name: string;
}

export interface ListMemosRequest {
}

export interface ListMemosResponse {
}

export interface GetMemoRequest {
}

function createBaseCreateMemoRequest(): CreateMemoRequest {
  return { name: "" };
}

export const CreateMemoRequest: MessageFns<CreateMemoRequest> = {
  encode(message: CreateMemoRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateMemoRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMemoRequest();
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

  create(base?: DeepPartial<CreateMemoRequest>): CreateMemoRequest {
    return CreateMemoRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CreateMemoRequest>): CreateMemoRequest {
    const message = createBaseCreateMemoRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseListMemosRequest(): ListMemosRequest {
  return {};
}

export const ListMemosRequest: MessageFns<ListMemosRequest> = {
  encode(_: ListMemosRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListMemosRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMemosRequest();
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

  create(base?: DeepPartial<ListMemosRequest>): ListMemosRequest {
    return ListMemosRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<ListMemosRequest>): ListMemosRequest {
    const message = createBaseListMemosRequest();
    return message;
  },
};

function createBaseListMemosResponse(): ListMemosResponse {
  return {};
}

export const ListMemosResponse: MessageFns<ListMemosResponse> = {
  encode(_: ListMemosResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListMemosResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListMemosResponse();
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

  create(base?: DeepPartial<ListMemosResponse>): ListMemosResponse {
    return ListMemosResponse.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<ListMemosResponse>): ListMemosResponse {
    const message = createBaseListMemosResponse();
    return message;
  },
};

function createBaseGetMemoRequest(): GetMemoRequest {
  return {};
}

export const GetMemoRequest: MessageFns<GetMemoRequest> = {
  encode(_: GetMemoRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetMemoRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMemoRequest();
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

  create(base?: DeepPartial<GetMemoRequest>): GetMemoRequest {
    return GetMemoRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<GetMemoRequest>): GetMemoRequest {
    const message = createBaseGetMemoRequest();
    return message;
  },
};

export type MemoServiceDefinition = typeof MemoServiceDefinition;
export const MemoServiceDefinition = {
  name: "MemoService",
  fullName: "api.v1.MemoService",
  methods: {
    /** CreateMemo creates a memo. */
    createMemo: {
      name: "CreateMemo",
      requestType: CreateMemoRequest,
      requestStream: false,
      responseType: Memo,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [new Uint8Array([15, 34, 13, 47, 97, 112, 105, 47, 118, 49, 47, 109, 101, 109, 111, 115])],
        },
      },
    },
    /** ListMemos lists memos with pagination and filter. */
    listMemos: {
      name: "ListMemos",
      requestType: ListMemosRequest,
      requestStream: false,
      responseType: ListMemosResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [new Uint8Array([15, 18, 13, 47, 97, 112, 105, 47, 118, 49, 47, 109, 101, 109, 111, 115])],
        },
      },
    },
    /** GetMemo gets a memo. */
    getMemo: {
      name: "GetMemo",
      requestType: GetMemoRequest,
      requestStream: false,
      responseType: Memo,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [new Uint8Array([17, 18, 15, 47, 97, 112, 105, 47, 118, 49, 47, 109, 101, 109, 111, 115, 47, 42])],
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
