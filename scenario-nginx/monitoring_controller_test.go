/*
Copyright 2024.

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

package controller

import (
	"context"
	"os"
	"testing"

	ctrl "sigs.k8s.io/controller-runtime"

	"k8s.io/client-go/kubernetes/scheme"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"

	mynginxv1 "blabla.com/api/v1"
)

var _ = Describe("Monitoring Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default", // TODO(user):Modify as needed
		}
		monitoring := &mynginxv1.Monitoring{}

		// == to tear up CP
		kubeconfig := os.Getenv("KUBECONFIG")

		if kubeconfig == "" {
			panic("KUBECONFIG environment variable must be set")
		}

		cfg, err := ctrl.GetConfig()

		if err != nil {
			panic(" ---- get config of KUBECONFIG is nil ")
		}

		k8sClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})

		if err != nil {
			panic("==9900==Cleanup the specific resource instance Monitoring")
		}

		BeforeEach(func() {
			By("creating the custom resource for the Kind Monitoring")
			err := k8sClient.Get(ctx, typeNamespacedName, monitoring)
			if err != nil && errors.IsNotFound(err) {
				resource := &mynginxv1.Monitoring{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					// TODO(user): Specify other spec details if needed.
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			// TODO(user): Cleanup logic after each test, like removing the resource instance.
			resource := &mynginxv1.Monitoring{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance Monitoring")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})

		It("should successfully reconcile the resource", func() {
			By("Reconciling the created resource")
			controllerReconciler := &MonitoringReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
			// Example: If you expect a certain status condition after reconciliation, verify it here.
		})
	})
})

func TestListIngress(t *testing.T) {

	// v1alpha1  -->  mynginxv1
	nginx := &mynginxv1.Monitoring{ObjectMeta: metav1.ObjectMeta{Name: "my-nginx", Namespace: "default"}}
	_ = nginx

}
