package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pradeepto/tektod/pkg/foobar"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("🐈 tektod")
	kubeconfig := loadKubeConfig()
	foobar.ListPipelines(kubeconfig, "default")
}

//LoadKubeConfig loads kubecofnig
func loadKubeConfig() *rest.Config {
	//path for kubeconfig
	homeDir := os.Getenv("HOME")
	kubeCfgFile := filepath.Join(homeDir, ".kube", "config")

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeCfgFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return kubeConfig
}
