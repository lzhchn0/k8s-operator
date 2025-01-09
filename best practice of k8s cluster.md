
- enableServiceLinks : false 

The enableServiceLinks parameter(in PodSpec v1) is enabled by default but is often unnecessary. Disabling it can help reduce performance issues at the Kubernetes level.


- admission controller

admission controller is a powerful policy enforcement tool, in a given cloud native environment 

<ins>Must check if informer gets synced before admission controller. </ins>

