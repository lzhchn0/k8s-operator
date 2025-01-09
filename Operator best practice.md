## Informer
We need to ensure the informer is synchronized before it is used, specifically prior to the Admission Controller.

### Explaination 1

Certainly! This statement is related to the use of **informers** in Kubernetes admission controllers, specifically in the context of **admission plug-ins**. Let me break it down for you:

### Key Concepts:
1. **Informers**:
   - Informers are part of the Kubernetes client-go library and are used to watch and cache resources (e.g., Pods, Deployments, ConfigMaps) in a Kubernetes cluster.
   - They provide a local cache of the cluster state and notify listeners when changes occur.

2. **Admission Plug-ins**:
   - Admission plug-ins are components in Kubernetes that intercept requests to the API server (e.g., creating or updating a resource) and can either allow or deny the request.
   - Examples include `ValidatingAdmissionWebhook` and `MutatingAdmissionWebhook`.

3. **Synced Informers**:
   - When an informer starts, it needs time to populate its cache with the current state of the cluster. Until this is done, the informer is not considered "synced."
   - A synced informer ensures that the local cache is up-to-date with the cluster's state.

### The Statement Explained:
- **"If informers are used in admission plug-ins, always check first that the informers are synced before using them in the actual Admit() or Validate() functions."**:
  - When writing admission plug-ins, you might rely on informers to make decisions (e.g., checking if a resource already exists or validating against other resources).
  - Before using the informer's cache in the `Admit()` or `Validate()` functions, you must ensure the informer has fully synced. If the informer is not synced, the cache might be incomplete or stale, leading to incorrect decisions.

- **"Reject requests with a Forbidden error before that is the case."**:
  - If the informer is not yet synced, you should reject the admission request with a `Forbidden` error. This is a safety measure to prevent incorrect decisions based on incomplete data.
  - For example, you might return an error like: `admission.Forbidden("informer cache not synced")`.

### Why Is This Important?
- **Data Consistency**:
  - Using an unsynced informer could lead to incorrect decisions, such as allowing a resource creation that violates a policy because the informer's cache was missing critical data.
- **Cluster Stability**:
  - Rejecting requests until the informer is synced ensures that your admission controller operates with accurate and consistent data, maintaining the integrity of the cluster.

### Example in Code:
Here’s a simplified example of how you might implement this:

```go
func (p *MyAdmissionPlugin) Admit(ctx context.Context, admissionSpec *admissionv1.AdmissionRequest) (*admissionv1.AdmissionResponse, error) {
    // Check if the informer is synced
    if !p.myInformer.HasSynced() {
        return &admissionv1.AdmissionResponse{
            Allowed: false,
            Result: &metav1.Status{
                Code:    http.StatusForbidden,
                Message: "Informer cache not synced",
            },
        }, nil
    }

    // Proceed with admission logic using the synced informer
    // ...
}
```

### Summary:
- Always ensure informers are synced before using their cache in admission logic.
- Reject requests with a `Forbidden` error if the informer is not yet synced to avoid making decisions based on incomplete or stale data.
- This practice ensures the reliability and correctness of your admission plug-ins.

### Explaination 2

The need to check if an **informer is synced** depends on the context in which the informer is being used. While it is **critical to check in admission webhooks**, there are other scenarios where you might or might not need to check. Let’s break this down:

---

### **When You Must Check if the Informer is Synced**
1. **Admission Webhooks**:
   - Admission webhooks are invoked synchronously during API requests (e.g., creating or updating a resource).
   - If the informer is not synced, the cache might be incomplete, leading to incorrect decisions (e.g., allowing a request that should be denied or vice versa).
   - Rejecting requests with a `Forbidden` error until the informer is synced ensures correctness and consistency.

2. **Synchronous Operations**:
   - If your code relies on the informer's cache to make immediate decisions (e.g., validating or mutating resources in real-time), you must check if the informer is synced.
   - Examples:
     - A controller that processes events and needs to make decisions based on the current state of the cluster.
     - Any logic that depends on the informer's cache being up-to-date.

---

### **When You Might Not Need to Check if the Informer is Synced**
1. **Asynchronous or Background Processes**:
   - If your code is not making immediate decisions based on the informer's cache, you might not need to check if the informer is synced.
   - For example:
     - A background process that periodically processes data from the informer's cache.
     - A reporting tool that aggregates data over time.

2. **Best-Effort Scenarios**:
   - If your application can tolerate stale or incomplete data (e.g., logging, monitoring, or non-critical operations), you might not need to check if the informer is synced.
   - In such cases, the informer will eventually sync, and your application will start working with accurate data.

3. **Startup Logic**:
   - If your application has a startup phase where it can wait for the informer to sync before performing any critical operations, you might not need to explicitly check in every function.
   - Instead, you can block or wait for the informer to sync during initialization.

---

### **Why Admission Webhooks Are Special**
Admission webhooks are special because:
- They are **synchronous** and **time-sensitive**. The API server waits for the webhook to respond before proceeding with the request.
- They are **critical to cluster operations**. Incorrect decisions (e.g., allowing invalid resources) can destabilize the cluster.
- They **must operate with accurate data**. Using an unsynced informer can lead to incorrect decisions, which is why checking for sync status is mandatory.

---

### **General Best Practices**
1. **Always Check in Critical Paths**:
   - If your code relies on the informer's cache to make decisions, always check if the informer is synced.

2. **Wait for Sync During Initialization**:
   - During application startup, wait for all informers to sync before proceeding with critical operations. This avoids the need to check repeatedly in every function.

   ```go
   if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
       log.Fatal("Failed to sync informer cache")
   }
   ```

3. **Handle Stale Data Gracefully**:
   - If your application can tolerate stale data, design it to handle incomplete or outdated information gracefully.

4. **Log Warnings for Debugging**:
   - If you choose not to check for sync status, log warnings when the informer is not synced to help with debugging.

---

### **Example: When Not to Check**
If you’re writing a background process that aggregates metrics from the informer's cache, you might not need to check if the informer is synced:

```go
func processMetrics(informer cache.SharedIndexInformer) {
    // No need to check if the informer is synced
    items := informer.GetStore().List()
    for _, item := range items {
        // Process items
    }
}
```

---

### **Summary**
- **In admission webhooks**, you **must** check if the informer is synced because decisions are made synchronously and have a direct impact on cluster operations.
- **In other scenarios**, the need to check depends on whether your application relies on the informer's cache for immediate decisions or can tolerate stale data.
- As a general rule, if your code depends on the informer's cache for correctness, always check if it’s synced. Otherwise, you might not need to.
