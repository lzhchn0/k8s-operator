
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
kubectl create deploy mydep --image=nginx 
```{{exec}}

view log from metacontroller
```shell
kubectl logs metacontroller-0  | jq -r '.stacktrace | gsub("\\n";"\n    ")|("Stack Trace:\n  \(.)")'
```{{exec}}


view stacktrace in log from metacontroller
```shell
kubectl logs metacontroller-0  | jq '.'
```{{exec}}

To update configmap and sync.py
```shell
kubectl replace -k manifest  $now
```{{exec}}

- check log from sync.py
```shell
kubetail mydep --since 33m
```{{exec}}

- 'GET' request to webhook
```shell
kubectl run tmp --restart=Never --rm --image=nginx:alpine -i -- curl   -X GET   http://mydep-controller.metacontroller/sync
```{{exec}}



- POST a command to webhook
```shell
curl -H 'Content-Type: application/json' \
       -d '{ "title":"foo","body":"bar", "id": 1}' \
       -X POST \
       http://10.96.3.237/sync  
```{{exec}}

- POST a command to webhook
```shell
kubectl run tmp --restart=Never --rm --image=nginx:alpine -i -- curl -d '{ "title":"foo","body":"bar", "id": 1}' -H 'Content-Type: application/json' -X POST   http://mydep-controller.metacontroller/sync
```{{exec}}


```shell
kubectl create deploy mydep --image=httpd:2.4.41-alpine
```{{exec}}



