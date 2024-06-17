setup Prometheus Operator


==GOOD
```
kubectl create namespace monitoring
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install prometheus-operator prometheus-community/kube-prometheus-stack -n monitoring
alias kx='f() { [ "$1" ] && kubectl config use-context $1 || kubectl config current-context ; } ; f'
alias kn='f() { [ "$1" ] && kubectl config set-context --current --namespace $1 || kubectl config view --minify | grep namespace | cut -d" " -f6 ; } ; f'
kn monitoring 
k get pods
```{{exec}}


> reset grafana credential --
```
/usr/share/grafana/bin/grafana-cli admin reset-admin-password
```{{exec}}
