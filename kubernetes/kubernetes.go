package kubernetes

import (
	"context"
	"log"
	"os"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getNamespace() string {
	// namespace will be in /var/run/secrets/kubernetes.io/serviceaccount/namespace
	data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}

func CreateWireguardServerSecret(secretName string, privateKey wgtypes.Key) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secret, err := clientset.CoreV1().Secrets(getNamespace()).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	if secret != nil {
		log.Println(secret)
	}
}
