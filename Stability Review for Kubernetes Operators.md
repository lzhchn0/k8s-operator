


### Differences Between DaemonSets and Other Controllers:
| Name of Op             | DaemonSet                          | Vendor                              |
|------------------------|------------------------------------|-------------------------------------|
| **Prometheus Op**       | One Pod per node                   | Pods scheduled based on replicas    |
| **Strimzi Op**          | Automatically scales with nodes    | Manually set replica count          |
| **Velero Op**           | System-level services              | Application workloads               |
| **Node Selectivity**    | Can target specific nodes          | Not node-specific                   |

---
