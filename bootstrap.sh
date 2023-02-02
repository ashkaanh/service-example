#!/bin/bash
cd k8s/charts/myapp
kubectl create ns dev-myapp
helm install dev -n dev-myapp .
