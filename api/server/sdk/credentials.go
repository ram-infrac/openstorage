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

	"github.com/libopenstorage/openstorage/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *VolumeServer) ProvideForAWS(
	ctx context.Context,
	req *api.ProvideCredentialsForAWSRequest,
) (*api.ProvideCredentialsForAWSResponse, error) {

	if len(req.GetAccessKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Access Key")
	}

	if len(req.GetSecretKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Secret Key")
	}

	if len(req.GetRegion()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Region Key")
	}

	if len(req.GetEndpoint()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Endpoint Key")
	}

	params := make(map[string]string)

	params[api.OptCredType] = req.GetCredType()
	params[api.OptCredRegion] = req.GetRegion()
	params[api.OptCredEndpoint] = req.GetEndpoint()
	params[api.OptCredAccessKey] = req.GetAccessKey()
	params[api.OptCredSecretKey] = req.GetSecretKey()

	uuid, err := s.driver.CredsCreate(params)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to create AWS S3 credentials: %v",
			err.Error())
	}
	return &api.ProvideCredentialsForAWSResponse{CredentialId: uuid}, nil

}

func (s *VolumeServer) ProvideForAzure(
	ctx context.Context,
	req *api.ProvideCredentialsForAzureRequest,
) (*api.ProvideCredentialsForAzureResponse, error) {

	return &api.ProvideCredentialsForAzureResponse{CredentialId: "dummy-azure"}, nil

}

func (s *VolumeServer) ProvideForGoogle(
	ctx context.Context,
	req *api.ProvideCredentialsForGoogleRequest,
) (*api.ProvideCredentialsForGoogleResponse, error) {

	return &api.ProvideCredentialsForGoogleResponse{CredentialId: "dummy-google"}, nil
}
