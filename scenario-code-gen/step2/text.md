Setup a appropriate environment.

step 1: foo-controller

```shell

mkdir foo-controller && cd foo-controller
go mod init foo-controller
go get k8s.io/apimachinery@v0.29.0
go get k8s.io/code-generator@v0.29.0
```{{exec}}

step 2
```shell
mkdir -p pkg/client/clientset
mkdir -p pkg/apis/foo.example.com/v1 && cd pkg/apis/foo.example.com/v1
touch doc.go
touch types.go
touch register.go
```{{exec}}


step 3
```shell
cd ~/foo-controller
go mod tidy
go mod vendor
chmod -R 777 hack
# vendor目录中没有code-generator目录，因为k8s.io/code-generator这个依赖在项目中并没有真正被引用过，所以使用go mod vendor是无法将这个依赖更新到vendor中。可以选择手动拷贝，注意修改对应目录。也可以用tools.go来手动导入这个包。
go env | grep GOMODCACHE
cd $GOMODCACHE/k8s.io
cp -r code-generator@v0.24.16 foo-controller/vendor/k8s.io/code-generator
chmod -R 777 vendor
cd hack
./update-codegen.sh
```{{exec}}



step 4
```shell
GOMODCACHE='/root/go/pkg/mod'
```{{exec}}

./update-codegen.sh: line 14: ../vendor/k8s.io/code-generator/generate-groups.sh: No such file or directory
step 5
```shell
cd /root/go/pkg/mod/k8s.io
cp /root/f/pkg/apis/foo/v1/*.go  /root/foo-controller/pkg/apis/foo/v1/
```{{exec}}



step 6
```shell
go get k8s.io/apimachinery@v0.31.0
go get k8s.io/code-generator@v0.31.0
```{{exec}}


step 61: copy genertor
```shell
 cp -r $(go env GOMODCACHE)/k8s.io/code-generator@v0.31.0   ~/foo-controller/vendor/k8s.io/code-generator
```{{exec}}


step 62 
```shell
go get k8s.io/code-generator@v0.29.0
```{{exec}}

step 63: 
```shell

```{{exec}}


step 64
```shell
 
```{{exec}}

step 642 -- GOOD
```shell
apt-get install tree
```{{exec}}



step 643
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```{{exec}}


step 643 -- GOOD
https://killercoda.com/walllnerryan/scenario/k8s-redis-sentinal-lab
```shell
helm upgrade -i redis-operator oci://ghcr.io/sap/redis-operator-helm/redis-operator
```{{exec}}

step 644 -- GOOD -- install a redis 
```shell
helm install redis -f values.yaml bitnami/redis
kubectl create -f app.yaml
```{{exec}}




step 65: install and verify the installation -- GOOD
```shell
mkdir boo && cd boo
# git init
git clone https://github.com/sportshead/codegen-demo.git
cd codegen-demo
kubectl apply -f crd.yaml
kubectl apply -f example/songs.yaml
go run example/main.go
```{{exec}}

step 65: build SAP/redis -- GOOD
```shell
mkdir boo && cd boo
git clone https://github.com/SAP/redis-operator.git
   
cd redis-operator
make all

```{{exec}}



step 66
```shell
cd boo/codegen-demo
mkdir -p pkg/apis/foo.example.com/v1
mkdir -p pkg/client/clientset

```{{exec}}



step 661:
```shell
go mod tidy
go list -m k8s.io/code-generator
```{{exec}}



step 662
```shell
chmod -R 777 hack
hack/update-codegen.sh
find .  -name "*.go"  -type f -print0 | xargs -0 cat >> prom.total
```{{exec}}




```shell
# Generate informers
./informer-gen
--input-dirs="github.com/codegen-demo/pkg/apis/foo.example.com/v1"
--versioned-clientset-package="github.com/codegen-demo/pkg/generated/clientset/versioned"
--listers-package="github.com/codegen-demo/pkg/generated/listers"
--output-package="github.com/codegen-demo/pkg/generated/informers"
```{{exec}}


```shell

```{{exec}}


```shell

```{{exec}}

```shell

```{{exec}}


step 655 GOOD - create tls.crt  tls.csr  tls.key
```shell
openssl genrsa -out tls.key 2048
openssl req -new -key tls.key -out tls.csr -subj "/CN=webhook-server"
openssl x509 -req -in tls.csr -signkey tls.key -out tls.crt -days 365
kubectl create secret tls webhook-server-cert --cert=tls.crt --key=tls.key -n <namespace>
```{{exec}}



```shell
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webhook-server
  namespace: <namespace>
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: webhook-server
    spec:
      containers:
      - name: webhook-server
        image: your-webhook-server-image
        volumeMounts:
        - name: webhook-certs
          mountPath: /tmp/k8s-webhook-server/serving-certs
          readOnly: true
      volumes:
      - name: webhook-certs
        secret:
          secretName: webhook-server-cert
```{{exec}}




```shell
cp tls*  /tmp/k8s-webhook-server/serving-certs/
```{{exec}}




```shell

```{{exec}}




```shell

```{{exec}}




```shell

```{{exec}}




```shell

```{{exec}}


## ERROR 1:36pm
controlplane $ curl -k  https://10.96.0.1:443/apis/coordination.k8s.io/v1/namespaces/default/leases/redis-operator.cs.sap.com  
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Failure",
  "message": "leases.coordination.k8s.io \"redis-operator.cs.sap.com\" is forbidden: User \"system:anonymous\" cannot get resource \"leases\" in API group \"coordination.k8s.io\" in the namespace \"default\"",
  "reason": "Forbidden",
  "details": {
    "name": "redis-operator.cs.sap.com",
    "group": "coordination.k8s.io",
    "kind": "leases"
  },
  "code": 403
}


## Pod ERROR
controlplane $ k logs redis-operator-74f74c75f6-zh65p 

{"level":"info","ts":"2024-10-03T17:26:26Z","logger":"controller-runtime.webhook","msg":"Registering webhook","path":"/admission/cache.cs.sap.com/v1alpha1/redis/validate"}
{"level":"info","ts":"2024-10-03T17:26:26Z","logger":"setup","msg":"starting manager"}
{"level":"info","ts":"2024-10-03T17:26:26Z","logger":"controller-runtime.metrics","msg":"Starting metrics server"}
{"level":"info","ts":"2024-10-03T17:26:26Z","logger":"controller-runtime.metrics","msg":"Serving metrics server","bindAddress":":8080","secure":false}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"starting server","kind":"health probe","addr":"[::]:8081"}
{"level":"info","ts":"2024-10-03T17:26:26Z","logger":"controller-runtime.webhook","msg":"Starting webhook server"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Stopping and waiting for non leader election runnables"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Stopping and waiting for leader election runnables"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Starting EventSource","controller":"redis","controllerGroup":"cache.cs.sap.com","controllerKind":"Redis","source":"kind source: *v1alpha1.Redis"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Starting Controller","controller":"redis","controllerGroup":"cache.cs.sap.com","controllerKind":"Redis"}
I1003 17:26:26.619309       1 leaderelection.go:250] attempting to acquire leader lease default/redis-operator.cs.sap.com...
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Starting workers","controller":"redis","controllerGroup":"cache.cs.sap.com","controllerKind":"Redis","worker count":3}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Shutdown signal received, waiting for all workers to finish","controller":"redis","controllerGroup":"cache.cs.sap.com","controllerKind":"Redis"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"All workers finished","controller":"redis","controllerGroup":"cache.cs.sap.com","controllerKind":"Redis"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Stopping and waiting for caches"}
{"level":"error","ts":"2024-10-03T17:26:26Z","logger":"controller-runtime.source.EventHandler","msg":"failed to get informer from cache","error":"Timeout: failed waiting for *v1alpha1.Redis Informer to sync","stacktrace":"sigs.k8s.io/controller-runtime/pkg/internal/source.(*Kind).Start.func1.1\n\t/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.17.2/pkg/internal/source/kind.go:68\nk8s.io/apimachinery/pkg/util/wait.loopConditionUntilContext.func1\n\t/go/pkg/mod/k8s.io/apimachinery@v0.29.3/pkg/util/wait/loop.go:53\nk8s.io/apimachinery/pkg/util/wait.loopConditionUntilContext\n\t/go/pkg/mod/k8s.io/apimachinery@v0.29.3/pkg/util/wait/loop.go:54\nk8s.io/apimachinery/pkg/util/wait.PollUntilContextCancel\n\t/go/pkg/mod/k8s.io/apimachinery@v0.29.3/pkg/util/wait/poll.go:33\nsigs.k8s.io/controller-runtime/pkg/internal/source.(*Kind).Start.func1\n\t/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.17.2/pkg/internal/source/kind.go:56"}
W1003 17:26:26.621991       1 reflector.go:539] pkg/mod/k8s.io/client-go@v0.29.3/tools/cache/reflector.go:229: failed to list *v1alpha1.Redis: client rate limiter Wait returned an error: context canceled
E1003 17:26:26.622010       1 reflector.go:147] pkg/mod/k8s.io/client-go@v0.29.3/tools/cache/reflector.go:229: Failed to watch *v1alpha1.Redis: failed to list *v1alpha1.Redis: client rate limiter Wait returned an error: context canceled
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Stopping and waiting for webhooks"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Stopping and waiting for HTTP servers"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"shutting down server","kind":"health probe","addr":"[::]:8081"}
{"level":"info","ts":"2024-10-03T17:26:26Z","logger":"controller-runtime.metrics","msg":"Shutting down metrics server with timeout of 1 minute"}
{"level":"info","ts":"2024-10-03T17:26:26Z","msg":"Wait completed, proceeding to shutdown the manager"}
E1003 17:26:26.622303       1 leaderelection.go:332] error retrieving resource lock default/redis-operator.cs.sap.com: Get "https://10.96.0.1:443/apis/coordination.k8s.io/v1/namespaces/default/leases/redis-operator.cs.sap.com": context canceled
{"level":"error","ts":"2024-10-03T17:26:26Z","logger":"setup","msg":"problem running manager","error":"open /tmp/k8s-webhook-server/serving-certs/tls.crt: no such file or directory","stacktrace":"main.main\n\t/workspace/main.go:120\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:271"}
