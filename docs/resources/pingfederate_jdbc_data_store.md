# Resource: pingfederate_jdbc_data_store

Provides a Jdbc Data Store.

## Example Usage
```hcl
resource "pingfederate_jdbc_data_store" "demo" {
  name           = "terraform"
  driver_class   = "org.hsqldb.jdbcDriver"
  user_name      = "sa"
  password       = ""
  max_pool_size  = 10
  connection_url = "jdbc:hsqldb:mem:mymemdb"
  connection_url_tags {
    connection_url = "jdbc:hsqldb:mem:mymemdb"
    default_source = true
  }
}
```

## Argument Attributes

The following arguments are supported:

- `allow_multi_value_attributes` - (Optional) Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute.

- `blocking_timeout` - (Optional) The amount of time in milliseconds a request waits to get a connection from the connection pool before it fails. Omitting this attribute will set the value to the connection pool default.

- `connection_url` - (Optional) The default location of the JDBC database. This field is required if no mapping for JDBC database location and tags are specified.

- `connection_url_tags` - (Optional) The set of connection URLs and associated tags for this JDBC data store.

- `driver_class` - (Required) The name of the driver class used to communicate with the source database.

- `encrypted_password` - (Optional) The encrypted password needed to access the database. If you do not want to update the stored value, this attribute should be passed back unchanged.

- `idle_timeout` - (Optional) The length of time in minutes the connection can be idle in the pool before it is closed. Omitting this attribute will set the value to the connection pool default.

- `mask_attribute_values` - (Optional) Whether attribute values should be masked in the log.

- `max_pool_size` - (Optional) The largest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.

- `min_pool_size` - (Optional) The smallest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.

- `name` - (Optional) The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the connection url and the username.

- `password` - (Optional) The password needed to access the database. GETs will not return this attribute. To update this field, specify the new value in this attribute.

- `type` - (Required) The data store type.

- `user_name` - (Required) The name that identifies the user when connecting to the database.

- `validate_connection_sql` - (Optional) A simple SQL statement used by PingFederate at runtime to verify that the database connection is still active and to reconnect if needed.

### Connection Url Tags

The `connection_url_tags` block - A JDBC data store's connection URLs and tags configuration. This is required if no default JDBC database location is specified.

- `connection_url` - (Required) The location of the JDBC database.

- `default_source` - (Optional) Whether this is the default connection. Defaults to false if not specified.

- `tags` - (Optional) Tags associated with this data source.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the jdbc data store.

## Import

Jdbc Data Stores can be imported using the id, e.g.

```
terraform import pingfederate_jdbc_data_store.example 123
```
