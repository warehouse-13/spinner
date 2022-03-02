# spinner

a tiny thing to smash the cpu.

## Getting Started
```
make up
```

or

```
minikube start --nodes 2 --addons metrics-server // Create a cluster with 2 nodes & a metrics-server
kubectl apply -f deployment.yaml // Apply the Deployment & Service manifest:
kubectl autoscale deployment spinner --cpu-percent=2 --min=1 --max=10
```
