# Makefile

test:
	@rm -f pingfederate/terraform.log
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

build:
	@go build  -o terraform-provider-pingfederate .

deploy-local:
	@mkdir -p ~/.terraform.d/plugins
	@cp terraform-provider-pingfederate ~/.terraform.d/plugins/

func-init:
	@rm -rf func-tests/.terraform
	@rm -rf func-tests/crash.log
	@rm -rf func-tests/run.log
	@cd func-tests && terraform init

func-plan:
	@cd func-tests && terraform plan

func-apply:
	@cd func-tests && terraform apply -auto-approve

func-cli-gen:
	@cd ../pingfederate-sdk-go-gen-cli/ && make generate

sonar:
	@sonar-scanner -X \
		-Dsonar.projectKey=github.com.iwarapter.terraform-provider-pingfederate \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=28d86a90f2e4ae9563b4501cbc99de7522219c88 \
		-Dsonar.go.coverage.reportPaths=coverage.out \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor/**/* \
		-Dsonar.tests="." \
		-Dsonar.test.inclusions="**/*_test.go"
		