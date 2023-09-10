- very similar to replica set definition files

- to create:

kubectl create -f deployment-definition.yaml

- to get deployments:

kubectl get deployments

- so if we get replicasets we will see a new replicaset in the name of the deployment

kubectl get replicasets

- and these replica sets create pods

kubectl get pods

- to see all the created objects at once run

kubectl get all
