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

// CreateForAWS method creates credential for AWS S3.
func (s *VolumeServer) CreateForAWS(
	ctx context.Context,
	req *api.CredentialCreateAWSRequest,
) (*api.CredentialCreateAWSResponse, error) {

	if len(req.GetCredential().GetAccessKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Access Key")
	}

	if len(req.GetCredential().GetSecretKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Secret Key")
	}

	if len(req.GetCredential().GetRegion()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Region Key")
	}

	if len(req.GetCredential().GetEndpoint()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Endpoint Key")
	}

	params := make(map[string]string)

	params[api.OptCredType] = "s3"
	params[api.OptCredRegion] = req.GetCredential().GetRegion()
	params[api.OptCredEndpoint] = req.GetCredential().GetEndpoint()
	params[api.OptCredAccessKey] = req.GetCredential().GetAccessKey()
	params[api.OptCredSecretKey] = req.GetCredential().GetSecretKey()

	uuid, err := s.driver.CredsCreate(params)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to create S3 credentials: %v",
			err.Error())
	}

	err = validateAndDelete(s, uuid)

	if err != nil {
		return nil, err
	}
	return &api.CredentialCreateAWSResponse{CredentialId: uuid}, nil

}

// CreateForAzure method creates credential for Azure.
func (s *VolumeServer) CreateForAzure(
	ctx context.Context,
	req *api.CredentialCreateAzureRequest,
) (*api.CredentialCreateAzureResponse, error) {

	if len(req.GetCredential().GetAccountKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Account Key")
	}

	if len(req.GetCredential().GetAccountName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Account name")
	}

	params := make(map[string]string)

	params[api.OptCredType] = "azure"
	params[api.OptCredAzureAccountKey] = req.GetCredential().GetAccountKey()
	params[api.OptCredAzureAccountName] = req.GetCredential().GetAccountName()

	uuid, err := s.driver.CredsCreate(params)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to create Azure credentials: %v",
			err.Error())
	}

	err = validateAndDelete(s, uuid)

	if err != nil {
		return nil, err
	}
	return &api.CredentialCreateAzureResponse{CredentialId: uuid}, nil
}

// CreateForGoogle method creates credential for Google.
func (s *VolumeServer) CreateForGoogle(
	ctx context.Context,
	req *api.CredentialCreateGoogleRequest,
) (*api.CredentialCreateGoogleResponse, error) {

	if len(req.GetCredential().GetJsonKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply JSON Key")
	}

	if len(req.GetCredential().GetProjectId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must supply Project ID")
	}

	params := make(map[string]string)

	params[api.OptCredType] = "google"
	params[api.OptCredGoogleProjectID] = req.GetCredential().GetProjectId()
	params[api.OptCredGoogleJsonKey] = req.GetCredential().GetJsonKey()

	uuid, err := s.driver.CredsCreate(params)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to create Google credentials: %v",
			err.Error())
	}

	err = validateAndDelete(s, uuid)

	if err != nil {
		return nil, err
	}

	return &api.CredentialCreateGoogleResponse{CredentialId: uuid}, nil
}

// CredentialValidate deletes a specified Credential.
func (s *VolumeServer) CredentialValidate(
	ctx context.Context,
	req *api.CredentialValidateRequest,
) (*api.CredentialValidateResponse, error) {

	if len(req.GetCredentialId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must provide credentials uuid")
	}

	validateReq := &api.CredentialValidateRequest{CredentialId: req.GetCredentialId()}

	err := s.driver.CredsValidate(validateReq.GetCredentialId())

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to validate credentials: %v",
			err.Error())
	}
	return &api.CredentialValidateResponse{}, nil

}

// CredentialDelete delete a specified credential
func (s *VolumeServer) CredentialDelete(
	ctx context.Context,
	req *api.CredentialDeleteRequest,
) (*api.CredentialDeleteResponse, error) {

	if len(req.GetCredentialId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Must provide credentials uuid")
	}

	err := s.driver.CredsDelete(req.GetCredentialId())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to  delete credentials: %v",
			err.Error())
	}

	return &api.CredentialDeleteResponse{}, nil
}

// EnumerateForAWS list credentials for AWS
func (s *VolumeServer) EnumerateForAWS(
	ctx context.Context,
	req *api.CredentialEnumerateAWSRequest,
) (*api.CredentialEnumerateAWSResponse, error) {

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

	return &api.CredentialEnumerateAWSResponse{Cred: creds}, nil
}

// EnumerateForAzure list credentials for AWS
func (s *VolumeServer) EnumerateForAzure(
	ctx context.Context,
	req *api.CredentialEnumerateAzureRequest,
) (*api.CredentialEnumerateAzureResponse, error) {
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
	return &api.CredentialEnumerateAzureResponse{Cred: creds}, nil
}

// EnumerateForGoogle list credentials for Google
func (s *VolumeServer) EnumerateForGoogle(
	ctx context.Context,
	req *api.CredentialEnumerateGoogleRequest,
) (*api.CredentialEnumerateGoogleResponse, error) {
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

	return &api.CredentialEnumerateGoogleResponse{Cred: creds}, nil
}

func validateAndDelete(s *VolumeServer, uuid string) error {
	// Validate if the credentials provided were correct or not
	req := &api.CredentialValidateRequest{CredentialId: uuid}

	validateErr := s.driver.CredsValidate(req.GetCredentialId())

	if validateErr != nil {
		deleteCred := &api.CredentialDeleteRequest{CredentialId: uuid}
		err := s.driver.CredsDelete(deleteCred.GetCredentialId())

		if err != nil {
			return status.Errorf(
				codes.Internal,
				"failed to delete invalid Google credentials: %v",
				err.Error())
		}

		return status.Errorf(
			codes.Internal,
			"credentials could not be validated: %v",
			validateErr.Error())
	}

	return nil
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
