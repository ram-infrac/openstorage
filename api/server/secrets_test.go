package server

import (
	"fmt"
	"testing"

	clusterclient "github.com/libopenstorage/openstorage/api/client/cluster"
	"github.com/libopenstorage/openstorage/secrets"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaultSecretKeySuccess(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	secretKey := "testkey"
	overrideFlag := true
	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		SetDefaultSecretKey(secretKey, overrideFlag).
		Return(nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.SetDefaultSecretKey(secretKey, overrideFlag)

	assert.NoError(t, err)
}

func TestSetDefaultSecretKeyFailed(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	secretKey := "testClusterKey"
	overrideFlag := false
	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		SetDefaultSecretKey(secretKey, overrideFlag).
		Return(fmt.Errorf("Not Implemented"))

		// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.SetDefaultSecretKey(secretKey, overrideFlag)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Not Implemented")
}

func TestGetDefaultSecretKeySuccess(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	defaultSecretTest := "testclusterkeyval"
	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		GetDefaultSecretKey().
		Return(defaultSecretTest, nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	resp, err := restClient.GetDefaultSecretKey()
	//	fmt.Println("resp", resp)
	//	convResp := resp.(*secrets.GetSecretResponse).SecretValue
	//convResp := resp.(string)
	//	fmt.Println("convResp", convResp)

	assert.NoError(t, err)
	assert.Equal(t, resp.(*secrets.GetSecretResponse).SecretValue, defaultSecretTest)
}

func TestGetDefaultSecretKeyFailed(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		GetDefaultSecretKey().
		Return(nil, fmt.Errorf("Not Implemented"))

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	resp, err := restClient.GetDefaultSecretKey()

	assert.Error(t, err)
	assert.Nil(t, resp.(*secrets.GetSecretResponse).SecretValue)
	assert.Contains(t, err.Error(), "500")
}

func TestGetSuccess(t *testing.T) {

	testKey := "testkey"
	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		Get(testKey).
		Return("testval", nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	getResp, err := restClient.Get(testKey)
	assert.NoError(t, err)
	assert.EqualValues(t, getResp.(secrets.GetSecretResponse).SecretValue, "testval")
}

func TestGetFailed(t *testing.T) {

	testKey := "testkey"
	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		Get(testKey).
		Return(nil, fmt.Errorf("Not Implemented"))

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	testVal, err := restClient.Get(testKey)

	assert.Error(t, err)
	assert.Nil(t, testVal.(secrets.GetSecretResponse).SecretValue)
	assert.Contains(t, err.Error(), "500")
}

func TestSetSuccess(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		Set("testkey", "testvalue").
		Return(nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.Set("testkey", "testvalue")
	assert.NoError(t, err)
}

func TestSetFailed(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		Set("testkey", "testvalue").
		Return(fmt.Errorf("Not Implemented"))

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.Set("testkey", "testvalue")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")

}

func TestVerifySuccess(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
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
}

func TestVerifyFailed(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		CheckLogin().
		Return(fmt.Errorf("Not Implemented"))

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.CheckLogin()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")
}

func TestSecretLoginSuccess(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		Login("teststore1", nil).
		Return(nil)

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.Login("teststore1", nil)
	assert.NoError(t, err)
}

func TestSecretLoginFailed(t *testing.T) {

	// Create a new global test cluster
	ts, tc := testClusterSecrets(t)
	defer ts.Close()
	defer tc.Finish()

	// mock the cluster response
	tc.MockClusterSecrets().
		EXPECT().
		Login("teststore1", nil).
		Return(fmt.Errorf("Not Implemented"))

	// create a cluster client to make the REST call
	c, err := clusterclient.NewClusterClient(ts.URL, "v1")
	assert.NoError(t, err)

	// make the REST call
	restClient := clusterclient.ClusterManager(c)
	err = restClient.Login("teststore1", nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "500")
}
