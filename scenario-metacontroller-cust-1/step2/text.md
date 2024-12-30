

```shell
 cd ..
 mkdir mydep
 cd mydep
 mkdir v1
 mkdir manifest
 cd  /root/metacontroller/examples/mydep
```{{exec}}

```shell
tar -xvf mydep*.tar
rm mydep*.tar
```{{exec}}

```shell
kubectl apply -k v1
kubectl apply -k manifest
kubectl apply -f mydep-ctrl-custom.yaml
kubectl create deploy mydep --image=nginx 
```{{exec}}


view stacktrace in log from metacontroller
```shell
kubectl logs metacontroller-0  | jq '.'
```{{exec}}


- check log from sync.py
```shell
kubetail mydep --since 33m
```{{exec}}

To update configmap and sync.py
```shell
kubectl replace -k manifest  $now
```{{exec}}


 
```shell
kubectl run my-tmp1 --image=busybox -it -- nslookup google.com
```{{exec}}
