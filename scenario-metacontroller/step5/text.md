Setup a bluegreen metacontroller.
 
```shell
git clone https://github.com/metacontroller/metacontroller.git 
```{{exec}}


```shell
cd /root/metacontroller/examples/bluegreen
```{{exec}}

- manifests
- nov 18
```shell
make build
make generate_crds
```{{exec}}

- install meta-controller
```shell
kubectl apply -k v1 
```{{exec}}


- change replicas here, 
```shell
kubectl apply -f my-bluegreen.yaml
```{{exec}}



```shell
cd /root/metacontroller/examples/bluegreen/manifest
k create configmap bluegreen-controller  -n metacontroller  --from-file=sync.js
k apply -f bluegreen-controller.yaml
```{{exec}}

