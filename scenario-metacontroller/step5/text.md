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
make build
```{{exec}}

- deploy bluegreen-controller
```shell
kubectl apply -k v1 
```{{exec}}

-- revised bluegreen-controller
```shell
kubectl replace -k v1   $now
```{{exec}}


- create a bluegreen-deployment
- change replicas here, 
```shell
kubectl apply -f my-bluegreen.yaml
```{{exec}}

```shell
kubectl -n metacontroller logs --tail=25 -l app=bluegreen-controller
```{{exec}}

```shell
ssh node01
watch crictl ps
```{{exec}}

```shell
cd /root/metacontroller/examples/bluegreen/manifest
k create configmap bluegreen-controller  -n metacontroller  --from-file=sync.js
k apply -f bluegreen-controller.yaml
```{{exec}}


```shell
watch "kubectl get pods -n default --show-labels  &&  kubectl get rs -n default && kubectl get svc -n default "
```{{exec}}


```shell
watch "kubectl get pods -n default -o yaml | grep 'image:'   "
```{{exec}}


```shell
kn  metacontroller
```{{exec}}


```shell
kn  default
```{{exec}}

upload api/types.go
upload my-bluegreendeployment.yaml

