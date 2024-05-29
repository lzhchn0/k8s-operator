# Change monitoring_controller.go

```
type MonitoringReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	Clientset *kubernetes.Clientset
	Config    *rest.Config
}
```

```
import (
	"context"
	"fmt"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	redisclusterv1 "blabla.com/api/v1"
)
```


```
func (r *MonitoringReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logf := log.FromContext(ctx)

	logf.Info("==== Begin ")

	podList := &corev1.PodList{}
	if err := r.List(ctx, podList, client.InNamespace(req.Namespace)); err != nil {
		// handle error
	}

	// -- start loop
	for _, pod := range podList.Items {
		fmt.Printf("Name: %s, Status: %s, Namespace: %s\n", pod.ObjectMeta.Name, pod.Status.Phase, pod.ObjectMeta.Namespace)
		if strings.HasPrefix(pod.ObjectMeta.Name, "redis-worker-5c8b468cf7-z") {

			// -- process pod

			podName := pod.ObjectMeta.Name // replace with your pod name
			namespace := pod.ObjectMeta.Namespace
			command := []string{"/bin/sh", "-c", "/usr/local/bin/redis-cli info server; /usr/local/bin/redis-cli info memory; /usr/local/bin/redis-cli  -p 6379 set fop 304"} // replace with your command
			req2 := r.Clientset.CoreV1().RESTClient().Post().
				Resource("pods").
				Name(podName).
				Namespace(namespace).
				SubResource("exec").
				VersionedParams(&corev1.PodExecOptions{
					Container: "worker",
					Command:   command,
					Stdin:     true,
					Stdout:    true,
					Stderr:    true,
					TTY:       true,
				}, scheme.ParameterCodec)

			exec, err := remotecommand.NewSPDYExecutor( // ctrl.Manager.GetConfig(),
				//r.Clientset.RESTConfig(),
				r.Config,
				"POST", req2.URL())
			if err != nil {
				// handle error
				logf.Error(err, "==**==2 Failed to exec reids command")
			}

			err = exec.Stream(remotecommand.StreamOptions{
				Stdin:  os.Stdin,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
				Tty:    true,
			})

			if err != nil {
				// handle error
				logf.Error(err, "==**==3 Failed to exec reids command")
			}


		}
	}

	return ctrl.Result{}, nil
}
```
