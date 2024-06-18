> To install the Redis Operator using Helm

0. **Namespace setup**:

```
kubectl create ns spota
kn spota
```{{exec}}
1. **Clone the repository**:

```
git clone https://github.com/spotahome/redis-operator.git
```{{exec}}

2. **Checkout the specific version**:

```
cd redis-operator
git checkout v1.2.4
```{{exec}}

3. **Install the Helm chart**:

```
helm install my-redis-op ./charts/redisoperator
```{{exec}}

In this example, `my-redis-op` is the name of the Helm release. You can replace it with any name you prefer.
