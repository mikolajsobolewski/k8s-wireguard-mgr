package main

import (
	"log"
	"os"

	"github.com/bryopsida/k8s-wireguard-mgr/kubernetes"
	"github.com/bryopsida/k8s-wireguard-mgr/wireguard"
)

func main() {
	log.Println("Starting Kubernetes Wireguard Manager")
	key := wireguard.GenerateWireguardKey()
	secretName := os.Getenv("K8S_WG_MGR_SERVER_SECRET_NAME")
	kubernetes.CreateWireguardServerSecret(secretName, key)
	log.Println("Finished")
}
