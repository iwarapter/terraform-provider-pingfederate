# Makefile
VERSION ?= 0.0.1-BETA
NAME=terraform-provider-pingfederate_v${VERSION}
OS_NAME := $(shell uname -s | tr A-Z a-z)
CURDATE := ${shell date +'%y%m%d'}

pf-init:
	@docker run -d --rm --hostname pingfederate \
		--name pingfederate \
		-e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) \
		-e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) \
		-e PING_IDENTITY_ACCEPT_EULA=YES \
		-e OPERATIONAL_MODE=CLUSTERED_CONSOLE \
		-e CLUSTER_BIND_ADDRESS=LINK_LOCAL \
		-e CLUSTER_NAME=COMPOSE_PF_CLUSTER \
		-e DNS_QUERY_LOCATION=pingfederate-admin \
		-e DNS_RECORD_TYPE=A \
		-e PF_LOG_LEVEL=DEBUG \
		-e TAIL_LOG_PARALLEL="y" \
		-e TAIL_LOG_FILES="/opt/out/instance/log/server.log /opt/out/instance/log/admin-api.log" \
		-e SERVER_PROFILE_URL=https://github.com/pingidentity/pingidentity-server-profiles.git \
		-e SERVER_PROFILE_PATH=getting-started/pingfederate \
		-e IMAGE_VERSION=pingfederate-alpine-az11-10.2.2-${CURDATE}-d9b5 \
		--publish 9999:9999 \
		--publish 9031:9031 \
		pingidentity/pingfederate:10.2.2-edge

checks:
	@go fmt ./...
	@go vet ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w pingfederate

unit-test:
	@go test -mod=vendor ./... -v


sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test ./pingfederate -v -sweep=all -timeout 60m

test:
	@rm -f pingfederate/terraform.log && TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v

test-and-report:
	@rm -f pingfederate/terraform.log coverage.out report.json
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

build:
	@go build -mod=vendor -o ${NAME} -gcflags "all=-trimpath=$GOPATH" .

deploy-local:
	@mkdir -p ~/.terraform.d/plugins
	@cp ${NAME} ~/.terraform.d/plugins/
	@mkdir -p ~/.terraform.d/plugins/registry.terraform.io/iwarapter/pingfederate/${VERSION}/${OS_NAME}_amd64
	@cp ${NAME} ~/.terraform.d/plugins/registry.terraform.io/iwarapter/pingfederate/${VERSION}/${OS_NAME}_amd64

func-init:
	@rm -rf func-tests/.terraform.lock.hcl
	@rm -rf func-tests/.terraform
	@rm -rf func-tests/crash.log
	@rm -rf func-tests/run.log
	@cd func-tests && terraform init -lock=true

func-plan:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform plan

func-apply:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform apply -auto-approve

func-destroy:
	@cd func-tests && terraform destroy -auto-approve
