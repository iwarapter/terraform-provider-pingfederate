# Resource: pingfederate_extended_properties

Managed PingFederate Extended Properties.

## Example Usage
```hcl
resource "pingfederate_extended_properties" "example" {
  property {
    name         = "example"
    description  = "something about this attribute"
    multi_valued = true
  }
}
```

## Argument Attributes

The following arguments are supported:

### Property

The `property` block - Extended Property definition that allows to store additional information about IdP/SP Connections and OAuth Clients.

- `name` - (Required) The property name.

- `description` - (Optional) The property description.

- `multi_valued` - (Optional) Indicates whether the property should allow multiple values (defaults to false).

## Import

-> The resource ID is fixed as `extended_properties` because this is a singleton resource.

Extended Properties can be imported using the id, e.g.

```
terraform import pingfederate_extended_properties.example extended_properties
```
