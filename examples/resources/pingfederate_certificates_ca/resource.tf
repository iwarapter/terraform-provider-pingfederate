resource "pingfederate_certificates_ca" "example" {
  certificate_id = "example"
  file_data      = filebase64("example.pem")
}
