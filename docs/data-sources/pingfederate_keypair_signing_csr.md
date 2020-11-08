# Data Source: pingfederate_keypair_signing_csr

Use this data source to get the CSR of a signing keypair in Ping Federate.

## Example Usage
```hcl
data "pingfederate_keypair_signing_csr" "csr" {
  id = pingfederate_keypair_signing.demo_generate.id
}
```
## Argument Attributes
The following arguments are supported:

- [`id`](#id) - (required) The ID for the keypair.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`cert_request_pem`](#cert_request_pem) - The keypairs's Certificate Signing Response.
