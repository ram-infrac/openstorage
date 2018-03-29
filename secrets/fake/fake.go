package fake

import (
	"errors"
	"fmt"
	"github.com/libopenstorage/openstorage/secrets"
)

var (
	// ErrAlreadyShutdown returned when driver is shutdown
	ErrNotImplemented = errors.New("Not Implemented")
)

type Fake struct {
	secrets.SecretManager
}

// New returns a new Fake cloud provider
func New() *Fake {
	return &Fake{}
}

func (f *Fake) SecretLogin(secretType int, secretConfig map[string]string) error {
	fmt.Println("fake secretLogin")
	return ErrNotImplemented
}

func (f *Fake) SetClusterSecretKey(secretKey string, override bool) error {
	fmt.Println("fake clusterkey")
	return ErrNotImplemented
}

func (f *Fake) CheckSecretLogin() error {
	fmt.Println("fake checksecretlogin")
	return ErrNotImplemented
}

func (f *Fake) SetSecret(secretKey string, secretValue string) error {
	fmt.Println(" Fake set secret")
	return ErrNotImplemented
}

func (f *Fake) GetSecret(secretKey string) (string, error) {
	fmt.Println(" Fake set secret")
	return "", ErrNotImplemented
}
