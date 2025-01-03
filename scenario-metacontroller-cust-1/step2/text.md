

```shell
 cd ..
 mkdir hawk
 cd hawk
 mkdir v1
 mkdir manifest
 cd  /root/metacontroller/examples/hawk
```{{exec}}

```shell
tar -xvf hawk*.tar
rm hawk*.tar
```{{exec}}

```shell
kubectl apply -k v1
kubectl apply -k manifest
kubectl apply -f hawk-ctrl-custom.yaml 
```{{exec}}

Label test-hawk
```shell
kubectl label hawk test-hawk    new4=Hello4
```{{exec}}

view stacktrace in log from metacontroller
```shell
kubectl logs metacontroller-0  | jq '.'
```{{exec}}


- check log from sync.py
```shell
kubetail hawk --since 33m
```{{exec}}


- Watch status of my operators'
```shell
watch "kubectl get pods,svc,deploy,cm"
```{{exec}}

- update a property of MyDep test-hawk
```shell
kubectl labels MyDep test-hawk new2=Hello2
```{{exec}}


To update configmap and sync.py
```shell
kubectl replace -k manifest  $now
```{{exec}}


- 'GET' request to webhook
```shell
kubectl run tmp --restart=Never --rm --image=nginx:alpine -i -- curl   -X GET   http://hawk-controller.metacontroller/sync
```{{exec}}


 
```shell
kubectl run my-tmp1 --image=busybox --rm -i  -- nslookup google.com
```{{exec}}



get extracted_json.log
```shell
kubectl logs hawk-controller-678f8cc6c5-nfzb2  > extracted_json.log
```{{exec}}
