# Makefile
VERSION ?= 0.0.1-BETA
NAME=terraform-provider-pingfederate_v${VERSION}
OS_NAME := $(shell uname -s | tr A-Z a-z)


pf-init:
	@docker run -d --rm --hostname pingfederate \
		--name pingfederate \
		-e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) \
		-e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) \
		-e PING_IDENTITY_ACCEPT_EULA=YES \
		-e PF_LOG_LEVEL=DEBUG \
		-e TAIL_LOG_PARALLEL="y" \
		-e TAIL_LOG_FILES="/opt/out/instance/log/server.log /opt/out/instance/log/admin-api.log" \
		-e SERVER_PROFILE_URL=https://github.com/pingidentity/pingidentity-server-profiles.git \
		-e SERVER_PROFILE_PATH=getting-started/pingfederate \
		--publish 9999:9999 \
		--publish 9031:9031 \
		pingidentity/pingfederate:10.0.6-edge

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
	go test ./... -v -sweep=all -timeout 60m

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
