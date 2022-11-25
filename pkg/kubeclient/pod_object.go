package kubeclient

import (
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getPodObject(name string) *coreV1.Pod {
	return &coreV1.Pod{
		TypeMeta: metaV1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name: name,
		},
		Spec: getPodSpecObject(name),
	}
}

func getPodSpecObject(name string) coreV1.PodSpec {
	user := int64(1001)
	group := int64(1002)
	diskSizeLimit := resource.MustParse("10Gi")

	return coreV1.PodSpec{
		SecurityContext: &coreV1.PodSecurityContext{
			RunAsUser:  &user,
			RunAsGroup: &group,
		},
		Containers: []coreV1.Container{
			{
				Name:  name,
				Image: "maxnrb/velocityxpapermc:1.0.1",
				VolumeMounts: []coreV1.VolumeMount{
					{
						Name:      "cache-volume",
						MountPath: "/data",
					},
				},
				Env: []coreV1.EnvVar{
					{
						Name:  "MC_RAM",
						Value: "1G",
					},
					{
						Name:  "WORLD_URL",
						Value: "http://fileserver:5000/world/lobby.zip",
					},
					{
						Name:  "PLUGINS_URL",
						Value: "http://fileserver:5000/plugin/NoEncryption-v4.3--1.19.2_only.jar",
					},
					{
						Name:  "VELOCITY_SECRET",
						Value: "Lh7UyFG1mcbR",
					},
				},
			},
		},
		Volumes: []coreV1.Volume{
			{
				Name: "cache-volume",
				VolumeSource: coreV1.VolumeSource{
					EmptyDir: &coreV1.EmptyDirVolumeSource{
						SizeLimit: &diskSizeLimit,
					},
				},
			},
		},
	}
}
