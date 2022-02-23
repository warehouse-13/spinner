# spinner

a tiny thing to smash the cpu.

## Getting Started
Open a terminal with `htop` on the right side to monitor the CPU usage of the local machine.

Create a kind cluster with 3 nodes:
```
kind create cluster –name ipv6 –config kind-config.yaml
```

Apply the Deployment & Service manifest:
```
k apply -f deployment.yaml
```

### Install a Metrics Server on the cluster. 
To meet the [Requirements](https://github.com/kubernetes-sigs/metrics-server#requirements), follow these steps:

TODO: Hack the upstream Node Bootstrap script to enable aggregation layer, so that we can deploy a metrics server: https://github.com/awslabs/amazon-eks-ami/blob/master/files/bootstrap.sh

Install the Metrics Server through the following yaml:
```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

### Scale Pod
`k autoscale deployment spinner –cpu-percent=10 –min=1 –max=10` 
