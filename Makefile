READY_STATE ?= true

image:
	docker build . -t ghcr.io/bryopsida/k8s-wireguard-mgr:local

run:
	docker run ghcr.io/bryopsida/k8s-wireguard-mgr:local

test:
	skaffold build -q > build_result.json
	skaffold deploy --load-images=true -a build_result.json
	skaffold verify -a build_result.json
