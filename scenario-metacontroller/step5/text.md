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
make
```{{exec}}

- 2. deploy bluegreen-controller
```shell
kubectl apply -k v1 
```{{exec}}

- 3. revised bluegreen-controller
```shell
kubectl replace -k v1   $now
```{{exec}}


- 1. create a bluegreen-deployment
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
kn  metacontroller
cd /root/metacontroller/examples/bluegreen/manifest
kubectl delete configmap bluegreen-controller  -n metacontroller  $now
kubectl create configmap bluegreen-controller  -n metacontroller  --from-file=sync.js
kubectl rollout restart deploy/bluegreen-controller
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


```shell
edit manifests/sync.js
```{{exec}}



```shell
find . -name '*.go' -mtime -2
```{{exec}}



```shell
kn  metacontroller
kubetail meta
```{{exec}}



```shell
kn  metacontroller
kubetail blue
```{{exec}}

```shell
kubectl get bluegreendeployments.ctl.enisoc.com nginx -o yaml | grep -i color
```{{exec}}

```shell
export bgd="bluegreendeployments"
kubectl patch $bgd nginx --type=merge -p '{"spec":{"bluetemplate":{"metadata":{"labels":{"new":"label"}}}}}'
kubectl patch $bgd nginx --type=merge -p '{"spec":{"greentemplate":{"metadata":{"labels":{"new":"label"}}}}}'
```{{exec}}


```shell
kubectl patch $bgd nginx --type=merge -p '{"spec":{"bluetemplate":{"metadata":{"labels":{"new2":"label2"}}}}}'
kubectl patch $bgd nginx --type=merge -p '{"spec":{"greentemplate":{"metadata":{"labels":{"new2":"label2"}}}}}'
```{{exec}}



```shell
grep creationTimestamp  tt2
"activeColor": "blue",
```{{exec}}



upload api/types.go
upload my-bluegreendeployment.yaml

