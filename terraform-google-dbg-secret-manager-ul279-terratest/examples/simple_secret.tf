module "simple_secret_manager" {
  source     = "../"
  project_id = local.project_id
  secrets = {
    simple-secret = [local.region]
  }
  labels = local.labels
}
