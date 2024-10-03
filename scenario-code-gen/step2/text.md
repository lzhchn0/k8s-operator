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
