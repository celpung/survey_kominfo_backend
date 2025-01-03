package user_usecase_implementation

import (
	"errors"

	user_repository "github.com/celpung/gocleanarch/domain/user/repository"
	user_usecase "github.com/celpung/gocleanarch/domain/user/usecase"
	"github.com/celpung/gocleanarch/entity"
	jwt_services "github.com/celpung/gocleanarch/services/jwt"
	password_services "github.com/celpung/gocleanarch/services/password"
)

type UserUsecaseStruct struct {
	UserRepository  user_repository.UserRepositoryInterface
	PasswordService *password_services.PasswordService
	JWTService      *jwt_services.JwtService
}

// Create implements user_usecase.UserUsecaseInterface.
func (u *UserUsecaseStruct) Create(user *entity.User) (*entity.UserHttpResponse, error) {
	if len(user.Password) < 8 {
		return nil, errors.New("password minimal 8 karakter")
	}
	// hashing password
	hashedPassword, err := u.PasswordService.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// set the hashed password into new user password
	user.Password = hashedPassword
	user.Active = true

	// perform create user
	user, userErr := u.UserRepository.Create(user)
	if userErr != nil {
		return nil, userErr
	}

	userResponse := &entity.UserHttpResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Active:   user.Active,
		Role:     user.Role,
	}

	return userResponse, nil
}

// Delete implements user_usecase.UserUsecaseInterface.
func (u *UserUsecaseStruct) Delete(userID uint) error {
	// perform delete user
	return u.UserRepository.Delete(userID)
}

// Read implements user_usecase.UserUsecaseInterface.
func (u *UserUsecaseStruct) Read(page, limit int) ([]*entity.UserHttpResponse, int64, error) {
	// perform read all user
	user, total, err := u.UserRepository.Read(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var userResponse []*entity.UserHttpResponse
	for _, v := range user {
		userResponse = append(userResponse, &entity.UserHttpResponse{
			ID:       v.ID,
			Name:     v.Name,
			Username: v.Username,
			Active:   v.Active,
			Role:     v.Role,
		})
	}
	return userResponse, total, nil

	// return u.UserRepository.Read()
}

// ReadByID implements user_usecase.UserUsecaseInterface.
func (u *UserUsecaseStruct) ReadByID(userID uint) (*entity.UserHttpResponse, error) {
	// perform read user by id
	user, userErr := u.UserRepository.ReadByID(userID)
	if userErr != nil {
		return nil, userErr
	}

	userResponse := &entity.UserHttpResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Active:   user.Active,
		Role:     user.Role,
	}

	return userResponse, nil
}

// Update implements user_usecase.UserUsecaseInterface.
func (u *UserUsecaseStruct) Update(user *entity.UserUpdate) (*entity.UserHttpResponse, error) {
	existingUser, err := u.UserRepository.ReadByID(user.ID)
	if err != nil {
		return nil, err
	}

	// Update only the non-zero fields
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Password != "" {
		if len(user.Password) < 8 {
			return nil, errors.New("password minimal 8 karakter")
		}
		hashedPassword, err := u.PasswordService.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		existingUser.Password = hashedPassword
	}
	if user.Active {
		existingUser.Active = user.Active
	}
	if user.Role > 0 {
		existingUser.Role = user.Role
	}

	// Perform the update operation
	updatedUser, err := u.UserRepository.Update(existingUser)
	if err != nil {
		return nil, err
	}

	userResponse := &entity.UserHttpResponse{
		ID:       updatedUser.ID,
		Name:     updatedUser.Name,
		Username: updatedUser.Username,
		Active:   updatedUser.Active,
		Role:     updatedUser.Role,
	}

	return userResponse, nil
}

// Login implements user_usecase.UserUsecaseInterface.
func (u *UserUsecaseStruct) Login(username, password string) (string, error) {
	// perform read user by email
	user, err := u.UserRepository.ReadByEmail(username, true)
	if err != nil {
		return "", err
	}

	// check is user active
	if !user.Active {
		return "", errors.New("user not active")
	}

	// verify hash password match plain password
	if err := u.PasswordService.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("wrong password")
	}

	// generate jwt token
	token, err := u.JWTService.JWTGenerator(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUserUsecase(repository user_repository.UserRepositoryInterface, passwordServive *password_services.PasswordService, jwtService *jwt_services.JwtService) user_usecase.UserUsecaseInterface {
	return &UserUsecaseStruct{
		UserRepository:  repository,
		PasswordService: passwordServive,
		JWTService:      jwtService,
	}
}
