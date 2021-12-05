package v1

import (
	"context"
	"net/http"

	"chegirma_api_gateway/ek_variables/ek_setting_service"
	"chegirma_api_gateway/genproto/setting_service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Router /v1/user [post]
// @Summary Create user
// @Description API for creating user
// @Tags user
// @Accept json
// @Produce json
// @Param user body ek_setting_service.CreateUpdateUserSwag  true "User"
// @Success 201 {object} setting_service.CreatedResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		user   ek_setting_service.CreateUpdateUserSwag
		userID = primitive.NewObjectID().Hex()
		// userInfo, _ = h.UserInfo(c, true)
	)

	if err := c.BindJSON(&user); HandleHTTPError(c, http.StatusBadRequest, "DiscussionLogicService.Action.Create.BindingAction", err) {
		return
	}
	resp, err := h.services.UserService().Create(
		context.Background(),
		&setting_service.User{
			Id:       userID,
			Name:     user.Name,
			Link:     user.Link,
			Password: user.Password,
			Login:    user.Login,
			ImageUrl: user.ImageUrl,
		},
	)

	if HandleRPCError(c, "SettingService.CreateUser.InterService", err) {
		return
	}
	// actionHistory := &ek_analytic_service.ActionHistory{
	// 	UserID:         userInfo.ID,
	// 	UserUniqueName: userInfo.Login,
	// 	Action:         ek_variables.EntityCreated,
	// 	EntityID:       userID,
	// 	EntityName:     "entity",
	// }
	// h.CreateActionHistory(c, actionHistory)

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/user/{user_id} [get]
// @Summary Get User
// @Description API for getting user
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string  true "user_id"
// @Success 200 {object} setting_service.User
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) GetUser(c *gin.Context) {
	var (
		userID = c.Param("user_id")
		_, err = primitive.ObjectIDFromHex(userID)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "SettingService.GetUser.ParseUserID", err) {
		return
	}
	user, err := h.services.UserService().Get(context.Background(), &setting_service.UserGetReq{
		Id: userID,
	})

	if HandleRPCError(c, "SettingService.GetUser.InterService", err) {
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Router /v1/user [get]
// @Summary Getting All users
// @Description API for getting all useres
// @Tags user
// @Accept json
// @Produce json
// @Param name query string  false "name"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} setting_service.GetAllUsersResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	var (
		name = c.Query("name")
	)
	page, err := ParseQueryParam(c, h.log, "page", "1")
	if err != nil {
		return
	}

	limit, err := ParseQueryParam(c, h.log, "limit", "20")
	if err != nil {
		return
	}
	users, err := h.services.UserService().GetAll(
		context.Background(),
		&setting_service.GetAllUsersRequest{
			Name:  name,
			Page:  uint32(page),
			Limit: uint32(limit),
		})

	if HandleRPCError(c, "SettingService.GetAllUsers.InternalService", err) {
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Router /v1/user/{user_id} [put]
// @Summary Update user
// @Description API for updating user
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string  true "user_id"
// @Param user body ek_setting_service.CreateUpdateUserSwag true "user"
// @Success 200 {object} setting_service.CreatedResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		user   setting_service.User
		userID = c.Param("user_id")
		_, err = primitive.ObjectIDFromHex(userID)
	)
	if HandleHTTPError(c, http.StatusBadRequest, "SettingService.UpdateUser.ParsingObjectID", err) {
		return
	}
	if err = c.ShouldBindJSON(&user); HandleHTTPError(c, http.StatusBadRequest, "SettingService.UpdateUser.InvalidRequest", err) {
		return
	}
	user.Id = userID

	resp, err := h.services.UserService().Update(
		context.Background(),
		&user)

	if HandleRPCError(c, "SettingService.UpdateUser.InternalService", err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /v1/user-exists	[post]
// @Summary Get User
// @Description API for getting user
// @Tags user
// @Accept json
// @Produce json
// @Param login query string  true "login"
// @Param password query string  true "password"
// @Success 200 {object} setting_service.UserExistsResponse
// @Failure 400 {object} ek_variables.FailureResponse
// @Failure 404 {object} ek_variables.FailureResponse
// @Failure 500 {object} ek_variables.FailureResponse
// @Failure 503 {object} ek_variables.FailureResponse
func (h *handlerV1) UserExists(c *gin.Context) {
	var (
		login    = c.Query("login")
		password = c.Query("password")
	)

	user, err := h.services.UserService().UserExists(context.Background(), &setting_service.UserExistsRequest{
		Login:    login,
		Password: password,
	})

	if HandleRPCError(c, "SettingService.UserExists.InterService", err) {
		return
	}

	c.JSON(http.StatusOK, user)
}
