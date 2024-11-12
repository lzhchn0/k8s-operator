Setup a appropriate environment.
 
```shell
git clone https://github.com/metacontroller/metacontroller.git 
```{{exec}}


```shell
cd /root/metacontroller/examples/indexedjob
```{{exec}}



```shell
kubectl apply -k v1 
```{{exec}}



```shell
kubectl apply -f my-indexedjob.yaml 
```{{exec}}


```shell
$ kubectl logs print-index-2 
```{{exec}}
