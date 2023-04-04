output "name" {
  description = "The compute instance name."
  value       = module.compute_instance.name
}

output "instance_name" {
  description = "The compute instance name."
  value       = module.compute_instance.instance_name
}

output "instance_id" {
  description = "The compute instance id."
  value       = module.compute_instance.instance_id
}

output "internal_ip" {
  description = "The internal IP address."
  value       = module.compute_instance.internal_ip
}