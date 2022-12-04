resource "pingfederate_keypair_signing" "demo_generate" {
  city                      = "Test"
  common_name               = "Test"
  country                   = "GB"
  key_algorithm             = "RSA"
  key_size                  = 2048
  organization              = "Test"
  organization_unit         = "Test"
  state                     = "Test"
  valid_days                = 365
  subject_alternative_names = ["foo", "bar"]
}
