package meta

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewObjectMeta() *ObjectMeta {
	return &ObjectMeta{}
}

type ObjectMeta struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}
