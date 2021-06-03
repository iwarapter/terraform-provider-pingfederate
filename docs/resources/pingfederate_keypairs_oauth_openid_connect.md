# Resource: pingfederate_keypairs_oauth_openid_connect

Provides a signing keypair.

## Example Usage

```hcl
resource "pingfederate_keypairs_oauth_openid_connect" "example" {
  static_jwks_enabled = true
  rsa_active_cert_ref {
    id = pingfederate_keypair_signing.new.id
  }
  rsa_previous_cert_ref {
    id = pingfederate_keypair_signing.old.id
  }
  rsa_publish_x5c_parameter = true
}
```

## Argument Attributes

The following arguments are supported:

- `static_jwks_enabled` (Required) Enable static keys.
- `rsa_active_cert_ref` - (Required) Reference to the RSA key currently active.
- `rsa_previous_cert_ref` - (Optional) Reference to the RSA key previously active.
- `rsa_publish_x5c_parameter` - (Optional) Enable publishing of the RSA certificate chain associated with the active key.
- `p256_active_cert_ref` - (Optional) Reference to the P-256 key currently active.
- `p256_previous_cert_ref` - (Optional) Reference to the P-256 key previously active.
- `p256_publish_x5c_parameter` - (Optional) Enable publishing of the P-256 certificate chain associated with the active key.
- `p384_active_cert_ref` - (Optional) Reference to the P-384 key currently active.
- `p384_previous_cert_ref` - (Optional) Reference to the P-384 key previously active.
- `p384_publish_x5c_parameter` - (Optional) Enable publishing of the P-384 certificate chain associated with the active key.
- `p521_active_cert_ref` - (Optional) Reference to the P-521 key currently active.
- `p521_previous_cert_ref` - (Optional) Reference to the P-521 key previously active.
- `p521_publish_x5c_parameter` - (Optional) Enable publishing of the P-521 certificate chain associated with the active key.
- `p256_decryption_active_cert_ref` - (Optional) Reference to the P-256 decryption key currently active.
- `p256_decryption_previous_cert_ref` - (Optional) Reference to the P-256 decryption key previously active.
- `p256_decryption_publish_x5c_parameter` - (Optional) Enable publishing of the P-256 certificate chain associated with the active key.
- `p384_decryption_active_cert_ref` - (Optional) Reference to the P-384 decryption key currently active.
- `p384_decryption_previous_cert_ref` - (Optional) Reference to the P-384 decryption key previously active.
- `p384_decryption_publish_x5c_parameter` - (Optional) Enable publishing of the P-384 certificate chain associated with the active key.
- `p521_decryption_active_cert_ref` - (Optional) Reference to the P-521 decryption key currently active.
- `p521_decryption_previous_cert_ref` - (Optional) Reference to the P-521 decryption key previously active.
- `p521_decryption_publish_x5c_parameter` - (Optional) Enable publishing of the P-521 certificate chain associated with the active key.
- `rsa_decryption_active_cert_ref` - (Optional) Reference to the RSA decryption key currently active.
- `rsa_decryption_previous_cert_ref` - (Optional) Reference to the RSA decryption key previously active.
- `rsa_decryption_publish_x5c_parameter` - (Optional) Enable publishing of the RSA certificate chain associated with the active key.

## Import

-> The resource ID is fixed as `oidc_keypairs` because this is a singleton resource.

OIDC KeyPairs can be imported using the id, e.g.

```
terraform import pingfederate_keypairs_oauth_openid_connect.demo oidc_keypairs
```
