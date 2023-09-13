## Deployments

- deployments are better than just a pod because they can help with self healing and scaling.

A deployment creates a replicaset (controller) which will be responsible for creating the pod(s), and keeping them runing

A controller ensures that the desired state is the same as the activee state

Q: what is the difference between a replica set and a deployment?

a replicaset is responsibe for autohealing, it is a controller, a rs is part of a deployment.

- lets say someone deletes a pod accidently (not part of a deployment)

[deployment reference](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

- you can always ssh into a deployment but without a service you cannot access it from the outside
