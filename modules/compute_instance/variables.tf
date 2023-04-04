variable "name" {
  description = "Name of the instance"
  type        = string
}

 
variable "machine_type" {
  description = "Type of machine"
  type        = string
}

variable "gcp_zone" {
  description = "GCP zone, e.g. us-east1-a"
  default     = "europe-west3-a"
}

variable "image" {
  description = "image to build instance from"
  type        = string
}

variable "networking_project" {
  description = "The network to deploy to"
  type        = string
}

variable "network_tier" {
  description = "The default network tier."
  type        = string
  default     = "PREMIUM"
}

variable "subnetwork" {
  description = "The subnetwork to deploy to"
  type        = string
}

variable "network_ip" {
  description = "The private IP address to assign to the instance. If empty, the address will be automatically assigned."
  type        = string
  default     = ""
}

variable "external_ip" {
  description = "The public IP address to assign to the instance. Instance must be whitelisted by Cloud team in Org-Policy. Public IP addresses are subject for strict security control."
  type        = bool
  default     = false
}

variable "can_ip_forward" {
  description = "Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false."
  type        = bool
  default     = false
}