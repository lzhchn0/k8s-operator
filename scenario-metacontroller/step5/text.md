Setup a bluegreen metacontroller.
 
```shell
git clone https://github.com/metacontroller/metacontroller.git 
```{{exec}}


```shell
cd /root/metacontroller/examples/bluegreen
```{{exec}}

- manifests

```shell
make generate_crds
```{{exec}}


```shell
kubectl apply -k v1 
```{{exec}}


- change replicas here, 
```shell
kubectl apply -f my-bluegreen.yaml
```{{exec}}



```shell
k create configmap bluegreen-controller  -n metacontroller  --from-file=sync.js
```{{exec}}

