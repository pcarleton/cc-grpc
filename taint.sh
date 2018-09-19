#!/usr/bin/env bash

set -eux

cd terraform/prod
terraform taint google_compute_instance.api
terraform apply
