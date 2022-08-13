package k8s

import "k8s.io/apimachinery/pkg/util/intstr"

func NewIntStr(v int) *intstr.IntOrString {
	is := intstr.FromInt(v)
	return &is
}
