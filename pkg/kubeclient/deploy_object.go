package kubeclient

import (
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getDeployObject(name string) *appsV1.Deployment {
	replicas := int32(1)

	return &appsV1.Deployment{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name: name,
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: getPodSpecObject(name),
			},
		},
	}
}
