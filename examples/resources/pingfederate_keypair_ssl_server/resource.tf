# importing a keypair using base64 encoded PKCS12 data
resource "pingfederate_keypair_ssl_server" "example_import" {
  file_data = filebase64("provider.p12")
  password  = "password"
}

# generating a keypair
resource "pingfederate_keypair_ssl_server" "example_generate" {
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
