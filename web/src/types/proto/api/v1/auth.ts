// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               unknown
// source: api/v1/auth.proto

/* eslint-disable */
import { Empty } from "../../google/protobuf/empty";
import { GetAuthStatusRequest, SignInRequest, SignOutRequest, SignUpRequest } from "./auth/auth";
import { User } from "./system/user";

export const protobufPackage = "api.v1";

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
