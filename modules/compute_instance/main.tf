resource "google_compute_instance" "demo" {
  name         = var.name
  machine_type = var.machine_type
  zone         = var.gcp_zone
  boot_disk {
    initialize_params {
      image = var.image
    }
  }

  can_ip_forward = var.can_ip_forward

  # networks to attach to the instance
  network_interface {
    subnetwork_project = var.networking_project
    subnetwork         = var.subnetwork
    network_ip         = var.network_ip != "" ? var.network_ip : ""

    # access configurations determines if instance gets public ip address
    dynamic "access_config" {
      for_each = var.external_ip == true ? [1] : []
      content {
        network_tier = var.network_tier
      }
    }
  }
}