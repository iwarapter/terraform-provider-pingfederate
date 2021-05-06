# Data Source: pingfederate_custom_data_store

Use this data source to get a custom data store in Ping Federate by its name.

## Example Usage
```hcl
data "pingfederate_custom_data_store" "example" {
  name = "example"
}
```

## Argument Attributes
The following arguments are supported:

- [`name`](#name) - (required) The name for the custom data store.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The data store's id.
