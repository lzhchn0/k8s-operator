# Describe best practice for go development from cloud native perspective

Essence of Cloud Native Computing
- Scalable: Handle growing workloads without degradation.
- Resilient: Recover quickly from failures.
- Efficient: Optimize resource usage and reduce costs.
- Harnessing the power of scale to solve problems while mitigating the inherent challenges of distributed systems.
- **Techniques and Technologies**
     - **Microservices**: Breaking applications into smaller, independent services that can be scaled and deployed independently.
     - **Containers**: Packaging applications and dependencies into lightweight, portable units.
     - **DevOps**: Integrating development and operations to improve collaboration and automation.
     - **Kubernetes**: Orchestrates containers and manages resources at scale.
     - **Service Meshes**: Handle communication between microservices (e.g., Istio, Linkerd).
     - **CI/CD Pipelines**: Automate testing and deployment to ensure rapid, reliable releases.
- Challenges: unpredictability, complexity.
- Benefits: scalability, flexibility.
- **Harness the power of scale**
     - **Elasticity**: Systems can scale up or down dynamically to handle varying workloads.
     - **Distributed computing**: Workloads are spread across multiple nodes, improving performance and fault tolerance.
     - **Automation**: Tools like Kubernetes automate resource management, reducing operational overhead.
- **Problems Caused by Scale**    
     - **Performance bottlenecks**: Increased traffic or data volume can overwhelm systems.
     - **Resource management**: Allocating and managing resources efficiently becomes harder.
     - **Failure points**: Larger systems have more components, increasing the likelihood of failures.
  
equivalent
- curl http://127.0.0.1:8080/apis/apps/v1
- kubectl get --raw /apis/apps/v1
- curl http://127.0.0.1:8080/apis/networking.k8s.io/v1 
- kubectl get --raw /apis/networking.k8s.io/v1 


Typed Client vs. Untyped Client (Dynamic Client)
- Typed client is Type safe and has more readability
- The typed clients are generated automatically from the tools like k8s.io/code-generator
- Typed Client: Works with strongly-typed objects. You need to know the resource type at compile time.
- Typed clients provide methods for creating, reading, updating, and deleting (CRUD) resources. 
- Dynamic Client: Works with unstructured data (map[string]interface{}). It is more flexible and can work with any resource type, but you lose type safety and compile-time checks.
- Example of using a typed client
```go
package main

import (
    "context"
    "fmt"
    "log"
    "path/filepath"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
)

func main() {
    // Load kubeconfig file
    kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatalf("Error building kubeconfig: %v", err)
    }

    // Create a typed client
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalf("Error creating clientset: %v", err)
    }

    // List Pods in the "default" namespace
    pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        log.Fatalf("Error listing pods: %v", err)
    }

    // Print Pod names
    for _, pod := range pods.Items {
        fmt.Println(pod.Name)
    }
}
```
- A **typed client** in `client-go` is a client that works with **strongly-typed objects**. This means that the client is aware of the specific Kubernetes resource types (e.g., `Pod`, `Deployment`, `Service`) and their associated Go structs. When you use a typed client, you interact with Kubernetes resources using these predefined structs, which ensures type safety and makes your code easier to understand and maintain.
- Untyped client (e.g., clientset.CoreV1().RESTClient())
- Typed client (e.g., clientset.CoreV1().Pods(namespace).Get(ctx, name, opts))

  
The convention for package structure is `pkg/apis/<group>/<version>`.


Advanced Features of Custom Resources
- Validation           - Enforce constraints on custom resource fields using OpenAPI v3 schema.      
- Defaulting           - Set default values for fields if not provided by the user.                  
- Subresources         - Enable `/status` and `/scale` subresources for advanced use cases.          
- Webhooks             - Implement custom logic for validation, mutation, and conversion.            
- Versioning           - Support multiple versions of a custom resource with conversion between them.
- Finalizers           - Implement custom cleanup logic during resource deletion.                    
- Controllers/Operators- Manage custom resources and reconcile desired state with actual state.

Internal Types and External Types 
- Multiple API Versions(External):	Support API evolution, backward compatibility, and gradual migration.
- Internal Types:	Provide a unified, efficient representation of resources for internal use.
- Conversion Mechanism:	Translate resources between external versions and internal types.


**Multiple API Versions (`v1`, `v1alpha1`, `v1beta1`)**

Kubernetes APIs evolve over time, and new features or changes are introduced in a controlled manner. To manage this evolution, Kubernetes uses **API versioning**. Here’s why multiple versions are necessary:

**API Evolution and Stability**
- **Stable Versions (`v1`)**: These are production-ready, stable APIs. They are well-tested, backward-compatible, and recommended for use in production environments.
- **Beta Versions (`v1beta1`)**: These APIs are feature-complete but may still undergo changes. They are suitable for testing and experimentation but not guaranteed to be stable.
- **Alpha Versions (`v1alpha1`)**: These APIs are experimental and may change or be removed entirely in future releases. They are not recommended for production use.

By supporting multiple versions, Kubernetes allows users to:
- Use stable APIs for production workloads.
- Experiment with new features in alpha or beta versions.
- Gradually migrate to newer versions as they become stable.

**Backward Compatibility**
- Kubernetes ensures that older API versions remain functional even as new versions are introduced. This allows users to continue using existing tools and configurations without breaking changes.
- For example, if a resource was created using `v1beta1`, it should still be accessible and usable even after `v1` becomes the stable version.

**Gradual Migration**
- Users can migrate their workloads from alpha/beta versions to stable versions at their own pace. Kubernetes provides tools like **API deprecation policies** and **conversion webhooks** to assist with this migration.


**Strictly follow the guidelines for both Global Tag and Local Tag**

The dynamic client uses 
- map[string]interface{}.
- []interface{}.
- string, bool, float64, or int64.

Example of the **Dynamic Client**
```go
package main

import (
    "context"
    "fmt"
    "log"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
    "k8s.io/apimachinery/pkg/runtime/schema"
    "k8s.io/client-go/dynamic"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    // Load kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
    if err != nil {
        log.Fatalf("Error building kubeconfig: %v", err)
    }

    // Create dynamic client
    dynamicClient, err := dynamic.NewForConfig(config)
    if err != nil {
        log.Fatalf("Error creating dynamic client: %v", err)
    }

    // Define the GroupVersionResource (GVR) for the custom resource
    gvr := schema.GroupVersionResource{
        Group:    "example.com",
        Version:  "v1",
        Resource: "myresources",
    }

    // Create an unstructured object for the custom resource
    obj := &unstructured.Unstructured{
        Object: map[string]interface{}{
            "apiVersion": "example.com/v1",
            "kind":       "MyResource",
            "metadata": map[string]interface{}{
                "name":      "my-resource",
                "namespace": "default",
            },
            "spec": map[string]interface{}{
                "replicas": 3,
                "image":    "nginx:1.14.2",
            },
        },
    }

    // Create the custom resource
    createdObj, err := dynamicClient.Resource(gvr).Namespace("default").Create(context.TODO(), obj, metav1.CreateOptions{})
    if err != nil {
        log.Fatalf("Error creating custom resource: %v", err)
    }

    // Access the created resource's data
    fmt.Printf("Created resource: %v\n", createdObj.UnstructuredContent())

    // Retrieve the custom resource
    retrievedObj, err := dynamicClient.Resource(gvr).Namespace("default").Get(context.TODO(), "my-resource", metav1.GetOptions{})
    if err != nil {
        log.Fatalf("Error retrieving custom resource: %v", err)
    }

    // Access the retrieved resource's data
    replicas, found, err := unstructured.NestedInt64(retrievedObj.Object, "spec", "replicas")
    if err != nil || !found {
        log.Fatalf("Error accessing replicas field: %v", err)
    }
    fmt.Printf("Retrieved resource replicas: %d\n", replicas)
}
```

Sample of json.Unmarshal

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	
	data := []byte(`{"name": "John", "age": 30}`)

	// Unmarshal to Person 
	var person Person
	err := json.Unmarshal(data, &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(person)
}
```


API Machinery is the foundation of the Kubernetes API, enabling it to handle resources in a flexible, efficient, and consistent manner. It is essential for building tools, controllers, and clients that interact with Kubernetes.
- Type Definitions:	Common types like ObjectMeta, TypeMeta, and ListMeta.
- Serialization:	Utilities for encoding/decoding resources (JSON, YAML, protobuf).
- Versioning and Conversion:	Tools for handling multiple API versions and converting between them.
- Dynamic Data:	Support for unstructured data (e.g., custom resources).
- Utilities:	Labels, selectors, field selectors, pagination, and watch utilities.

Sample of API Machinery
```go
package main

import (
    "fmt"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/serializer/json"
    "k8s.io/client-go/kubernetes/scheme"
    corev1 "k8s.io/api/core/v1"
)

func main() {
    // Create a Pod object
    pod := &corev1.Pod{
        TypeMeta: metav1.TypeMeta{
            APIVersion: "v1",
            Kind:       "Pod",
        },
        ObjectMeta: metav1.ObjectMeta{
            Name:      "my-pod",
            Namespace: "default",
        },
        Spec: corev1.PodSpec{
            Containers: []corev1.Container{
                {
                    Name:  "nginx",
                    Image: "nginx:1.14.2",
                },
            },
        },
    }

    // Serialize the Pod to JSON
    serializer := json.NewSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, false)
    encoded, err := runtime.Encode(serializer, pod)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Serialized Pod:\n%s\n", string(encoded))

    // Deserialize the JSON back into a Pod object
    decoded, err := runtime.Decode(scheme.Codecs.UniversalDeserializer(), encoded)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Deserialized Pod: %+v\n", decoded)
}
```
