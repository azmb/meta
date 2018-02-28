# Reproducing gRPC issue

This minimal implementation helps to reproduce gRPC issue. When we are trying to send `\n` in gRPC metadata this causes:
```
RPC call failed: rpc error: code = Internal desc = stream terminated by RST_STREAM with error code: PROTOCOL_ERROR
```
