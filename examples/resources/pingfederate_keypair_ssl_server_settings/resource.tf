resource "pingfederate_keypair_ssl_server_settings" "example" {
  admin_server_cert           = pingfederate_keypair_ssl_server.example.id
  runtime_server_cert         = pingfederate_keypair_ssl_server.example.id
  active_runtime_server_certs = [pingfederate_keypair_ssl_server.example.id]
  active_admin_server_certs   = [pingfederate_keypair_ssl_server.example.id]
}
