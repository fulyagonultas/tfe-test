locals {
  project_id = "dbg-tfe-modules-95de4071"
  region     = "europe-west3"
  labels = { simple-secret = {
    "name"           = "simple-secret"
    "environment"    = "dev"
    "product"        = "web"
    "productlineid"  = "3001"
    "creator"        = "admin"
    "supportgroupid" = "admin"
    "applicationid"  = "4001"
    "costcenter"     = "11001"
    }
  }
  europe_west3_encryption_key = "projects/dbg-tfe-modules-95de4071/locations/europe-west3/keyRings/test-secret-encryption/cryptoKeys/test-encryption"
}
