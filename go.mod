module github.com/mixflowtech/go-librt

replace (
	github.com/Sirupsen/logrus v1.3.0 => github.com/sirupsen/logrus v1.3.0
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b => github.com/golang/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/text v0.3.1-0.20180805044716-cb6730876b98 => github.com/golang/text v0.3.1-0.20180805044716-cb6730876b98
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2 => github.com/golang/time v0.0.0-20180412165947-fbb02b2291d2
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 => github.com/golang/tools v0.0.0-20180828015842-6cd1fcedba52
	google.golang.org/api v0.0.0-20180910000450-7ca32eb868bf => github.com/google/google-api-go-client v0.0.0-20180910000450-7ca32eb868bf
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.16.0 => github.com/grpc/grpc-go v1.16.0
	gopkg.in/yaml.v2 v2.2.2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
)

require (
	github.com/keybase/go-framed-msgpack-rpc a27a4f7712dd5021fa6f4f09f2f96f905e8b7747
	github.com/keybase/go-logging f3c7c3c1605e29ebcce430bb1e0d048faf4081ce
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20170602164621-9e8dc3f972df // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
