// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/workspace.proto

package v1

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	workspace "github.com/smartmemos/memos/internal/proto/model/workspace"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// WorkspaceServiceName is the fully-qualified name of the WorkspaceService service.
	WorkspaceServiceName = "api.v1.WorkspaceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WorkspaceServiceGetWorkspaceProfileProcedure is the fully-qualified name of the
	// WorkspaceService's GetWorkspaceProfile RPC.
	WorkspaceServiceGetWorkspaceProfileProcedure = "/api.v1.WorkspaceService/GetWorkspaceProfile"
	// WorkspaceServiceGetWorkspaceSettingProcedure is the fully-qualified name of the
	// WorkspaceService's GetWorkspaceSetting RPC.
	WorkspaceServiceGetWorkspaceSettingProcedure = "/api.v1.WorkspaceService/GetWorkspaceSetting"
	// WorkspaceServiceSetWorkspaceSettingProcedure is the fully-qualified name of the
	// WorkspaceService's SetWorkspaceSetting RPC.
	WorkspaceServiceSetWorkspaceSettingProcedure = "/api.v1.WorkspaceService/SetWorkspaceSetting"
	// WorkspaceServiceListIdentityProvidersProcedure is the fully-qualified name of the
	// WorkspaceService's ListIdentityProviders RPC.
	WorkspaceServiceListIdentityProvidersProcedure = "/api.v1.WorkspaceService/ListIdentityProviders"
	// WorkspaceServiceGetIdentityProviderProcedure is the fully-qualified name of the
	// WorkspaceService's GetIdentityProvider RPC.
	WorkspaceServiceGetIdentityProviderProcedure = "/api.v1.WorkspaceService/GetIdentityProvider"
	// WorkspaceServiceListInboxesProcedure is the fully-qualified name of the WorkspaceService's
	// ListInboxes RPC.
	WorkspaceServiceListInboxesProcedure = "/api.v1.WorkspaceService/ListInboxes"
	// WorkspaceServiceUpdateInboxProcedure is the fully-qualified name of the WorkspaceService's
	// UpdateInbox RPC.
	WorkspaceServiceUpdateInboxProcedure = "/api.v1.WorkspaceService/UpdateInbox"
	// WorkspaceServiceDeleteInboxProcedure is the fully-qualified name of the WorkspaceService's
	// DeleteInbox RPC.
	WorkspaceServiceDeleteInboxProcedure = "/api.v1.WorkspaceService/DeleteInbox"
)

// WorkspaceServiceClient is a client for the api.v1.WorkspaceService service.
type WorkspaceServiceClient interface {
	// GetProfile returns the workspace profile.
	GetWorkspaceProfile(context.Context, *connect.Request[GetWorkspaceProfileRequest]) (*connect.Response[workspace.Profile], error)
	// GetWorkspaceSetting returns the setting by name.
	GetWorkspaceSetting(context.Context, *connect.Request[GetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error)
	// SetWorkspaceSetting updates the setting.
	SetWorkspaceSetting(context.Context, *connect.Request[SetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error)
	// ListIdentityProviders lists identity providers.
	ListIdentityProviders(context.Context, *connect.Request[ListIdentityProvidersRequest]) (*connect.Response[ListIdentityProvidersResponse], error)
	// GetIdentityProvider gets an identity provider.
	GetIdentityProvider(context.Context, *connect.Request[GetIdentityProviderRequest]) (*connect.Response[workspace.IdentityProvider], error)
	// ListInboxes lists inboxes for a user.
	ListInboxes(context.Context, *connect.Request[ListInboxesRequest]) (*connect.Response[ListInboxesResponse], error)
	// UpdateInbox updates an inbox.
	UpdateInbox(context.Context, *connect.Request[UpdateInboxRequest]) (*connect.Response[workspace.Inbox], error)
	// DeleteInbox deletes an inbox.
	DeleteInbox(context.Context, *connect.Request[DeleteInboxRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewWorkspaceServiceClient constructs a client for the api.v1.WorkspaceService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkspaceServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) WorkspaceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	workspaceServiceMethods := File_api_v1_workspace_proto.Services().ByName("WorkspaceService").Methods()
	return &workspaceServiceClient{
		getWorkspaceProfile: connect.NewClient[GetWorkspaceProfileRequest, workspace.Profile](
			httpClient,
			baseURL+WorkspaceServiceGetWorkspaceProfileProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("GetWorkspaceProfile")),
			connect.WithClientOptions(opts...),
		),
		getWorkspaceSetting: connect.NewClient[GetWorkspaceSettingRequest, workspace.Setting](
			httpClient,
			baseURL+WorkspaceServiceGetWorkspaceSettingProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("GetWorkspaceSetting")),
			connect.WithClientOptions(opts...),
		),
		setWorkspaceSetting: connect.NewClient[SetWorkspaceSettingRequest, workspace.Setting](
			httpClient,
			baseURL+WorkspaceServiceSetWorkspaceSettingProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("SetWorkspaceSetting")),
			connect.WithClientOptions(opts...),
		),
		listIdentityProviders: connect.NewClient[ListIdentityProvidersRequest, ListIdentityProvidersResponse](
			httpClient,
			baseURL+WorkspaceServiceListIdentityProvidersProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("ListIdentityProviders")),
			connect.WithClientOptions(opts...),
		),
		getIdentityProvider: connect.NewClient[GetIdentityProviderRequest, workspace.IdentityProvider](
			httpClient,
			baseURL+WorkspaceServiceGetIdentityProviderProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("GetIdentityProvider")),
			connect.WithClientOptions(opts...),
		),
		listInboxes: connect.NewClient[ListInboxesRequest, ListInboxesResponse](
			httpClient,
			baseURL+WorkspaceServiceListInboxesProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("ListInboxes")),
			connect.WithClientOptions(opts...),
		),
		updateInbox: connect.NewClient[UpdateInboxRequest, workspace.Inbox](
			httpClient,
			baseURL+WorkspaceServiceUpdateInboxProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("UpdateInbox")),
			connect.WithClientOptions(opts...),
		),
		deleteInbox: connect.NewClient[DeleteInboxRequest, emptypb.Empty](
			httpClient,
			baseURL+WorkspaceServiceDeleteInboxProcedure,
			connect.WithSchema(workspaceServiceMethods.ByName("DeleteInbox")),
			connect.WithClientOptions(opts...),
		),
	}
}

// workspaceServiceClient implements WorkspaceServiceClient.
type workspaceServiceClient struct {
	getWorkspaceProfile   *connect.Client[GetWorkspaceProfileRequest, workspace.Profile]
	getWorkspaceSetting   *connect.Client[GetWorkspaceSettingRequest, workspace.Setting]
	setWorkspaceSetting   *connect.Client[SetWorkspaceSettingRequest, workspace.Setting]
	listIdentityProviders *connect.Client[ListIdentityProvidersRequest, ListIdentityProvidersResponse]
	getIdentityProvider   *connect.Client[GetIdentityProviderRequest, workspace.IdentityProvider]
	listInboxes           *connect.Client[ListInboxesRequest, ListInboxesResponse]
	updateInbox           *connect.Client[UpdateInboxRequest, workspace.Inbox]
	deleteInbox           *connect.Client[DeleteInboxRequest, emptypb.Empty]
}

// GetWorkspaceProfile calls api.v1.WorkspaceService.GetWorkspaceProfile.
func (c *workspaceServiceClient) GetWorkspaceProfile(ctx context.Context, req *connect.Request[GetWorkspaceProfileRequest]) (*connect.Response[workspace.Profile], error) {
	return c.getWorkspaceProfile.CallUnary(ctx, req)
}

// GetWorkspaceSetting calls api.v1.WorkspaceService.GetWorkspaceSetting.
func (c *workspaceServiceClient) GetWorkspaceSetting(ctx context.Context, req *connect.Request[GetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error) {
	return c.getWorkspaceSetting.CallUnary(ctx, req)
}

// SetWorkspaceSetting calls api.v1.WorkspaceService.SetWorkspaceSetting.
func (c *workspaceServiceClient) SetWorkspaceSetting(ctx context.Context, req *connect.Request[SetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error) {
	return c.setWorkspaceSetting.CallUnary(ctx, req)
}

// ListIdentityProviders calls api.v1.WorkspaceService.ListIdentityProviders.
func (c *workspaceServiceClient) ListIdentityProviders(ctx context.Context, req *connect.Request[ListIdentityProvidersRequest]) (*connect.Response[ListIdentityProvidersResponse], error) {
	return c.listIdentityProviders.CallUnary(ctx, req)
}

// GetIdentityProvider calls api.v1.WorkspaceService.GetIdentityProvider.
func (c *workspaceServiceClient) GetIdentityProvider(ctx context.Context, req *connect.Request[GetIdentityProviderRequest]) (*connect.Response[workspace.IdentityProvider], error) {
	return c.getIdentityProvider.CallUnary(ctx, req)
}

// ListInboxes calls api.v1.WorkspaceService.ListInboxes.
func (c *workspaceServiceClient) ListInboxes(ctx context.Context, req *connect.Request[ListInboxesRequest]) (*connect.Response[ListInboxesResponse], error) {
	return c.listInboxes.CallUnary(ctx, req)
}

// UpdateInbox calls api.v1.WorkspaceService.UpdateInbox.
func (c *workspaceServiceClient) UpdateInbox(ctx context.Context, req *connect.Request[UpdateInboxRequest]) (*connect.Response[workspace.Inbox], error) {
	return c.updateInbox.CallUnary(ctx, req)
}

// DeleteInbox calls api.v1.WorkspaceService.DeleteInbox.
func (c *workspaceServiceClient) DeleteInbox(ctx context.Context, req *connect.Request[DeleteInboxRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteInbox.CallUnary(ctx, req)
}

// WorkspaceServiceHandler is an implementation of the api.v1.WorkspaceService service.
type WorkspaceServiceHandler interface {
	// GetProfile returns the workspace profile.
	GetWorkspaceProfile(context.Context, *connect.Request[GetWorkspaceProfileRequest]) (*connect.Response[workspace.Profile], error)
	// GetWorkspaceSetting returns the setting by name.
	GetWorkspaceSetting(context.Context, *connect.Request[GetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error)
	// SetWorkspaceSetting updates the setting.
	SetWorkspaceSetting(context.Context, *connect.Request[SetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error)
	// ListIdentityProviders lists identity providers.
	ListIdentityProviders(context.Context, *connect.Request[ListIdentityProvidersRequest]) (*connect.Response[ListIdentityProvidersResponse], error)
	// GetIdentityProvider gets an identity provider.
	GetIdentityProvider(context.Context, *connect.Request[GetIdentityProviderRequest]) (*connect.Response[workspace.IdentityProvider], error)
	// ListInboxes lists inboxes for a user.
	ListInboxes(context.Context, *connect.Request[ListInboxesRequest]) (*connect.Response[ListInboxesResponse], error)
	// UpdateInbox updates an inbox.
	UpdateInbox(context.Context, *connect.Request[UpdateInboxRequest]) (*connect.Response[workspace.Inbox], error)
	// DeleteInbox deletes an inbox.
	DeleteInbox(context.Context, *connect.Request[DeleteInboxRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewWorkspaceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkspaceServiceHandler(svc WorkspaceServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	workspaceServiceMethods := File_api_v1_workspace_proto.Services().ByName("WorkspaceService").Methods()
	workspaceServiceGetWorkspaceProfileHandler := connect.NewUnaryHandler(
		WorkspaceServiceGetWorkspaceProfileProcedure,
		svc.GetWorkspaceProfile,
		connect.WithSchema(workspaceServiceMethods.ByName("GetWorkspaceProfile")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceGetWorkspaceSettingHandler := connect.NewUnaryHandler(
		WorkspaceServiceGetWorkspaceSettingProcedure,
		svc.GetWorkspaceSetting,
		connect.WithSchema(workspaceServiceMethods.ByName("GetWorkspaceSetting")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceSetWorkspaceSettingHandler := connect.NewUnaryHandler(
		WorkspaceServiceSetWorkspaceSettingProcedure,
		svc.SetWorkspaceSetting,
		connect.WithSchema(workspaceServiceMethods.ByName("SetWorkspaceSetting")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceListIdentityProvidersHandler := connect.NewUnaryHandler(
		WorkspaceServiceListIdentityProvidersProcedure,
		svc.ListIdentityProviders,
		connect.WithSchema(workspaceServiceMethods.ByName("ListIdentityProviders")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceGetIdentityProviderHandler := connect.NewUnaryHandler(
		WorkspaceServiceGetIdentityProviderProcedure,
		svc.GetIdentityProvider,
		connect.WithSchema(workspaceServiceMethods.ByName("GetIdentityProvider")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceListInboxesHandler := connect.NewUnaryHandler(
		WorkspaceServiceListInboxesProcedure,
		svc.ListInboxes,
		connect.WithSchema(workspaceServiceMethods.ByName("ListInboxes")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceUpdateInboxHandler := connect.NewUnaryHandler(
		WorkspaceServiceUpdateInboxProcedure,
		svc.UpdateInbox,
		connect.WithSchema(workspaceServiceMethods.ByName("UpdateInbox")),
		connect.WithHandlerOptions(opts...),
	)
	workspaceServiceDeleteInboxHandler := connect.NewUnaryHandler(
		WorkspaceServiceDeleteInboxProcedure,
		svc.DeleteInbox,
		connect.WithSchema(workspaceServiceMethods.ByName("DeleteInbox")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.WorkspaceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WorkspaceServiceGetWorkspaceProfileProcedure:
			workspaceServiceGetWorkspaceProfileHandler.ServeHTTP(w, r)
		case WorkspaceServiceGetWorkspaceSettingProcedure:
			workspaceServiceGetWorkspaceSettingHandler.ServeHTTP(w, r)
		case WorkspaceServiceSetWorkspaceSettingProcedure:
			workspaceServiceSetWorkspaceSettingHandler.ServeHTTP(w, r)
		case WorkspaceServiceListIdentityProvidersProcedure:
			workspaceServiceListIdentityProvidersHandler.ServeHTTP(w, r)
		case WorkspaceServiceGetIdentityProviderProcedure:
			workspaceServiceGetIdentityProviderHandler.ServeHTTP(w, r)
		case WorkspaceServiceListInboxesProcedure:
			workspaceServiceListInboxesHandler.ServeHTTP(w, r)
		case WorkspaceServiceUpdateInboxProcedure:
			workspaceServiceUpdateInboxHandler.ServeHTTP(w, r)
		case WorkspaceServiceDeleteInboxProcedure:
			workspaceServiceDeleteInboxHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWorkspaceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkspaceServiceHandler struct{}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceProfile(context.Context, *connect.Request[GetWorkspaceProfileRequest]) (*connect.Response[workspace.Profile], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.GetWorkspaceProfile is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetWorkspaceSetting(context.Context, *connect.Request[GetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.GetWorkspaceSetting is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) SetWorkspaceSetting(context.Context, *connect.Request[SetWorkspaceSettingRequest]) (*connect.Response[workspace.Setting], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.SetWorkspaceSetting is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) ListIdentityProviders(context.Context, *connect.Request[ListIdentityProvidersRequest]) (*connect.Response[ListIdentityProvidersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.ListIdentityProviders is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) GetIdentityProvider(context.Context, *connect.Request[GetIdentityProviderRequest]) (*connect.Response[workspace.IdentityProvider], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.GetIdentityProvider is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) ListInboxes(context.Context, *connect.Request[ListInboxesRequest]) (*connect.Response[ListInboxesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.ListInboxes is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) UpdateInbox(context.Context, *connect.Request[UpdateInboxRequest]) (*connect.Response[workspace.Inbox], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.UpdateInbox is not implemented"))
}

func (UnimplementedWorkspaceServiceHandler) DeleteInbox(context.Context, *connect.Request[DeleteInboxRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.WorkspaceService.DeleteInbox is not implemented"))
}
