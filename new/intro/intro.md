### Introduction

#### Containers

- Are ephemeral in nature, ie they have a short life

#### Kubernetes solves the main problems of:

- A single host
- Autoscaling
- Autohealing
- Enterprise level support

#### Elaboration

- By default K8S is a cluster
- it is installed in a master node architecture
- when we practice we usually use one node on our local machine, i.e. minikube
- you can scale by using replica sets or HPA, Hpa can look and see if a node is reaching a defined limit lets say 80% it will spin up another container
- Autohealing: whenever there is damage k8s controls ansd fixes the damage. Lets say a container goes down, for many a reason, k8s will start a new container even before it fully goes down, the api server detects the container going down, the api server will rollout a new container.
- Enterprise Nature: Docker is never used in production, even k8s isnt enough but can be extended.
