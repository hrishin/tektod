package foobar

import (
	"fmt"
	"os"

	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

//ListPipelines lists pipelines
func ListPipelines(kubeConfig *rest.Config, namespace string) {
	tektonCient, err := versioned.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client from config %s  %s", kubeConfig, err)
		os.Exit(1)
	}
	fmt.Println("tekton client is created")

	pipelineAPI := tektonCient.TektonV1alpha1().Pipelines("")
	ps, err := pipelineAPI.List(metav1.ListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to llist pipelines %v", err)
		return
	}
	fmt.Printf("Got %v pipeline(s) \n", len(ps.Items))

	for _, p := range ps.Items {
		fmt.Println(p)
		// do your stuff here
	}
}
