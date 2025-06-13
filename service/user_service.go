package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"fp_mbd/dto"
	"fp_mbd/entity"
	"fp_mbd/helpers"
	"fp_mbd/repository"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
		GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error)
		Update(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error)
		Delete(ctx context.Context, userId string) error
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.TokenResponse, error)
		// SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error
		// VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error)
		// Verify(ctx context.Context, req dto.UserLoginRequest) (dto.TokenResponse, error)
		// RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.TokenResponse, error)
	}

	userService struct {
		userRepo         repository.UserRepository
		refreshTokenRepo repository.RefreshTokenRepository
		jwtService       JWTService
		db               *gorm.DB
	}
)

func NewUserService(
	userRepo repository.UserRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	jwtService JWTService,
	db *gorm.DB,
) UserService {
	return &userService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtService:       jwtService,
		db:               db,
	}
}

const (
	LOCAL_URL          = "http://localhost:3000"
	VERIFY_EMAIL_ROUTE = "register/verify_email"
)

func SafeRollback(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
		// TODO: Do you think that we should panic here?
		// panic(r)
	}
}

func (s *userService) Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {

	_, flag, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.UserResponse{}, err
	}

	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	role := helpers.CheckRole(req.UserID)

	user := entity.User{
		Name:        req.Name,
		UserID:      req.UserID,
		Email:       req.Email,
		Role:        role,
		ContactInfo: req.ContactInfo,
		Password:    req.Password,
	}

	userReg, err := s.userRepo.Register(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		UserID:      userReg.UserID,
		Name:        userReg.Name,
		Email:       userReg.Email,
		Role:        userReg.Role,
		ContactInfo: userReg.ContactInfo,
	}, nil

}

// func (s *userService) SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error {
// 	user, err := s.userRepo.GetUserByEmail(ctx, nil, req.Email)
// 	if err != nil {
// 		return dto.ErrEmailNotFound
// 	}

// 	draftEmail, err := makeVerificationEmail(user.Email)
// 	if err != nil {
// 		return err
// 	}

// 	err = utils.SendMail(user.Email, draftEmail["subject"], draftEmail["body"])
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *userService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error) {
// 	decryptedToken, err := utils.AESDecrypt(req.Token)
// 	if err != nil {
// 		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
// 	}

// 	if !strings.Contains(decryptedToken, "_") {
// 		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
// 	}

// 	decryptedTokenSplit := strings.Split(decryptedToken, "_")
// 	email := decryptedTokenSplit[0]
// 	expired := decryptedTokenSplit[1]

// 	now := time.Now()
// 	expiredTime, err := time.Parse("2006-01-02 15:04:05", expired)
// 	if err != nil {
// 		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
// 	}

// 	if expiredTime.Sub(now) < 0 {
// 		return dto.VerifyEmailResponse{
// 			Email:      email,
// 			IsVerified: false,
// 		}, dto.ErrTokenExpired
// 	}

// 	user, err := s.userRepo.GetUserByEmail(ctx, nil, email)
// 	if err != nil {
// 		return dto.VerifyEmailResponse{}, dto.ErrUserNotFound
// 	}

// 	if user.IsVerified {
// 		return dto.VerifyEmailResponse{}, dto.ErrAccountAlreadyVerified
// 	}

// 	updatedUser, err := s.userRepo.Update(
// 		ctx, nil, entity.User{
// 			ID:         user.ID,
// 			IsVerified: true,
// 		},
// 	)
// 	if err != nil {
// 		return dto.VerifyEmailResponse{}, dto.ErrUpdateUser
// 	}

// 	return dto.VerifyEmailResponse{
// 		Email:      email,
// 		IsVerified: updatedUser.IsVerified,
// 	}, nil
// }

func (s *userService) GetAllUserWithPagination(
	ctx context.Context,
	req dto.PaginationRequest,
) (dto.UserPaginationResponse, error) {
	dataWithPaginate, err := s.userRepo.GetAllUserWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserPaginationResponse{}, err
	}

	var datas []dto.UserResponse
	for _, user := range dataWithPaginate.Users {
		data := dto.UserResponse{
			UserID:      user.UserID,
			Name:        user.Name,
			Email:       user.Email,
			Role:        user.Role,
			ContactInfo: user.ContactInfo,
		}

		datas = append(datas, data)
	}

	return dto.UserPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		UserID:      user.UserID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		ContactInfo: user.ContactInfo,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	emails, err := s.userRepo.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserByEmail
	}

	return dto.UserResponse{
		UserID:      emails.UserID,
		Name:        emails.Name,
		Email:       emails.Email,
		Role:        emails.Role,
		ContactInfo: emails.ContactInfo,
	}, nil
}

func (s *userService) Update(ctx context.Context, req dto.UserUpdateRequest, userId string) (
	dto.UserUpdateResponse,
	error,
) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUserNotFound
	}

	data := entity.User{
		UserID:      user.UserID,
		Name:        req.Name,
		Email:       req.Email,
		Role:        user.Role,
		ContactInfo: req.ContactInfo,
	}

	userUpdate, err := s.userRepo.Update(ctx, nil, data)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUpdateUser
	}

	return dto.UserUpdateResponse{
		UserID:      userUpdate.UserID,
		Name:        userUpdate.Name,
		Email:       userUpdate.Email,
		Role:        userUpdate.Role,
		ContactInfo: userUpdate.ContactInfo,
	}, nil
}

func (s *userService) Delete(ctx context.Context, userId string) error {
	tx := s.db.Begin()
	defer SafeRollback(tx)

	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.ErrUserNotFound
	}

	err = s.userRepo.Delete(ctx, nil, user.UserID)
	if err != nil {
		return dto.ErrDeleteUser
	}

	return nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.TokenResponse, error) {
	tx := s.db.Begin()
	defer SafeRollback(tx)

	user, err := s.userRepo.GetUserByEmail(ctx, tx, req.Email)
	if err != nil {
		tx.Rollback()
		return dto.TokenResponse{}, errors.New("invalid email or password")
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		tx.Rollback()
		return dto.TokenResponse{}, errors.New("invalid email or password")
	}

	accessToken := s.jwtService.GenerateAccessToken(user.UserID, user.Role)

	refreshTokenString, expiresAt := s.jwtService.GenerateRefreshToken()

	hashedToken, err := helpers.HashPassword(refreshTokenString)
	if err != nil {
		tx.Rollback()
		return dto.TokenResponse{}, err
	}

	if err := s.refreshTokenRepo.DeleteByUserID(ctx, tx, user.UserID); err != nil {
		tx.Rollback()
		return dto.TokenResponse{}, err
	}

	refreshToken := entity.RefreshToken{
		UserID:    user.UserID,
		Token:     hashedToken,
		ExpiresAt: expiresAt,
	}

	if _, err := s.refreshTokenRepo.Create(ctx, tx, refreshToken); err != nil {
		tx.Rollback()
		return dto.TokenResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return dto.TokenResponse{}, err
	}

	return dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
		Role:         user.Role,
	}, nil
}

// func (s *userService) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.TokenResponse, error) {
// 	tx := s.db.Begin()
// 	defer SafeRollback(tx)

// 	// Find the refresh token in the database
// 	dbToken, err := s.refreshTokenRepo.FindByToken(ctx, tx, req.RefreshToken)
// 	if err != nil {
// 		tx.Rollback()
// 		return dto.TokenResponse{}, errors.New(dto.MESSAGE_FAILED_INVALID_REFRESH_TOKEN)
// 	}

// 	if time.Now().After(dbToken.ExpiresAt) {
// 		tx.Rollback()
// 		return dto.TokenResponse{}, errors.New(dto.MESSAGE_FAILED_EXPIRED_REFRESH_TOKEN)
// 	}

// 	user, err := s.userRepo.GetUserById(ctx, tx, dbToken.UserID.String())
// 	if err != nil {
// 		tx.Rollback()
// 		return dto.TokenResponse{}, dto.ErrUserNotFound
// 	}

// 	accessToken := s.jwtService.GenerateAccessToken(user.ID.String(), user.Role)

// 	refreshTokenString, expiresAt := s.jwtService.GenerateRefreshToken()

// 	hashedToken, err := helpers.HashPassword(refreshTokenString)
// 	if err != nil {
// 		tx.Rollback()
// 		return dto.TokenResponse{}, err
// 	}

// 	if err := s.refreshTokenRepo.DeleteByUserID(ctx, tx, user.ID.String()); err != nil {
// 		tx.Rollback()
// 		return dto.TokenResponse{}, err
// 	}

// 	refreshToken := entity.RefreshToken{
// 		UserID:    user.ID,
// 		Token:     hashedToken,
// 		ExpiresAt: expiresAt,
// 	}

// 	if _, err := s.refreshTokenRepo.Create(ctx, tx, refreshToken); err != nil {
// 		tx.Rollback()
// 		return dto.TokenResponse{}, err
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		return dto.TokenResponse{}, err
// 	}

// 	return dto.TokenResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshTokenString,
// 		Role:         user.Role,
// 	}, nil
// }

// func (s *userService) RevokeRefreshToken(ctx context.Context, userID string) error {
// 	tx := s.db.Begin()
// 	defer SafeRollback(tx)

// 	// Check if user exists
// 	_, err := s.userRepo.GetUserById(ctx, tx, userID)
// 	if err != nil {
// 		tx.Rollback()
// 		return dto.ErrUserNotFound
// 	}

// 	// Delete all refresh tokens for the user
// 	if err := s.refreshTokenRepo.DeleteByUserID(ctx, tx, userID); err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// Commit transaction
// 	if err := tx.Commit().Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
