up:
	minikube start --nodes 2 --addons metrics-server
	kubectl apply -f ./deployment.yaml
	kubectl autoscale deployment spinner --cpu-percent=2 --min=1 --max=10

eks-up:
	kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
	kubectl get deployment metrics-server -n kube-system
	kubectl autoscale deployment spinner --cpu-percent=2 --min=1 --max=10

