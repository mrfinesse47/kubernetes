Replica sets

kubectl create -f replicates-definition.yaml

kubectl get replicaset

kubectl get pods

// how to scale up say if you already have 3 pods running

kubectl replace -f replicase-definition.yaml // in which you define 6

// or by command line

kubectl scale -replicas=6 -f replicase-definition.yaml

![Replica Set](./replica-set-example.jpg 'Replica Set Example')

Useful commands

![Replica Set](./useful-commands.jpg 'Useful Commands')

// to edit a replicaset by command line

kubectl edit replicaset myapp-replicaset

// or to scale by command

kubectl scale replicaset myapp-replicaset --replicas=2

// if you need to look into a pod description

kubectl describe pod [podname]
