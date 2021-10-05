# Glog

__glog__ is an abstraction for various Go language log libraries.

The ```Logger``` interface is in maintenance-mode to prevent broken changes to projects based on it. But we will add more different implementations (wrappers) for various log libraries. 

Logger doesn't implements "_Fatal_" level because logger shouldn't terminate a process. You can use Panic of Panicf methods without os.Exit(1) and terminate process clearly.

```context.go``` implements helper functions to attach logger over a context. From context will return empty usable logger if there is no logger attached. It prevents app chash if context doesn't conains Logger implementations.

```grpc_interceptors.go``` implemets gRPC interceptors to replace embeeded gRPC logger by the Logger implementations.

All interface implementations are located in different directories. You can use any of it or implement your own one.

This repository will be archived when Go implements it's own standard logger interface.