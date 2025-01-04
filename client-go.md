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
