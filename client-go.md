## Clients
- **RESTClient**: A low-level client for making raw HTTP requests to the Kubernetes API.
- **Clientset**: A high-level, type-safe client for interacting with core Kubernetes resources (e.g., Pods, Deployments).
- **Dynamic Client**: A client for interacting with unstructured data (e.g., custom resources).
- **Discovery Client**: A client for discovering API resources and versions supported by the Kubernetes API server.


## Informers and Caching	
- Introduced **informers** and **shared informer factories** for efficient resource watching.
- **Informers**: Watch and cache Kubernetes resources, reducing API server load.
- **SharedInformerFactory**: Allows multiple controllers to share the same cache.
- **Lister**: Provides read-only access to cached resources.

## Dynamic Client	
Added support for unstructured data and custom resources.

## Code Generation	
Introduced tools for generating clients and informers for custom resources.

## Testing Utilities	
Added fake clients and testing utilities.

## Context Support	
Integrated context.Context for better cancellation and timeout handling.

## Operator Pattern	
Became the foundation for building Kubernetes operators.

Sample of using client-go (Type-Safe Client)
```go
clientset, err := kubernetes.NewForConfig(config)
if err != nil {
    log.Fatal(err)
}

pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
if err != nil {
    log.Fatal(err)
}

for _, pod := range pods.Items {
    fmt.Println(pod.Name)
}
```
## Workqueues: 
For handling events and reconciling desired states.

## Tools and Plugins:
Tools like kubectl, Helm, and **cluster management tools** rely on client-go.

## important packages 
- "k8s.io/client-go"
- "k8s.io/client-go/discovery"
- "k8s.io/client-go/dynamic"
- "k8s.io/client-go/dynamic/dynamiclister"
- "k8s.io/client-go/dynamic/fake"
- "k8s.io/client-go/kubernetes"
- "k8s.io/client-go/rest"
- "k8s.io/client-go/tools/cache"
- "k8s.io/client-go/tools/leaderelection/resourcelock"
- "k8s.io/client-go/tools/record"
- "k8s.io/client-go/util/retry"
- "k8s.io/client-go/util/workqueue"
- clientgo_cache "k8s.io/client-go/tools/cache"
- clientgotesting "k8s.io/client-go/testing"
- clientsetscheme "k8s.io/client-go/kubernetes/scheme"
- **discovery** "k8s.io/client-go/discovery"
- fakeclientset "k8s.io/client-go/kubernetes/fake"
- fakediscovery "k8s.io/client-go/discovery/fake"
- flowcontrol "k8s.io/client-go/util/flowcontrol"
- rest "k8s.io/client-go/rest"
- typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

## REST Mapping
- **GVK**            Group-Version-Kind: Identifies the type of resource (e.g., `apps/v1.Deployment`).
- **GVR**            Group-Version-Resource: Identifies the REST endpoint for the resource (e.g., `GVR: apps/v1/deployments; Scope: Namespaced`).
- **REST Mapping**   The process of mapping a GVK to a GVR and determining the resource's scope.
- **RESTMapper**     A tool in `client-go` that performs REST mapping.

## **Kind vs. Resource** (Pluralization Rules)

The pluralization of Kinds to Resources follows standard English rules, but there are some exceptions. Here are a few examples:

| **Kind**       | **Resource**   |
|----------------|----------------|
| `Deployment`   | `deployments`  |
| `Pod`          | `pods`         |
| `Service`      | `services`     |
| `Namespace`    | `namespaces`   |
| `ConfigMap`    | `configmaps`   |
| `Ingress`      | `ingresses`    |

- Kind used in YAML/JSON manifests and Go code.
- Resource used in REST API endpoints.
- Kind is used in YAML/JSON manifests and Go code, where readability and clarity are important. PascalCase is more readable and aligns with programming conventions.
- Resource is used in REST API endpoints, where lowercase and pluralization are standard conventions for URLs.



## Sample of using RESTMapper
```go
package main

import (
    "fmt"
    "log"

    "k8s.io/apimachinery/pkg/api/meta"
    "k8s.io/client-go/discovery"
    "k8s.io/client-go/discovery/cached/memory"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    // Load kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
    if err != nil {
        log.Fatalf("Error building kubeconfig: %v", err)
    }

    // Create a discovery client
    discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
    if err != nil {
        log.Fatalf("Error creating discovery client: %v", err)
    }

    // Create a RESTMapper
    restMapper := memory.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))

    // Define a GVK
    gvk := schema.GroupVersionKind{
        Group:   "apps",
        Version: "v1",
        Kind:    "Deployment",
    }

    // Perform REST mapping
    mapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
    if err != nil {
        log.Fatalf("Error performing REST mapping: %v", err)
    }

    // Print the GVR and scope
    fmt.Printf("GVR: %s/%s/%s\n", mapping.Resource.Group, mapping.Resource.Version, mapping.Resource.Resource)
    fmt.Printf("Scope: %s\n", mapping.Scope.Name())
}
```

## **TypeMeta**
### **Fields**
- **`APIVersion`**: The API version of the object (e.g., `apps/v1`, `batch/v1`).
- **`Kind`**: The type of the object (e.g., `Pod`, `Deployment`, `Service`).

### **Example**
```go
type TypeMeta struct {
    APIVersion string `json:"apiVersion,omitempty"`
    Kind       string `json:"kind,omitempty"`
}
```

## **ObjectMeta** 


### **Fields**
- **`Name`**: The name of the object.
- **`Namespace`**: The namespace to which the object belongs (if namespaced).
- **`Labels`**: Key-value pairs used for identifying and organizing objects.
- **`Annotations`**: Key-value pairs used for storing non-identifying metadata.
- **`OwnerReferences`**: References to the object’s owners (e.g., a Deployment owning a ReplicaSet).
- **`ResourceVersion`**: A version identifier used for optimistic concurrency control.
- **`UID`**: A unique identifier for the object.

### **Example**
```go
type ObjectMeta struct {
    Name            string            `json:"name,omitempty"`
    Namespace       string            `json:"namespace,omitempty"`
    Labels          map[string]string `json:"labels,omitempty"`
    Annotations     map[string]string `json:"annotations,omitempty"`
    OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty"`
    ResourceVersion string            `json:"resourceVersion,omitempty"`
    UID             types.UID         `json:"uid,omitempty"`
}
```

### **Key Differences**

| **Aspect**            | **TypeMeta**                                      | **ObjectMeta**                                      |
|------------------------|---------------------------------------------------|-----------------------------------------------------|
| **Purpose**            | Describes the **type** of the object.             | Describes the **instance** of the object.           |
| **Fields**             | `APIVersion`, `Kind`.                             | `Name`, `Namespace`, `Labels`, `Annotations`, etc.  |
| **Embedded In**        | Top-level structure of Kubernetes objects.        | `metadata` field of Kubernetes objects.             |
| **Used For**           | Identifying and processing the object type.       | Managing the object’s lifecycle and identity.       |
| **Example**            | `apiVersion: apps/v1`, `kind: Deployment`.        | `name: my-deployment`, `namespace: default`.        |
  



## **`client-go`-based applications remain compatible with a Kubernetes cluster**

| **Aspect**               | **Details**                                                                 |
|---------------------------|-----------------------------------------------------------------------------|
| **Version Alignment**     | `client-go` versions align with Kubernetes versions (e.g., `v0.22.x` for `v1.22.x`). |
| **Backward Compatibility**| `client-go` can work with older Kubernetes versions within the same major version. |
| **Forward Compatibility** | `client-go` is not forward-compatible with newer Kubernetes versions.       |
| **Skew Policy**           | `client-go` can be up to two minor versions older than the Kubernetes API server. |
| **Best Practices**        | Pin `client-go` version, test with target Kubernetes version, stay updated. |


