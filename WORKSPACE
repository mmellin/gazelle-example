workspace(name = "gazelle-example")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
    ],
)
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)


load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:nogo_vet",
    version = "1.17.5"
)

gazelle_dependencies()

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

#######  @com_google_protobuf package and dependencies

http_archive(
    name = "com_google_protobuf",
    sha256 = "87407cd28e7a9c95d9f61a098a53cf031109d451a7763e7dd1253abf8b4df422",
    strip_prefix = "protobuf-3.19.1",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.19.1.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.19.1.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    sha256 = "e30f2aa357e4cf3dc9b347ee5703ce0346f05b127a10f3da5b359748e01c65f1",
    strip_prefix = "grpc-gateway-2.7.1",
    urls = ["https://github.com/grpc-ecosystem/grpc-gateway/archive/refs/tags/v2.7.1.tar.gz"],
)

load("@com_github_grpc_ecosystem_grpc_gateway_v2//:repositories.bzl", grpc_gateway_go_deps = "go_repositories")

grpc_gateway_go_deps()
