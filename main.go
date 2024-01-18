package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	namespace = "ns-two"
)

func main() {
	kubeConfig := flag.String("kubeconfig", "/home/saima/.kube/config", "location to kube config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		fmt.Println("error in building config is::", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("error in gettin config from cluster is::", err.Error())
		}
		// handle error
	}
	//fmt.Println(config)
	//co:qnfig.Timeout = 120 * time.Second
	// gives the list of clientset that can interact with k8 components
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("error in building client set is::", err.Error())
	}
	// informerFactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	//fmt.Println("client set are:::", clientset)
	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Println("error in listing pods:::", err)
	}
	fmt.Println("no of pods are ::::", len(pods.Items))
	for i := range pods.Items {
		fmt.Println("pods are:::", pods.Items[i].Name)
	}

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Println("error in listing deployments:::", err)
	}

	fmt.Println("no of deployments are ::::", len(deployments.Items))
	for i := range deployments.Items {
		fmt.Println("deployment is:::", deployments.Items[i].Name)
	}
	// fmt.Println("pods are:::", pods)

}
