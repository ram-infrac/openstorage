package secrets

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	ErrSecretsNotFound = errors.New("Secrets Implementor not registerd")
	secretKeyOkMsg     = "Secret Key set successfully"
	secretLoginOkMsg   = "Login completd"
)

type SecretManager interface {
	// Login create session with secret store
	SecretLogin(secretType int, secretConfig map[string]string) error
	// SetClusterSecretKey sets the cluster wide secret key
	SetClusterSecretKey(secretKey string, override bool) error
	//CheckSecretLogin validates session with secret store
	CheckSecretLogin() error
	//SetSecret sets secret key against data
	SetSecret(key string, value string) error
	// GetSecret retrieves the data for the given feature and key
	GetSecret(key string) (string, error)
	//Routes to the secret API
	Routes() []*Route
}

type Manager struct {
	secret SecretManager
	lock   sync.Mutex
}

func NewSecretManager(sec SecretManager) *Manager {
	return &Manager{
		secret: sec,
	}
}

func (m *Manager) Routes() []*Route {
	return []*Route{
		{Verb: "GET", Path: secretVersion("checksecretslogin", APIVersion), Fn: m.secretLoginCheck},
		{Verb: "GET", Path: secretVersion("getsecrets", APIVersion), Fn: m.getSecrets},
		{Verb: "PUT", Path: secretVersion("setsecrets", APIVersion), Fn: m.setSecrets},
		{Verb: "POST", Path: secretVersion("setclustersecretkey", APIVersion), Fn: m.setClusterSecretKey},
		{Verb: "POST", Path: secretVersion("secretslogin", APIVersion), Fn: m.secretsLogin},
	}
}
func secretVersion(route, version string) string {
	if version == "" {
		return "/" + route
	} else {
		return "/" + version + "/" + route
	}
}

//TODO: Add swagger yaml
func (m *Manager) setClusterSecretKey(w http.ResponseWriter, r *http.Request) {
	//	method := "setClustersecretKey"
	params := r.URL.Query()
	secretKey := params["clustersecretkey"][0]
	override, _ := strconv.ParseBool(params["override"][0])

	if secretKey == "" {
		//m.sendError(m.name, method, w, "Missing cluster key", http.StatusBadRequest)
		http.Error(w, "Missing cluster key", http.StatusNotImplemented)
		return
	}
	fmt.Println("secretekey", secretKey, "over", override)
	err := m.secret.SetClusterSecretKey(secretKey, override)
	//err := m.secret.SetClusterSecretKey(secretKey, false)
	if err != nil {
		//	m.sendError(m.name, method, w, err.Error(), http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	w.Write([]byte("cluster" + secretKeyOkMsg + "\n"))
}

//TODO: Add swagger yaml
func (m *Manager) secretsLogin(w http.ResponseWriter, r *http.Request) {
	var dcReq SecretLoginRequest
	//	method := "secretsLogin"
	params := r.URL.Query()
	if err := json.NewDecoder(r.Body).Decode(&dcReq); err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//		m.sendError(m.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	secretStore, convErr := strconv.ParseInt(params["secret"][0], 10, 64)
	if convErr != nil {
		//m.sendError(m.name, method, w, "Missing secret store type", http.StatusInternalServerError)
		http.Error(w, convErr.Error(), http.StatusNotImplemented)
		return
	}
	fmt.Println("store", secretStore)
	err := m.secret.SecretLogin(int(secretStore), dcReq.SecretConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//		m.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(secretLoginOkMsg + "\n"))
}

//TODO: Add swagger yaml
func (m *Manager) setSecrets(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	secretKey := params["secretid"][0]
	secretValue := params["secretvalue"][0]
	if secretKey == "" {
		http.Error(w, "Missing secret key", http.StatusNotImplemented)
		//m.sendError(m.name, method, w, "Missing secret key", http.StatusInternalServerError)
		return
	}
	if secretValue == "" {
		http.Error(w, "Missing secret value", http.StatusNotImplemented)
		//		m.sendError(m.name, method, w, "Missing secret value", http.StatusInternalServerError)
		return
	}
	fmt.Println("secretType", secretKey, "value", secretValue)

	err := m.secret.SetSecret(secretKey, secretValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(secretKeyOkMsg + "\n"))
}

//TODO: Add swagger yaml
func (m *Manager) getSecrets(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	secretid := params["secretid"][0]
	if secretid == "" {
		http.Error(w, "Missing Secret ID", http.StatusNotImplemented)
		//		m.sendError(m.name, method, w, "Missing secret key", http.StatusInternalServerError)
		return
	}

	secretValue, err := m.secret.GetSecret(secretid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//c.sendError(c.name, method, w, "Invalid secret id", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(secretValue))
}

//TODO: Add swagger yaml
func (m *Manager) secretLoginCheck(w http.ResponseWriter, r *http.Request) {

	err := m.secret.CheckSecretLogin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//m.sendError(m.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("err", err)
	w.Write([]byte("ok"))
}
