#!/bin/bash

echo -n "admin:"
kubectl -n prometheus get secret prometheus-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
