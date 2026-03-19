package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/dto"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/model"
	"github.com/maxcore25/proga-po-go-transport-service/internal/auth/repository"
	"github.com/maxcore25/proga-po-go-transport-service/internal/shared/utils"
)

type AuthService interface {
	Register(req dto.RegisterRequest) (*dto.AuthTokens, error)
	Login(req dto.LoginRequest) (*dto.AuthTokens, error)
	Refresh(refreshToken string) (*dto.AuthTokens, error)
	Logout(refreshToken string) error
}

type authService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
	jwt         *utils.JWTManager
}

func NewAuthService(u repository.UserRepository, r repository.RefreshTokenRepository, j *utils.JWTManager) AuthService {
	return &authService{userRepo: u, refreshRepo: r, jwt: j}
}

func (s *authService) Register(req dto.RegisterRequest) (*dto.AuthTokens, error) {
	// 1. Check if user already exists
	existing, _ := s.userRepo.GetByUsername(req.Username)
	if existing != nil {
		return nil, errors.New("user with this username already exists")
	}

	// 2. Hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 3. Create user model
	user := &model.User{
		ID:           uuid.New(),
		Username:     req.Username,
		PasswordHash: hashed,
		FullName:     req.FullName,
		IsAdmin:      false,
		IsActive:     true,
	}

	// 4. Save to DB
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// 5. Generate JWTs
	access, err := s.jwt.GenerateAccessToken(user.ID, "user")
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwt.GenerateRefreshToken(user.ID, "user")
	if err != nil {
		return nil, err
	}

	// 6. Store refresh token
	tokenModel := &model.RefreshToken{
		UserID:    user.ID,
		Token:     refresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}
	if err := s.refreshRepo.Save(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthTokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *authService) Login(req dto.LoginRequest) (*dto.AuthTokens, error) {
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid username or password")
	}

	role := "user"
	if user.IsAdmin {
		role = "admin"
	}

	access, err := s.jwt.GenerateAccessToken(user.ID, role)
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwt.GenerateRefreshToken(user.ID, role)
	if err != nil {
		return nil, err
	}

	tokenModel := &model.RefreshToken{
		UserID:    user.ID,
		Token:     refresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshRepo.Save(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthTokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *authService) Refresh(refreshToken string) (*dto.AuthTokens, error) {
	claims, err := s.jwt.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	role := claims.Role

	access, err := s.jwt.GenerateAccessToken(userID, role)
	if err != nil {
		return nil, err
	}

	newRefresh, err := s.jwt.GenerateRefreshToken(userID, role)
	if err != nil {
		return nil, err
	}

	_ = s.refreshRepo.Delete(refreshToken)

	tokenModel := &model.RefreshToken{
		UserID:    userID,
		Token:     newRefresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshRepo.Save(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthTokens{
		AccessToken:  access,
		RefreshToken: newRefresh,
	}, nil
}

func (s *authService) Logout(refreshToken string) error {
	return s.refreshRepo.Delete(refreshToken)
}
