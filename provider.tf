terraform {
  cloud {
    organization = "deutsche-boerse"
    hostname     = "tfe-dev.deutsche-boerse.de"
    workspaces {
      name = "terraform-snippets"
    }
  }
  required_version = ">= 1.2.7"
  required_providers {
    google = {
      version = "4.50.0"
    }
    vault = {
      version = "3.14.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.4.3"
    }
  }
}

provider "google" {
  project      = var.gcp_project
  region       = var.gcp_region
  access_token = data.vault_generic_secret.gcp_tf_token.data["token"]
}

provider "vault" {
  address         = "https://vault-dev.deutsche-boerse.de"
  skip_tls_verify = true
}