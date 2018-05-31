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
	"errors"
	"reflect"

	"github.com/libopenstorage/openstorage/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Delete specfied credentials
func (s *VolumeServer) CredentialsDelete(
	ctx context.Context,
	req *api.CredentialsDeleteRequest,
) (*api.CredentialsDeleteResponse, error) {

	if len(req.GetCredentailId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must provide credentials uuid")
	}

	err := s.driver.CredsDelete(req.GetCredentailId())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to  delete credentials: %v",
			err.Error())
	}

	return &api.CredentialsDeleteResponse{}, nil
}

func getCredentialMap(credList map[string]interface{}, credType string) (map[string]interface{}, error) {
	creds := make(map[string]interface{})

	for k, v := range credList {
		c, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("Error parsing credentials %v" +
				reflect.TypeOf(v).String())
		}

		// Look for only one type
		switch c[api.OptCredType] {
		case credType:
			creds[k] = v
		default:
			return nil, errors.New("Could not find credentials stored for " + credType)
		}
	}

	return creds, nil
}

// EnumerateCredentialsForAWS list credentials for AWS
func (s *VolumeServer) EnumerateForAWS(
	ctx context.Context,
	req *api.EnumerateCredentialsForAWSRequest,
) (*api.EnumerateCredentialsForAWSResponse, error) {

	credList, err := s.driver.CredsEnumerate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to enumerate credentials AWS: %v",
			err.Error())
	}

	// By defaultcredList will have all credential details, we will extract for
	// respective cloud provider and return result
	// this may not be expected behaviour, we have to do this since
	// `interface` can't be mapped directly with other lang
	s3Creds, err := getCredentialMap(credList, "s3")
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to enumerate credentials AWS: %v",
			err.Error())
	}

	// Fill up s3 credential resonse
	creds := []*api.S3Credential{}
	for k, v := range s3Creds {
		cred, ok := v.(map[string]interface{})
		if !ok {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to enumerate credentials AWS: %v",
				reflect.TypeOf(v).String())
		}

		credResp := &api.S3Credential{
			CredentialId: k,
			AccessKey:    cred[api.OptCredAccessKey].(string),
			Endpoint:     cred[api.OptCredEndpoint].(string),
			Region:       cred[api.OptCredRegion].(string),
		}
		creds = append(creds, credResp)
	}

	return &api.EnumerateCredentialsForAWSResponse{Cred: creds}, nil
}

// EnumerateCredentialsForAWS list credentials for AWS
func (s *VolumeServer) EnumerateForAzure(
	ctx context.Context,
	req *api.EnumerateCredentialsForAzureRequest,
) (*api.EnumerateCredentialsForAzureResponse, error) {
	credList, err := s.driver.CredsEnumerate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to enumerate credentials: %v",
			err.Error())
	}

	// By defaultcredList will have all credential details, we will extract for
	// respective cloud provider and return result
	// this may not be expected behaviour, we have to do this since
	// `interface` can't be mapped directly with other lang
	azureCreds, err := getCredentialMap(credList, "azure")
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to enumerate credentials Azure: %v",
			err.Error())
	}

	// Fill up azure credential resonse
	creds := []*api.AzureCredential{}
	for k, v := range azureCreds {
		cred, ok := v.(map[string]interface{})
		if !ok {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to enumerate credentials AWS: %v",
				reflect.TypeOf(v).String())
		}

		credResp := &api.AzureCredential{
			CredentialId: k,
			AccountName:  cred[api.OptCredAzureAccountName].(string),
			AccountKey:   cred[api.OptCredAzureAccountKey].(string),
		}
		creds = append(creds, credResp)
	}
	return &api.EnumerateCredentialsForAzureResponse{Cred: creds}, nil
}

// EnumerateCredentialsForAWS list credentials for Google
func (s *VolumeServer) EnumerateForGoogle(
	ctx context.Context,
	req *api.EnumerateCredentialsForGoogleRequest,
) (*api.EnumerateCredentialsForGoogleResponse, error) {
	credList, err := s.driver.CredsEnumerate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to enumerate credentials: %v",
			err.Error())
	}

	// By defaultcredList will have all credential details, we will extract for
	// respective cloud provider and return result
	// this may not be expected behaviour, we have to do this since
	// `interface` can't be mapped directly with other lang
	googleCreds, err := getCredentialMap(credList, "google")
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to enumerate credentials Azure: %v",
			err.Error())
	}

	// Fill up google credential resonse
	creds := []*api.GoogleCredential{}
	for k, v := range googleCreds {
		cred, ok := v.(map[string]interface{})
		if !ok {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to enumerate credentials AWS: %v",
				reflect.TypeOf(v).String())
		}

		credResp := &api.GoogleCredential{
			CredentialId: k,
			ProjectId:    cred[api.OptCredGoogleProjectID].(string),
		}
		creds = append(creds, credResp)
	}

	return &api.EnumerateCredentialsForGoogleResponse{Cred: creds}, nil
}
