module "encrypt_secret" {
  source     = "../"
  project_id = local.project_id
  secrets = {
    cmek-encrypted-secret = [local.region]
  }
  encryption_key = {
    europe-west3 = local.europe_west3_encryption_key
  }
}
