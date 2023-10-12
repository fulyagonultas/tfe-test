variable "iam" {
  description = "IAM bindings in {SECRET => {ROLE => [MEMBERS]}} format."
  type        = map(map(list(string)))
  default     = {}
}

variable "labels" {
  description = "Optional labels for each secret."
  type        = map(map(string))
  default     = {}
}

variable "project_id" {
  description = "Project id where the secret will be created."
  type        = string
}

variable "secrets" {
  description = "Map of secrets to manage and their locations. Locations should not be kept null."
  type        = map(list(string))
  default     = {}
  validation {
    condition = alltrue(
      [for secret in values(var.secrets) : secret != null]
    )
    error_message = "Each secret should contain a location"
  }
}

variable "encryption_key" {
  description = "Self link of the KMS keys in {LOCATION => KEY} format. A key must be provided for all replica locations."
  type        = map(string)
  default     = null
}
