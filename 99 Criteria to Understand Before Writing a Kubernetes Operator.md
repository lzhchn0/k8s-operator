## client-go library

## controller-runtime library

## deeocopy-gen
## lister-gen
## informer-gen

## Global Tags vs. Local Tags

- **Global Tags Definition:**
  - Global tags are annotations applied at the **package level** (typically in a `pkg/apis/group/version/doc.go` file before "package v1alpha1).
  - They affect all types within the package unless overridden by local tags.

- **Global Tags Purpose:**
  - They provide default behavior for code generation tools like `deepcopy-gen`, `informer-gen`, and `client-gen` across the entire package.

- **Common Global Tags:**
  - `+k8s:deepcopy-gen=package`: Enable deepcopy generation for all types in the package.
  - `+groupName=<group>`: Specify the API group for the package.
  - `+k8s:openapi-gen=true`: Enable OpenAPI schema generation for the package.
  - `+k8s:conversion-gen=<package>`: Enable conversion functions for the package.

- **Local Tags Definition:**
  - Local tags are annotations applied to **individual types or fields** within a Go file.
  - They override or supplement the behavior specified by global tags.

- **Local Tags Purpose:**
  - They allow fine-grained control over code generation for specific types or fields.

- **Common Local Tags:**
  - `+k8s:deepcopy-gen=true/false`: Enable or disable deepcopy generation for a specific type.
  - `+k8s:deepcopy-gen:interfaces=<interface>`: Specify that the type implements a particular interface (e.g., `runtime.Object`).
  - `+kubebuilder:validation:Required`: Mark a field as required in the OpenAPI schema.
  - `+kubebuilder:subresource:status`: Enable a status subresource for the type.


## Let Informer to watch resources in a limited scope. 

## Observability is crucial.
- Once you have full visibility into your operator, nothing will stand in the way of your success.
- Spew is an excellent tool for gaining the visibility you need.
- 
## It's better to rewrite your operator than to spend time debugging it.

## Opt for incremental builds rather than full builds in Operator development.

## As an escalation path for operator development, begin with Metacontroller, then progress to Operator SDK, and ultimately Kubebuilder. This approach will grant you progressively greater control over every aspect of Custom Resources.

## Run webhook in your cluster, not outside the cluster, Otherwise there will be
1. **Networking Challenges**:
   - The API server needs to reach the webhook's IP address and port, which may not be routable from the cluster.
   - Firewalls or network policies may block incoming connections.

2. **Certificate Challenges**:
   - The webhook endpoint must use HTTPS with a valid TLS certificate.
   - Self-signed certificates are not trusted by the API server unless explicitly configured.

3. **Dynamic Environments**:
   - If the operator is running on a laptop or in a CI/CD pipeline, the IP address or hostname may change frequently, making it difficult to configure the webhook URL.

## Operator SDK: All your defined custom resources may be reprocessed on startup

## The resourceVersion field is used to detect conflicts and ensure consistency
### **1. How Kubernetes Implements Optimistic Locking**
Kubernetes uses the `resourceVersion` field in the metadata of a resource to implement optimistic locking. Here’s how it works:

#### **a. The `resourceVersion` Field**
- Every Kubernetes resource (e.g., Deployment, Pod, CustomResource) has a `metadata.resourceVersion` field.
- This field is a string that represents the **current version of the resource**.
- It is updated by the Kubernetes API server every time the resource is modified.

#### **b. How Updates Work**
1. **Read the Resource**:
   - When you fetch a resource (e.g., using `kubectl get` or the Kubernetes API), you get the current state of the resource, including its `resourceVersion`.

2. **Modify the Resource**:
   - You make changes to the resource (e.g., update a field in a Deployment).

3. **Send the Update Request**:
   - When you send the updated resource back to the Kubernetes API server, the `resourceVersion` field is included in the request.

4. **Conflict Detection**:
   - The API server compares the `resourceVersion` in your request with the current `resourceVersion` of the resource in its database.
     - If the `resourceVersion` matches, it means no one else has modified the resource since you last read it. The update is applied, and the `resourceVersion` is incremented.
     - If the `resourceVersion` does **not** match, it means someone else has modified the resource in the meantime. The API server rejects your update with a **409 Conflict** error.
### **2. Practical Implications**
When working with Kubernetes resources (especially in controllers or operators), you need to handle the `resourceVersion` field correctly to avoid conflicts and ensure smooth updates.

#### **a. Always Use the Latest `resourceVersion`**
- Before updating a resource, always fetch the latest version of the resource from the API server.
- Use the latest `resourceVersion` in your update request to avoid conflicts.

#### **b. Handle Conflicts Gracefully**
- If you receive a **409 Conflict** error, it means someone else has modified the resource. In this case:
  1. Fetch the latest version of the resource.
  2. Reapply your changes to the latest version.
  3. Retry the update.

#### **c. Example in Code**
Here’s an example of how you might handle updates in a Kubernetes controller or operator:

```go
func updateDeployment(client kubernetes.Interface, deployment *appsv1.Deployment) error {
    for {
        // Fetch the latest version of the Deployment
        latestDeployment, err := client.AppsV1().Deployments(deployment.Namespace).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
        if err != nil {
            return err
        }

        // Apply your changes to the latest Deployment
        latestDeployment.Spec.Replicas = deployment.Spec.Replicas

        // Try to update the Deployment
        _, err = client.AppsV1().Deployments(deployment.Namespace).Update(context.TODO(), latestDeployment, metav1.UpdateOptions{})
        if err == nil {
            // Update succeeded
            return nil
        }

        // Check if the error is a conflict
        if !apierrors.IsConflict(err) {
            // Handle other errors
            return err
        }

        // If it's a conflict, retry the update
    }
}
```
