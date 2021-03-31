workspace(name = "com_rogchap_v8go")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "69de5c704a05ff37862f7e0f5534d4f479418afc21806c887db544a316f3cb6b",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.27.0/rules_go-v0.27.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.16.2")

http_archive(
    name = "bazel_gazelle",
    sha256 = "6148b0430093ccff298f17c1bcf555f449fa35ea90e0c96d3324cd97c75d48f7",
    strip_prefix = "bazel-gazelle-50712ce78b4843fc8620278075383e1ca53dfd74",
    url = "https://github.com/yext/bazel-gazelle/archive/50712ce78b4843fc8620278075383e1ca53dfd74.zip",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()
