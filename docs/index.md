The PingFederate provider is used to interact with the many resources supported by the PingFederate admin API. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.


### Example Usage
```terraform
# Configure the PingFederate Provider
provider "pingfederate" {
  username = "Administrator"
  password = "2Access"
  base_url = "https://localhost:9999"
  context  = "/pf-admin-api/v1"
}

# Create a site
resource "pingaccess_site" "site" {
  # ...
}
```

### Authentication

The PingFederate provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials

- Environment variables

### Static credentials
Static credentials can be provided by adding an `username` and `password` in-line in the PingFederate provider block:

Usage:
```terraform
provider "pingfederate" {
  username = "Administrator"
  password = "2Access"
  base_url = "https://localhost:9999"
  context  = "/pf-admin-api/v1"
}
```

### Environment variables
You can provide your credentials via the `PINGFEDERATE_USERNAME`, `PINGFEDERATE_PASSWORD`, `PINGFEDERATE_CONTEXT` and `PINGFEDERATE_BASEURL` environment variables.

```terraform
provider "pingfederate" {}
```

Usage:
```bash
$ export PINGFEDERATE_USERNAME="Administrator"
$ export PINGFEDERATE_PASSWORD="top_secret"
$ export PINGFEDERATE_CONTEXT="/pf-admin-api/v1"
$ export PINGFEDERATE_BASEURL="https://myadmin.server:9999"
$ terraform plan
```
