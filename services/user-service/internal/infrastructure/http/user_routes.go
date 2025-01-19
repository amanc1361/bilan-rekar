package http

// UserRoutes defines routes for user-related operations
func UserRoutes(handler *UserHandler) []Route {
	return []Route{
		// GetUsers godoc
		// @Summary Get all users
		// @Description Retrieve a list of all users
		// @Tags Users
		// @Produce json
		// @Success 200 {array} models.User
		// @Failure 500 {object} gin.H
		// @Router /users [get]
		{
			Method:      "GET",
			Path:        "/users",
			HandlerFunc: handler.GetUsers,
		},
		// GetUserByID godoc
		// @Summary Get a user by ID
		// @Description Retrieve details of a single user by ID
		// @Tags Users
		// @Produce json
		// @Param id path int true "User ID"
		// @Success 200 {object} models.User
		// @Failure 404 {object} gin.H
		// @Router /users/{id} [get]
		{
			Method:      "GET",
			Path:        "/users/:id",
			HandlerFunc: handler.GetUserByID,
		},
	}
}
