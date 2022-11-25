package kubeclient

import (
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func getServiceObject(name string) *coreV1.Service {
	return &coreV1.Service{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name: name,
		},
		Spec: coreV1.ServiceSpec{
			Type: coreV1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": name,
			},
			Ports: []coreV1.ServicePort{
				{
					Port:       25565,
					TargetPort: intstr.FromInt(25565),
				},
			},
		},
	}
}
