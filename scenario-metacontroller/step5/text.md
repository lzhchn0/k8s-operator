Setup a appropriate environment.
 
```shell
git clone https://github.com/metacontroller/metacontroller.git 
```{{exec}}


```shell
cd /root/metacontroller/examples/bluegreen
```{{exec}}


```shell
make generate_crds
```{{exec}}


```shell
kubectl apply -k v1 
```{{exec}}

