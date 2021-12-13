# Example Bazel repo with Gazelle

The purpose of this repo is to learn how to configure [Gazelle](https://github.com/bazelbuild/bazel-gazelle) to auto-generate `BUILD` files and rules for a Golang application which includes `.proto` files. 

## Setup

The following is required to be installed in order to build this repo:

- [Golang 1.17.3](https://go.dev/doc/install)
- [Bazelisk](https://docs.bazel.build/versions/main/install-bazelisk.html) version `1.10.1`
- Linux-based 64-bit system (recommended)

## Build app

```
$ bazel build //...
```

## Useful commands

- Create/update all `BUILD.bazel` files automatically:
    ```
    bazel run //:gazelle
    ```
- Add a new `go_repository()` dependency:
    ```
    bazel run //:gazelle -- update-repos -build_file_proto_mode=disable_global -to_macro=deps.bzl%go_dependencies <repo url>
    ```

## Compilers

The default gRPC compiler is `@io_bazel_rules_go//proto:go_grpc`. 

In order to support creation of gRPC `*.pb.gw.go` files, I've added another compiler via a Gazelle directive:

```
# gazelle:go_grpc_compilers @io_bazel_rules_go//proto:go_grpc,@com_github_grpc_ecosystem_grpc_gateway//protoc-gen-grpc-gateway:go_gen_grpc_gateway
```

I've also needed to use the `update-repos` Gazelle command to add this new repo and its deps to the `deps.bzl` using by `WORKSPACE` for dependency resolution.

```
$ bazel run //:gazelle -- update-repos -build_file_proto_mode=disable_global -to_macro=deps.bzl%go_dependencies github.com/grpc-ecosystem/grpc-gateway 

$ bazel run //:gazelle -- update-repos -build_file_proto_mode=disable_global -to_macro=deps.bzl%go_dependencies github.com/ghodss/yaml
```

See more about directives: https://github.com/bazelbuild/bazel-gazelle/tree/v0.24.0#directives

## Bazel Versions

The version of Bazel is set in the `.bazeliskrc` file.