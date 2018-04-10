package secrets

import (
	"errors"
	"sync"
)

var (
	// ErrNotImplemented default secrets in OSD
	ErrNotImplemented = errors.New("Not Implemented")
)

type Secrets interface {
	// Login create session with secret store
	Login(secretType string, secretConfig map[string]string) error
	// DefaultSecretKey  sets the cluster wide secret key
	SetDefaultSecretKey(secretKey string, override bool) error
	// GetDefaultSecretKey returns cluster wide secret key's value
	GetDefaultSecretKey() (interface{}, error)
	// CheckSecretLogin validates session with secret store
	CheckLogin() error
	// SetSecret sets secret key against data
	Set(key string, value interface{}) error
	// GetSecret retrieves the data for key
	Get(key string) (interface{}, error)
}

type Manager struct {
	Secrets
	lock sync.Mutex
}

func NewSecretManager(sec Secrets) *Manager {
	return &Manager{
		Secrets: sec,
	}
}

type nullSecretManager struct {
}

// New returns a new Default secrets implementation
func New() Secrets {
	return &nullSecretManager{}
}

func (f *nullSecretManager) Login(secretType string, secretConfig map[string]string) error {
	return ErrNotImplemented
}

func (f *nullSecretManager) SetDefaultSecretKey(secretKey string, override bool) error {
	return ErrNotImplemented
}

func (f *nullSecretManager) GetDefaultSecretKey() (interface{}, error) {
	return nil, ErrNotImplemented
}

func (f *nullSecretManager) CheckLogin() error {
	return ErrNotImplemented
}

func (f *nullSecretManager) Set(secretKey string, secretValue interface{}) error {
	return ErrNotImplemented
}

func (f *nullSecretManager) Get(secretKey string) (interface{}, error) {
	return nil, ErrNotImplemented
}
