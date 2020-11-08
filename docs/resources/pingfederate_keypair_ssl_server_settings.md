# Resource: pingfederate_keypair_ssl_server_settings

Provides a ssl server keypair settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops terraform tracking and updating the configuration.

## Example Usage

```hcl
resource "pingfederate_keypair_ssl_server_settings" "example" {
	admin_server_cert           = pingfederate_keypair_ssl_server.demo.id
	runtime_server_cert         = pingfederate_keypair_ssl_server.demo.id
	active_runtime_server_certs = [pingfederate_keypair_ssl_server.demo.id]
	active_admin_server_certs   = [pingfederate_keypair_ssl_server.demo.id]
}
```

## Argument Attributes

The following arguments are supported:

- `admin_server_cert` - (Required) Reference to the default SSL Server Certificate Key pair active for PF Administrator Console.
- `runtime_server_cert` - (Required) Reference to the default SSL Server Certificate Key pair active for Runtime Server.
- `active_admin_server_certs` - (Required) The CSR response file data in PKCS7 format or as an X.509 certificate. PEM encoding (with or without the header and footer lines) is required. New line characters should be omitted or encoded in this value.
- `active_runtime_server_certs` - (Required) The active SSL Server Certificate Key pairs for PF Administrator Console.
