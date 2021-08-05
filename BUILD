load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/jlamb1/gin-api
gazelle(name = "gazelle")

go_library(
    name = "gin-api_lib",
    srcs = ["main.go"],
    importpath = "github.com/jlamb1/gin-api",
    visibility = ["//visibility:private"],
    deps = ["@com_github_gin_gonic_gin//:gin"],
)

go_binary(
    name = "gin-api",
    embed = [":gin-api_lib"],
    visibility = ["//visibility:public"],
)
