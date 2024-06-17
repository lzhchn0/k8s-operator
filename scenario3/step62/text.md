spotahome setup 


```
git init
git clone https://github.com/spotahome/redis-operator.git
```{{exec}}



```
k apply -f basic.yaml
k apply -f bootstrapping-with-port.yaml
k apply -f bootstrapping-with-sentinels.yaml
k apply -f bootstrapping.yaml
k apply -f container-security-context.yaml
k apply -f control-label-propagation.yaml
k apply -f custom-annotations.yaml
k apply -f custom-command.yaml
k apply -f custom-config.yaml
k apply -f custom-image.yaml
k apply -f custom-port.yaml
k apply -f custom-renames.yaml
k apply -f custom-shutdown.yaml
k apply -f enable-exporter.yaml
k apply -f extravolumes-mounts.yaml
k apply -f minimum.yaml
k apply -f node-affinity.yaml
k apply -f persistent-storage-no-pvc-deletion.yaml
k apply -f persistent-storage.yaml
k apply -f pmem.yaml
k apply -f pod-anti-affinity.yaml
k apply -f security-context.yaml
k apply -f sidecars.yaml
k apply -f tolerations.yaml
k apply -f topology-spread-contraints.yaml
```{{exec}}


```
helm repo add redis-operator https://spotahome.github.io/redis-operator
helm repo update
helm install redis-operator redis-operator/redis-operator
```{{exec}}


```
REDIS_OPERATOR_VERSION=v1.3.0
kubectl replace -f https://raw.githubusercontent.com/spotahome/redis-operator/${REDIS_OPERATOR_VERSION}/manifests/databases.spotahome.com_redisfailovers.yaml
```{{exec}}

```
k apply -f ./redis-operator/example/redisfailover
```{{exec}}






```

```{{exec}}



```

```{{exec}}



```

```{{exec}}
