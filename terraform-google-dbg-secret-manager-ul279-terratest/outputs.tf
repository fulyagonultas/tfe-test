output "ids" {
  description = "Fully qualified secret ids."
  value       = module.secret_manager.ids
}

output "secrets" {
  description = "Secret resources."
  value       = module.secret_manager.secrets
}
