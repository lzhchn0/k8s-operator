# Util 

```
apt install graphviz
```{{exec}}

> goimportdot

```
go get -u github.com/yqylovy/goimportdot
```{{exec}}

> kustomize_v4.0.5
```shell
curl -LO https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv4.0.5/kustomize_v4.0.5_linux_amd64.tar.gz
```{{exec}}

> kustomize_v4.5.6
```shell
curl -LO https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv4.5.6/kustomize_v4.5.6_linux_amd64.tar.gz
```{{exec}}


> go-callvis

```
go install github.com/ofabry/go-callvis@latest
```{{exec}}

```
kubectl port-forward --address 0.0.0.0  svc/prometheus-operator-grafana 8080:80
```{{exec}}

> Now access Web Selector
{{TRAFFIC_SELECTOR}}

> Now access  go-callvis:4444
{{TRAFFIC_HOST1_4444}}

> Now access  go-callvis:7878
[ACCESS go-callvis]({{TRAFFIC_HOST1_7878}})

> Now access  port:80

[ACCESS go-callvis]({{TRAFFIC_HOST1_80}})
