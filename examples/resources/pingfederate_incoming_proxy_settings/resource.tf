resource "pingfederate_incoming_proxy_settings" "settings" {
  client_cert_chain_sslheader_name  = "foo"
  client_cert_sslheader_name        = "foo"
  forwarded_host_header_index       = "LAST"
  forwarded_host_header_name        = "X-Forwarded-Host"
  forwarded_ip_address_header_index = "LAST"
  forwarded_ip_address_header_name  = "X-Forwarded-For"
  proxy_terminates_https_conns      = false
}
