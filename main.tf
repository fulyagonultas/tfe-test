resource "random_string" "random" {
  length  = 4
  special = false
}

module "compute_instance" {
  source             = "./modules/compute_instance"
  name               = "${var.instance_name}-${lower(random_string.random.result)}"
  machine_type       = var.machine_type
  gcp_zone           = var.gcp_zone
  image              = var.image
  networking_project = var.networking_project
  network_tier       = var.network_tier
  subnetwork         = var.subnetwork
  network_ip         = var.network_ip
  external_ip        = var.external_ip
  can_ip_forward     = var.can_ip_forward
}