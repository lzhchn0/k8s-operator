
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

Informer behind metacontroller
- The informer periodically flushes the cache and refreshes the objects from the API server. Disabling this mechanism is not feasible.
- The updated cache will trigger the reconciliation loop for each object in the cache. This mechanism helps the controller prevent any drift between the actual state and the desired state.
- the configurable resync period


Webhook
- It is good practice to ensure that your webhooks are always fully idempotent.

Worker in metacontroller
- A worker is a process that handles a given webhook request. command line parameter is "--workers"
- By default, the queue is processed by 5 workers.
