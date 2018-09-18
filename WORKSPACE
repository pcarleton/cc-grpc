load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.1/rules_go-0.15.1.tar.gz"],
    sha256 = "5f3b0304cdf0c505ec9e5b3c4fc4a87b5ca21b13d8ecc780c97df3d1809b9ce6",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# Container dependencies
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "29d109605e0d6f9c892584f07275b8c9260803bf0c6fcb7de2623b2bedc910bd",
    strip_prefix = "rules_docker-0.5.1",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.5.1.tar.gz"],
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)

# This is NOT needed when going through the language lang_image
# "repositories" function(s).
container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

# Imported via Gazelle

go_repository(
    name = "org_golang_x_oauth2",
    commit = "d2e6202438beef2727060aa7cabdd924d92ebfd9",
    importpath = "golang.org/x/oauth2",
)

go_repository(
    name = "com_github_nmrshll_oauth2_noserver",
    commit = "16b622b98a45a4bb67e24e61e76892580026c9c9",
    importpath = "github.com/nmrshll/oauth2-noserver",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "8991bc29aa16c548c550c7ff78260e27b9ab7c73",
    importpath = "github.com/davecgh/go-spew",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    commit = "06ea1031745cb8b3dab3f6a236daf2b0aa468b7e",
    importpath = "github.com/dgrijalva/jwt-go",
)

go_repository(
    name = "com_github_fatih_color",
    commit = "5b77d2a35fb0ede96d138fc9a99f5c9b6aef11b4",
    importpath = "github.com/fatih/color",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "aa810b61a9c79d51363740d207bb46cf8e620ed5",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_gorilla_context",
    commit = "08b5f424b9271eedf6f9f0ce86cb9396ed337a42",
    importpath = "github.com/gorilla/context",
)

go_repository(
    name = "com_github_gorilla_securecookie",
    commit = "e59506cc896acb7f7bf732d4fdf5e25f7ccd8983",
    importpath = "github.com/gorilla/securecookie",
)

go_repository(
    name = "com_github_gorilla_sessions",
    commit = "03b6f63cc43ef9c7240a635a5e22b13180e822b8",
    importpath = "github.com/gorilla/sessions",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "167de6bfdfba052fa6b2d3664c8f5272e23c9072",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_nmrshll_rndm_go",
    commit = "8da3024e53de85cef9aac652d4144717f37ada86",
    importpath = "github.com/nmrshll/rndm-go",
)

go_repository(
    name = "com_github_palantir_stacktrace",
    commit = "78658fd2d1772b755720ed8c44367d11ee5380d6",
    importpath = "github.com/palantir/stacktrace",
)

go_repository(
    name = "com_github_skratchdot_open_golang",
    commit = "75fb7ed4208cf72d323d7d02fd1a5964a7a9073c",
    importpath = "github.com/skratchdot/open-golang",
)

go_repository(
    name = "org_golang_google_appengine",
    commit = "b1f26356af11148e710935ed1ac8a7f5702c7612",
    importpath = "google.golang.org/appengine",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "11092d34479b07829b72e10713b159248caf5dad",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "32fb0ac620c32ba40a4626ddf94d90d12cce3455",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_net",
    commit = "8a410e7b638dca158bf9e766925842f6651ff828",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "fa5fdf94c78965f1aa8423f0cc50b8b8d728b05a",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_googleapis_gax_go",
    commit = "317e0006254c44a0ac427cc52a0e083ff0b9622f",
    importpath = "github.com/googleapis/gax-go",
)

go_repository(
    name = "com_google_cloud_go",
    commit = "c728a003b238b26cef9ab6753a5dc424b331c3ad",
    importpath = "cloud.google.com/go",
)

go_repository(
    name = "io_opencensus_go",
    commit = "79993219becaa7e29e3b60cb67f5b8e82dee11d6",
    importpath = "go.opencensus.io",
)

go_repository(
    name = "io_opencensus_go_contrib_exporter_stackdriver",
    commit = "2b93072101d466aa4120b3c23c2e1b08af01541c",
    importpath = "contrib.go.opencensus.io/exporter/stackdriver",
)

go_repository(
    name = "org_golang_google_api",
    commit = "19ff8768a5c0b8e46ea281065664787eefc24121",
    importpath = "google.golang.org/api",
)

go_repository(
    name = "com_github_pcarleton_sheets",
    commit = "16c0ae2058d3f80cd4d1046a429ce64d7ad8001a",
    importpath = "github.com/pcarleton/sheets",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    commit = "c2828203cd70a50dcccfb2761f8b1f8ceef9a8e9",
    importpath = "github.com/fsnotify/fsnotify",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    commit = "8cb6e5b959231cc1119e43259c4a608f9c51a241",
    importpath = "github.com/hashicorp/hcl",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    commit = "76626ae9c91c4f2a10f34cad8ce83ea42c93bb75",
    importpath = "github.com/inconshreveable/mousetrap",
)

go_repository(
    name = "com_github_magiconair_properties",
    commit = "c2353362d570a7bfa228149c62842019201cfb71",
    importpath = "github.com/magiconair/properties",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    commit = "ae18d6b8b3205b561c79e8e5f69bff09736185f4",
    importpath = "github.com/mitchellh/go-homedir",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "fa473d140ef3c6adf42d6b391fe76707f1f243c8",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    commit = "c01d1270ff3e442a8a57cddc1c92dc1138598194",
    importpath = "github.com/pelletier/go-toml",
)

go_repository(
    name = "com_github_spf13_afero",
    commit = "d40851caa0d747393da1ffb28f7f9d8b4eeffebd",
    importpath = "github.com/spf13/afero",
)

go_repository(
    name = "com_github_spf13_cast",
    commit = "8965335b8c7107321228e3e3702cab9832751bac",
    importpath = "github.com/spf13/cast",
)

go_repository(
    name = "com_github_spf13_cobra",
    commit = "ef82de70bb3f60c65fb8eebacbb2d122ef517385",
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    commit = "4a4406e478ca629068e7768fc33f3f044173c0a6",
    importpath = "github.com/spf13/jwalterweatherman",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "9a97c102cda95a86cec2345a6f09f55a939babf5",
    importpath = "github.com/spf13/pflag",
)

go_repository(
    name = "com_github_spf13_viper",
    commit = "8fb642006536c8d3760c99d4fa2389f5e2205631",
    importpath = "github.com/spf13/viper",
)
