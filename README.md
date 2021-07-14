# throttling-test

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#requirements">Requirements</a></li>
    <li><a href="#payload">Payload</a></li>
    <li><a href="#run-locally">Run Locally</a></li>
    <li><a href="#run-on-kubernetes">Run on Kubernetes</a></li>
    <li><a href="#expected-response">Expected Response</a></li>
  </ol>
</details>

## prerequisites

  1. [Docker](https://docs.docker.com/docker-for-mac/install/)

  2. [Kubectl](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-kubectl/)

  3. [Go 1.16](https://golang.org/doc/install?download=go1.16.darwin-amd64.pkg)

## requirements

```
 ❯ maximum of 2 request per second and regardless of the x-secret-key

 ❯ maximum of 10 request per minute and per x-secret-key
```

## payload

```
curl -X POST \
    http://localhost:8080/ping \
    -H 'content-type: application/json' \
    -H 'x-secret-key: ahs98h' \
    -d '{"request":"ping"}'
```

## run locally

```
❯ go mod init github.com/fabmorais/throttling-test

❯ go mod tidy -v

❯ go run main.go
```

## run on kubernetes

```
❯ docker build -t ping-pong-app:latest .

❯ kubectl apply -f kube/deployment.yaml

❯ kubectl apply -f kube/service.yaml
```

## expected response

```

(not throttled):
 ❯ {"response":"pong"}

(throttled)
 ❯ {"message":"request throttled request","throttle_age":"10"}

```
