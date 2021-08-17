package gosrc

import (
	"encoding/json"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var urlTests = []struct {
		expected bool   // expected result
		in       string // input
	}{
		{true, `go-mod-graph-charting github.com/gin-gonic/gin@v1.6.3
go-mod-graph-charting github.com/go-playground/validator/v10@v10.4.1
go-mod-graph-charting github.com/golang/protobuf@v1.4.3
go-mod-graph-charting github.com/json-iterator/go@v1.1.10
go-mod-graph-charting github.com/leodido/go-urn@v1.2.1
go-mod-graph-charting github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd
go-mod-graph-charting github.com/modern-go/reflect2@v1.0.1
go-mod-graph-charting github.com/ugorji/go@v1.2.2
go-mod-graph-charting golang.org/x/crypto@v0.0.0-20201221181555-eec23a3978ad
go-mod-graph-charting golang.org/x/sys@v0.0.0-20201223074533-0d417f636930
go-mod-graph-charting google.golang.org/protobuf@v1.25.0
go-mod-graph-charting gopkg.in/yaml.v2@v2.4.0
github.com/leodido/go-urn@v1.2.1 github.com/stretchr/testify@v1.6.1
github.com/gin-gonic/gin@v1.6.3 github.com/gin-contrib/sse@v0.1.0
github.com/gin-gonic/gin@v1.6.3 github.com/go-playground/validator/v10@v10.2.0
github.com/gin-gonic/gin@v1.6.3 github.com/golang/protobuf@v1.3.3
github.com/gin-gonic/gin@v1.6.3 github.com/json-iterator/go@v1.1.9
github.com/gin-gonic/gin@v1.6.3 github.com/mattn/go-isatty@v0.0.12
github.com/gin-gonic/gin@v1.6.3 github.com/stretchr/testify@v1.4.0
github.com/gin-gonic/gin@v1.6.3 github.com/ugorji/go/codec@v1.1.7
github.com/gin-gonic/gin@v1.6.3 gopkg.in/yaml.v2@v2.2.8
github.com/stretchr/testify@v1.4.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.4.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.4.0 github.com/stretchr/objx@v0.1.0
github.com/stretchr/testify@v1.4.0 gopkg.in/yaml.v2@v2.2.2
github.com/ugorji/go/codec@v1.1.7 github.com/ugorji/go@v1.1.7
github.com/stretchr/testify@v1.6.1 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.6.1 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.6.1 github.com/stretchr/objx@v0.1.0
github.com/stretchr/testify@v1.6.1 gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
golang.org/x/crypto@v0.0.0-20201221181555-eec23a3978ad golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
golang.org/x/crypto@v0.0.0-20201221181555-eec23a3978ad golang.org/x/sys@v0.0.0-20191026070338-33540a1f6037
golang.org/x/crypto@v0.0.0-20201221181555-eec23a3978ad golang.org/x/term@v0.0.0-20201117132131-f5c789dd3221
google.golang.org/protobuf@v1.25.0 github.com/golang/protobuf@v1.4.1
google.golang.org/protobuf@v1.25.0 github.com/google/go-cmp@v0.5.0
google.golang.org/protobuf@v1.25.0 google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013
github.com/golang/protobuf@v1.4.1 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.1 google.golang.org/protobuf@v1.22.0
github.com/ugorji/go@v1.2.2 github.com/ugorji/go/codec@v1.2.2
gopkg.in/yaml.v2@v2.2.8 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
github.com/gin-contrib/sse@v0.1.0 github.com/stretchr/testify@v1.3.0
github.com/json-iterator/go@v1.1.10 github.com/davecgh/go-spew@v1.1.1
github.com/json-iterator/go@v1.1.10 github.com/google/gofuzz@v1.0.0
github.com/json-iterator/go@v1.1.10 github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
github.com/json-iterator/go@v1.1.10 github.com/modern-go/reflect2@v0.0.0-20180701023420-4b7aa43c6742
github.com/json-iterator/go@v1.1.10 github.com/stretchr/testify@v1.3.0
github.com/golang/protobuf@v1.4.3 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.3 google.golang.org/protobuf@v1.23.0
golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3 golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2
golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3 golang.org/x/text@v0.3.0
gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
github.com/json-iterator/go@v1.1.9 github.com/davecgh/go-spew@v1.1.1
github.com/json-iterator/go@v1.1.9 github.com/google/gofuzz@v1.0.0
github.com/json-iterator/go@v1.1.9 github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
github.com/json-iterator/go@v1.1.9 github.com/modern-go/reflect2@v0.0.0-20180701023420-4b7aa43c6742
github.com/json-iterator/go@v1.1.9 github.com/stretchr/testify@v1.3.0
github.com/stretchr/testify@v1.3.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.3.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.3.0 github.com/stretchr/objx@v0.1.0
github.com/google/go-cmp@v0.5.0 golang.org/x/xerrors@v0.0.0-20191204190536-9bdfabe68543
github.com/google/go-cmp@v0.4.0 golang.org/x/xerrors@v0.0.0-20191204190536-9bdfabe68543
github.com/ugorji/go@v1.1.7 github.com/ugorji/go/codec@v1.1.7
gopkg.in/yaml.v2@v2.4.0 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
github.com/mattn/go-isatty@v0.0.12 golang.org/x/sys@v0.0.0-20200116001909-b77594299b42
github.com/go-playground/validator/v10@v10.4.1 github.com/go-playground/assert/v2@v2.0.1
github.com/go-playground/validator/v10@v10.4.1 github.com/go-playground/locales@v0.13.0
github.com/go-playground/validator/v10@v10.4.1 github.com/go-playground/universal-translator@v0.17.0
github.com/go-playground/validator/v10@v10.4.1 github.com/leodido/go-urn@v1.2.0
github.com/go-playground/validator/v10@v10.4.1 golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9
github.com/leodido/go-urn@v1.2.0 github.com/stretchr/testify@v1.4.0
golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9 golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
golang.org/x/crypto@v0.0.0-20200622213623-75b288015ac9 golang.org/x/sys@v0.0.0-20190412213103-97732733099d
github.com/go-playground/validator/v10@v10.2.0 github.com/go-playground/assert/v2@v2.0.1
github.com/go-playground/validator/v10@v10.2.0 github.com/go-playground/locales@v0.13.0
github.com/go-playground/validator/v10@v10.2.0 github.com/go-playground/universal-translator@v0.17.0
github.com/go-playground/validator/v10@v10.2.0 github.com/leodido/go-urn@v1.2.0
github.com/go-playground/universal-translator@v0.17.0 github.com/go-playground/locales@v0.13.0
google.golang.org/protobuf@v1.23.0 github.com/golang/protobuf@v1.4.0
google.golang.org/protobuf@v1.23.0 github.com/google/go-cmp@v0.4.0
github.com/ugorji/go/codec@v1.2.2 github.com/ugorji/go@v1.2.2
github.com/golang/protobuf@v1.4.0 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.0 google.golang.org/protobuf@v1.21.0
google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013 github.com/golang/protobuf@v1.4.1
google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013 golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013 golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135
google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013 google.golang.org/grpc@v1.27.0
google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013 google.golang.org/protobuf@v1.23.1-0.20200526195155-81db48ad09cc
google.golang.org/genproto@v0.0.0-20200526211855-cb27e3aa2013 honnef.co/go/tools@v0.0.0-20190523083050-ea95bdfd59fc
google.golang.org/protobuf@v1.21.0 github.com/golang/protobuf@v1.4.0-rc.4.0.20200313231945-b860323f09d0
google.golang.org/protobuf@v1.21.0 github.com/google/go-cmp@v0.4.0
golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3 golang.org/x/tools@v0.0.0-20190311212946-11955173bddd
golang.org/x/term@v0.0.0-20201117132131-f5c789dd3221 golang.org/x/sys@v0.0.0-20191026070338-33540a1f6037
github.com/go-playground/locales@v0.13.0 golang.org/x/text@v0.3.2
golang.org/x/text@v0.3.2 golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2 golang.org/x/sys@v0.0.0-20190215142949-d0b11bdaac8a
github.com/golang/protobuf@v1.4.0-rc.4.0.20200313231945-b860323f09d0 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.0-rc.4.0.20200313231945-b860323f09d0 google.golang.org/protobuf@v1.20.1-0.20200309200217-e05f789c0967
golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135 golang.org/x/sync@v0.0.0-20190423024810-112230192c58
golang.org/x/tools@v0.0.0-20190311212946-11955173bddd golang.org/x/net@v0.0.0-20190311183353-d8887717615a
google.golang.org/protobuf@v1.23.1-0.20200526195155-81db48ad09cc github.com/golang/protobuf@v1.4.0
google.golang.org/protobuf@v1.23.1-0.20200526195155-81db48ad09cc github.com/google/go-cmp@v0.4.0
gopkg.in/yaml.v2@v2.2.2 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
google.golang.org/grpc@v1.27.0 github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473
google.golang.org/grpc@v1.27.0 github.com/envoyproxy/protoc-gen-validate@v0.1.0
google.golang.org/grpc@v1.27.0 github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b
google.golang.org/grpc@v1.27.0 github.com/golang/mock@v1.1.1
google.golang.org/grpc@v1.27.0 github.com/golang/protobuf@v1.3.2
google.golang.org/grpc@v1.27.0 github.com/google/go-cmp@v0.2.0
google.golang.org/grpc@v1.27.0 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
google.golang.org/grpc@v1.27.0 golang.org/x/oauth2@v0.0.0-20180821212333-d2e6202438be
google.golang.org/grpc@v1.27.0 golang.org/x/sys@v0.0.0-20190215142949-d0b11bdaac8a
google.golang.org/grpc@v1.27.0 google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55
google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55 github.com/golang/protobuf@v1.3.2
google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55 golang.org/x/exp@v0.0.0-20190121172915-509febef88a4
google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55 golang.org/x/lint@v0.0.0-20190227174305-5b3e6a55c961
google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55 golang.org/x/tools@v0.0.0-20190226205152-f727befe758c
google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55 google.golang.org/grpc@v1.19.0
google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55 honnef.co/go/tools@v0.0.0-20190102054323-c2f93a96b099
google.golang.org/grpc@v1.19.0 cloud.google.com/go@v0.26.0
google.golang.org/grpc@v1.19.0 github.com/BurntSushi/toml@v0.3.1
google.golang.org/grpc@v1.19.0 github.com/client9/misspell@v0.3.4
google.golang.org/grpc@v1.19.0 github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b
google.golang.org/grpc@v1.19.0 github.com/golang/mock@v1.1.1
google.golang.org/grpc@v1.19.0 github.com/golang/protobuf@v1.2.0
google.golang.org/grpc@v1.19.0 golang.org/x/lint@v0.0.0-20181026193005-c67002cb31c3
google.golang.org/grpc@v1.19.0 golang.org/x/net@v0.0.0-20180826012351-8a410e7b638d
google.golang.org/grpc@v1.19.0 golang.org/x/oauth2@v0.0.0-20180821212333-d2e6202438be
google.golang.org/grpc@v1.19.0 golang.org/x/sync@v0.0.0-20180314180146-1d60e4601c6f
google.golang.org/grpc@v1.19.0 golang.org/x/sys@v0.0.0-20180830151530-49385e6e1522
google.golang.org/grpc@v1.19.0 golang.org/x/text@v0.3.0
google.golang.org/grpc@v1.19.0 golang.org/x/tools@v0.0.0-20190114222345-bf090417da8b
google.golang.org/grpc@v1.19.0 google.golang.org/appengine@v1.1.0
google.golang.org/grpc@v1.19.0 google.golang.org/genproto@v0.0.0-20180817151627-c66870c02cf8
google.golang.org/grpc@v1.19.0 honnef.co/go/tools@v0.0.0-20190102054323-c2f93a96b099
golang.org/x/net@v0.0.0-20190311183353-d8887717615a golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2
golang.org/x/net@v0.0.0-20190311183353-d8887717615a golang.org/x/text@v0.3.0
github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473 github.com/census-instrumentation/opencensus-proto@v0.2.1
github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473 github.com/envoyproxy/protoc-gen-validate@v0.1.0
github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473 github.com/golang/protobuf@v1.3.2
github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473 github.com/prometheus/client_model@v0.0.0-20190812154241-14fe0d1b01d4
github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473 google.golang.org/genproto@v0.0.0-20190819201941-24fa4b261c55
github.com/envoyproxy/go-control-plane@v0.9.1-0.20191026205805-5f8ba28d4473 google.golang.org/grpc@v1.23.0
golang.org/x/tools@v0.0.0-20190226205152-f727befe758c golang.org/x/net@v0.0.0-20190213061140-3a22650c66bd
golang.org/x/tools@v0.0.0-20190226205152-f727befe758c golang.org/x/sync@v0.0.0-20181108010431-42b317875d0f
golang.org/x/tools@v0.0.0-20190226205152-f727befe758c google.golang.org/appengine@v1.4.0
google.golang.org/appengine@v1.4.0 github.com/golang/protobuf@v1.2.0
google.golang.org/appengine@v1.4.0 golang.org/x/net@v0.0.0-20180724234803-3673e40ba225
google.golang.org/appengine@v1.4.0 golang.org/x/text@v0.3.0
google.golang.org/protobuf@v1.20.1-0.20200309200217-e05f789c0967 github.com/golang/protobuf@v1.4.0-rc.2
google.golang.org/protobuf@v1.20.1-0.20200309200217-e05f789c0967 github.com/google/go-cmp@v0.4.0
google.golang.org/protobuf@v1.22.0 github.com/golang/protobuf@v1.4.0
google.golang.org/protobuf@v1.22.0 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.0-rc.2 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.0-rc.2 google.golang.org/protobuf@v0.0.0-20200228230310-ab0ca4ff8a60
google.golang.org/protobuf@v0.0.0-20200228230310-ab0ca4ff8a60 github.com/golang/protobuf@v1.4.0-rc.1.0.20200221234624-67d41d38c208
google.golang.org/protobuf@v0.0.0-20200228230310-ab0ca4ff8a60 github.com/google/go-cmp@v0.4.0
github.com/prometheus/client_model@v0.0.0-20190812154241-14fe0d1b01d4 github.com/golang/protobuf@v1.2.0
github.com/prometheus/client_model@v0.0.0-20190812154241-14fe0d1b01d4 golang.org/x/sync@v0.0.0-20181108010431-42b317875d0f
golang.org/x/lint@v0.0.0-20190227174305-5b3e6a55c961 golang.org/x/tools@v0.0.0-20190226205152-f727befe758c
github.com/golang/protobuf@v1.4.0-rc.1.0.20200221234624-67d41d38c208 github.com/google/go-cmp@v0.4.0
github.com/golang/protobuf@v1.4.0-rc.1.0.20200221234624-67d41d38c208 google.golang.org/protobuf@v0.0.0-20200221191635-4d8936d0db64
google.golang.org/grpc@v1.23.0 cloud.google.com/go@v0.26.0
google.golang.org/grpc@v1.23.0 github.com/BurntSushi/toml@v0.3.1
google.golang.org/grpc@v1.23.0 github.com/client9/misspell@v0.3.4
google.golang.org/grpc@v1.23.0 github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b
google.golang.org/grpc@v1.23.0 github.com/golang/mock@v1.1.1
google.golang.org/grpc@v1.23.0 github.com/golang/protobuf@v1.2.0
google.golang.org/grpc@v1.23.0 github.com/google/go-cmp@v0.2.0
google.golang.org/grpc@v1.23.0 golang.org/x/lint@v0.0.0-20190313153728-d0100b6bd8b3
google.golang.org/grpc@v1.23.0 golang.org/x/net@v0.0.0-20190311183353-d8887717615a
google.golang.org/grpc@v1.23.0 golang.org/x/oauth2@v0.0.0-20180821212333-d2e6202438be
google.golang.org/grpc@v1.23.0 golang.org/x/sys@v0.0.0-20190215142949-d0b11bdaac8a
google.golang.org/grpc@v1.23.0 golang.org/x/tools@v0.0.0-20190524140312-2c0ae7006135
google.golang.org/grpc@v1.23.0 google.golang.org/appengine@v1.1.0
google.golang.org/grpc@v1.23.0 google.golang.org/genproto@v0.0.0-20180817151627-c66870c02cf8
google.golang.org/grpc@v1.23.0 honnef.co/go/tools@v0.0.0-20190523083050-ea95bdfd59fc
google.golang.org/protobuf@v0.0.0-20200221191635-4d8936d0db64 github.com/golang/protobuf@v1.4.0-rc.1
google.golang.org/protobuf@v0.0.0-20200221191635-4d8936d0db64 github.com/google/go-cmp@v0.3.1
github.com/golang/protobuf@v1.4.0-rc.1 github.com/google/go-cmp@v0.3.1
github.com/golang/protobuf@v1.4.0-rc.1 google.golang.org/protobuf@v0.0.0-20200109180630-ec00e32a8dfd
google.golang.org/protobuf@v0.0.0-20200109180630-ec00e32a8dfd github.com/google/go-cmp@v0.3.0`},
	}

	for _, tt := range urlTests {
		actual, _, _, actual2 := BuildTree(tt.in)
		if (nil != actual) != tt.expected {
			t.Errorf("BuildTree(%s) = %v, %v, expected %t", tt.in, actual, actual2, tt.expected)
		} else {
			v1, _ :=json.Marshal(actual)
			v2, _ :=json.Marshal(actual2)
			t.Logf("tree: %s\n an-tree: %s\n", v1, v2)
		}
	}
}
