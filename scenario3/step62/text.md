spotahome setup 


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



```

```{{exec}}



```

```{{exec}}
