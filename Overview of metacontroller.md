
Two kinds of controllers 
- Composite Controller
- Decorator Controller

Decorator Controller has three examples
- clusteredparent
- service-per-pod (jsonnet)
- crd-roles

Here are the key differences between Decorator Controllers and Composite Controllers in Metacontroller:

Composite Controller:
- Creates and manages complex resources that don't exist directly in Kubernetes
- Manages entire lifecycle of custom resources
- Can create, update, and delete related resources based on custom logic
- More powerful and flexible than Decorator Controllers
- Implements full control loops for custom resource management
- Often used for creating higher-level abstractions or implementing complex operational patterns

Decorator Controller:
- Modifies or adds behavior to existing Kubernetes resources without changing their core definition
- Intercepts and can mutate resources before they are persisted
- Typically used for adding labels, annotations, or making minor transformations
- Operates in a "wrapper" or "enhancement" mode
- Does not fundamentally change the resource's core functionality
- Best for lightweight, non-invasive modifications

The constraints of metacontroller
- Do not support operations across multiple clusters
- Do not support dependencies between complex stateful resources
- Kubebuilder watches bigger scope than metacontroller does.
  

Informer behind metacontroller
- The informer periodically flushes the cache and refreshes the objects from the API server. Disabling this mechanism is not feasible.
- The updated cache will trigger the reconciliation loop for each object in the cache. This mechanism helps the controller prevent any drift between the actual state and the desired state.
- The configurable resync period
- Never attempt to modify an object owned by an informer; instead, create a deepcopy of it.

Manage multiple CRDs 
- As a rule of best practice, a single controller typically manages one CRD, but a metacontroller runtime instance can manage multiple CRDs at the same time.

Webhook
- It is good practice to ensure that your webhooks are always fully idempotent.

The sync webhook is triggered when:
1. A change occurs in the parent resource (create/update/delete)
2. A change occurs in any of the child resources
3. A change occurs in any of the related resources specified in the controller configuration

Sync webhook request sample
```json
{
    "controller": {
        "apiVersion": "metacontroller.k8s.io/v1alpha1",
        "kind": "CompositeController",
        "metadata": {
            "name": "my-controller",
            "uid": "1234abcd-5678-efgh..."
        }
    },
    "parent": {
        "apiVersion": "example.com/v1",
        "kind": "MyResource",
        "metadata": {
            "name": "example-resource",
            "namespace": "default",
            "uid": "5678efgh-1234-abcd...",
            "generation": 1
        },
        "spec": {
            # ... parent resource spec
        },
        "status": {
            # ... current status if any
        }
    },
    "children": {
        "Deployment.apps/v1": {
            "deployment-name": {
                "apiVersion": "apps/v1",
                "kind": "Deployment",
                "metadata": {
                    "name": "deployment-name",
                    "namespace": "default"
                },
                "spec": {
                    # ... deployment spec
                }
            }
        }
    },
    "related": {
        "ConfigMap.v1": {
            "config-name": {
                "apiVersion": "v1",
                "kind": "ConfigMap",
                "metadata": {
                    "name": "config-name",
                    "namespace": "default"
                },
                "data": {
                    # ... configmap data
                }
            }
        }
    }
    "finalizing": {
      # ... 
    }
}
```

Sync webhook response

Through mysterious webhook to develop controller
- multiple patterns/nodes behind controller

Array is the key
- in message, CompositeHookResponse.children of type []*unstructured.Unstructured, array is the key. 

If there is an update to the child, metacontroller-0 will log an error message.
- "msg": "Failed to create child, child object already exists",
- Solution is, when updating an existing object, you must include the resourceVersion field in the metadata. This ensures that Kubernetes can handle concurrent updates correctly. The resourceVersion is obtained from the existing object.
- [Update Methods, updateStrategy](https://metacontroller.github.io/metacontroller/api/decoratorcontroller.html?highlight=inplace#attachment-update-methods)

How to debug metacontroller
- check logs in metacontroller-0 pod, search keyword like "error", "can't find resource" 
- check logs in parent resources
- check logs in child resources
- check logs in related resources

Worker in metacontroller
- A worker is a process that handles a given webhook request. Its command line parameter is "--workers", refer to [Here](https://metacontroller.github.io/metacontroller/guide/configuration.html).
- By default, the queue is processed by 5 workers.

Additional Readings
- [hypergig](https://github.com/GoogleCloudPlatform/metacontroller/issues/190)
