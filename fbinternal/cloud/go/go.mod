module magma/fbinternal/cloud/go

replace (
	magma/gateway => ./../../../orc8r/gateway/go
	magma/lte/cloud/go => ./../../../lte/cloud/go
	magma/orc8r/cloud/go => ./../../../orc8r/cloud/go
	magma/orc8r/lib/go => ./../../../orc8r/lib/go
	magma/orc8r/lib/go/protos => ./../../../orc8r/lib/go/protos
)

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3
	github.com/Masterminds/squirrel v1.1.1-0.20190513200039-d13326f0be73
	github.com/aws/aws-sdk-go v1.19.6
	github.com/coreos/go-semver v0.2.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/strfmt v0.19.4
	github.com/go-openapi/swag v0.18.0
	github.com/go-openapi/validate v0.18.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.3
	github.com/labstack/echo v0.0.0-20181123063414-c54d9e8eed6c
	github.com/lib/pq v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_model v0.2.0
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.4.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.27.1
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
	magma/lte/cloud/go v0.0.0-00010101000000-000000000000
	magma/orc8r/cloud/go v0.0.0
	magma/orc8r/lib/go v0.0.0
	magma/orc8r/lib/go/protos v0.0.0
)

go 1.12
