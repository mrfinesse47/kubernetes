for mac osx to use for more advanced concepts
do not use with the docker driver instead

> minikube start --memory=4096 --driver=hyperkit

-or

> minikube start --driver=hyperkit

if not using advanced features

> minikube start

is fine

- to log into a cluster

> minikube ssh

- you can ping non exposed pods from the inside or curl
- if not for minikube we would ssh into the master node to do this

> minikube status
