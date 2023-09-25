To rollback a deployment to a previous replica set in Kubernetes, you can use the `kubectl rollout` command. Here are the steps to perform a rollback:

1. View the rollout history of the deployment to identify the revision you want to roll back to. You can use the following command to see the revision history:

   ```bash
   kubectl rollout history deployment <deployment-name>
   ```

   Replace `<deployment-name>` with the name of your deployment.

2. Roll back to a specific revision by using the `kubectl rollout undo` command. Specify the revision number you want to roll back to:

   ```bash
   kubectl rollout undo deployment <deployment-name> --to-revision=<revision-number>
   ```

   Replace `<deployment-name>` with the name of your deployment, and `<revision-number>` with the revision you want to roll back to.

3. Monitor the deployment's status to ensure that the rollback is successful:

   ```bash
   kubectl rollout status deployment <deployment-name>
   ```

   Replace `<deployment-name>` with the name of your deployment.

After executing these commands, Kubernetes will roll back the deployment to the specified revision, and the previous replica set will be used. You can verify the status of the rollout and ensure that your application is functioning as expected.
