helm repo add redis-operator https://spotahome.github.io/redis-operator
helm repo update
helm install redis-operator redis-operator/redis-operator
