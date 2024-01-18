package main

import (
	"flag"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
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

	}

	// gives the list of clientset that can interact with k8 components
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("error in building client set is::", err.Error())
	}
	informerFactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	informers.NewFilteredSharedInformerFactory(clientset, 10*time.Minute, namespace, func(lo *v1.ListOptions) {})
	podinformer := informerFactory.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add func is called")
		},

		DeleteFunc: func(obj1 interface{}) {
			fmt.Println("delete func is called")
		},

		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("update func is called")
		},
	})

	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)
	pod, err := podinformer.Lister().Pods(namespace).Get(namespace)
	if err != nil {
		fmt.Println("error is ::::", err)
	}

	fmt.Println("pod is::", pod)
}
