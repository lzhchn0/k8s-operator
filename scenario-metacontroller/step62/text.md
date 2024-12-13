configmap propagation 


kubectl create cm configmap-propagation-controller --from-file=sync.py


kubectl apply -f configmap-propagation.yaml 
