# Makefile
VERSION ?= 0.0.1-BETA
NAME=terraform-provider-pingfederate_v${VERSION}

pf-init:
	@docker run --rm -d --hostname pingfederate -v `pwd`/pingfederate/pingfederate-data.zip:/opt/in/instance/server/default/data/drop-in-deployer/data.zip --name pingfederate -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9999:9999 --publish 9031:9031 pingidentity/pingfederate:10.0.2-edge

checks:
	@go fmt ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w pingfederate

unit-test:
	@go test -mod=vendor ./... -v

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

func-init:
	@rm -rf func-tests/.terraform
	@rm -rf func-tests/crash.log
	@rm -rf func-tests/run.log
	@cd func-tests && terraform init

func-plan:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform plan

func-apply:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform apply -auto-approve

func-destroy:
	@cd func-tests && terraform destroy -auto-approve
