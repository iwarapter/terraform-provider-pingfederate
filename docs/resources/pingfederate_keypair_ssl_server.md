# Resource: pingfederate_keypair_ssl_server

Provides a ssl server keypair.

## Example Usage

### Generating a KeyPair
```hcl
resource "pingfederate_keypair_ssl_server" "example" {
  city                      = "London"
  common_name               = "Example"
  country                   = "GB"
  key_algorithm             = "RSA"
  key_size                  = 2048
  organization              = "Testing"
  organization_unit         = "Test"
  state                     = "Test"
  valid_days                = 365
  subject_alternative_names = ["examle.com", "another.com"]
}
```

### Importing a KeyPair
```hcl
resource "pingfederate_keypair_ssl_server" "example" {
  file_data = filebase64("keystore.p12")
  password = "changeit"
}
```

## Argument Attributes

The following arguments are supported:

- Importing a KeyPair
    - `crypto_provider` - (Optional) Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.
    - `file_data` - (Required) Base64 encoded PKCS12 file data. New line characters should be omitted or encoded in this value.
    - `password` - (Required) Password for the PKCS12 file.

- Generating a KeyPair
    - `city` - (Optional) City.
    - `common_name` - (Required) Common name for key pair subject.
    - `country` - (Required) Country.
    - `crypto_provider` - (Optional) Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.
    - `key_algorithm` - (Required) Key generation algorithm. Supported algorithms are available through the /keyPairs/keyAlgorithms endpoint.
    - `key_size` - (Optional) Key size, in bits. If this property is unset, the default size for the key algorithm will be used. Supported key sizes are available through the /keyPairs/keyAlgorithms endpoint.
    - `organization` - (Required) Organization.
    - `organization_unit` - (Optional) Organization unit.
    - `signature_algorithm` - (Optional) Signature algorithm. If this property is unset, the default signature algorithm for the key algorithm will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint.
    - `state` - (Optional) State.
    - `subject_alternative_names` - (Optional) The subject alternative names (SAN).
    - `valid_days` - (Required) Number of days the key pair will be valid for.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The keypair ID (`keypair_id`)
- `crypto_provider` - Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.
- `expires` - The end date up until which the item is valid, in ISO 8601 format (UTC).
- `issuer_dn` - The issuer's distinguished name.
- `key_algorithm` - The public key algorithm.
- `key_size` - The public key size.
- `serial_number` - The serial number assigned by the CA.
- `sha1_fingerprint` - SHA-1 fingerprint in Hex encoding.
- `sha256_fingerprint` - SHA-256 fingerprint in Hex encoding.
- `signature_algorithm` - The signature algorithm.
- `status` - Status of the item.
- `subject_alternative_names` - The subject alternative names (SAN).
- `subject_dn` - The subject's distinguished name.
- `valid_from` - The start date from which the item is valid, in ISO 8601 format (UTC).
- `version` - The X.509 version to which the item conforms.

## Import

-> This is currently only supported for generated KeyPairs.

SSL Server KeyPairs can be imported using the id, e.g.

```
terraform import pingfederate_keypair_ssl_server.demo 123
```
