load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "map.go",
        "merge.go",
        "mergo.go",
    ],
    importmap = "github.com/kubernetes-sigs/poseidon/vendor/github.com/imdario/mergo",
    importpath = "github.com/imdario/mergo",
    visibility = ["//visibility:public"],
)
