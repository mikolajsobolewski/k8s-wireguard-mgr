package kubernetes

import (
	"context"
	"errors"
	"log"
	"os"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getNamespace() string {
	//namespace will be in /var/run/secrets/kubernetes.io/serviceaccount/namespace
	data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}

type apiStatus interface {
	Status() metav1.Status
}

func reasonForError(err error) metav1.StatusReason {
	if status, ok := err.(apiStatus); ok || errors.As(err, &status) {
		return status.Status().Reason
	}
	return metav1.StatusReasonUnknown
}

func kubernetesErrorIsAlreadyExists(err error) bool {
	return reasonForError(err) == metav1.StatusReasonAlreadyExists
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
	// try and create the secret, if it already exists we will get an error
	// check the error result, if it's the error about it already existing
	// exit with status 0, this is done because we do not want to request to read access to the secret in our role
	// binding but it is ok to request the ability to create one as this can't be used to read existing secret values
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: getNamespace(),
		},
		StringData: map[string]string{
			"privatekey": privateKey.String(),
		},
		Type: corev1.SecretTypeOpaque,
	}
	_, err = clientset.CoreV1().Secrets(getNamespace()).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		if !kubernetesErrorIsAlreadyExists(err) {
			panic(err.Error())
		} else {
			log.Println("Secret already exists. Skipping creation.")
		}
	} else {
		log.Println("Secret created successfully.")
	}
}
