resource "pingfederate_authentication_policy_contract" "demo" {
  name                = "demo"
  extended_attributes = ["foo", "bar"]
}
