package secrets

const (
	DefaultSecretKey = "defaultsecretkey"
	OverrideSecrets  = "override"
	SecretKey        = "id"
	SecretValue      = "secretvalue"
	SecretType       = "store"
	APIVersion       = "v1"
)

const (
	TypeNone   = "None"
	TypeKvdb   = "kvdb"
	TypeVault  = "vault"
	TypeAWS    = "aws-kms"
	TypeDocker = "docker"
	TypeK8s    = "k8s"
	TypeDCOS   = "dcos"
)

// SecretLoginRequest specify secret store and config to initiate
// secret store session
// swagger: parameters secret
type SecretLoginRequest struct {
	SecretType   string
	SecretConfig map[string]string
}

// SecretLoginResponse whether login is successful or not
type SecretStatusResponse struct {
	Status string
}

// ClusterSecretKeyRequest  specify request to set cluster secret key
// swagger: parameters clusterKey
type DefaultSecretKeyRequest struct {
	DefaultSecretKey string
	Override         bool
}

// SetsecretsLogin setsecrets
// swagger: parameters secret
type SetSecretRequest struct {
	SecretValue interface{}
}

// GetSecretResponse gets secret value for given key
type GetSecretResponse struct {
	SecretValue interface{}
}

// GetSecretResponse gets secret value for given key
type GetDefaultSecretKeyResponse struct {
	SecretValue interface{}
}
