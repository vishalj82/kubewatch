#!/bin/bash
cd  ./../service/deployment
pwd
kubectl delete  mutatingwebhookconfigurations mutating-webhook-configuration
kubectl delete -f deployment.yaml
kubectl apply -f namespace.yaml
kubectl apply -f service.yaml
kubectl apply -f deployment.yaml
