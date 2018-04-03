package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/libopenstorage/openstorage/secrets"
)

const (
	secretKeyOkMsg   = "Secret Key set successfully"
	secretLoginOkMsg = "Login completd"
)

//TODO: Add swagger yaml
func (c *clusterApi) setClusterSecretKey(w http.ResponseWriter, r *http.Request) {
	//  method := "setClustersecretKey"
	params := r.URL.Query()
	secretKey := params["clustersecretkey"][0]
	override, _ := strconv.ParseBool(params["override"][0])

	if secretKey == "" {
		//m.sendError(m.name, method, w, "Missing cluster key", http.StatusBadRequest)
		http.Error(w, "Missing cluster key", http.StatusNotImplemented)
		return
	}
	fmt.Println("secretekey", secretKey, "over", override)
	err := SecretManager.Secret.SetClusterSecretKey(secretKey, override)
	rt := GetClusterAPIRoutes()
	fmt.Println("changed", rt)
	//err := m.secret.SetClusterSecretKey(secretKey, false)
	if err != nil {
		//  m.sendError(m.name, method, w, err.Error(), http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	w.Write([]byte("cluster" + secretKeyOkMsg + "\n"))
}

//TODO: Add swagger yaml
func (c *clusterApi) secretsLogin(w http.ResponseWriter, r *http.Request) {
	var dcReq secrets.SecretLoginRequest
	//  method := "secretsLogin"
	params := r.URL.Query()
	if err := json.NewDecoder(r.Body).Decode(&dcReq); err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//      m.sendError(m.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	secretStore, convErr := strconv.ParseInt(params["secret"][0], 10, 64)
	if convErr != nil {
		//m.sendError(m.name, method, w, "Missing secret store type", http.StatusInternalServerError)
		http.Error(w, convErr.Error(), http.StatusNotImplemented)
		return
	}
	fmt.Println("store", secretStore)
	err := SecretManager.Secret.SecretLogin(int(secretStore), dcReq.SecretConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//      m.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(secretLoginOkMsg + "\n"))
}

//TODO: Add swagger yaml
func (c *clusterApi) setSecrets(w http.ResponseWriter, r *http.Request) {
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
		//      m.sendError(m.name, method, w, "Missing secret value", http.StatusInternalServerError)
		return
	}
	fmt.Println("secretType", secretKey, "value", secretValue)

	err := SecretManager.Secret.SetSecret(secretKey, secretValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(secretKeyOkMsg + "\n"))
}

//TODO: Add swagger yaml
func (c *clusterApi) getSecrets(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	secretid := params["secretid"][0]
	if secretid == "" {
		http.Error(w, "Missing Secret ID", http.StatusNotImplemented)
		//      m.sendError(m.name, method, w, "Missing secret key", http.StatusInternalServerError)
		return
	}

	secretValue, err := SecretManager.Secret.GetSecret(secretid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//c.sendError(c.name, method, w, "Invalid secret id", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(secretValue))
}

//TODO: Add swagger yaml
func (c *clusterApi) secretLoginCheck(w http.ResponseWriter, r *http.Request) {

	err := SecretManager.Secret.CheckSecretLogin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		//m.sendError(m.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("err", err)
	w.Write([]byte("ok"))
}
