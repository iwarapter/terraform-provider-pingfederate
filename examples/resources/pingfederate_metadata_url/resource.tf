resource "pingfederate_metadata_url" "test" {
  name               = "test"
  url                = "https://sptest.iamshowcase.com/testsp_metadata.xml"
  validate_signature = false
}
