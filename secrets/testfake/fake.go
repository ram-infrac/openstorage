package testfake

import (
	"errors"
	"github.com/libopenstorage/openstorage/secrets"
)

var (
	// ErrNotImplemented default secrets in OSD
	ErrNotImplemented = errors.New("Not Implemented123")
)

type TestFake struct {
	secrets.SecretManager
}

// New returns a new TestFake secret implementation
func New() *TestFake {
	return &TestFake{}
}

func (f *TestFake) SecretLogin(secretType int, secretConfig map[string]string) error {
	return ErrNotImplemented
}

func (f *TestFake) SetClusterSecretKey(secretKey string, override bool) error {
	return ErrNotImplemented
}

func (f *TestFake) CheckSecretLogin() error {
	return ErrNotImplemented
}

func (f *TestFake) SetSecret(secretKey string, secretValue string) error {
	return ErrNotImplemented
}

func (f *TestFake) GetSecret(secretKey string) (string, error) {
	return "", ErrNotImplemented
}
