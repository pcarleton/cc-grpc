#!/usr/bin/env bash
set -eux 

#bazel run //server/bin/run_server:debug -- --norun
docker tag bazel/server/bin/run_server:debug $CONTAINER_NAME
docker push $CONTAINER_NAME

gcloud compute instances stop grpc-server                                                       
gcloud compute instances start grpc-server                                                       

