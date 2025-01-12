## client-go library

## controller-runtime library

## deeocopy-gen
## lister-gen
## informer-gen

## Global Tags vs. Local Tags


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
### **How Kubernetes Implements Optimistic Locking**
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
