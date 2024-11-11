Setup a Composite Controller.
 
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



```shell
 
```{{exec}}

