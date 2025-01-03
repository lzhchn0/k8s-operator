describe best practice for go development from cloud native perspective

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
