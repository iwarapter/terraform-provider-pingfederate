# Resource: pingfederate_keypair_signing_csr

Provides a ssl server keypair csr.

## Example Usage

### Signing a CSR with an example tls signer

```hcl
resource "pingfederate_keypair_signing" "test_generate" {
  city              = "Test"
  common_name       = "Test"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 2048
  organization      = "Test"
  organization_unit = "Test"
  state             = "Test"
  valid_days        = 365
}

data "pingfederate_keypair_signing_csr" "csr" {
  id = pingfederate_keypair_signing.test_generate.id
}

resource "pingfederate_keypair_signing_csr" "test" {
  keypair_id = pingfederate_keypair_signing.test_generate.id
  file_data  = base64encode(tls_locally_signed_cert.example.cert_pem)
}

resource "pingfederate_certificates_ca" "demo" {
  file_data = base64encode(tls_self_signed_cert.example.cert_pem)
}

resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_locally_signed_cert" "example" {
  cert_request_pem   = data.pingfederate_keypair_signing_csr.csr.cert_request_pem
  ca_key_algorithm   = "RSA"
  ca_private_key_pem = tls_private_key.example.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.example.cert_pem

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "tls_self_signed_cert" "example" {
  key_algorithm   = "RSA"
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = "example.com"
    organization = "ACME Examples, Inc"
  }

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}
```

## Argument Attributes

The following arguments are supported:

- `keypair_id` - (Required) The ID for the keypair
- `file_data` - (Required) The CSR response file data in PKCS7 format or as an X.509 certificate. PEM encoding (with or without the header and footer lines) is required. New line characters should be omitted or encoded in this value.

## Attributes Reference

None
