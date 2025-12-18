package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(user model.CreateUserParameter) (model.User, error)
	FindByUsername(username string) model.User
	IsUserRegistered(userID uint64, username string) bool
	IsDuplicateUsername(username string) bool
	CheckExisting(user model.User) (userOutput model.User, err error)
	SendMail(mail model.Mail) (res gin.H, err error)
	SendMail2(mail model.Mail2) (res gin.H, err error)
	UpdateDataPassword(user model.ResetPasswordParameter, username string, email string) (userOutput model.User, err error)
	UpdateDataRequest(user model.RequestPasswordChangeParameter, username string, email string) (userOutput model.User, err error)
	UploadDocument(file *multipart.FileHeader, fileName string) (*s3manager.UploadOutput, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username)
	if v, ok := res.(model.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return res
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (service *authService) CreateUser(user model.CreateUserParameter) (model.User, error) {
	newUser := model.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		return newUser, err
	} else {
		res, err := service.userRepository.InsertUser(newUser)
		return res, err
	}
}

func (service *authService) FindByUsername(username string) model.User {
	return service.userRepository.FindByUsername(username)
}

func (service *authService) IsUserRegistered(userID uint64, username string) bool {
	res := service.userRepository.IsUserRegistered(userID, username)
	return !(res.Error == nil)
}

func (service *authService) IsDuplicateUsername(username string) bool {
	res := service.userRepository.IsDuplicateUsername(username)
	return !(res.Error == nil)
}

func (service *authService) CheckExisting(user model.User) (userOutput model.User, err error) {
	newUser := model.User{}
	err1 := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err1 != nil {
		return newUser, err1
	}
	res, err := service.userRepository.CheckExisting(newUser)
	return res, err
}

func (service *authService) SendMail(mail model.Mail) (res gin.H, err error) {
	newMail := model.Mail{}
	err1 := smapping.FillStruct(&newMail, smapping.MapFields(&mail))
	if err1 != nil {
		return res, err1
	}
	res, err = service.userRepository.SendMail(newMail)
	return res, err
}

func (service *authService) SendMail2(mail model.Mail2) (res gin.H, err error) {
	newMail := model.Mail2{}
	err1 := smapping.FillStruct(&newMail, smapping.MapFields(&mail))
	if err1 != nil {
		return res, err1
	}
	res, err = service.userRepository.SendMail2(newMail)
	return res, err
}

func (service *authService) UpdateDataPassword(user model.ResetPasswordParameter, username string, email string) (userOutput model.User, err error) {
	newUser := model.User{}
	err1 := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err1 != nil {
		return newUser, err1
	}
	res, err := service.userRepository.UpdateDataPassword(newUser, username, email)
	return res, err
}

func (service *authService) UpdateDataRequest(user model.RequestPasswordChangeParameter, username string, email string) (userOutput model.User, err error) {
	newUser := model.User{}
	err1 := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err1 != nil {
		return newUser, err1
	}
	res, err := service.userRepository.UpdateDataRequest(newUser, username, email)
	return res, err
}

func (service *authService) UploadDocument(file *multipart.FileHeader, fileName string) (*s3manager.UploadOutput, error) {
	fileBody, openFileErr := file.Open()

	if openFileErr != nil {
		return nil, openFileErr
	}

	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}

	region := os.Getenv("AWS_REGION")
	aws_id := os.Getenv("AWS_ID")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")
	bucket := os.Getenv("AWS_BUCKET")

	// Save file to root directory:
	s3Config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(aws_id, aws_secret_key, ""),
	}

	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)
	MyBucket := bucket

	contentType := "application/pdf"

	splitFile := strings.Split(fileName, "/")
	contentDisposition := fmt.Sprintf("inline; filename=\"%s\"", splitFile[len(splitFile)-1])

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(MyBucket),
		Key:                aws.String(fileName),
		Body:               fileBody,
		ContentType:        &contentType,
		ContentDisposition: &contentDisposition,
	})

	return up, err
}
