Setup a Hello Controller.


step 3

```shell
kubectl create namespace hello
 
```{{exec}}


```shell

apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: helloworlds.example.com
spec:
  group: example.com
  names:
    kind: HelloWorld
    plural: helloworlds
    singular: helloworld
  scope: Namespaced
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              who:
                type: string
    subresources:
     status: {}
```{{exec}}


^^ 
```shell
kubectl apply -f crd.yaml
 
```{{exec}}

^^ 
```shell
apiVersion: metacontroller.k8s.io/v1alpha1
kind: CompositeController
metadata:
  name: hello-controller
spec:
  generateSelector: true
  parentResource:
    apiVersion: example.com/v1
    resource: helloworlds
  childResources:
  - apiVersion: v1
    resource: pods
    updateStrategy:
      method: Recreate
  hooks:
    sync:
      webhook:
        url: http://hello-controller.hello/sync
```{{exec}}

^^ 
```shell
kubectl apply -f controller.yaml
 
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
