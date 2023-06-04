// handlers.go

package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"packages/services"
)

// POST /user - Create User
func createUserHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Call the CreateUser function from the UserService
	err := userService.CreateUser(name, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GET /user/:id - Get User Info
func getUserHandler(c *gin.Context) {
	userID := c.Param("id")

	// Convert the user ID to an integer
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the GetUserByID function from the UserService
	user, err := userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE /user/:id - Remove User
func deleteUserHandler(c *gin.Context) {
	userID := c.Param("id")

	// Convert the user ID to an integer
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the DeleteUserByID function from the UserService
	err = userService.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed successfully"})
}

// PATCH /user/:id - Edit User
func editUserHandler(c *gin.Context) {
	userID := c.Param("id")

	// Convert the user ID to an integer
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get the updated user data from the request payload
	var updatedUser User
	err = c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the EditUserByID function from the UserService
	err = userService.EditUserByID(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User edited successfully"})
}

// GET /users - List All Users
func listUsersHandler(c *gin.Context) {
	// Call the ListUsers function from the UserService
	users, err := userService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

