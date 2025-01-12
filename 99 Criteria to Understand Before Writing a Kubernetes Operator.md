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
