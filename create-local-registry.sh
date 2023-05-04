#!/bin/bash
set -e

k run registry --image docker.io/registry
k expose pod/registry --port 5000   
skopeo copy tarball:/tmp/cryptofaker.tar.gz docker://localhost:5000/cryptofaker:latest --dest-tls-verify=false
k run cryptofaker --image registry.default:5000/cryptofaker:latest
