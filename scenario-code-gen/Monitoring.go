package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Monitoring struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   MonitoringSpec   `json:"spec,omitempty"`
    Status MonitoringStatus `json:"status,omitempty"`
}

type MonitoringSpec struct {
    Field1 string `json:"field1,omitempty"`
    Field2 NestedField `json:"field2,omitempty"`
}

type NestedField struct {
    SubField1 string `json:"subField1,omitempty"`
    SubField2 string `json:"subField2,omitempty"`
}

type MonitoringStatus struct {
    Condition string `json:"condition,omitempty"`
}

type DeploymentStatus struct {
    Name string `json:"name"`
}


type ServiceStatus struct {
    Name string `json:"name"`
    IPs []string `json:"ips,omitempty"`
    Hostnames []string `json:"hostnames,omitempty"`
}


type IngressStatus struct {
    Name string `json:"name"`
    IPs []string `json:"ips,omitempty"`
    Hostnames []string `json:"hostnames,omitempty"`

}

