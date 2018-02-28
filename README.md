# Reproducing gRPC issue

This minimal implementation helps to reproduce gRPC issue https://github.com/grpc/grpc-go/issues/1888. 
When we are trying to send `\n` in gRPC metadata this causes:
```
rpc error: code = Internal desc = stream terminated by RST_STREAM with error code: PROTOCOL_ERROR
```
