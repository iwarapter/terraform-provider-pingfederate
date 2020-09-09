# Resource: pingfederate_certificates_ca

Provides a certificate ca.

## Example Usage
```terraform
resource "pingfederate_certificates_ca" "demo" {
  certificate_id = "example"
  file_data = base64encode(file("test_cases/amazon_root_ca1.pem"))
}
```

## Argument Attributes

The following arguments are supported:

- [`certificate_id`](#certificate_id) -  The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.

- [`file_data`](#file_data) - (Required) The certificate data in PEM format. New line characters should be omitted or encoded in this value.

- [`crypto_provider`](#crypto_provider) - ['LOCAL' or 'HSM']: Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the certificate (`certificate_id`).
- [`serial_number`](#serial_number) - The serial number assigned by the CA.
- [`subject_dn`](#subject_dn) - The subject's distinguished name.
- [`subject_alternative_names`](#subject_alternative_names) - The subject alternative names (SAN).
- [`issuer_dn`](#issuer_dn) - The issuer's distinguished name.
- [`valid_from`](#valid_from) - The start date from which the item is valid, in ISO 8601 format (UTC).
- [`expires`](#expires) - The end date up until which the item is valid, in ISO 8601 format (UTC).
- [`key_algorithm`](#key_algorithm) - The public key algorithm.
- [`key_size`](#key_size) - The public key size.
- [`signature_algorithm`](#signature_algorithm) - The signature algorithm.
- [`version`](#version) - The X.509 version to which the item conforms.
- [`sha1_fingerprint`](#sha1_fingerprint) - SHA-1 fingerprint in Hex encoding.
- [`sha256_fingerprint`](#sha256_fingerprint) - SHA-256 fingerprint in Hex encoding.
- [`status`](#status) - ['VALID' or 'EXPIRED' or 'NOT_YET_VALID' or 'REVOKED']: Status of the item.

## Import

Certificate CA can be imported using the id, e.g.

```bash
$ terraform import pingfederate_certificates_ca.demo 123
```
