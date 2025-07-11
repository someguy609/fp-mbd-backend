package controller

import (
	"net/http"

	"fp_mbd/dto"

	"fp_mbd/service"

	"fp_mbd/utils"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Me(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
		GetUserByUserId(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		GetUserProjects(ctx *gin.Context)
		// Logout(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.Register(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetAllUser(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user_id := ctx.GetString("user_id")
	if user_id == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.GetAllUserWithPagination(ctx.Request.Context(), req, user_id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}
func (c *userController) GetUserByUserId(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	if userId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.GetUserByUserId(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Me(ctx *gin.Context) {

	println("Me called")

	userId := ctx.GetString("user_id")
	if userId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	println("userId:", userId)

	result, err := c.userService.GetUserByUserId(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Login(ctx *gin.Context) {

	println("Login called")

	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.userService.Login(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, res)
}

// func (c *userController) Logout(ctx *gin.Context) {
// 	userId := ctx.GetString("user_id")

// 	if err := c.userService.Logout(ctx.Request.Context(), userId); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, nil)
// 	ctx.JSON(http.StatusOK, res)
// }

// func (c *userController) SendVerificationEmail(ctx *gin.Context) {
// 	var req dto.SendVerificationEmailRequest
// 	if err := ctx.ShouldBind(&req); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	err := c.userService.SendVerificationEmail(ctx.Request.Context(), req)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS, nil)
// 	ctx.JSON(http.StatusOK, res)
// }

// func (c *userController) VerifyEmail(ctx *gin.Context) {
// 	var req dto.VerifyEmailRequest
// 	if err := ctx.ShouldBind(&req); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	result, err := c.userService.VerifyEmail(ctx.Request.Context(), req)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_EMAIL, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_VERIFY_EMAIL, result)
// 	ctx.JSON(http.StatusOK, res)
// }

func (c *userController) Update(ctx *gin.Context) {
	var req dto.UserUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userId := ctx.GetString("user_id")

	if userId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.Update(ctx.Request.Context(), req, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Delete(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	if userId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userIdParam := ctx.Param("user_id")
	if userIdParam == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.userService.Delete(ctx.Request.Context(), userId, userIdParam); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	ctx.JSON(http.StatusOK, res)
}

// func (c *userController) Refresh(ctx *gin.Context) {
// 	var req dto.RefreshTokenRequest
// 	if err := ctx.ShouldBind(&req); err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	result, err := c.userService.RefreshToken(ctx.Request.Context(), req)
// 	if err != nil {
// 		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REFRESH_TOKEN, err.Error(), nil)
// 		ctx.JSON(http.StatusUnauthorized, res)
// 		return
// 	}

// 	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REFRESH_TOKEN, result)
// 	ctx.JSON(http.StatusOK, res)
// }

func (c *userController) GetUserProjects(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	if userId == "" {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, "user_id is required", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.GetUserProjects(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_PROJECTS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER_PROJECTS, result)
	ctx.JSON(http.StatusOK, res)
}
