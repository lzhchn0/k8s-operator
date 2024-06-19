# Build Redis-Cluster

This will create a Redis cluster with one master and three workers. 

- Create a Redis master Deployment

```
kubectl apply -f - << EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-master
  namespace: default
spec:
  selector:
    matchLabels:
      app: redis
      role: master
  replicas: 1
  template:
    metadata:
      labels:
        app: redis
        role: master
    spec:
      containers:
      - name: master
        image: redis:5.0.3
        resources:
          limits:
            cpu: "0.1"
        ports:
        - containerPort: 6379
EOF
```{{exec}}


- Create a Redis master Service

```
kubectl apply -f - << EOF
apiVersion: v1
kind: Service
metadata:
  name: redis-master
  namespace: default  
spec:
  ports:
  - port: 6379
  selector:
    app: redis
    role: master
EOF
```{{exec}}

- Create a Redis worker Deployment

```
kubectl apply -f - << EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-worker
  namespace: default    
spec:
  selector:
    matchLabels:
      app: redis
      role: worker
  replicas: 3
  template:
    metadata:
      labels:
        app: redis
        role: worker
    spec:
      containers:
      - name: worker
        image: redis:5.0.3
        resources:
          limits:
            cpu: "0.1"
        env:
        - name: REDIS_MASTER_SERVICE_HOST
          value: redis-master
        ports:
        - containerPort: 6379
EOF
```{{exec}}


- Create a Redis worker Service.
```
kubectl apply -f - << EOF
apiVersion: v1
kind: Service
metadata:
  name: redis-worker
  namespace: default    
spec:
  ports:
  - port: 6379
  selector:
    app: redis
    role: worker
EOF
```{{exec}}


- To import data using redis-cli 
```
kubectl exec reids-2 -- redis-cli -h your-redis-host -p your-redis-port  set foo 10
```{{exec}}


- To create the first event
```
k apply  -f config/samples/rediscluster_v1_monitoring.yaml
```{{exec}}


- To create a new event
```
k replace --grace-period=0  --force  -f config/samples/rediscluster_v1_monitoring.yaml
```{{exec}}
