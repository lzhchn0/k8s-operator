Setup a appropriate environment.

step 1: foo-controller

```shell

mkdir foo-controller && cd foo-controller
go mod init foo-controller
go get k8s.io/apimachinery@v0.24.16
go get k8s.io/code-generator@v0.24.16
```{{exec}}

step 2
```shell
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


step 62 -- BAD
```shell
go get k8s.io/code-generator@v0.28.1
```{{exec}}

step 63: install and verify the installation -- GOOD
```shell
go mod tidy
go list -m k8s.io/code-generator
```{{exec}}


step 64
```shell

```{{exec}}


step 65
```shell
mkdir boo && cd boo
git init
git clone https://github.com/sportshead/codegen-demo.git
cd codegen-demo
kubectl apply -f crd.yaml
kubectl apply -f example/songs.yaml
go run example/main.go
```{{exec}}



step 66
```shell
mkdir -p pkg/apis/foo.example.com/v1

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




```shell

```{{exec}}
