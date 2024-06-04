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

> Go1.1.1 (2018-08-24 released)

Setup Go path
```shell
export GOPATH=/root/go
export GOROOT=/usr/local/go
export PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
alias h=history
```{{exec}}

Download/install kubebuilder v2.0.1
```shell
curl -L -o kubebuilder.tar.gz https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.0.1/kubebuilder_2.0.1_linux_amd64.tar.gz
tar -xvzf ./kubebuilder.tar.gz  -C kubebuilder201
```{{exec}}


Download/install kubebuilder for latest 
```shell
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```{{exec}}

Release		Release date	Maintenance end
go1		3/28/2012	12/1/2013
go1.1		5/13/2013	6/18/2014
go1.2		12/1/2013	12/10/2014
go1.3		6/18/2014	8/19/2015
go1.4		12/10/2014	2/17/2016
go1.5		8/19/2015	8/15/2016
go1.6		2/17/2016	2/16/2017
go1.7		8/15/2016	8/24/2017
go1.8		2/16/2017	2/16/2018
go1.9		8/24/2017	8/24/2018
go1.10		2/16/2018	2/25/2019
go1.11		8/24/2018	9/3/2019
go1.12		2/25/2019	2/25/2020
go1.13		9/3/2019	8/11/2020
go1.14		2/25/2020	2/16/2021
go1.15		8/11/2020	8/16/2021
go1.16		2/16/2021	3/15/2022
go1.17		8/16/2021	8/2/2022
go1.18		3/15/2022	2/1/2023
go1.19		8/2/2022	8/8/2023
go1.20		2/1/2023	Q1 2024
go1.21		8/8/2023	Q3 2024
go1.22		2/6/2024	Q1 2025



