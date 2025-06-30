package controller

import (
	"bankSystem/dto"
	"bankSystem/mapper"
	"bankSystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// @Summary      Create new user
// @Description  Create a new user with login, name, sex and hair color
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserRequest  true  "User to create"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router       /user/create [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	sex, hair, err := mapper.StringToEnum(req.Sex, req.HairColor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid arguments sex or hair color"})
		return
	}

	user, err := uc.userService.NewUser(req.Login, req.Name, sex, hair)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user_login": user.Login})
}

// @Summary      Add friend
// @Description  Add a friend using user login and friend login
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        friendship  body      dto.FriendsRequest  true  "Friendship info"
// @Success      200         {object}  map[string]interface{}
// @Failure      400         {object}  map[string]interface{}
// @Failure      409         {object}  map[string]interface{}
// @Router       /user/add_friend [post]
func (uc *UserController) AddFriend(c *gin.Context) {
	var req dto.FriendsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := uc.userService.AddFriend(req.UserLogin, req.FriendLogin); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "users are friends"})
}

// @Summary      Delete user's friend
// @Description  Delete user's friend
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        friendship  body      dto.FriendsRequest  true  "Friendship info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router       /user/remove_friend [post]
func (uc *UserController) RemoveFriend(c *gin.Context) {
	var req dto.FriendsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := uc.userService.RemoveFriend(req.UserLogin, req.FriendLogin); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "users have deleted"})
}

// @Summary      Get user
// @Description  Get user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param         login  query     string  true  "User login"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router       /user/get_user [get]
func (uc *UserController) GetUser(c *gin.Context) {
	userLogin := c.Query("login")
	if userLogin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login query parameter is required"})
		return
	}

	user, err := uc.userService.GetUser(userLogin)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"login": user.Login, "name": user.Name, "sex": user.Sex,
		"hair_color": user.HairColor, "friends": user.Friends, "accounts": user.Accounts})
}

// @Summary      Delete user
// @Description  Delete user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param         login  query     string  true  "User deleted"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router       /user/delete [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	userLogin := c.Query("login")
	if userLogin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login query parameter is required"})
		return
	}

	if err := uc.userService.DeleteUser(userLogin); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	return
}
