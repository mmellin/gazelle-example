load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_bazel_gomock",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/jmhodges/bazel_gomock",
        sum = "h1:d3jFA1JzkFVXN61mrTKGI4NzujCyMmjiEBGx53EyDyA=",
        version = "v0.0.0-20210217072139-fde78c91cf17",
    )
    # go_repository(
    #     name = "com_github_ghodss_yaml",
    #     build_file_proto_mode = "disable_global",
    #     importpath = "github.com/ghodss/yaml",
    #     sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    #     version = "v1.0.0",
    # )

    # go_repository(
    #     name = "com_github_grpc_ecosystem_grpc_gateway",
    #     build_file_proto_mode = "disable_global",
    #     importpath = "github.com/grpc-ecosystem/grpc-gateway",
    #     sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    #     version = "v1.16.0",
    # )
    # go_repository(
    #     name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    #     build_file_proto_mode = "disable_global",
    #     importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
    #     sum = "h1:p5m7GOEGXyoq6QWl4/RRMsQ6tWbTpbQmAnkxXgWSprY=",
    #     version = "v2.7.1",
    # )
