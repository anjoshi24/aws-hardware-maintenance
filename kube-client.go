package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// path-to-kubeconfig
	config, _ := clientcmd.BuildConfigFromFlags("", "/home/anjoshi/work/aws-hardware-maint_bak/anjoshi-test-01")

	// creates the clientset
	clientset, _ := kubernetes.NewForConfig(config)

	nodes, err := clientset.Core().Nodes().List(context.TODO(), v1.ListOptions{})
	fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
}
