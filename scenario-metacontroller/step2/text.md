Install metacontroller.

- clusterrole
- user is 'admin'
- pwd is 'admin' 
```shell
kubectl create clusterrolebinding my-cluster-admin-binding --clusterrole=cluster-admin --user=admin@admin

```{{exec}}


Apply all set of production resources defined in kustomization.yaml in `production` directory 

```shell
 kubectl apply -k https://github.com/metacontroller/metacontroller/manifests/production
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
 
```shell
 
```{{exec}}

