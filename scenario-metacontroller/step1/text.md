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
 
alias kx='f() { [ "$1" ] && kubectl config use-context $1 || kubectl config current-context ; } ; f'
alias kn='f() { [ "$1" ] && kubectl config set-context --current --namespace $1 || kubectl config view --minify | grep namespace | cut -d" " -f6 ; } ; f'

PS1='$(pwd)$ '

export now="--grace-period=0 --force"
export do="-o yaml --dry-run=client"
export d="describe"
 
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

- Setup kubetail
```shell
curl -LO https://raw.githubusercontent.com/johanhaleby/kubetail/master/kubetail
chmod +x kubetail
sudo mv kubetail /usr/local/bin/
```{{exec}}

- Setup yq
```shell
sudo add-apt-repository ppa:rmescandon/yq
sudo apt update
sudo apt install yq -y
```{{exec}}

- delete 'description' fields in yaml file crdv1.yaml
```shell
yq eval -o=json crdv1.yaml  | jq 'walk(if type == "object" then del(.description) else . end)'   > tt0.yaml
yq eval tt0.yaml -P  > crdv1-new.yaml
```{{exec}}

nl usage
```shell
apt install coreutils
nl -ba -s": " sync.py
```{{exec}}

yq retrieve first-level elements and 2nd level 
```shell
yq eval '. | keys' crdv1.yaml
yq eval '.metadata | keys' crdv1.yaml
```{{exec}}



```shell
yq eval '.' crdv1.yaml
yq eval '.<first_level_key>' yourfile.yaml
```{{exec}}

```shell
yq eval '.[]' crdv1.yaml 
```{{exec}}


