package server

import (
	//"fmt"
	//	"errors"
	"testing"

	clusterclient "github.com/libopenstorage/openstorage/api/client/cluster"
	//"github.com/libopenstorage/openstorage/secrets"
	"github.com/stretchr/testify/assert"
)

/*
func TestSetDefaultSecretKey(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterServer(t)
	defer ts.Close()
	defer tc.Finish()

	secretKey := "testkey"
	overrideFlag := true
	// mock the cluster response
	tc.MockCluster().
		EXPECT().
		SetDefaultSecretKey(secretKey, overrideFlag).
		Return(nil)

		// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.SetDefaultSecretKey("testsecretkey", false)
	assert.NoError(t, err)
	//assert.Contains(t, err.Error(), "Not Implemented")
}
*/
/*
func TestGetDefaultSecretKey(t *testing.T) {

	// Create a new global test cluster
	var secretKey interface{}
	ts, tc := testClusterServer(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockCluster().
		EXPECT().
		GetDefaultSecretKey().
		Return(secretKey, nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	secretKey, err = restClient.GetDefaultSecretKey()
	//	assert.NoError(t, err)
	fmt.Println("key", secretKey)
	assert.Contains(t, err.Error(), "500")
}
*/
/*
func TestGet(t *testing.T) {

	// Create a new global test cluster
	var secretValue interface{}
	ts, tc := testClusterServer(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockCluster().
		EXPECT().
		Get("test").
		Return(nil, nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	secretValue, err = restClient.Get("test")
	assert.NoError(t, err)
	fmt.Println("val", secretValue)
	//assert.Contains(t, err.Error(), "500")
}*/
/*
func TestSet(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterServer(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockCluster().
		EXPECT().
		Set("test", nil).
		Return(nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	secretValue, err = restClient.Set("test", nil)
	assert.NoError(t, err)
	//fmt.Println("val", secretValue)
	//assert.Contains(t, err.Error(), "500")
}*/
/*
func TestSet(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterServer(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockCluster().
		EXPECT().
		Set("test", nil).
		Return(nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	secretValue, err = restClient.Set("test", nil)
	assert.NoError(t, err)
	//fmt.Println("val", secretValue)
	//assert.Contains(t, err.Error(), "500")
}*/

func TestVerify(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterServer(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockCluster().
		EXPECT().
		CheckLogin().
		Return(nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.CheckLogin()
	assert.NoError(t, err)
	//fmt.Println("val", secretValue)
	//assert.Contains(t, err.Error(), "500")
}
