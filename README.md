# Node Down Webhook

This is a webhook server meant to receive calls from the Prometheus AlertManager, whenever a node is in the NotReady status for longer than 15min. It will then delete the node from the cluster.

## Deploy to your Kubernetes cluster

This section assumes you have a local multi-node Kubernetes cluster running, created by
running this command:

```bash
kind create cluster --config ./testdata/kind.yaml
```

Build and deploy the webhook server:

```bash
make docker-build kind-load deploy IMG=node-down-webhook:v0.0.1-$(date +%s)
```

You can check the webhook server's logs like so

```bash
kubectl logs -f node-down-handler-66c665df7c-rcph2
```
