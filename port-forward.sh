#!/bin/bash

set -e

kubectl -n prometheus port-forward svc/prometheus-grafana 3000:80 > /dev/null &
