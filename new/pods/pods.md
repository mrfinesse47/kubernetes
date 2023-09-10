### Pods

- A pod is a definition on how to run a container
- A pod is usually a single container

#### advantages of multi container pods

- shared networking (talk to each other using localhost)
- shared storage

! this is very rare

- a more common case are side car containers, which will be covered later

#### notes

- a pod is a wrapper for a container to make the lives of dev ops engineers easier

[reference](https://kubernetes.io/docs/concepts/workloads/pods/)

- even the pros look at the reference to get example yaml files

#### to get pods

> kubectl get pods

#### to get detailed pods

> kubectl get pods -o wide
