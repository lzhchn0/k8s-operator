

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
```{{exec}}

Label test-mydep
```shell
kubectl label mydep test-mydep    new4=Hello4
```{{exec}}

view stacktrace in log from metacontroller
```shell
kubectl logs metacontroller-0  | jq '.'
```{{exec}}


- check log from sync.py
```shell
kubetail mydep --since 33m
```{{exec}}


- Watch status of my operators'
```shell
watch "kubectl get pods,svc,deploy,cm"
```{{exec}}

- update a property of MyDep test-mydep
```shell
kubectl labels MyDep test-mydep new2=Hello2
```{{exec}}


To update configmap and sync.py
```shell
kubectl replace -k manifest  $now
```{{exec}}


- 'GET' request to webhook
```shell
kubectl run tmp --restart=Never --rm --image=nginx:alpine -i -- curl   -X GET   http://mydep-controller.metacontroller/sync
```{{exec}}


 
```shell
kubectl run my-tmp1 --image=busybox --rm -i  -- nslookup google.com
```{{exec}}



get extracted_json.log
```shell
kubectl logs mydep-controller-678f8cc6c5-nfzb2  > extracted_json.log
```{{exec}}
