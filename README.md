# spinner

a tiny thing to smash the cpu.

## Getting Started
locally:
```
make up
```

or

```
minikube start --nodes 2 --addons metrics-server // Create a cluster with 2 nodes & a metrics-server
kubectl apply -f deployment.yaml // Apply the Deployment & Service manifest:
kubectl autoscale deployment spinner --cpu-percent=2 --min=1 --max=10
```

EKS:

Create a cluster with OIDC enabled:
```
eksctl create cluster --with-oidc
```

Enable the metrics server and the autoscaler:
```
make eks-up
```

or

```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
kubectl get deployment metrics-server -n kube-system
kubectl autoscale deployment spinner --cpu-percent=2 --min=1 --max=10
```
