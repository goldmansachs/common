#load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@bazel_rules//area:def.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        build_file_proto_mode = "disable",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        build_file_proto_mode = "disable",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_go_kit_log",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        build_file_proto_mode = "disable",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_kr_pretty",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/text",
        sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/matttproud/golang_protobuf_extensions/v2",
        sum = "h1:jWpvCLoY8Z/e3VKvlsiIGKtc+UG6U5vzxaoagmhXfyg=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        build_file_proto_mode = "disable",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        build_file_proto_mode = "disable",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:HzFfmkOzH5Q8L8G+kSJKUx5dtG87sewO+FoDDqP5Tbk=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:VQw1hfvPvk3Uv6Qf29VrPF32JB6rtbgI6cYPYQjL0Qw=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:jluTpSng7V9hY0O2R9DzzJHYb2xULk9VTR1V1R/k6Bo=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/testify",
        sum = "h1:+h33VjcLVPDHtOdpUCuF+7gSuG3yGIftsP1YvFihtJ8=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:6aKEtlUiwEpJzM001l0yFkpXmUVXaN8W+fbkb2AZNbg=",
        version = "v1.20.1",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "org_golang_google_appengine",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/appengine",
        sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
        version = "v1.6.7",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/protobuf",
        sum = "h1:pPC6BG5ex8PDFnkbrGU3EixyhKcQ2aDuBS36lqK/C7I=",
        version = "v1.32.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/crypto",
        sum = "h1:PGVlW0xEltQnzFZ55hkuX5+KLyrMYhHld1YHO4AKcdc=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mod",
        sum = "h1:LUYupSeNrTNCGzR/hVBk2NHZO4hXcVaW1k4Qx7rjPx8=",
        version = "v0.8.0",
    )
    go_repository(
        name = "org_golang_x_net",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/net",
        sum = "h1:aCL9BSgETF1k+blQaYUBx9hJ9LOGP3gAVemcZlf1Kpo=",
        version = "v0.20.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/oauth2",
        sum = "h1:aDkGMBSYxElaoP81NpoUoz2oo2R2wHdZpGToUxfyQrQ=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sync",
        sum = "h1:ftCYgMx6zT/asHUrPw8BLLscYtGznsLAnjq5RH9P66E=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sys",
        sum = "h1:xWw16ngr6ZMtmxDyKyIgsE93KNKz5HKmMa3b8ALHidU=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_term",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/term",
        sum = "h1:m+B6fahuftsE9qjo0VWp2FW0mB3MTJvR0BaMQrq0pmE=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_text",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/tools",
        sum = "h1:BOw41kyTf3PuCW1pVQf8+Cyg8pMlkYB1oo9iJ6D/lKM=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/xerrors",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        version = "v0.0.0-20191204190536-9bdfabe68543",
    )
