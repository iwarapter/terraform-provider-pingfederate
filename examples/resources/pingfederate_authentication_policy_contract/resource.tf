resource "pingfederate_authentication_policy_contract" "example" {
  name                = "example"
  extended_attributes = ["foo", "bar"]
}
