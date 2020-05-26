provider "cloudflare" {
  version = "~> 2.6.0"
}
provider "digitalocean" {
  version = "~> 1.18"
}
provider "kubernetes" {
  version          = "~> 1.11"
  load_config_file = false
  host             = digitalocean_kubernetes_cluster.cluster.endpoint
  token            = digitalocean_kubernetes_cluster.cluster.kube_config[0].token
  cluster_ca_certificate = base64decode(
    digitalocean_kubernetes_cluster.cluster.kube_config[0].cluster_ca_certificate
  )
}

provider "helm" {
  version = "~> 1.2"
  kubernetes {
    load_config_file = false
    host             = digitalocean_kubernetes_cluster.cluster.endpoint
    token            = digitalocean_kubernetes_cluster.cluster.kube_config[0].token
    cluster_ca_certificate = base64decode(
      digitalocean_kubernetes_cluster.cluster.kube_config[0].cluster_ca_certificate
    )
  }
}

provider "github" {
  version = "~> 2.8"
}

provider "google-beta" {
  version = "~> 3.23"
  region  = "europe-west1"
  project = "wspieramprzyrode"
}
