load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

# gazelle:prefix github.com/jlamb1/gin-api
gazelle(name = "gazelle")

go_binary(
    name = "api",
    srcs = ["main.go"],
)