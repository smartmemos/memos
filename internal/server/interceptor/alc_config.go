package interceptor

var authenticationAllowlistMethods = map[string]bool{
	// "/api.v1.AuthService/GetAuthStatus": true,
	"/api.v1.AuthService/SignIn":                   true,
	"/api.v1.WorkspaceService/GetWorkspaceProfile": true,
	"/api.v1.WorkspaceService/GetWorkspaceSetting": true,
	"/api.v1.AuthService/SignOut":                  true,
	"/api.v1.AuthService/SignUp":                   true,
}

// isUnauthorizeAllowedMethod returns whether the method is exempted from authentication.
func isUnauthorizeAllowedMethod(fullMethodName string) bool {
	return authenticationAllowlistMethods[fullMethodName]
}

var allowedMethodsOnlyForAdmin = map[string]bool{
	"/api.v1.UserService/CreateUser": true,
}

// isOnlyForAdminAllowedMethod returns true if the method is allowed to be called only by admin.
func isOnlyForAdminAllowedMethod(methodName string) bool {
	return allowedMethodsOnlyForAdmin[methodName]
}
