
Two kinds of controllers 
- Composite Controller
- Decorator Controller


Her are the key differences between Decorator Controllers and Composite Controllers in Metacontroller:

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

the limitation of metacontroller
