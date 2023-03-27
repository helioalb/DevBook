package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Handler:               controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Handler:               controllers.GetAllUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Handler:               controllers.GetUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Handler:               controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Handler:               controllers.DeleteUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Handler:               controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		Handler:               controllers.UnfollowUser,
		RequireAuthentication: true,
	},
}
