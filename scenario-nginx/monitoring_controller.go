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
	"fmt"
	"reflect"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	nv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	mv1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/record"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mynginxv1 "blabla.com/api/v1"

	"github.com/davecgh/go-spew/spew"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// MonitoringReconciler reconciles a Monitoring object
type MonitoringReconciler struct {
	client.Client
	Scheme           *runtime.Scheme
	Event1           record.EventRecorder
	AnnotationFilter labels.Selector
	Log              logr.Logger
}

var scs3 = spew.ConfigState{
	Indent:                  "    ", // 索引为 Tab
	DisableMethods:          true,
	DisablePointerMethods:   true,
	DisablePointerAddresses: true,
	MaxDepth:                4, // 设置打印深度为 1

}

// +kubebuilder:rbac:groups=mynginx.opencanon.io,resources=monitorings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mynginx.opencanon.io,resources=monitorings/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=mynginx.opencanon.io,resources=monitorings/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Monitoring object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *MonitoringReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logr := log.FromContext(ctx)

	var instance = &mynginxv1.Monitoring{}

	err := r.Client.Get(ctx, req.NamespacedName, instance)

	if apierrors.IsNotFound(err) {
		logr.Info("== Nginx resource not found, skipping reconcile")
		return ctrl.Result{}, nil

	}

	if err != nil {

		logr.Error(err, "== Unable to get the instance ")
		return ctrl.Result{}, err

	}

	// TODO(user): your logic here
	err = r.reconcileMonitor(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	if err := r.refreshStatus(ctx, instance); err != nil {
		logr.Error(err, "== Fail to refresh status subresource")
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

func (r *MonitoringReconciler) reconcileMonitor(ctx context.Context, m *mynginxv1.Monitoring) error {
	err := r.reconcileDeployment(ctx, m)
	err = r.reconcileService(ctx, m)
	err = r.reconcileIngress(ctx, m)

	return err
}

func (r *MonitoringReconciler) reconcileService(ctx context.Context, m *mynginxv1.Monitoring) error {
	// Fetch the Service instance

	// r.Event1.Eventf(m, corev1.EventTypeNormal, "ServiceUpdated", "service updated successfully")
	logr := log.FromContext(ctx)

	// address violation
	if r != nil && m != nil && r.Event1 != nil {
		r.Event1.Event(m, corev1.EventTypeNormal, "ServiceUpdated", "== service updated successfully")
	}

	newService := &corev1.Service{

		ObjectMeta: mv1.ObjectMeta{
			Name:      "my-new-service",
			Namespace: "default",
			Labels: map[string]string{
				"app": "nginx",
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "nginx",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.ProtocolTCP,
					Port:     80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
				},
			},
		},
	}

	svc := &corev1.Service{}
	err := r.Get(ctx, client.ObjectKey{
		Namespace: "default",
		Name:      "my-service",
	}, svc)

	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			return err
		}
		// Service not found, could have been deleted after reconcile request
		logr.Info("== 666 not found")

		err = r.Client.Create(ctx, newService)

		return nil
	}

	// Check and update the Service if necessary
	updated := false
	if svc.Spec.Type != corev1.ServiceTypeClusterIP {
		svc.Spec.Type = corev1.ServiceTypeClusterIP
		updated = true
	}

	if updated {
		err = r.Update(ctx, svc)
		if err != nil {
			// return ctrl.Result{}, err
			return err
		}
	}

	// Reconcile logic for other resources can go here

	// return ctrl.Result{}, nil
	return nil

}
func (r *MonitoringReconciler) reconcileIngress(ctx context.Context, m *mynginxv1.Monitoring) error {

	// Define the Ingress resource
	mying := &nv1.Ingress{
		ObjectMeta: mv1.ObjectMeta{
			Name:      "example-ingress",
			Namespace: "default",
			Labels: map[string]string{
				"app": "nginx",
				"env": "coop-sit6",
			},
		},
		Spec: nv1.IngressSpec{
			Rules: []nv1.IngressRule{
				{
					Host: "example.com",
					IngressRuleValue: nv1.IngressRuleValue{
						HTTP: &nv1.HTTPIngressRuleValue{
							Paths: []nv1.HTTPIngressPath{
								{
									Path: "/",
									PathType: func() *nv1.PathType {
										pt := nv1.PathTypePrefix
										return &pt
									}(),
									Backend: nv1.IngressBackend{
										Service: &nv1.IngressServiceBackend{
											Name: "example-service",
											Port: nv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	logr := log.FromContext(ctx)
	logr.Info("== 333-Ingress ")

	currentIngress := &nv1.Ingress{}
	err := r.Client.Get(ctx, client.ObjectKey{Name: mying.Name, Namespace: mying.Namespace}, currentIngress)
	if apierrors.IsNotFound(err) {

		return r.Client.Create(ctx, mying)
	}

	return nil
}

func int32Ptr(i int32) *int32 {
	return &i
}

// , nginx *nginxv1alpha1.Nginx

func (r *MonitoringReconciler) reconcileDeployment(ctx context.Context, m *mynginxv1.Monitoring) error {

	logr := log.FromContext(ctx)

	labels := map[string]string{"app": "nginx"}

	newD := &appsv1.Deployment{

		ObjectMeta: mv1.ObjectMeta{

			Name:      "my-service",
			Namespace: "default",
			Labels: map[string]string{
				"app": "nginx",
				"env": "coop-deploy-sit6",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &mv1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: mv1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.14.2",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	_ = newD
	_ = m

	// Print the Deployment spec for debugging
	fmt.Printf("Deployment Spec: %+v\n", newD.Spec)

	foundD := &appsv1.Deployment{}

	err := r.Get(ctx, client.ObjectKey{Name: newD.Name, Namespace: newD.Namespace}, foundD)

	if err == nil {
		return err
	} else {

		if apierrors.IsNotFound(err) {
			logr.Info("== not found Deploy ")
			// err = r.Create(ctx, newD)
			err = r.Client.Create(ctx, newD)

			logr.Error(err, "== my first error message")
		} else {

		}
	}

	// _ = foundD

	if reflect.DeepEqual(newD.Spec, foundD.Spec) {
		return nil
	} else {
		foundD.Spec = newD.Spec
		err = r.Update(ctx, foundD)
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MonitoringReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mynginxv1.Monitoring{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&nv1.Ingress{}).
		Complete(r)
}

func (r *MonitoringReconciler) refreshStatus(ctx context.Context, m *mynginxv1.Monitoring) error {

	deploys, err := r.listDeployments(ctx, r.Client, m)
	if err != nil {
		return err
	}

	svcs, err := r.listServices(ctx, r.Client, m)
	if err != nil {
		return err
	}

	ings, err := r.listIngresses(ctx, r.Client, m)
	if err != nil {
		return err
	}

	status := mynginxv1.MonitoringStatus{
		CurrentReplicas: 5,
		//PodSelector:     k8s.LabelsForNginxString(nginx.Name),
		Deployments: deploys,
		Services:    svcs,
		Ingresses:   ings,
	}

	_ = deploys
	_ = status
	return nil
}

func (r *MonitoringReconciler) listDeployments(ctx context.Context, c client.Client, m *mynginxv1.Monitoring) ([]mynginxv1.DeploymentStatus, error) {
	//var deployList appsv1.DeploymentList
	// Define the label selector
	labelSelector := client.MatchingLabels{"app": "nginx"}

	logr := log.FromContext(ctx)

	logr.Info("== list--Deployments ")

	// Define the namespace
	namespace := "default"

	// List the deployments
	deploymentList := &appsv1.DeploymentList{}
	err := r.List(ctx, deploymentList, client.InNamespace(namespace), labelSelector)
	if err != nil {
		// logger.Error(err, "Failed to list deployments")
		return nil, err
	}

	logr.Info("==DUMP-- deploymentList")
	scs3.Dump(deploymentList)

	// Print the names of the deployments
	for _, deployment := range deploymentList.Items {
		fmt.Printf("== Deployment Name: %s\n", deployment.Name)
	}

	deploys := deploymentList.Items

	var deployStatuses []mynginxv1.DeploymentStatus
	var replicas int32
	for _, d := range deploys {
		replicas += d.Status.Replicas
		deployStatuses = append(deployStatuses, mynginxv1.DeploymentStatus{Name: d.Name})
	}

	logr.Info("==DUMP-- deploy status")
	scs3.Dump(deployStatuses)

	return deployStatuses, nil
}

func (r *MonitoringReconciler) listServices(ctx context.Context, c client.Client, m *mynginxv1.Monitoring) ([]mynginxv1.ServiceStatus, error) {
	// Define the label selector
	labelSelector := client.MatchingLabels{"app": "nginx"}

	// Define the namespace
	namespace := "default"

	logr := log.FromContext(ctx)
	logr.Info("== list--Services ")

	// List the services
	serviceList := &corev1.ServiceList{}
	err := r.List(ctx, serviceList, client.InNamespace(namespace), labelSelector)
	if err != nil {
		// logger.Error(err, "Failed to list services")
		return nil, err
	}

	// Print the names of the services
	for _, service := range serviceList.Items {
		fmt.Printf("Service Name: %s\n", service.Name)
	}

	var svcs []mynginxv1.ServiceStatus
	for _, s := range serviceList.Items {
		svc := mynginxv1.ServiceStatus{
			Name: s.Name,
		}

		svcs = append(svcs, svc)
	}

	logr.Info("== DUMP--serviceList ")
	scs3.Dump(svcs)

	return svcs, nil
}

func (r *MonitoringReconciler) listIngresses(ctx context.Context, c client.Client, m *mynginxv1.Monitoring) ([]mynginxv1.IngressStatus, error) {

	logr := log.FromContext(ctx)
	logr.Info("== 333--list--Ingress ")

	// Define the label selector
	labelSelector := client.MatchingLabels{"app": "nginx"}

	// Define the namespace
	namespace := "default"

	// List the ingresses
	ingressList := &nv1.IngressList{}
	err := r.List(ctx, ingressList, client.InNamespace(namespace), labelSelector)
	if err != nil {
		//logger.Error(err, "Failed to list ingresses")
		return nil, err
	}

	logr.Info("== DUMP -- ingressList ")
	scs3.Dump(ingressList)

	// Print the names of the ingresses
	for _, ingress := range ingressList.Items {
		fmt.Printf("Ingress Name: %s\n", ingress.Name)
	}

	logr.Info("== DUMP -- ingressList ")
	scs3.Dump(ingressList)

	var ingresses []mynginxv1.IngressStatus
	for _, i := range ingressList.Items {
		ing := mynginxv1.IngressStatus{Name: i.Name}

		ingresses = append(ingresses, ing)
	}

	logr.Info("==DUMP -- ingresses ")
	scs3.Dump(ingresses)

	return ingresses, nil
}
