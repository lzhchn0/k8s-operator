
```shell
 mkdir mydep
 cd mydep
 mkdir v1
 mkdir manifest
```{{exec}}

- enter folder mydep
```shell
cd  /root/metacontroller/examples/mydep
```{{exec}}




```shell
kubectl apply -k v1
kubectl apply -k manifest
kubectl apply -f mydep-ctrl-custom.yaml 
```{{exec}}

To update configmap and sync.py
```shell
kubectl replace -k manifest  $now
```{{exec}}

- check log from sync.py
```shell
kubetail mydep --since 33m
```{{exec}}



```shell
kubectl create deploy mydep --image=httpd:2.4.41-alpine
```{{exec}}



