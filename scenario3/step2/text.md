# Build Redis-Cluster

This will create a Redis cluster with one master and three workers. 

1.Create a Redis master Deployment

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


2. Create a Redis master Service

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

3. Create a Redis worker Deployment

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


4. Create a Redis worker Service.
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


