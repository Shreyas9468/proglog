module github.com/Shreyas9468/proglog

go 1.25.0

require (
	github.com/casbin/casbin/v3 v3.11.0
	github.com/gorilla/mux v1.8.1
	github.com/hashicorp/raft v1.1.1
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea
	github.com/hashicorp/serf v0.10.4
	github.com/stretchr/testify v1.11.1
	github.com/tysonmote/gommap v0.0.3
	go.opencensus.io v0.22.2
	go.uber.org/zap v1.28.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260819154853-08b0e4226688
	google.golang.org/grpc v1.83.1
)

require (
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/casbin/govaluate v1.3.0 // indirect
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-hclog v0.9.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.6.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-msgpack/v2 v2.1.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-sockaddr v1.0.7 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/memberlist v0.6.0 // indirect
	github.com/miekg/dns v1.1.72 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.37.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	golang.org/x/tools v0.47.0 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/travisjeffery/go-dynaport v1.0.0
	google.golang.org/protobuf v1.36.12
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/hashicorp/raft-boltdb => github.com/travisjeffery/raft-boltdb v1.0.0
