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
	"bytes"
	"context"
	"embed"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"
	"testing"

	ctrl "sigs.k8s.io/controller-runtime"

	"k8s.io/client-go/kubernetes/scheme"
	"text/template"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"k8s.io/apimachinery/pkg/runtime/schema"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtime "k8s.io/apimachinery/pkg/runtime"
	_      "k8s.io/apimachinery/pkg/runtime/schema"

	"sigs.k8s.io/controller-runtime/pkg/client"

	mynginxv1 "blabla.com/api/v1"
)

var (

	//go:embed manifests/*
	manifests embed.FS

	appsScheme = runtime.NewScheme()
	appsCodecs = serializer.NewCodecFactory(appsScheme)
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

func TestListService(t *testing.T) {
	fmt.Println("== 009 ListService  " )

	// v1alpha1  -->  mynginxv1
	nginx := &mynginxv1.Monitoring{ObjectMeta: metav1.ObjectMeta{Name: "monitoring-sample", Namespace: "default"}}
	_ = nginx

	s1, err := getService("default", "my-nginx-service", 9091)
	fmt.Println("==099-0 dump Service  " )
	scs3.Dump(s1)

	_ = err
	var resources []runtime.Object
	resources = append(resources, s1)

}

// Ref : https://itnext.io/kubernetes-custom-controllers-recipes-for-beginners-bbc286c05ef8
//       Kubernetes Custom Controllers Recipes for Beginners
//       Recipe.8 - Create objects from templates

func getService(namespace string, name string, port int) (*corev1.Service, error) {
	metadata := struct {
		Namespace string
		Name      string
		Port      int
	}{
		Namespace: namespace,
		Name:      name,
		Port:      port,
	}

	object, err := getObject("service", corev1.SchemeGroupVersion, metadata)
	if err != nil {
		return nil, err
	}

	return object.(*corev1.Service), nil
}

func getObject(name string, gv schema.GroupVersion, metadata any) (runtime.Object, error) {
	parse, err := getTemplate(name)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = parse.Execute(&buffer, metadata)

defer	scs3.Dump(err)
defer	scs3.Dump(buffer)
defer 	fmt.Println("==099-obj dump Object's buffer" ) 


	if err != nil {
		return nil, err
	}
	
	// step 1
	corev1.AddToScheme(appsScheme)
	// step 2
	gv = schema.GroupVersion{Group: "", Version: "v1"}

	object, err := runtime.Decode(
		appsCodecs.UniversalDecoder(gv),
		buffer.Bytes(),
	)


defer	scs3.Dump(err)
defer	scs3.Dump(object)
defer 	fmt.Println("==099-obj dump Object" ) 



	return object, nil
}

func getTemplate(name string) (*template.Template, error) {
	manifestBytes, err := manifests.ReadFile(fmt.Sprintf("manifests/%s.yaml", name))
	if err != nil {
		return nil, err
	}

	tmp := template.New(name)
	parse, err := tmp.Parse(string(manifestBytes))
	if err != nil {
		return nil, err
	}

	return parse, nil
}
