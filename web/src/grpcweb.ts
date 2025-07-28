import { createChannel, createClientFactory, FetchTransport } from "nice-grpc-web";
import { ActivityServiceDefinition } from "./types/proto/api/v1/activity_service";
import { IdentityProviderServiceDefinition } from "./types/proto/api/v1/idp_service";
import { InboxServiceDefinition } from "./types/proto/api/v1/inbox_service";
import { MarkdownServiceDefinition } from "./types/proto/api/v1/markdown_service";
import { ResourceServiceDefinition } from "./types/proto/api/v1/resource_service";
import { WebhookServiceDefinition } from "./types/proto/api/v1/webhook_service";
import { WorkspaceSettingServiceDefinition } from "./types/proto/api/v1/workspace_setting_service";

const channel = createChannel(
  window.location.origin,
  FetchTransport({
    credentials: "include",
  }),
);

const clientFactory = createClientFactory();

export const workspaceSettingServiceClient = clientFactory.create(WorkspaceSettingServiceDefinition, channel);

export const resourceServiceClient = clientFactory.create(ResourceServiceDefinition, channel);

export const inboxServiceClient = clientFactory.create(InboxServiceDefinition, channel);

export const activityServiceClient = clientFactory.create(ActivityServiceDefinition, channel);

export const webhookServiceClient = clientFactory.create(WebhookServiceDefinition, channel);

export const markdownServiceClient = clientFactory.create(MarkdownServiceDefinition, channel);

export const identityProviderServiceClient = clientFactory.create(IdentityProviderServiceDefinition, channel);
