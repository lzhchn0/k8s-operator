Setup a Composite Controller.

Create Custom resource definition for parent resource
```shell
wget https://raw.githubusercontent.com/gtanand1994/ideas2crd/main/gt-apps/controller.yaml
 
```{{exec}}



```shell
kubectl create configmap -n metacontroller gt-app-controller --from-file=app
kubectl create configmap -n metacontroller gt-app-templates --from-file=templates
 
```{{exec}}



```shell
wget https://raw.githubusercontent.com/gtanand1994/ideas2crd/main/gt-apps/crd.yaml 
```{{exec}}



```shell
kubectl get crd|grep gtapp
kubectl api-resources|grep gtapp 
```{{exec}}


```shell
git clone https://github.com/gtanand1994/ideas2crd.git
```{{exec}}



```shell
cd ideas2crd/gt-apps
```{{exec}}


```shell
kn metacontroller
```{{exec}}



```shell
k apply -f controller.yaml
k apply -f crd.yaml 
```{{exec}}


Bluegreen deployment


```shell
make build 
```{{exec}}


```shell
kubectl apply -f my-bluegreen.yaml
```{{exec}}




```shell
make unit-test
```{{exec}}



```shell
k create configmap bluegreen-controller  -n metacontroller  --from-file=sync.js
```{{exec}}

