package dto

import (
	"errors"

	"fp_mbd/entity"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER      = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER      = "failed get list user"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "token not found"
	MESSAGE_FAILED_GET_USER           = "failed get user"
	MESSAGE_FAILED_LOGIN              = "failed login"
	MESSAGE_FAILED_UPDATE_USER        = "failed update user"
	MESSAGE_FAILED_DELETE_USER        = "failed delete user"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS      = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL       = "failed verify email"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"

	MESSAGE_FAILED_GET_USER_PROJECTS  = "failed get user projects"
	MESSAGE_SUCCESS_GET_USER_PROJECTS = "success get user projects"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByEmail         = errors.New("failed to get user by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrUnauthorizedAccess     = errors.New("unauthorized access")
)

type (
	UserCreateRequest struct {
		UserID      string `json:"user_id" form:"user_id" binding:"required,min=2,max=15"`
		Name        string `json:"name" form:"name" binding:"required,min=2,max=100"`
		Email       string `json:"email" form:"email" binding:"required,email"`
		Password    string `json:"password" form:"password" binding:"required,min=8"`
		ContactInfo string `json:"contact_info" form:"contact_info" binding:"omitempty,min=8,max=100"`
	}

	UserResponse struct {
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		ContactInfo string `json:"contact_info"`
		Role        string `json:"role"`
	}

	UserPaginationResponse struct {
		Data []UserResponse `json:"data"`
		PaginationResponse
	}

	GetAllUserRepositoryResponse struct {
		Users []entity.User `json:"users"`
		PaginationResponse
	}

	// bagian di bawah belum ku update

	UserUpdateRequest struct {
		Name        string `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
		Email       string `json:"email" form:"email" binding:"omitempty,email"`
		ContactInfo string `json:"contact_info" form:"contact_info" binding:"omitempty,min=8,max=100"`
	}

	UserUpdateResponse struct {
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		Role        string `json:"role"`
		ContactInfo string `json:"contact_info"`
	}

	SendVerificationEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
)
