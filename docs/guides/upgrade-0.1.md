---
layout: ""
page_title: "PingFederate Provider 0.1 Upgrade Guide"
description: |-
  The PingFederate provider provides resources to interact with a PingFederate admin API.
---

# Terraform PingFederate Provider Version 0.1.0 Upgrade Guide

The provider is currently being re-written to use the new [terraform plugin framework](https://developer.hashicorp.com/terraform/plugin/framework).
This enables several language benefits on the previous SDK for how PingFederate resources are defined. With the new framework all resources are now generated complete from the PingFederate swagger API.

The following resources have been converted and their changes are documented below:
- `pingfederate_authentication_policy_contract`
- `pingfederate_oauth_client`

## pingfederate_authentication_policy_contract

- The previous versions supported a custom field `policy_contract_id` used to set the ID of the PingFederate resource, the resource now supports setting the `id` field directly.

## pingfederate_oauth_client

The oauth client has changes to several previous configurations since migration to the framework:

The following fields have been changed from a block to a nested attribute:
- client_auth
- oidc_policy
- jwks_settings

See below example:
#### Before
```hcl
resource "pingfederate_oauth_client" "example" {
  client_auth {
    secret = "super_top_secret"
    type   = "SECRET"
  }
}
```
#### After
```hcl
resource "pingfederate_oauth_client" "example" {
  client_auth = {
    secret = "super_top_secret"
    type   = "SECRET"
  }
}
```

The following resource references have been collapsed from blocks to attributes:
- default_access_token_manager_ref
- request_policy_ref
- token_exchange_processor_policy_ref
- oidc_policy.policy_group

See below example:
#### Before
```hcl
resource "pingfederate_oauth_client" "example" {
  default_access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }
}
```
#### After
```hcl
resource "pingfederate_oauth_client" "example" {
  default_access_token_manager_ref = pingfederate_oauth_access_token_manager.example.id
}
```

The `extended_properties` block has been renamed to `extended_parameters` inline with the API name and converted to a map:

#### Before
```hcl
resource "pingfederate_oauth_client" "example" {
  extended_properties {
    key_name = "foo"
    values   = ["foobar"]
  }
  extended_properties {
    key_name = "bar"
    values   = ["barfoo"]
  }
}
```

#### After
```hcl
resource "pingfederate_oauth_client" "example" {
  extended_parameters = {
    "foo" = {
      values = ["foobar"]
    },
    "bar" = {
      values = ["barfoo"]
    }
  }
}
```
