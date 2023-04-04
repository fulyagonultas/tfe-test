output "name" {
  description = "The compute instance name."
  value       = google_compute_instance.demo.name
}

output "instance_name" {
  description = "The compute instance name."
  value       = google_compute_instance.demo.*.name
}

output "instance_id" {
  description = "The compute instance id."
  value       = google_compute_instance.demo.*.instance_id
}

output "internal_ip" {
  description = "The internal IP address."
  value       = google_compute_instance.demo.*.network_interface.0.network_ip
}