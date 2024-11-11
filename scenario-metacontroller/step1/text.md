Setup a appropriate environment.

Go version 1.22.3
```shell
curl -OL  https://go.dev/dl/go1.22.3.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
```{{exec}}

Setup Go path
```shell
export GOPATH=/root/go
export GOROOT=/usr/local/go
export PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
alias h=history
alias kb=kubebuilder
```{{exec}}

Setup k8s commands env
```shell
alias kx='f() { [ "$1" ] && kubectl config use-context $1 || kubectl config current-context ; } ; f'
alias kn='f() { [ "$1" ] && kubectl config set-context --current --namespace $1 || kubectl config view --minify | grep namespace | cut -d" " -f6 ; } ; f'

export now="--grace-period=0 --force"
export do="-o yaml --dryrun=client"
```{{exec}}

```shell
alias finds="find . -name '*.go' -type f -exec grep -inH "
alias findy="find . -name '*.yaml' -type f -exec grep -inH "

alias  trim='sed "s/^[\t ]*//g"i'
```{{exec}}


Setup vim
```shell
chmod +w ~/.vimrc
echo ""  >> ~/.vimrc
echo "set number"  >> ~/.vimrc
```{{exec}}



# Download/install rabbitmq operator
```shell

wget https://charts.bitnami.com/bitnami/rabbitmq-cluster-operator-4.3.23.tgz
helm install my-rab ./rabbitmq-cluster-operator-4.3.23.tgz --namespace rabbit --create-namespace

```{{exec}}



# Download/install rabbitmq 
```shell

kubectl create namespace rabbit-system

helm repo add stable https://charts.helm.sh/stable

helm install rabbit --set service.type=NodePort stable/rabbitmq --namespace rabbit-system

```{{exec}}


# show the credential 
```shell
echo $(kubectl get secret --namespace rabbit mu-rabbit-rabbitmq -o jsonpath="{.data.rabbitmq-password}" | base64 --decode)
```{{exec}}


# forward the port 5672 to external 
```shell
kubectl port-forward --namespace rabbit svc/mu-rabbit-rabbitmq 5672:5672
```{{exec}}


# forward the port 15672 to external 
```shell
kubectl port-forward --namespace rabbit svc/mu-rabbit-rabbitmq 15672:15672
```{{exec}}
