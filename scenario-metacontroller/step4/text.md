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




```shell
kubectl apply -f my-indexedjob.yaml 
```{{exec}}


```shell
$ kubectl logs print-index-2 
```{{exec}}




```shell
make unit-test
```{{exec}}



