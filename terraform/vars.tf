variable "k8s_cluster_region" {
  type    = string
  default = "fra1"
}
variable "k8s_cluster_version" {
  type    = string
  default = "1.17.5-do.0"
}
variable "k8s_cluster_instance_type" {
  type    = string
  default = "s-1vcpu-2gb"
}
variable "cloudflare_api_key" {
  type = string
}
variable "cloudflare_email" {
  type = string
}
variable "github_repository" {
  type = string
}
variable "ambassador_licence" {
  type = string
}

variable "letsencrypt_email" {
  type = string
}
