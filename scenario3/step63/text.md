setup Prometheus Operator

```
alias kx='f() { [ "$1" ] && kubectl config use-context $1 || kubectl config current-context ; } ; f'
alias kn='f() { [ "$1" ] && kubectl config set-context --current --namespace $1 || kubectl config view --minify | grep namespace | cut -d" " -f6 ; } ; f'
```{{exec}}


==GOOD
```
kubectl create namespace monitoring
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install prometheus-operator prometheus-community/kube-prometheus-stack -n monitoring
kn monitoring 
k get pods
```{{exec}}

> forward port of grafana to browser
```
 kubectl port-forward svc/prometheus-operator-grafana --address=0.0.0.0  8080:80
 curl http://192.168.2.45:8099
```{{exec}}

```
 kubectl port-forward svc/prometheus-operator-grafana --address=0.0.0.0  8099:80
```{{exec}}

> reset grafana credential --
```
/usr/share/grafana/bin/grafana-cli admin reset-admin-password
```{{exec}}
