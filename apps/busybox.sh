#!/bin/bash
kubectl run busybox --image docker.io/busybox -- sleep infinity  

kubectl exec -it busybox -- busybox sh
