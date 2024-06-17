setup Prometheus Operator



```
kubectl create namespace monitoring
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm init
helm install prometheus-operator prometheus-community/kube-prometheus-stack -n monitoring
alias kx='f() { [ "$1" ] && kubectl config use-context $1 || kubectl config current-context ; } ; f'
alias kn='f() { [ "$1" ] && kubectl config set-context --current --namespace $1 || kubectl config view --minify | grep namespace | cut -d" " -f6 ; } ; f'
kn monitoring 
k get pods
```{{exec}}



```

```{{exec}}
