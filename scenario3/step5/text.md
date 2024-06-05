# Change main.go
When create controller in main.go,

```
import ( "k8s.io/client-go/kubernetes" )
```{{copy}}

124:
```
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	config := mgr.GetConfig()
	clientset, err := kubernetes.NewForConfig(config)
```{{copy}}

```
	if err = (&controller.MonitoringReconciler{
		Client:    mgr.GetClient(),
		Scheme:    mgr.GetScheme(),
		Clientset: clientset,
		Config:    mgr.GetConfig(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Monitoring")
		os.Exit(1)
	}
```{{copy}}
