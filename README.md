# Cashcoach GRPC

This is a GRPC server for interacting with Cashcoach.  Cachcoach is a personal finance API server that helps categorize and settle up your transactions.

This repo is based on github.com/pcarleton/grpc-starter

# Getting Started

## Local Development

To run locally, first [install bazel](https://docs.bazel.build/versions/master/install.html).  Then run:

```
# In one shell Start the server
# (This will probably take awhile to build the first time)
bazel run //server/bin/run_server -- -insecure


# Start the client
bazel run //server/bin/run_client -- -insecure -address localhost:5001
```
