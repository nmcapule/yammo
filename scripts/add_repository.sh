#!/bin/bash
bazel run //:gazelle -- update-repos $1 -to_macro=repositories.bzl%go_repositories
