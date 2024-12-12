Setup a indexedJob metacontroller.
 
```shell
git clone https://github.com/metacontroller/metacontroller.git 
```{{exec}}


```shell
cd /root/metacontroller/examples/indexedjob
kn  metacontroller
```{{exec}}

- 1. apply manifest 
```shell
cd /root/metacontroller/examples/indexedjob/manifest
kubectl create configmap indexedjob-controller -n metacontroller --from-file=sync.py
kubectl apply -f indexedjob-controller.yaml
```{{exec}}


- 2. deploy v1
```shell
cd /root/metacontroller/examples/indexedjob
kubectl apply -k v1 
```{{exec}}

- 3. deploy customized indexed-job 
```shell
cd /root/metacontroller/examples/indexedjob
kubectl  apply -f my-indexedjob.yaml
```{{exec}}


- 4. check logs 

```shell
watch " kubectl logs print-index-0;kubectl logs print-index-1;kubectl logs print-index-2;kubectl logs print-index-3;kubectl logs print-index-4;kubectl logs print-index-5;
kubectl logs print-index-6;kubectl logs print-index-7;kubectl logs print-index-8;kubectl logs print-index-9; "
```{{exec}}
