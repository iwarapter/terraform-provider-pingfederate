# Signing a CSR with an example tls signer
# Generate our keypair
resource "pingfederate_keypair_ssl_server" "test_generate" {
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

# Get the generated keypair CSR
data "pingfederate_keypair_ssl_server_csr" "csr" {
  id = pingfederate_keypair_ssl_server.test_generate.id
}

# Import the signed certificate for our keypair
resource "pingfederate_keypair_ssl_server_csr" "test" {
  keypair_id = pingfederate_keypair_ssl_server.test_generate.id
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
  cert_request_pem      = data.pingfederate_keypair_ssl_server_csr.csr.cert_request_pem
  ca_key_algorithm      = "RSA"
  ca_private_key_pem    = tls_private_key.example.private_key_pem
  ca_cert_pem           = tls_self_signed_cert.example.cert_pem
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
