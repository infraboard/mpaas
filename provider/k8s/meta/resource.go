package meta

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewApiResourceList(items []*metav1.APIResourceList) *ApiResourceList {
	return &ApiResourceList{
		Items: items,
	}
}

type ApiResourceList struct {
	Items []*metav1.APIResourceList
}

func (l *ApiResourceList) GetResourceByName(name string) metav1.APIResource {
	for i := range l.Items {
		group := l.Items[i]
		for ri := range group.APIResources {
			resource := group.APIResources[ri]
			if resource.Name == name {
				return resource
			}
		}
	}

	return metav1.APIResource{}
}
