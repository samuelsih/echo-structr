package auth

import (
	"context"
	"github.com/samuelsih/echo-structr/business"
	"github.com/samuelsih/echo-structr/entity"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type Provider interface {
	Get(email string) (entity.User, error)
	Insert(name, email, password string) error
}

type Uploader interface {
	Upload() error
}

type Authenticator interface {
	Verify() error
	Generate() (string, error)
}

type Service struct {
	provider   Provider
	uploader   Uploader
	jwtChecker Authenticator
}

func NewAuthService(provider Provider, uploader Uploader, jwtChecker Authenticator) Service {
	return Service{
		provider:   provider,
		uploader:   uploader,
		jwtChecker: jwtChecker,
	}
}

type LoginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOut struct {
	business.CommonResponse
}

func (s Service) Login(ctx context.Context, in LoginIn, cr business.CommonRequest) LoginOut {
	var out LoginOut

	user, err := s.provider.Get(in.Email)
	if err != nil {
		out.Set(400, err.Error())
		return out
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		out.Set(400, "bad credentials")
	}

	return out
}
