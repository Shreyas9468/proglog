.PHONY: compile test init gencert

CONFIG_PATH := $(USERPROFILE)\.proglog

compile:
	protoc internal/api/v1/log.proto --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --proto_path=.

init:
	if not exist "$(CONFIG_PATH)" mkdir "$(CONFIG_PATH)"

$(CONFIG_PATH)\model.conf: test\model.conf
	if not exist "$(CONFIG_PATH)" mkdir "$(CONFIG_PATH)"
	copy /Y test\model.conf "$(CONFIG_PATH)\model.conf"

$(CONFIG_PATH)\policy.csv: test\policy.csv
	if not exist "$(CONFIG_PATH)" mkdir "$(CONFIG_PATH)"
	copy /Y test\policy.csv "$(CONFIG_PATH)\policy.csv"

test: $(CONFIG_PATH)\policy.csv $(CONFIG_PATH)\model.conf
	go test -race ./...

gencert:
	cfssl gencert -initca test/ca-csr.json | cfssljson -bare ca
	cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=test/ca-config.json -profile=server test/server-csr.json | cfssljson -bare server
	cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=test/ca-config.json -profile=client -cn="root" test/client-csr.json | cfssljson -bare root-client
	cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=test/ca-config.json -profile=client -cn="nobody" test/client-csr.json | cfssljson -bare nobody-client
	if not exist "$(CONFIG_PATH)" mkdir "$(CONFIG_PATH)"
	move ca.pem "$(CONFIG_PATH)\ca.pem"
	move ca-key.pem "$(CONFIG_PATH)\ca-key.pem"
	move server.pem "$(CONFIG_PATH)\server.pem"
	move server-key.pem "$(CONFIG_PATH)\server-key.pem"
	move server.csr "$(CONFIG_PATH)\server.csr"
	move root-client.pem "$(CONFIG_PATH)\root-client.pem"
	move root-client-key.pem "$(CONFIG_PATH)\root-client-key.pem"
	move root-client.csr "$(CONFIG_PATH)\root-client.csr"
	move nobody-client.pem "$(CONFIG_PATH)\nobody-client.pem"
	move nobody-client-key.pem "$(CONFIG_PATH)\nobody-client-key.pem"
	move nobody-client.csr "$(CONFIG_PATH)\nobody-client.csr"

.PHONY: compile test init gencert build-docker push-docker

TAG ?= 0.0.1
GITHUB_USER ?= shreyas9468
GHCR_IMAGE := ghcr.io/$(GITHUB_USER)/proglog

build-docker:
		docker build -t proglog:$(TAG) -t $(GHCR_IMAGE):$(TAG) -t $(GHCR_IMAGE):latest .

push-docker: build-docker
		docker push $(GHCR_IMAGE):$(TAG)
		docker push $(GHCR_IMAGE):latest