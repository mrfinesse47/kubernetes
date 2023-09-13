## services

why do we need services?

if there were no services multiple pods have different ip addresses, and their addresses can change

so instead of giving users a bunch of different ip addresses that can change all the time, so instead they will access by the payment service name such as:

payment.default.svc

### a service can

- 1, the service will also act as a load balancer,
- 2, the service also provides service discovery, it does this with labels and selectors.
- 3, the service will expose to the outside world

## there are 3 defualt types of services\

- 1. Cluster IP -> discovery and load balancing internal to cluster, nothing much changes?
- 2. Node Port -> allows our app to be accessed inside our organization
- 3. Load Balancer -> will expose our app to the external world, it is usually from the cloud controller and provided by aws etc.
