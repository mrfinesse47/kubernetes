## namespaces

- to get pods from a namespace

kubectl get pods --namespace=default

- to set the namespace to a diferent one

kubectl config set-context --current --namespace=my-namespace

- to view pods in all name spaces

kubectl get pods --all-namespace or kubectl get pods -A

a service in the dev name space could be reached by

db-service.dev.svc.cluster.local

- cluster.local is the default domain name of the k8s cluster
- svc is the sub domain for service
- dev is the namespace
- db-service is the service name

### for a definition file.

- add the namespace under the meta tag
