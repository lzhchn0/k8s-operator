
```shell
 mkdir mydep
 cd mydep
 mkdir v1
 mkdir manifest
```{{exec}}



```shell
kubectl apply -k v1
kubectl apply -k manifest
kubectl apply -f mydep-ctrl-custom.yaml 
```{{exec}}



```shell
kubectl create deploy mydep --image=httpd:2.4.41-alpine
```{{exec}}



