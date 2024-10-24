Deploy, manage, and scale a simple Go web app on a local Kubernetes cluster created using minikube using the following tutorial: https://coding-bootcamps.com/build-containerized-applications-with-golang-on-kubernetes

## Create docker image
Dockerfile References: https://docs.docker.com/engine/reference/builder

```
export DOCKER_BUILDKIT=0
export COMPOSE_DOCKER_CLI_BUILD=0

docker build -t go-kubernetes .                            # Build the docker image
docker tag go-kubernetes <username>/go-name-store:1.0.0    # Tag the image
docker login                                               # Login to docker
docker push <username>/go-name-store:1.0.0                 # Push the image to docker hub
```

## Creating a kubernetes cluster locally
```
minikube start
kubectl apply -f k8s-deployment.yml
kubectl get deployments
kubectl get pods
kubectl port-forward go-name-store-7494cf9955-9vjx8 8080:8080
kubectl logs -f go-name-store-7494cf9955-9vjx8
```

## Creating a service
```
# kubectl apply -f k8s-deployment.yml
# kubectl get services 
# minikube service go-name-store-service --url             # in a seperate terminal, keep terminal open 
```

## Using the servive
```
curl -s -X PUT 'localhost:8080?name=Hasti'
```
