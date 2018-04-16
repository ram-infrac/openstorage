/*
Copyright 2017 Portworx

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

package sanity

import (
	"github.com/libopenstorage/openstorage/api/client"
	clusterclient "github.com/libopenstorage/openstorage/api/client/cluster"

	"github.com/libopenstorage/openstorage/cluster"
	//	"github.com/libopenstorage/openstorage/secrets"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//TODO : what to expect default implementation returns error
var _ = Describe("Secrets [Secrets Tests]", func() {
	var (
		restClient *client.Client
		manager    cluster.Cluster
	)

	BeforeEach(func() {
		var err error
		restClient, err = clusterclient.NewClusterClient(osdAddress, cluster.APIVersion)
		Expect(err).ToNot(HaveOccurred())
		manager = clusterclient.ClusterManager(restClient)
	})

	AfterEach(func() {
	})

	Describe("Set Cluster Secret Key", func() {
		It("Should have cluster key set status", func() {
			By("setting cluster wide secret key")
			err := manager.SetDefaultSecretKey("osd-sanity-cluster-key", true)
			Expect(err).NotTo(HaveOccurred())
			//Expect(status).TO(BeEquivalentTo("Cluster Secret Key set successfully"))

		})
	})

	Describe("Get Cluster Secret Key", func() {
		It("Should have cluster key set status", func() {
			By("setting cluster wide secret key")
			_, err := manager.GetDefaultSecretKey()
			Expect(err).NotTo(HaveOccurred())
			//Expect(status).TO(BeEquivalentTo("Cluster Secret Key set successfully"))

		})
	})

	Describe("Secret Login", func() {
		It("Should have secret login successful message", func() {
			By("login to secrets store")
			err := manager.Login("aws", nil)
			Expect(err).NotTo(HaveOccurred())

		})
	})

	Describe("Check Secret Login Status", func() {
		It("Should return no error if secrets session is vaild", func() {
			By("check secrets store login session")
			err := manager.CheckLogin()
			Expect(err).NotTo(HaveOccurred())

		})
	})

	Describe("Get Secrets", func() {
		It("Should return secret value for key", func() {
			By("get secrets for given key")
			secretValue, err := manager.Get("osd-id")
			Expect(err).NotTo(HaveOccurred())
			Expect(secretValue).To(BeEquivalentTo("osd-value"))
		})
	})

	Describe("Set Secrets", func() {
		It("Should return set secrets in secret store", func() {
			By("set secrets for given key")
			err := manager.Set("osd-id", "osd-value")
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
