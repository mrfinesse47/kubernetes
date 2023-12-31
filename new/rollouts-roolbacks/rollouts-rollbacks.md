Sure, here is the provided text formatted as Markdown (.md):

````markdown
# Updating a Deployment

Here are some handy examples related to updating a Kubernetes Deployment:

## Creating a deployment, checking the rollout status and history:

In the example below, we will first create a simple deployment and inspect the rollout status and the rollout history:

```shell
controlplane $ kubectl create deployment nginx --image=nginx:1.16
deployment.apps/nginx created

controlplane $ kubectl rollout status deployment nginx
Waiting for deployment "nginx" rollout to finish: 0 of 1 updated replicas are available...
deployment "nginx" successfully rolled out

controlplane $

controlplane $ kubectl rollout history deployment nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>

controlplane $
```
````

Using the `--revision` flag:

Here revision 1 is the first version where the deployment was created.

You can check the status of each revision individually by using the `--revision` flag:

```shell
controlplane $ kubectl rollout history deployment nginx --revision=1
deployment.apps/nginx with revision #1
Pod Template:
Labels:       app=nginx
pod-template-hash=78449c65d4
Containers:
  nginx:
    Image:      nginx:1.16
    Port:       <none>
    Host Port:  <none>
    Environment: <none>
    Mounts:      <none>
Volumes:

controlplane $
```

Using the `--record` flag:

You would have noticed that the “change-cause” field is empty in the rollout history output. We can use the `--record` flag to save the command used to create/update a deployment against the revision number.

```shell
controlplane $ kubectl set image deployment nginx nginx=nginx:1.17 --record
Flag --record has been deprecated, --record will be removed in the future
deployment.apps/nginx image updated

controlplane $

controlplane $ kubectl rollout history deployment nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>
2         kubectl set image deployment nginx nginx=nginx:1.17 --record=true

controlplane $
```

You can now see that the change-cause is recorded for the revision 2 of this deployment.

Let's make some more changes. In the example below, we are editing the deployment and changing the image from `nginx:1.17` to `nginx:latest` while making use of the `--record` flag.

```shell
controlplane $ kubectl edit deployments.apps nginx --record
Flag --record has been deprecated, --record will be removed in the future
deployment.apps/nginx edited

controlplane $ kubectl rollout history deployment nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>
2         kubectl set image deployment nginx nginx=nginx:1.17 --record=true
3         kubectl edit deployments.apps nginx --record=true

controlplane $ kubectl rollout history deployment nginx --revision=3
deployment.apps/nginx with revision #3
Pod Template:
  Labels:       app=nginx
  pod-template-hash=787f54657b
  Annotations:  kubernetes.io/change-cause: kubectl edit deployments.apps nginx --record=true
  Containers:
   nginx:
    Image:      nginx
    Port:       <none>
    Host Port:  <none>
    Environment:  <none>
    Mounts:     <none>
  Volumes:

controlplane $
```

## Undo a change:

Let's now rollback to the previous revision:

```shell
controlplane $ kubectl rollout history deployment nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1
3         kubectl edit deployments.apps nginx --record=true
4         kubectl set image deployment nginx nginx=nginx:1.17 --record=true

controlplane $ kubectl rollout history deployment nginx --revision=3
deployment.apps/nginx with revision #3
Pod Template:
  Labels:       app=nginx
  pod-template-hash=787f54657b
  Annotations:  kubernetes.io/change-cause: kubectl edit deployments.apps nginx --record=true
  Containers:
   nginx:
    Image:      nginx:latest
    Port:
    Host Port:
    Environment:
    Mounts:
  Volumes:

controlplane $ kubectl describe deployments. nginx | grep -i image:
    Image:        nginx:1.17

controlplane $
```

With this, we have rolled back to the previous version of the deployment with the image `nginx:1.17`.

```shell
controlplane $ kubectl rollout history deployment nginx --revision=1
deployment.apps/nginx with revision #1
Pod Template:
  Labels:       app=nginx
  pod-template-hash=78449c65d4
  Containers:
   nginx:
    Image:      nginx:1.16
    Port:       <none>
    Host Port:  <none>
    Environment: <none>
    Mounts:     <none>
  Volumes:

controlplane $ kubectl rollout undo deployment nginx --to-revision=1
deployment.apps/nginx rolled back
```

To rollback to a specific revision, we will use the `--to-revision` flag. With `--to-revision=1`, it will be rolled back with the first image we used to create a deployment as we can see in the rollout history output.

```shell
controlplane $ kubectl describe deployments. nginx | grep -i image:
Image: nginx:1.16
```

```

I hope this helps! Let me know if you have any other questions.
```

In Markdown format, you can use triple backticks (```) to format code blocks and single backticks (`) for inline code. Additionally, headers are formatted with the pound (#) symbol, and bullet points can be created with asterisks (\*).
