module "secret_iam_binding" {
  source     = "../"
  project_id = local.project_id
  secrets = {
    secret-iam-binding = [local.region]
  }
  iam = {
    secret-iam-binding = {
      "roles/secretmanager.secretAccessor" = ["serviceAccount:dbg-tfe-modules-95de4071@appspot.gserviceaccount.com"]
    }
  }
}
