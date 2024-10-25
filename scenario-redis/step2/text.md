Setup a appropriate environment.

step 1: create a Redis cluster

```shell
mkdir foo && cd foo
helm repo add bitnami  https://charts.bitnami.com/bitnami
helm repo update
```{{exec}}

```shell
cat <<EOF  > values.yaml
# Create a values.yaml file for Redis cluster configuration
architecture: replication
auth:
  enabled: true
  password: "pwd" 
  
cluster:
  enabled: true
  slaveCount: 3

metrics:
  enabled: true

persistence:
  enabled: true
  size: 8Gi

service:
  type: ClusterIP

# Redis cluster configuration
master:
  resources:
    requests:
      memory: 100Mi
      cpu: 100m
    limits:
      memory: 100Mi
      cpu: 200m
  persistence:
    enabled: true
    size: 8Gi

replica:
  replicaCount: 3
  resources:
    requests:
      memory: 100Mi
      cpu: 100m
    limits:
      memory: 100Mi
      cpu: 200m
  persistence:
    enabled: true
    size: 8Gi

EOF
```{{exec}} 

```shell
helm install my-redis bitnami/redis -f values.yaml -n redis --create-namespace
```{{exec}}

Check Health
```shell
kn redis
kubectl get pods 
```{{exec}} 


upload redis operator from ~/boo/redis-operator-000/bin/manager
```shell
```{{exec}}

NAME: my-redis
LAST DEPLOYED: Fri Oct 25 14:08:32 2024
NAMESPACE: redis
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: redis
CHART VERSION: 20.2.1
APP VERSION: 7.4.1

** Please be patient while the chart is being deployed **

Redis&reg; can be accessed on the following DNS names from within your cluster:

    my-redis-master.redis.svc.cluster.local for read/write operations (port 6379)
    my-redis-replicas.redis.svc.cluster.local for read-only operations (port 6379)



To get your password run:

    export REDIS_PASSWORD=$(kubectl get secret --namespace redis my-redis -o jsonpath="{.data.redis-password}" | base64 -d)

To connect to your Redis&reg; server:

1. Run a Redis&reg; pod that you can use as a client:

   kubectl run --namespace redis redis-client --restart='Never'  --env REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:7.4.1-debian-12-r0 --command -- sleep infinity

   Use the following command to attach to the pod:

   kubectl exec --tty -i redis-client \
   --namespace redis -- bash

2. Connect using the Redis&reg; CLI:
   REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h my-redis-master
   REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h my-redis-replicas

To connect to your database from outside the cluster execute the following commands:

    kubectl port-forward --namespace redis svc/my-redis-master 6379:6379 &
    REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h 127.0.0.1 -p 6379

WARNING: There are "resources" sections in the chart not set. Using "resourcesPreset" is not recommended for production. For production installations, please set the following values according to your workload needs:
  - metrics.resources
+info https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


```shell
```{{exec}} 

```shell
```{{exec}}

```shell
```{{exec}} 

```shell
```{{exec}} 
