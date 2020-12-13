# Resource: pingfederate_authentication_api_application

Provides an authentication API application.

## Example Usage
```hcl
resource "pingfederate_authentication_api_application" "example" {
  name                       = "myapp"
  url                        = "https://example.com"
  description                = "This is an auth api app"
  additional_allowed_origins = ["https://foo.bar", "https://bar.foo"]
}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The Authentication API Application Name. Name must be unique.

- [`url`](#url) - (Required) The Authentication API Application redirect URL.

- [`description`](#description) - The Authentication API Application description.

- [`additional_allowed_origins`](#additional_allowed_origins) - The domain in the redirect URL is always whitelisted. This field contains a list of additional allowed origin URL's for cross-origin resource sharing.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the application.

## Import

Authentication API Applications can be imported using the id, e.g.

```
terraform import pingfederate_authentication_api_application.example 123
```
