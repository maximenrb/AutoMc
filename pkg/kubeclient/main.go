package kubeclient

import (
	"context"
	"flag"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type Client struct {
	*kubernetes.Clientset
}

func New() *Client {
	var kubeConfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	flag.Parse()

	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientSet
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &Client{
		Clientset: client,
	}
}

func (client *Client) GetPods() {
	pods, err := client.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func (client *Client) CreateNewServer(name string) {
	client.createPod(name)
}

func (client *Client) CreateNewExposedServer(name string) (int32, bool) {
	client.createDeployment(name)
	client.waitMinecraftServerUp()
	return client.createService(name)
}

func (client *Client) waitMinecraftServerUp() {
	logs := client.CoreV1().Pods("seedbox").GetLogs("lala", &coreV1.PodLogOptions{})
	watch, err := logs.Watch(context.TODO())
	if err != nil {
		fmt.Println("Watch Error: ", err)
		return
	}

	fmt.Println("Watch: ")
	fmt.Println(watch)
}

func (client *Client) createPod(name string) {
	_, err := client.CoreV1().Pods("seedbox").Create(context.TODO(), getPodObject(name), metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("Unable to create pod", err)
		return
	}
	fmt.Println("Pod successfully created")
}

func (client *Client) createDeployment(name string) {
	deployment, err := client.AppsV1().Deployments("seedbox").Create(context.TODO(), getDeployObject(name), metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("Unable to create deployment", err)
		return
	}
	fmt.Println("Deployment successfully created", deployment)
}

func (client *Client) createService(name string) (int32, bool) {
	service, err := client.CoreV1().Services("seedbox").Create(context.TODO(), getServiceObject(name), metaV1.CreateOptions{})
	if err != nil {
		fmt.Println("Unable to create service", err)
		return 0, false
	}
	fmt.Println("Service successfully created, ports: ", service.Spec.Ports[0].NodePort)
	return service.Spec.Ports[0].NodePort, true
}
