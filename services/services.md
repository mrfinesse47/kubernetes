##Service

#### Node Port

- to create a service from a definition file

kubectl create -f service-definition.yaml

- to get services

kubectl get services

- then you can reach things by using curl

curl http://[ipaddress]:[nodeport]

- an interestng thing is the service is shared across nodes so if you use one address it will load balance across all of them.

- to get the minikube external ip for a service

minikube service myapp-service --url
