module github.com/iwarapter/terraform-provider-pingfederate

go 1.14

require (
	github.com/aws/aws-sdk-go v1.30.12 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.0
	github.com/iwarapter/pingfederate-sdk-go v0.0.0-20211021210954-a29e9b626dca
	github.com/stretchr/testify v1.7.0
	golang.org/x/tools v0.1.0 // indirect
)

replace github.com/go-git/go-git-fixtures/v4 v4.2.1 => github.com/go-git/go-git-fixtures/v4 v4.0.1
