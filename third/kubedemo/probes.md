probes

-- see probe with wrong port

now if you describe the pod you will see in the events

Warning Unhealthy 2s (x2 over 12s) kubelet Readiness probe failed: Get "http://10.244.0.104:9999/": dial tcp 10.244.0.104:9999: connect: connection refused

-- the liveness probe is useful because you can set how many fails are allowed before its considered to be a fail
