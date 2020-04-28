# My Own Caddy
## Version
**v2.0.0-rc.3** `h1:z2H/QnaRscip6aZJxwTbghu3zhC88Vo8l/K57WUce4Q=`
## Modules
- [admin.api.load](https://caddyserver.com/docs/modules/admin.api.load)
- [caddy.adapters.caddyfile](https://caddyserver.com/docs/modules/caddy.adapters.caddyfile)
- [caddy.listeners.tls](https://caddyserver.com/docs/modules/caddy.listeners.tls)
- [caddy.logging.encoders.console](https://caddyserver.com/docs/modules/caddy.logging.encoders.console)
- [caddy.logging.encoders.filter](https://caddyserver.com/docs/modules/caddy.logging.encoders.filter)
- [caddy.logging.encoders.filter.delete](https://caddyserver.com/docs/modules/caddy.logging.encoders.filter.delete)
- [caddy.logging.encoders.filter.ip_mask](https://caddyserver.com/docs/modules/caddy.logging.encoders.filter.ip_mask)
- [caddy.logging.encoders.json](https://caddyserver.com/docs/modules/caddy.logging.encoders.json)
- [caddy.logging.encoders.logfmt](https://caddyserver.com/docs/modules/caddy.logging.encoders.logfmt)
- [caddy.logging.encoders.single_field](https://caddyserver.com/docs/modules/caddy.logging.encoders.single_field)
- [caddy.logging.writers.discard](https://caddyserver.com/docs/modules/caddy.logging.writers.discard)
- [caddy.logging.writers.file](https://caddyserver.com/docs/modules/caddy.logging.writers.file)
- [caddy.logging.writers.net](https://caddyserver.com/docs/modules/caddy.logging.writers.net)
- [caddy.logging.writers.stderr](https://caddyserver.com/docs/modules/caddy.logging.writers.stderr)
- [caddy.logging.writers.stdout](https://caddyserver.com/docs/modules/caddy.logging.writers.stdout)
- [caddy.storage.file_system](https://caddyserver.com/docs/modules/caddy.storage.file_system)
- [http](https://caddyserver.com/docs/modules/http)
- [http.authentication.hashes.bcrypt](https://caddyserver.com/docs/modules/http.authentication.hashes.bcrypt)
- [http.authentication.hashes.scrypt](https://caddyserver.com/docs/modules/http.authentication.hashes.scrypt)
- [http.authentication.providers.http_basic](https://caddyserver.com/docs/modules/http.authentication.providers.http_basic)
- [http.encoders.gzip](https://caddyserver.com/docs/modules/http.encoders.gzip)
- [http.encoders.zstd](https://caddyserver.com/docs/modules/http.encoders.zstd)
- [http.handlers.authentication](https://caddyserver.com/docs/modules/http.handlers.authentication)
- [http.handlers.encode](https://caddyserver.com/docs/modules/http.handlers.encode)
- [http.handlers.error](https://caddyserver.com/docs/modules/http.handlers.error)
- [http.handlers.file_server](https://caddyserver.com/docs/modules/http.handlers.file_server)
- [http.handlers.headers](https://caddyserver.com/docs/modules/http.handlers.headers)
- [http.handlers.request_body](https://caddyserver.com/docs/modules/http.handlers.request_body)
- [http.handlers.reverse_proxy](https://caddyserver.com/docs/modules/http.handlers.reverse_proxy)
- [http.handlers.rewrite](https://caddyserver.com/docs/modules/http.handlers.rewrite)
- [http.handlers.static_response](https://caddyserver.com/docs/modules/http.handlers.static_response)
- [http.handlers.subroute](https://caddyserver.com/docs/modules/http.handlers.subroute)
- [http.handlers.templates](https://caddyserver.com/docs/modules/http.handlers.templates)
- [http.handlers.vars](https://caddyserver.com/docs/modules/http.handlers.vars)
- [http.matchers.expression](https://caddyserver.com/docs/modules/http.matchers.expression)
- [http.matchers.file](https://caddyserver.com/docs/modules/http.matchers.file)
- [http.matchers.header](https://caddyserver.com/docs/modules/http.matchers.header)
- [http.matchers.header_regexp](https://caddyserver.com/docs/modules/http.matchers.header_regexp)
- [http.matchers.host](https://caddyserver.com/docs/modules/http.matchers.host)
- [http.matchers.method](https://caddyserver.com/docs/modules/http.matchers.method)
- [http.matchers.not](https://caddyserver.com/docs/modules/http.matchers.not)
- [http.matchers.path](https://caddyserver.com/docs/modules/http.matchers.path)
- [http.matchers.path_regexp](https://caddyserver.com/docs/modules/http.matchers.path_regexp)
- [http.matchers.protocol](https://caddyserver.com/docs/modules/http.matchers.protocol)
- [http.matchers.query](https://caddyserver.com/docs/modules/http.matchers.query)
- [http.matchers.remote_ip](https://caddyserver.com/docs/modules/http.matchers.remote_ip)
- [http.matchers.vars](https://caddyserver.com/docs/modules/http.matchers.vars)
- [http.matchers.vars_regexp](https://caddyserver.com/docs/modules/http.matchers.vars_regexp)
- [http.reverse_proxy.circuit_breakers.internal](https://caddyserver.com/docs/modules/http.reverse_proxy.circuit_breakers.internal)
- [http.reverse_proxy.selection_policies.first](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.first)
- [http.reverse_proxy.selection_policies.header](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.header)
- [http.reverse_proxy.selection_policies.ip_hash](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.ip_hash)
- [http.reverse_proxy.selection_policies.least_conn](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.least_conn)
- [http.reverse_proxy.selection_policies.random](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.random)
- [http.reverse_proxy.selection_policies.random_choose](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.random_choose)
- [http.reverse_proxy.selection_policies.round_robin](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.round_robin)
- [http.reverse_proxy.selection_policies.uri_hash](https://caddyserver.com/docs/modules/http.reverse_proxy.selection_policies.uri_hash)
- [http.reverse_proxy.transport.fastcgi](https://caddyserver.com/docs/modules/http.reverse_proxy.transport.fastcgi)
- [http.reverse_proxy.transport.http](https://caddyserver.com/docs/modules/http.reverse_proxy.transport.http)
- [pki](https://caddyserver.com/docs/modules/pki)
- [tls](https://caddyserver.com/docs/modules/tls)
- [tls.certificates.automate](https://caddyserver.com/docs/modules/tls.certificates.automate)
- [tls.certificates.load_files](https://caddyserver.com/docs/modules/tls.certificates.load_files)
- [tls.certificates.load_folders](https://caddyserver.com/docs/modules/tls.certificates.load_folders)
- [tls.certificates.load_pem](https://caddyserver.com/docs/modules/tls.certificates.load_pem)
- [tls.handshake_match.sni](https://caddyserver.com/docs/modules/tls.handshake_match.sni)
- [tls.issuance.acme](https://caddyserver.com/docs/modules/tls.issuance.acme)
- [tls.issuance.internal](https://caddyserver.com/docs/modules/tls.issuance.internal)
- [tls.stek.distributed](https://caddyserver.com/docs/modules/tls.stek.distributed)
- [tls.stek.standard](https://caddyserver.com/docs/modules/tls.stek.standard)
## Info
```text
path: caddy
main: caddy (devel) 
dependencies:
cloud.google.com/go v0.54.0 h1:3ithwDMr7/3vpAMXiH+ZQnYbuIsh+OPhUPMFC9enmn0=
github.com/AndreasBriese/bbloom v0.0.0-20190306092124-e2d15f34fcf9 h1:HD8gA2tkByhMAwYaFAX9w2l7vxvBQ5NMoxDrkhqhtn4=
github.com/Masterminds/goutils v1.1.0 h1:zukEsf/1JZwCMgHiK3GZftabmxiCw4apj3a28RPBiVg=
github.com/Masterminds/semver/v3 v3.0.3 h1:znjIyLfpXEDQjOIEWh+ehwpTU14UzUPub3c3sm36u14=
github.com/Masterminds/sprig/v3 v3.0.2 h1:wz22D0CiSctrliXiI9ZO3HoNApweeRGftyDN+BQa3B8=
github.com/alecthomas/chroma v0.7.2-0.20200305040604-4f3623dce67a h1:3v1NrYWWqp2S72e4HLgxKt83B3l0lnORDholH/ihoMM=
github.com/antlr/antlr4 v0.0.0-20190819145818-b43a4c3a8015 h1:StuiJFxQUsxSCzcby6NFZRdEhPkXD5vxN7TZ4MD6T84=
github.com/caddyserver/caddy/v2 v2.0.0-rc.3 h1:z2H/QnaRscip6aZJxwTbghu3zhC88Vo8l/K57WUce4Q=
github.com/caddyserver/certmagic v0.10.11 h1:70wN1rNTc1EVwWA2i1IwYg8jJcVFsE7cvy2UT0uVA5k=
github.com/cenkalti/backoff/v4 v4.0.2 h1:JIufpQLbh4DkbQoii76ItQIUFzevQSqOLZca4eamEDs=
github.com/cheekybits/genny v1.0.0 h1:uGGa4nei+j20rOSeDeP5Of12XVm7TGUd4dJA9RDitfE=
github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=
github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=
github.com/cpuguy83/go-md2man/v2 v2.0.0 h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=
github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 h1:y5HC9v93H5EPKqaS1UYVg1uYah5Xf51mBfIoWehClUQ=
github.com/dgraph-io/badger v1.5.3 h1:5oWIuRvwn93cie+OSt1zSnkaIQ1JFQM8bGlIv6O6Sts=
github.com/dgryski/go-farm v0.0.0-20190423205320-6a90982ecee2 h1:tdlZCpZ/P9DhczCTSixgIKmwPv6+wP5DGjqLYw5SUiA=
github.com/dlclark/regexp2 v1.2.0 h1:8sAhBGEM0dRWogWqWyQeIJnxjWO6oIjl8FKqREDsGfk=
github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac h1:opbrjaN/L8gg6Xh5D04Tem+8xVcz6ajZlGCs49mQgyg=
github.com/francoispqt/gojay v1.2.13 h1:d2m3sFjloqoIUQU3TsHBgj6qg/BVGlTBeHDUmyJnXKk=
github.com/go-acme/lego/v3 v3.5.0 h1:/0+NJQK+hNwRznhCi+19lbEa4xufhe7wJZOVd5j486s=
github.com/go-sql-driver/mysql v1.4.1 h1:g24URVg0OFbNUTx9qqY1IRZ9D9z3iPyi5zKhQZpNwpA=
github.com/gogo/protobuf v1.3.1 h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=
github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e h1:1r7pUrabqp18hOBcwBwiTsbnFeTZHV9eER/QT5JVZxY=
github.com/golang/protobuf v1.3.4 h1:87PNWwrRvUSnqS4dlcBU/ftvOIBep4sYuBLlh6rX2wk=
github.com/google/cel-go v0.4.1 h1:2kqc5arTucvtLJzXVUbmiUh7n2xjizwZijPrpEsagAE=
github.com/google/go-cmp v0.4.0 h1:xsAVV57WRhGj6kEIi8ReJzQlHHqcBYCElAvkovg3B/4=
github.com/google/uuid v1.1.1 h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=
github.com/googleapis/gax-go/v2 v2.0.5 h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=
github.com/huandu/xstrings v1.2.0 h1:yPeWdRnmynF7p+lLYz0H2tthW9lqhMJrQV/U7yy4wX0=
github.com/imdario/mergo v0.3.9 h1:UauaLniWCFHWd+Jp9oCEkTBj8VO/9DKg3PV3VCNMDIg=
github.com/jsternberg/zap-logfmt v1.2.0 h1:1v+PK4/B48cy8cfQbxL4FmmNZrjnIMr2BsnyEmXqv2o=
github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a h1:FaWFmfWdAUKbSCtOU2QjDaorUexogfaMgbipgYATUMU=
github.com/klauspost/compress v1.10.4 h1:jFzIFaf586tquEB5EhzQG0HwGNSlgAJpG53G6Ss11wc=
github.com/klauspost/cpuid v1.2.3 h1:CCtW0xUnWGVINKvE/WWOYKdsPV6mawAtvQuSl8guwQs=
github.com/lucas-clemente/quic-go v0.15.3 h1:i6n4Jr7673z9TlurAjc87+GlE/BN10++r9XZIPS9j6I=
github.com/lunixbochs/vtclean v1.0.0 h1:xu2sLAri4lGiovBDQKxl5mrXyESr3gUr5m5SM5+LVb8=
github.com/mailgun/timetools v0.0.0-20141028012446-7e6055773c51 h1:Kg/NPZLLC3aAFr1YToMs98dbCdhootQ1hZIvZU28hAQ=
github.com/manifoldco/promptui v0.7.0 h1:3l11YT8tm9MnwGFQ4kETwkzpAwY2Jt9lCrumCUW4+z4=
github.com/marten-seemann/qpack v0.1.0 h1:/0M7lkda/6mus9B8u34Asqm8ZhHAAt9Ho0vniNuVSVg=
github.com/marten-seemann/qtls v0.9.0 h1:8Zguhc72eS+DH5EAb0BpAPIy3HDXYcihQi4xoDZOnjQ=
github.com/mattn/go-colorable v0.1.4 h1:snbPLB8fVfU9iwbbo30TPtbLRzwWu6aJS6Xh4eaaviA=
github.com/mattn/go-isatty v0.0.11 h1:FxPOTFNqGkuDUGi3H/qkUbQO4ZiBa2brKq5r0l8TGeM=
github.com/miekg/dns v1.1.29 h1:xHBEhR+t5RzcFJjBLJlax2daXOrTYtr9z4WdKEfWFzg=
github.com/mitchellh/copystructure v1.0.0 h1:Laisrj+bAB6b/yJwB5Bt3ITZhGJdqmxquMKeZ+mmkFQ=
github.com/mitchellh/reflectwalk v1.0.0 h1:9D+8oIskB4VJBN5SFlmc27fSlIBZaov1Wpk/IfikLNY=
github.com/naoina/go-stringutil v0.1.0 h1:rCUeRUHjBjGTSHl0VC00jUPLz8/F9dDzYI70Hzifhks=
github.com/naoina/toml v0.1.1 h1:PT/lllxVVN0gzzSqSlHEmP8MJB4MY2U7STGxiouV4X8=
github.com/pkg/errors v0.8.1 h1:iURUrRGxPUNPdy5/HRSm+Yj6okJ6UtLINN0Q9M4+h3I=
github.com/russross/blackfriday/v2 v2.0.1 h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=
github.com/samfoo/ansi v0.0.0-20160124022901-b6bd2ded7189 h1:CmSpbxmewNQbzqztaY0bke1qzHhyNyC29wYgh17Gxfo=
github.com/shurcooL/sanitized_anchor_name v1.0.0 h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=
github.com/smallstep/certificates v0.14.1 h1:dkNs0woJHUEOk7KYMsZ2eTViwPE+TkBGyseD5ctvhQo=
github.com/smallstep/cli v0.14.1 h1:iD068BTHnPYS8/ZMtaDju2x4M9c4aLklPOgwM69lCBo=
github.com/smallstep/nosql v0.2.0 h1:IscXK9m9hRyl5GoYgn+Iml//5Bpad3LyIj6R0dZosKM=
github.com/smallstep/truststore v0.9.5 h1:KQ6bFXUadu3PG57sFSIBsu2pb/35NqO+MyS2Pvi62bA=
github.com/spf13/cast v1.3.1 h1:nFm6S0SMdyzrzcmThSipiEubIDy8WEXKNZ0UOgiRpng=
github.com/urfave/cli v1.22.2 h1:gsqYFH8bb9ekPA12kRo0hfjngWQjkJPlN9R0N78BoUo=
github.com/vulcand/oxy v1.1.0 h1:DbBijGo1+6cFqR9jarkMxasdj0lgWwrrFtue6ijek4Q=
github.com/yuin/goldmark v1.1.27 h1:nqDD4MMMQA0lmWq03Z2/myGPYLQoXtmi0rGVs95ntbo=
github.com/yuin/goldmark-highlighting v0.0.0-20200307114337-60d527fdb691 h1:VWSxtAiQNh3zgHJpdpkpVYjTPqRE3P6UZCOPa1nRDio=
go.etcd.io/bbolt v1.3.2 h1:Z/90sZLPOeCy2PwprqkFa25PdkusRzaj9P8zm/KNyvk=
go.opencensus.io v0.22.3 h1:8sGtKOrtQqkN1bp2AtX+misvLIlOmsEsNd+9NIcPEm8=
go.uber.org/atomic v1.6.0 h1:Ezj3JGmsOnG1MoRWQkPBsKLe9DwWD9QeXzTRzzldNVk=
go.uber.org/multierr v1.5.0 h1:KCa4XfM8CWFCpxXRGok+Q0SS/0XBhMDbHHGABQLvD2A=
go.uber.org/zap v1.14.1 h1:nYDKopTbvAPq/NrUVZwT15y2lpROBiLLyoRTbXOYWOo=
golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 h1:DOmugCavvUtnUD114C1Wh+UgTgQZ4pMLzXxi1pSt+/Y=
golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e h1:3G+cUijn7XD+S4eJFddp53Pv7+slrESplyjG25HgL+k=
golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d h1:TzXSXBo42m9gQenoE3b9BGiEpg5IG2JkU5FkPIawgtw=
golang.org/x/sys v0.0.0-20200409092240-59c9f1ba88fa h1:mQTN3ECqfsViCNBgq+A40vdwhkGykrrQlYe3mPj6BoU=
golang.org/x/text v0.3.2 h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=
google.golang.org/api v0.20.0 h1:jz2KixHX7EcCPiQrySzPdnYT7DbINAypCqKZ1Z7GM40=
google.golang.org/genproto v0.0.0-20200409111301-baae70f3302d h1:I7Vuu5Ejagca+VcgfBINHke3xwjCTYnIG4Q57fv0wYY=
google.golang.org/grpc v1.27.1 h1:zvIju4sqAGvwKspUQOhwnpcqSbzi7/H6QomNNjTL4sk=
gopkg.in/natefinch/lumberjack.v2 v2.0.0 h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=
gopkg.in/square/go-jose.v2 v2.4.1 h1:H0TmLt7/KmzlrDOpa1F+zr0Tk90PbJYBfsVUmRLrf9Y=
gopkg.in/yaml.v2 v2.2.8 h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=
howett.net/plist v0.0.0-20181124034731-591f970eefbb h1:jhnBjNi9UFpfpl8YZhA9CrOqpnJdvzuiHsl/dnxl11M=
```
