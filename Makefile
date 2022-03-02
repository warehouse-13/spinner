up:
	minikube start --nodes 2 --addons metrics-server
	kubectl apply -f ./deployment.yaml
	kubectl autoscale deployment spinner --cpu-percent=2 --min=1 --max=10