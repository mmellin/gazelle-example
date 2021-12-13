load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_bazel_gomock",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/jmhodges/bazel_gomock",
        sum = "h1:d3jFA1JzkFVXN61mrTKGI4NzujCyMmjiEBGx53EyDyA=",
        version = "v0.0.0-20210217072139-fde78c91cf17",
    )
