package interceptor

var authenticationAllowlistMethods = map[string]bool{
	"/memos.api.v1.AuthService/GetAuthStatus": true,
	"/memos.api.v1.AuthService/SignIn":        true,
	"/memos.api.v1.AuthService/SignOut":       true,
	"/memos.api.v1.AuthService/SignUp":        true,
}

// isUnauthorizeAllowedMethod returns whether the method is exempted from authentication.
func isUnauthorizeAllowedMethod(fullMethodName string) bool {
	return authenticationAllowlistMethods[fullMethodName]
}

var allowedMethodsOnlyForAdmin = map[string]bool{
	"/memos.api.v1.UserService/CreateUser": true,
}

// isOnlyForAdminAllowedMethod returns true if the method is allowed to be called only by admin.
func isOnlyForAdminAllowedMethod(methodName string) bool {
	return allowedMethodsOnlyForAdmin[methodName]
}
