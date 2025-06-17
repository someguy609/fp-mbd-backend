package controller

import (
	"net/http"
	"strconv"

	"fp_mbd/dto"

	"fp_mbd/service"

	"fp_mbd/utils"

	"github.com/gin-gonic/gin"
)

type (
	DocumentController interface {
		Upload(ctx *gin.Context)
		GetDocument(ctx *gin.Context)
		GetAllDocument(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	documentController struct {
		documentService service.DocumentService
	}
)

func NewDocumentController(ds service.DocumentService) DocumentController {
	return &documentController{
		documentService: ds,
	}
}

func (c *documentController) Upload(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	var document dto.UploadDocumentRequest
	if err := ctx.ShouldBind(&document); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	document.File = file

	result, err := c.documentService.Upload(ctx.Request.Context(), document, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPLOAD_DOCUMENT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPLOAD_DOCUMENT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *documentController) GetAllDocument(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.documentService.GetAllDocumentWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_DOCUMENT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_DOCUMENT,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *documentController) GetDocument(ctx *gin.Context) {
	// userId := ctx.MustGet("user_id").(string)

	documentId_, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		// todo: change the error message
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	documentId := uint(documentId_)

	result, err := c.documentService.GetDocumentById(ctx.Request.Context(), documentId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DOCUMENT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DOCUMENT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *documentController) Update(ctx *gin.Context) {
	var req dto.DocumentUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// userId := ctx.MustGet("user_id").(string)
	result, err := c.documentService.Update(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_DOCUMENT, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_DOCUMENT, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *documentController) Delete(ctx *gin.Context) {
	documentId_, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		// todo: change the error message
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	documentId := uint(documentId_)
	// userId := ctx.MustGet("user_id").(string)

	if err := c.documentService.Delete(ctx.Request.Context(), documentId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_DOCUMENT, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_DOCUMENT, nil)
	ctx.JSON(http.StatusOK, res)
}
