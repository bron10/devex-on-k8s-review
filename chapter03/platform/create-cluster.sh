#!/bin/sh

set -euo pipefail

echo "\n📦 Initializing Kubernetes cluster..."

kind delete cluster --name devex-cluster

kind create cluster --config kind-config.yml

echo "\n🔌 Installing Contour Ingress..."

kubectl apply -f https://projectcontour.io/quickstart/contour.yaml 

sleep 10

kubectl wait --namespace projectcontour \
  --for=condition=ready pod \
  --selector=app=contour \
  --timeout=60s

kubectl wait --namespace projectcontour \
  --for=condition=ready pod \
  --selector=app=envoy \
  --timeout=60s

echo "\n🐘 Installing CloudNativePG..."

kubectl apply --server-side -f \
  https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/release-1.26/releases/cnpg-1.26.0.yaml

sleep 10

kubectl wait --namespace cnpg-system \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/name=cloudnative-pg \
  --timeout=60s

echo "\n⛵ Happy Sailing!\n"
