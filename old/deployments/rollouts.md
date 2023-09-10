--get the status of a deployment

kubectl rollout status deployment [deployment name]

-- get the rollout history

kubectl rollout history deployment [deployment name]

- to apply changes ater changing a deployment file

kubectl apply -f [deployment definition]

- this will create a new rollout and revisiton

-- to undo a change

kubectl rollout undo deployment/myapp-deployment

- this will destroy the pods in the new replica set and create the pods from the old replica set.

- to record the cause of change

kubectl create -f deployment.yaml record

- to set the image of a deployment

kubectl set image deployment my-app deployment nginx=nginx:1.18-perl
