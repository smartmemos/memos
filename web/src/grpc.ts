import type { Interceptor, StreamResponse, UnaryResponse } from '@connectrpc/connect';
import { Code, ConnectError, createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import type { DescMessage } from '@bufbuild/protobuf';
import { AuthService } from './types/proto2/api/v2/auth_pb';
import { UserService } from './types/proto2/api/v2/user_pb';
import { WorkspaceService } from './types/proto2/api/v2/workspace_pb';
import { InboxService } from './types/proto2/api/v2/inbox_pb';
import { MemoService } from './types/proto2/api/v2/memo_pb';

// const grpcServiceUrl = import.meta.env.VITE_PUBLIC_GRPC_SERVICE_URL || '/grpc';
const grpcServiceUrl = 'http://localhost:3001';

const logInterceptor: Interceptor = next => async req => {
  console.log('request message: ', req.message);
  const res = await next(req);
  // console.log('response: ', res);
  return res;
};


const authInterceptor: Interceptor =
  next =>
    async (req): Promise<UnaryResponse<DescMessage, DescMessage> | StreamResponse<DescMessage, DescMessage>> => {
      const doRequest: () => Promise<
        UnaryResponse<DescMessage, DescMessage> | StreamResponse<DescMessage, DescMessage>
      > = async () => {
        //   const token = getToken();
        //   if (token) {
        //     req.header.set('Authorization', `Bearer ${token}`);
        //   }
        return await next(req);
      };
      try {
        return await doRequest();
      } catch (err) {
        // if (
        //   err instanceof ConnectError &&
        //   err.code === Code.Unauthenticated &&
        //   !(req.method.name === 'RefreshToken' && req.method.parent.typeName === 'api.v1.AuthService') &&
        //   !req.stream
        // ) {
        //   const success = await handleRefreshToken();
        //   if (!success) throw err;
        //   return await doRequest();
        // }
        throw err;
      }
    };

const transport = createConnectTransport({
  baseUrl: grpcServiceUrl,
  interceptors: [logInterceptor, authInterceptor],
  useBinaryFormat: false
  // useBinaryFormat: true
});

export const authServiceClient = createClient(AuthService, transport);
export const userServiceClient = createClient(UserService, transport);
export const memoServiceClient = createClient(MemoService, transport);
export const workspaceServiceClient = createClient(WorkspaceService, transport);
export const inboxServiceClient = createClient(InboxService, transport);
