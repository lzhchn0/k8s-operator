Setup a appropriate environment is the most important step across all projects, choose it wisely. 

Go version 1.22.3
```shell
curl -OL  https://go.dev/dl/go1.22.3.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
```{{exec}}

Go version 1.21.4
```shell
curl -OL  https://go.dev/dl/go1.21.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.20.4
```shell
curl -OL  https://go.dev/dl/go1.20.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.19.4
```shell
curl -OL  https://go.dev/dl/go1.19.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz
```{{exec}}

Go version 1.18.4
```shell
curl -OL  https://go.dev/dl/go1.18.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz
```{{exec}}

Go version 1.16.4
```shell
curl -OL  https://go.dev/dl/go1.16.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.15.4
```shell
curl -OL  https://go.dev/dl/go1.15.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.15.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.14.4
```shell
curl -OL  https://go.dev/dl/go1.14.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.13.4
```shell
curl -OL  https://go.dev/dl/go1.13.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.12.4
```shell
curl -OL  https://go.dev/dl/go1.12.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.12.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.11.4
```shell
curl -OL  https://go.dev/dl/go1.11.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
```{{exec}}


Go version 1.10.4
```shell
curl -OL  https://go.dev/dl/go1.10.4.linux-amd64.tar.gz \
&&  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.10.4.linux-amd64.tar.gz
```{{exec}}

Setup Go path
```shell
export GOPATH=/root/go
export PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
alias h=history
```{{exec}}



Download/install kubebuilder
```shell
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```{{exec}}

