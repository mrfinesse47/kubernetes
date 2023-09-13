#### General info

[a good cheatsheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

### useful commands

> kubectl logs [name]

- to get all pods in every namespace

> kubectl get pods -A

- to get expanded pod details

> kubectl get pods -o wide

- to watch pods live

> kubectl get pods -w : try this while deleting a pod in a deployment

- never forget you can exec into a pod

> kubectl exec -it [pod] sh
