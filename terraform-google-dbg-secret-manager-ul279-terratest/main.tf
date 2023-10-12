module "secret_manager" {
  source         = "git::https://github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/secret-manager?ref=v26.0.0"
  project_id     = var.project_id
  secrets        = var.secrets
  iam            = var.iam
  encryption_key = var.encryption_key
  labels         = var.labels
}
