---
mapped_pages:
  - https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-dependencies.html
navigation_title: current
applies_to:
  deployment:
    eck: preview
---
% Generated documentation. Please do not edit.

# Third-party dependencies for Elastic Cloud on Kubernetes[k8s-dependencies]

This page lists the third-party dependencies used to build {{eck}} from the main branch.

## Direct dependencies [k8s-dependencies-direct]

| Name | Version | Licence |
| --- | --- | --- |
| [dario.cat/mergo](https://dario.cat/mergo) | v1.0.2 | BSD-3-Clause |
| [github.com/Masterminds/sprig/v3](https://github.com/Masterminds/sprig) | v3.3.0 | MIT |
| [github.com/blang/semver/v4](https://github.com/blang/semver) | v4.0.0 | MIT |
| [github.com/davecgh/go-spew](https://github.com/davecgh/go-spew) | v1.1.2-0.20180830191138-d8f796af33cc | ISC |
| [github.com/elastic/go-ucfg](https://github.com/elastic/go-ucfg) | v0.8.9-0.20250307075119-2a22403faaea | Apache-2.0 |
| [github.com/ghodss/yaml](https://github.com/ghodss/yaml) | v1.0.0 | MIT |
| [github.com/gkampitakis/go-snaps](https://github.com/gkampitakis/go-snaps) | v0.5.13 | MIT |
| [github.com/go-logr/logr](https://github.com/go-logr/logr) | v1.4.3 | Apache-2.0 |
| [github.com/go-test/deep](https://github.com/go-test/deep) | v1.1.1 | MIT |
| [github.com/gobuffalo/flect](https://github.com/gobuffalo/flect) | v1.0.3 | MIT |
| [github.com/google/go-cmp](https://github.com/google/go-cmp) | v0.7.0 | BSD-3-Clause |
| [github.com/google/go-containerregistry](https://github.com/google/go-containerregistry) | v0.20.6 | Apache-2.0 |
| [github.com/google/uuid](https://github.com/google/uuid) | v1.6.0 | BSD-3-Clause |
| [github.com/hashicorp/go-multierror](https://github.com/hashicorp/go-multierror) | v1.1.1 | MPL-2.0 |
| [github.com/hashicorp/golang-lru/v2](https://github.com/hashicorp/golang-lru) | v2.0.7 | MPL-2.0 |
| [github.com/hashicorp/vault/api](https://github.com/hashicorp/vault) | v1.20.0 | MPL-2.0 |
| [github.com/magiconair/properties](https://github.com/magiconair/properties) | v1.8.10 | BSD-2-Clause |
| [github.com/pkg/errors](https://github.com/pkg/errors) | v0.9.1 | BSD-2-Clause |
| [github.com/pmezard/go-difflib](https://github.com/pmezard/go-difflib) | v1.0.1-0.20181226105442-5d4384ee4fb2 | BSD-3-Clause |
| [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) | v1.22.0 | Apache-2.0 |
| [github.com/prometheus/common](https://github.com/prometheus/common) | v0.65.0 | Apache-2.0 |
| [github.com/sethvargo/go-password](https://github.com/sethvargo/go-password) | v0.3.1 | MIT |
| [github.com/spf13/cobra](https://github.com/spf13/cobra) | v1.9.1 | Apache-2.0 |
| [github.com/spf13/pflag](https://github.com/spf13/pflag) | v1.0.7 | BSD-3-Clause |
| [github.com/spf13/viper](https://github.com/spf13/viper) | v1.20.1 | MIT |
| [github.com/stretchr/testify](https://github.com/stretchr/testify) | v1.10.0 | MIT |
| [go.elastic.co/apm/module/apmelasticsearch/v2](https://go.elastic.co/apm/module/apmelasticsearch/v2) | v2.7.1 | Apache-2.0 |
| [go.elastic.co/apm/module/apmhttp/v2](https://go.elastic.co/apm/module/apmhttp/v2) | v2.7.1 | Apache-2.0 |
| [go.elastic.co/apm/module/apmzap/v2](https://go.elastic.co/apm/module/apmzap/v2) | v2.7.1 | Apache-2.0 |
| [go.elastic.co/apm/v2](https://go.elastic.co/apm/v2) | v2.7.1 | Apache-2.0 |
| [go.uber.org/automaxprocs](https://go.uber.org/automaxprocs) | v1.6.0 | MIT |
| [go.uber.org/zap](https://go.uber.org/zap) | v1.27.0 | MIT |
| [golang.org/x/crypto](https://golang.org/x/crypto) | v0.40.0 | BSD-3-Clause |
| [golang.org/x/exp](https://golang.org/x/exp) | v0.0.0-20240808152545-0cdaa3abc0fa | BSD-3-Clause |
| [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) | v3.0.1 | MIT |
| [k8s.io/api](https://github.com/kubernetes/api) | v0.33.3 | Apache-2.0 |
| [k8s.io/apimachinery](https://github.com/kubernetes/apimachinery) | v0.33.3 | Apache-2.0 |
| [k8s.io/client-go](https://github.com/kubernetes/client-go) | v0.33.3 | Apache-2.0 |
| [k8s.io/klog/v2](https://github.com/kubernetes/klog) | v2.130.1 | Apache-2.0 |
| [k8s.io/utils](https://github.com/kubernetes/utils) | v0.0.0-20241104100929-3ea5e8cea738 | Apache-2.0 |
| [sigs.k8s.io/controller-runtime](https://sigs.k8s.io/controller-runtime) | v0.21.0 | Apache-2.0 |
| [sigs.k8s.io/controller-tools](https://sigs.k8s.io/controller-tools) | v0.18.0 | Apache-2.0 |


##  Indirect dependencies [k8s-dependencies-indirect]

| Name | Version | Licence |
| --- | --- | --- |
| [cel.dev/expr](https://cel.dev/expr) | v0.19.1 | Apache-2.0 |
| [github.com/Masterminds/goutils](https://github.com/Masterminds/goutils) | v1.1.1 | Apache-2.0 |
| [github.com/Masterminds/semver/v3](https://github.com/Masterminds/semver) | v3.3.0 | MIT |
| [github.com/antlr4-go/antlr/v4](https://github.com/antlr4-go/antlr) | v4.13.0 | BSD-3-Clause |
| [github.com/armon/go-radix](https://github.com/armon/go-radix) | v1.0.0 | MIT |
| [github.com/armon/go-socks5](https://github.com/armon/go-socks5) | v0.0.0-20160902184237-e75332964ef5 | MIT |
| [github.com/beorn7/perks](https://github.com/beorn7/perks) | v1.0.1 | MIT |
| [github.com/cenkalti/backoff/v4](https://github.com/cenkalti/backoff) | v4.3.0 | MIT |
| [github.com/cespare/xxhash/v2](https://github.com/cespare/xxhash) | v2.3.0 | MIT |
| [github.com/containerd/stargz-snapshotter/estargz](https://github.com/containerd/stargz-snapshotter) | v0.16.3 | Apache-2.0 |
| [github.com/docker/cli](https://github.com/docker/cli) | v28.2.2+incompatible | Apache-2.0 |
| [github.com/docker/distribution](https://github.com/docker/distribution) | v2.8.3+incompatible | Apache-2.0 |
| [github.com/docker/docker-credential-helpers](https://github.com/docker/docker-credential-helpers) | v0.9.3 | MIT |
| [github.com/elastic/go-sysinfo](https://github.com/elastic/go-sysinfo) | v1.15.2 | Apache-2.0 |
| [github.com/elastic/go-windows](https://github.com/elastic/go-windows) | v1.0.2 | Apache-2.0 |
| [github.com/emicklei/go-restful/v3](https://github.com/emicklei/go-restful) | v3.12.1 | MIT |
| [github.com/evanphx/json-patch](https://github.com/evanphx/json-patch) | v5.6.0+incompatible | BSD-3-Clause |
| [github.com/evanphx/json-patch/v5](https://github.com/evanphx/json-patch) | v5.9.11 | BSD-3-Clause |
| [github.com/fatih/color](https://github.com/fatih/color) | v1.18.0 | MIT |
| [github.com/felixge/httpsnoop](https://github.com/felixge/httpsnoop) | v1.0.4 | MIT |
| [github.com/frankban/quicktest](https://github.com/frankban/quicktest) | v1.14.6 | MIT |
| [github.com/fsnotify/fsnotify](https://github.com/fsnotify/fsnotify) | v1.8.0 | BSD-3-Clause |
| [github.com/fxamacker/cbor/v2](https://github.com/fxamacker/cbor) | v2.7.0 | MIT |
| [github.com/gkampitakis/ciinfo](https://github.com/gkampitakis/ciinfo) | v0.3.2 | MIT |
| [github.com/gkampitakis/go-diff](https://github.com/gkampitakis/go-diff) | v1.3.2 | MIT |
| [github.com/go-jose/go-jose/v4](https://github.com/go-jose/go-jose) | v4.0.5 | Apache-2.0 |
| [github.com/go-logr/stdr](https://github.com/go-logr/stdr) | v1.2.2 | Apache-2.0 |
| [github.com/go-logr/zapr](https://github.com/go-logr/zapr) | v1.3.0 | Apache-2.0 |
| [github.com/go-openapi/jsonpointer](https://github.com/go-openapi/jsonpointer) | v0.21.0 | Apache-2.0 |
| [github.com/go-openapi/jsonreference](https://github.com/go-openapi/jsonreference) | v0.21.0 | Apache-2.0 |
| [github.com/go-openapi/swag](https://github.com/go-openapi/swag) | v0.23.0 | Apache-2.0 |
| [github.com/go-task/slim-sprig/v3](https://github.com/go-task/slim-sprig) | v3.0.0 | MIT |
| [github.com/go-viper/mapstructure/v2](https://github.com/go-viper/mapstructure) | v2.3.0 | MIT |
| [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml) | v1.18.0 | MIT |
| [github.com/gogo/protobuf](https://github.com/gogo/protobuf) | v1.3.2 | BSD-3-Clause |
| [github.com/golang/protobuf](https://github.com/golang/protobuf) | v1.5.4 | BSD-3-Clause |
| [github.com/google/btree](https://github.com/google/btree) | v1.1.3 | Apache-2.0 |
| [github.com/google/cel-go](https://github.com/google/cel-go) | v0.23.2 | Apache-2.0 |
| [github.com/google/gnostic-models](https://github.com/google/gnostic-models) | v0.6.9 | Apache-2.0 |
| [github.com/google/gofuzz](https://github.com/google/gofuzz) | v1.2.0 | Apache-2.0 |
| [github.com/google/pprof](https://github.com/google/pprof) | v0.0.0-20241029153458-d1b30febd7db | Apache-2.0 |
| [github.com/gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.4-0.20250319132907-e064f32e3674 | BSD-2-Clause |
| [github.com/grpc-ecosystem/grpc-gateway/v2](https://github.com/grpc-ecosystem/grpc-gateway) | v2.24.0 | BSD-3-Clause |
| [github.com/hashicorp/errwrap](https://github.com/hashicorp/errwrap) | v1.1.0 | MPL-2.0 |
| [github.com/hashicorp/go-cleanhttp](https://github.com/hashicorp/go-cleanhttp) | v0.5.2 | MPL-2.0 |
| [github.com/hashicorp/go-hclog](https://github.com/hashicorp/go-hclog) | v1.6.3 | MIT |
| [github.com/hashicorp/go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) | v0.7.7 | MPL-2.0 |
| [github.com/hashicorp/go-rootcerts](https://github.com/hashicorp/go-rootcerts) | v1.0.2 | MPL-2.0 |
| [github.com/hashicorp/go-secure-stdlib/parseutil](https://github.com/hashicorp/go-secure-stdlib) | v0.1.6 | MPL-2.0 |
| [github.com/hashicorp/go-secure-stdlib/strutil](https://github.com/hashicorp/go-secure-stdlib) | v0.1.2 | MPL-2.0 |
| [github.com/hashicorp/go-sockaddr](https://github.com/hashicorp/go-sockaddr) | v1.0.2 | MPL-2.0 |
| [github.com/hashicorp/hcl](https://github.com/hashicorp/hcl) | v1.0.1-vault-7 | MPL-2.0 |
| [github.com/huandu/xstrings](https://github.com/huandu/xstrings) | v1.5.0 | MIT |
| [github.com/inconshreveable/mousetrap](https://github.com/inconshreveable/mousetrap) | v1.1.0 | Apache-2.0 |
| [github.com/josharian/intern](https://github.com/josharian/intern) | v1.0.0 | MIT |
| [github.com/json-iterator/go](https://github.com/json-iterator/go) | v1.1.12 | MIT |
| [github.com/klauspost/compress](https://github.com/klauspost/compress) | v1.18.0 | Apache-2.0 |
| [github.com/kr/pretty](https://github.com/kr/pretty) | v0.3.1 | MIT |
| [github.com/kr/text](https://github.com/kr/text) | v0.2.0 | MIT |
| [github.com/kylelemons/godebug](https://github.com/kylelemons/godebug) | v1.1.0 | Apache-2.0 |
| [github.com/mailru/easyjson](https://github.com/mailru/easyjson) | v0.7.7 | MIT |
| [github.com/maruel/natural](https://github.com/maruel/natural) | v1.1.1 | Apache-2.0 |
| [github.com/mattn/go-colorable](https://github.com/mattn/go-colorable) | v0.1.13 | MIT |
| [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) | v0.0.20 | MIT |
| [github.com/mitchellh/copystructure](https://github.com/mitchellh/copystructure) | v1.2.0 | MIT |
| [github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir) | v1.1.0 | MIT |
| [github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) | v1.5.0 | MIT |
| [github.com/mitchellh/reflectwalk](https://github.com/mitchellh/reflectwalk) | v1.0.2 | MIT |
| [github.com/moby/spdystream](https://github.com/moby/spdystream) | v0.5.0 | Apache-2.0 |
| [github.com/modern-go/concurrent](https://github.com/modern-go/concurrent) | v0.0.0-20180306012644-bacd9c7ef1dd | Apache-2.0 |
| [github.com/modern-go/reflect2](https://github.com/modern-go/reflect2) | v1.0.2 | Apache-2.0 |
| [github.com/munnerz/goautoneg](https://github.com/munnerz/goautoneg) | v0.0.0-20191010083416-a7dc8b61c822 | BSD-3-Clause |
| [github.com/mxk/go-flowrate](https://github.com/mxk/go-flowrate) | v0.0.0-20140419014527-cca7078d478f | BSD-3-Clause |
| [github.com/nxadm/tail](https://github.com/nxadm/tail) | v1.4.8 | MIT |
| [github.com/onsi/ginkgo](https://github.com/onsi/ginkgo) | v1.16.5 | MIT |
| [github.com/onsi/ginkgo/v2](https://github.com/onsi/ginkgo) | v2.22.0 | MIT |
| [github.com/onsi/gomega](https://github.com/onsi/gomega) | v1.37.0 | MIT |
| [github.com/opencontainers/go-digest](https://github.com/opencontainers/go-digest) | v1.0.0 | Apache-2.0 |
| [github.com/opencontainers/image-spec](https://github.com/opencontainers/image-spec) | v1.1.1 | Apache-2.0 |
| [github.com/pelletier/go-toml/v2](https://github.com/pelletier/go-toml) | v2.2.3 | MIT |
| [github.com/prashantv/gostub](https://github.com/prashantv/gostub) | v1.1.0 | MIT |
| [github.com/prometheus/client_model](https://github.com/prometheus/client_model) | v0.6.2 | Apache-2.0 |
| [github.com/prometheus/procfs](https://github.com/prometheus/procfs) | v0.16.0 | Apache-2.0 |
| [github.com/rogpeppe/go-internal](https://github.com/rogpeppe/go-internal) | v1.13.1 | BSD-3-Clause |
| [github.com/ryanuber/go-glob](https://github.com/ryanuber/go-glob) | v1.0.0 | MIT |
| [github.com/sagikazarmark/locafero](https://github.com/sagikazarmark/locafero) | v0.7.0 | MIT |
| [github.com/shopspring/decimal](https://github.com/shopspring/decimal) | v1.4.0 | MIT |
| [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) | v1.9.3 | MIT |
| [github.com/sourcegraph/conc](https://github.com/sourcegraph/conc) | v0.3.0 | MIT |
| [github.com/spf13/afero](https://github.com/spf13/afero) | v1.12.0 | Apache-2.0 |
| [github.com/spf13/cast](https://github.com/spf13/cast) | v1.7.1 | MIT |
| [github.com/stoewer/go-strcase](https://github.com/stoewer/go-strcase) | v1.3.0 | MIT |
| [github.com/stretchr/objx](https://github.com/stretchr/objx) | v0.5.2 | MIT |
| [github.com/subosito/gotenv](https://github.com/subosito/gotenv) | v1.6.0 | MIT |
| [github.com/tidwall/gjson](https://github.com/tidwall/gjson) | v1.18.0 | MIT |
| [github.com/tidwall/match](https://github.com/tidwall/match) | v1.1.1 | MIT |
| [github.com/tidwall/pretty](https://github.com/tidwall/pretty) | v1.2.1 | MIT |
| [github.com/tidwall/sjson](https://github.com/tidwall/sjson) | v1.2.5 | MIT |
| [github.com/vbatts/tar-split](https://github.com/vbatts/tar-split) | v0.12.1 | BSD-3-Clause |
| [github.com/x448/float16](https://github.com/x448/float16) | v0.8.4 | MIT |
| [go.elastic.co/fastjson](https://go.elastic.co/fastjson) | v1.5.1 | MIT |
| [go.opentelemetry.io/auto/sdk](https://go.opentelemetry.io/auto/sdk) | v1.1.0 | Apache-2.0 |
| [go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp](https://go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp) | v0.61.0 | Apache-2.0 |
| [go.opentelemetry.io/otel](https://go.opentelemetry.io/otel) | v1.36.0 | Apache-2.0 |
| [go.opentelemetry.io/otel/exporters/otlp/otlptrace](https://go.opentelemetry.io/otel/exporters/otlp/otlptrace) | v1.33.0 | Apache-2.0 |
| [go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc](https://go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc) | v1.33.0 | Apache-2.0 |
| [go.opentelemetry.io/otel/metric](https://go.opentelemetry.io/otel/metric) | v1.36.0 | Apache-2.0 |
| [go.opentelemetry.io/otel/sdk](https://go.opentelemetry.io/otel/sdk) | v1.36.0 | Apache-2.0 |
| [go.opentelemetry.io/otel/sdk/metric](https://go.opentelemetry.io/otel/sdk/metric) | v1.36.0 | Apache-2.0 |
| [go.opentelemetry.io/otel/trace](https://go.opentelemetry.io/otel/trace) | v1.36.0 | Apache-2.0 |
| [go.opentelemetry.io/proto/otlp](https://go.opentelemetry.io/proto/otlp) | v1.4.0 | Apache-2.0 |
| [go.uber.org/goleak](https://go.uber.org/goleak) | v1.3.0 | MIT |
| [go.uber.org/multierr](https://go.uber.org/multierr) | v1.11.0 | MIT |
| [golang.org/x/mod](https://golang.org/x/mod) | v0.25.0 | BSD-3-Clause |
| [golang.org/x/net](https://golang.org/x/net) | v0.41.0 | BSD-3-Clause |
| [golang.org/x/oauth2](https://golang.org/x/oauth2) | v0.30.0 | BSD-3-Clause |
| [golang.org/x/sync](https://golang.org/x/sync) | v0.16.0 | BSD-3-Clause |
| [golang.org/x/sys](https://golang.org/x/sys) | v0.34.0 | BSD-3-Clause |
| [golang.org/x/term](https://golang.org/x/term) | v0.33.0 | BSD-3-Clause |
| [golang.org/x/text](https://golang.org/x/text) | v0.27.0 | BSD-3-Clause |
| [golang.org/x/time](https://golang.org/x/time) | v0.9.0 | BSD-3-Clause |
| [golang.org/x/tools](https://golang.org/x/tools) | v0.34.0 | BSD-3-Clause |
| [gomodules.xyz/jsonpatch/v2](https://gomodules.xyz/jsonpatch/v2) | v2.4.0 | Apache-2.0 |
| [google.golang.org/genproto/googleapis/api](https://google.golang.org/genproto/googleapis/api) | v0.0.0-20241209162323-e6fa225c2576 | Apache-2.0 |
| [google.golang.org/genproto/googleapis/rpc](https://google.golang.org/genproto/googleapis/rpc) | v0.0.0-20241223144023-3abc09e42ca8 | Apache-2.0 |
| [google.golang.org/grpc](https://google.golang.org/grpc) | v1.68.1 | Apache-2.0 |
| [google.golang.org/protobuf](https://google.golang.org/protobuf) | v1.36.6 | BSD-3-Clause |
| [gopkg.in/check.v1](https://gopkg.in/check.v1) | v1.0.0-20201130134442-10cb98267c6c | BSD-2-Clause |
| [gopkg.in/evanphx/json-patch.v4](https://gopkg.in/evanphx/json-patch.v4) | v4.12.0 | BSD-3-Clause |
| [gopkg.in/inf.v0](https://gopkg.in/inf.v0) | v0.9.1 | BSD-3-Clause |
| [gopkg.in/tomb.v1](https://gopkg.in/tomb.v1) | v1.0.0-20141024135613-dd632973f1e7 | BSD-3-Clause |
| [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) | v2.4.0 | Apache-2.0 |
| [gotest.tools/v3](https://gotest.tools/v3) | v3.0.3 | Apache-2.0 |
| [howett.net/plist](https://gitlab.howett.net/go/plist) | v1.0.1 | BSD-2-Clause |
| [k8s.io/apiextensions-apiserver](https://github.com/kubernetes/apiextensions-apiserver) | v0.33.0 | Apache-2.0 |
| [k8s.io/apiserver](https://github.com/kubernetes/apiserver) | v0.33.0 | Apache-2.0 |
| [k8s.io/code-generator](https://github.com/kubernetes/code-generator) | v0.33.0 | Apache-2.0 |
| [k8s.io/component-base](https://github.com/kubernetes/component-base) | v0.33.0 | Apache-2.0 |
| [k8s.io/gengo/v2](https://github.com/kubernetes/gengo) | v2.0.0-20250207200755-1244d31929d7 | Apache-2.0 |
| [k8s.io/kube-openapi](https://github.com/kubernetes/kube-openapi) | v0.0.0-20250318190949-c8a335a9a2ff | Apache-2.0 |
| [sigs.k8s.io/apiserver-network-proxy/konnectivity-client](https://sigs.k8s.io/apiserver-network-proxy/konnectivity-client) | v0.31.2 | Apache-2.0 |
| [sigs.k8s.io/json](https://sigs.k8s.io/json) | v0.0.0-20241010143419-9aa6b5e7a4b3 | Apache-2.0 |
| [sigs.k8s.io/randfill](https://sigs.k8s.io/randfill) | v1.0.0 | Apache-2.0 |
| [sigs.k8s.io/structured-merge-diff/v4](https://sigs.k8s.io/structured-merge-diff/v4) | v4.6.0 | Apache-2.0 |
| [sigs.k8s.io/yaml](https://sigs.k8s.io/yaml) | v1.4.0 | Apache-2.0 |

