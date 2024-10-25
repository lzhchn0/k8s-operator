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

```shell
```{{exec}} 

```shell
```{{exec}}

```shell
```{{exec}} 

```shell
```{{exec}}

```shell
```{{exec}} 

```shell
```{{exec}} 
