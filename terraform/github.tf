resource "github_actions_secret" "kuberenetes_github_repo_secret" {
  repository      = var.github_repository
  secret_name     = "KUBE_CONFIG"
  plaintext_value = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
}
