# Change main.go
When create controller, 

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
```
