resource "pingfederate_extended_properties" "example" {
  property {
    name         = "example"
    description  = "something about this attribute"
    multi_valued = true
  }
}
