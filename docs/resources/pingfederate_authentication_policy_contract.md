# Resource: pingfederate_authentication_policy_contract

Provides a authentication policy contract.

## Example Usage
```terraform
resource "pingfederate_authentication_policy_contract" "demo" {
  name                = "demo"
  extended_attributes = ["foo", "bar"]
}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The Authentication Policy Contract Name. Name is unique.

- [`extended_attributes`](#extended_attributes) - A list of additional attributes as needed.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The access token validator's ID.

- [`core_attributes`](#core_attributes) - A list of read-only assertion attributes (for example, subject) that are automatically populated by PingFederate.

## Import

Authentication Policy Contracts can be imported using the id, e.g.

```bash
$ terraform import pingfederate_authentication_policy_contract.demo 123
```
