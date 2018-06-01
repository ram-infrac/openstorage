/*
Package sdk is the gRPC implementation of the SDK gRPC server
Copyright 2018 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sdk

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/libopenstorage/openstorage/api"
)

func TestSdkCredentialCreateSuccess(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.ProvideCredentialsForAWSRequest{
		CredType:  "s3",
		AccessKey: "dummy-access",
		SecretKey: "dummy-secret",
		Endpoint:  "dummy-endpoint",
		Region:    "dummy-region",
	}

	params := make(map[string]string)

	params[api.OptCredType] = req.GetCredType()
	params[api.OptCredRegion] = req.GetRegion()
	params[api.OptCredEndpoint] = req.GetEndpoint()
	params[api.OptCredAccessKey] = req.GetAccessKey()
	params[api.OptCredSecretKey] = req.GetSecretKey()

	uuid := "good-uuid"
	s.MockDriver().
		EXPECT().
		CredsCreate(params).
		Return(uuid, nil)

	s.MockDriver().
		EXPECT().
		CredsValidate(uuid).
		Return(nil)

		// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Attach Volume
	_, err := c.ProvideForAWS(context.Background(), req)
	assert.NoError(t, err)
}

func TestSdkCredentialCreateFailed(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.ProvideCredentialsForAWSRequest{
		CredType:  "s3",
		AccessKey: "dummy-access",
		SecretKey: "dummy-secret",
		Endpoint:  "dummy-endpoint",
		Region:    "dummy-region",
	}

	params := make(map[string]string)

	params[api.OptCredType] = req.GetCredType()
	params[api.OptCredRegion] = req.GetRegion()
	params[api.OptCredEndpoint] = req.GetEndpoint()
	params[api.OptCredAccessKey] = req.GetAccessKey()
	params[api.OptCredSecretKey] = req.GetSecretKey()

	uuid := "bad-uuid"
	s.MockDriver().
		EXPECT().
		CredsCreate(params).
		Return(uuid, nil)

	s.MockDriver().
		EXPECT().
		CredsValidate(uuid).
		Return(fmt.Errorf("Invalid credentials"))

	s.MockDriver().
		EXPECT().
		CredsDelete(uuid).
		Return(nil)

		// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Attach Volume
	_, err := c.ProvideForAWS(context.Background(), req)
	assert.NoError(t, err)
}
func TestSdkCredentialEnumerateAWSSuccess(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.EnumerateCredentialsForAWSRequest{CredentialId: "test"}

	enumS3 := map[string]interface{}{
		api.OptCredType:      "s3",
		api.OptCredAccessKey: "test-access",
		api.OptCredSecretKey: "test-secret",
		api.OptCredEndpoint:  "test-endpoint",
		api.OptCredRegion:    "test-region",
	}
	enumerateData := map[string]interface{}{
		api.OptCredUUID: enumS3,
	}

	s.MockDriver().
		EXPECT().
		CredsEnumerate().
		Return(enumerateData, nil)

		// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Attach Volume
	resp, err := c.EnumerateForAWS(context.Background(), req)
	assert.NoError(t, err)

	assert.Equal(t, resp.GetCred()[0].AccessKey, enumS3[api.OptCredAccessKey])
	assert.Equal(t, resp.GetCred()[0].Endpoint, enumS3[api.OptCredEndpoint])
}

func TestSdkCredentialEnumerateAWSFailed(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.EnumerateCredentialsForAWSRequest{CredentialId: "test"}

	s.MockDriver().
		EXPECT().
		CredsEnumerate().
		Return(nil, fmt.Errorf("Failed to get credenntials details"))

	// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// EnumerateCredentials for AWS
	resp, err := c.EnumerateForAWS(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)

}

func TestSdkCredentialEnumerateAzureSuccess(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.EnumerateCredentialsForAzureRequest{CredentialId: "test"}

	enumAzure := map[string]interface{}{
		api.OptCredType:             "azure",
		api.OptCredAzureAccountName: "test-azure-account",
		api.OptCredAzureAccountKey:  "test-azure-account",
	}
	enumerateData := map[string]interface{}{
		api.OptCredUUID: enumAzure,
	}

	s.MockDriver().
		EXPECT().
		CredsEnumerate().
		Return(enumerateData, nil)

		// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Attach Volume
	resp, err := c.EnumerateForAzure(context.Background(), req)
	assert.NoError(t, err)

	assert.Equal(t, resp.GetCred()[0].AccountName, enumAzure[api.OptCredAzureAccountName])
	assert.Equal(t, resp.GetCred()[0].AccountKey, enumAzure[api.OptCredAzureAccountKey])

}

func TestSdkCredentialEnumerateAzureFailed(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.EnumerateCredentialsForAzureRequest{CredentialId: "test"}

	s.MockDriver().
		EXPECT().
		CredsEnumerate().
		Return(nil, fmt.Errorf("Failed to get credenntials details"))

	// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// EnumerateCredentials for AWS
	resp, err := c.EnumerateForAzure(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestSdkCredentialEnumerateGoogleSuccess(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.EnumerateCredentialsForGoogleRequest{CredentialId: "test"}

	enumGoogle := map[string]interface{}{
		api.OptCredType:            "google",
		api.OptCredGoogleProjectID: "test-google-project-id",
	}
	enumerateData := map[string]interface{}{
		api.OptCredUUID: enumGoogle,
	}

	s.MockDriver().
		EXPECT().
		CredsEnumerate().
		Return(enumerateData, nil)

		// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Attach Volume
	resp, err := c.EnumerateForGoogle(context.Background(), req)
	assert.NoError(t, err)

	assert.Equal(t, resp.GetCred()[0].ProjectId, enumGoogle[api.OptCredGoogleProjectID])
}

func TestSdkCredentialEnumerateGoogleFailed(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	req := &api.EnumerateCredentialsForGoogleRequest{CredentialId: "test"}

	s.MockDriver().
		EXPECT().
		CredsEnumerate().
		Return(nil, fmt.Errorf("Failed to get credenntials details"))

	// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// EnumerateCredentials for AWS
	resp, err := c.EnumerateForGoogle(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestSdkCredentialsDeleteSuccess(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	cred_id := "myid"
	req := &api.CredentialsDeleteRequest{
		CredentailId: cred_id,
	}
	s.MockDriver().
		EXPECT().
		CredsDelete(cred_id).
		Return(nil)

	// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Delete Credentials
	_, err := c.CredentialsDelete(context.Background(), req)
	assert.NoError(t, err)
}

func TestSdkCredentialsDeleteBadArgument(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	cred_id := ""
	req := &api.CredentialsDeleteRequest{
		CredentailId: cred_id,
	}

	// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Delete Credentials
	_, err := c.CredentialsDelete(context.Background(), req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Must provide credentials uuid")
}

func TestSdkCredentialsDeleteFailed(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	cred_id := "myid"
	req := &api.CredentialsDeleteRequest{
		CredentailId: cred_id,
	}
	s.MockDriver().
		EXPECT().
		CredsDelete(cred_id).
		Return(fmt.Errorf("Error deleting credentials"))

	// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Delete Credentials
	_, err := c.CredentialsDelete(context.Background(), req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Error deleting credentials")
}
