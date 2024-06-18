> To install the Redis Operator using Helm

0. **Namespace setup**:

```bash
kubectl create ns spota
kn spota
```
1. **Clone the repository**:

```bash
git clone https://github.com/spotahome/redis-operator.git
```

2. **Checkout the specific version**:

```bash
cd redis-operator
git checkout v1.2.4
```

3. **Install the Helm chart**:

```bash
helm install my-redis-op ./charts/redisoperator
```

In this example, `my-redis-op` is the name of the Helm release. You can replace it with any name you prefer.
