# canary-deployment-poc

<img width="1225" alt="simple-split" src="https://github.com/kwakubiney/canary-deployment-poc/assets/71296367/488070d0-84d6-4013-b185-63b64ebd031c">


# Demo

- Follow [these steps](https://github.com/nginxinc/nginx-gateway-fabric/blob/main/docs/installation.md) to install NGINX Gateway Fabric in your infrastructure.

- Build the docker images with the `Dockerfile` in both the `v1` and `v2` directories and apply the deployments after with `kubectl apply -f <deployment filename>`

- Run `kubectl apply -f gateway.yml` to deploy the gateway

- Run `kubectl apply -f services.yml` to deploy both `v1` and `v2` services

- Run `kubectl apply -f traffic-split.yml` to deploy the split routing rules. This default configuration routes all traffic to v1.

- Run `curl.sh <Gateway IP> <Gateway Port>` to simulate requests to the backends with a returned number of requests to both backends.


# The fun part

- [Demo](https://www.youtube.com/watch?v=IWJyc0p1UwY)
