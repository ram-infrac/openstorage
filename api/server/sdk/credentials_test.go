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

	s.MockDriver().
		EXPECT().
		CredsCreate(params).
		Return("good-uuid", nil)

		// Setup client
	c := api.NewOpenStorageCredentialsClient(s.Conn())

	// Attach Volume
	_, err := c.ProvideForAWS(context.Background(), req)
	assert.NoError(t, err)
}
