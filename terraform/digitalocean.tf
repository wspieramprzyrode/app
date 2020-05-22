resource "digitalocean_tag" "project_tag" {
  name = "wspieramprzyrode"
}
resource "digitalocean_tag" "cluster_tag" {
  name = "wspieramprzyrode-k8s"
}
resource "digitalocean_project" "do_project" {
  name = "wspieramprzyrode"
}

resource "digitalocean_kubernetes_cluster" "cluster" {
  name    = "wspieramprzyrode"
  region  = var.k8s_cluster_region
  version = var.k8s_cluster_version
  tags    = [digitalocean_tag.project_tag.id, digitalocean_tag.cluster_tag.id]
  node_pool {
    name       = "workers"
    size       = var.k8s_cluster_instance_type
    auto_scale = true
    min_nodes  = 1
    max_nodes  = 3
    tags       = [digitalocean_tag.project_tag.id, digitalocean_tag.cluster_tag.id]
  }
}
