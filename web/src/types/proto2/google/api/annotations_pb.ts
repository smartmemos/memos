// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-es v2.5.1 with parameter "target=ts"
// @generated from file google/api/annotations.proto (package google.api, syntax proto3)
/* eslint-disable */

import type { GenExtension, GenFile } from "@bufbuild/protobuf/codegenv2";
import { extDesc, fileDesc } from "@bufbuild/protobuf/codegenv2";
import type { HttpRule } from "./http_pb";
import { file_google_api_http } from "./http_pb";
import type { MethodOptions } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_descriptor } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file google/api/annotations.proto.
 */
export const file_google_api_annotations: GenFile = /*@__PURE__*/
  fileDesc("Chxnb29nbGUvYXBpL2Fubm90YXRpb25zLnByb3RvEgpnb29nbGUuYXBpOksKBGh0dHASHi5nb29nbGUucHJvdG9idWYuTWV0aG9kT3B0aW9ucxiwyrwiIAEoCzIULmdvb2dsZS5hcGkuSHR0cFJ1bGVSBGh0dHBCbgoOY29tLmdvb2dsZS5hcGlCEEFubm90YXRpb25zUHJvdG9QAVpBZ29vZ2xlLmdvbGFuZy5vcmcvZ2VucHJvdG8vZ29vZ2xlYXBpcy9hcGkvYW5ub3RhdGlvbnM7YW5ub3RhdGlvbnOiAgRHQVBJYgZwcm90bzM", [file_google_api_http, file_google_protobuf_descriptor]);

/**
 * See `HttpRule`.
 *
 * @generated from extension: google.api.HttpRule http = 72295728;
 */
export const http: GenExtension<MethodOptions, HttpRule> = /*@__PURE__*/
  extDesc(file_google_api_annotations, 0);

