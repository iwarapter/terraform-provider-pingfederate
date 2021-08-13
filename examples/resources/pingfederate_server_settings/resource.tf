resource "pingfederate_server_settings" "settings" {
  federation_info {
    base_url         = "https://localhost:9031"
    saml2_entity_id  = "testing"
    saml1x_issuer_id = "testing"
    wsfed_realm      = "testing"
  }
  # roles_and_protocols is deprecated in 10.1 and all roles are enabled by default
  roles_and_protocols {
    enable_idp_discovery = true
    idp_role {
      enable                       = true
      enable_outbound_provisioning = true
      enable_saml10                = true
      enable_saml11                = true
      enable_ws_fed                = true
      enable_ws_trust              = true
      saml20_profile {
        enable = true
      }
    }
    oauth_role {
      enable_oauth          = true
      enable_openid_connect = true
    }
    sp_role {
      enable                      = true
      enable_inbound_provisioning = true
      enable_openid_connect       = true
      enable_saml10               = true
      enable_saml11               = true
      enable_ws_fed               = true
      enable_ws_trust             = true
      saml20_profile {
        enable      = true
        enable_xasp = true
      }
    }
  }
}
