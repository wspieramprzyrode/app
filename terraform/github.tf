resource "github_actions_secret" "kuberenetes_github_repo_secret" {
  repository      = var.github_repository
  secret_name     = "KUBE_CONFIG"
  plaintext_value = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
}

resource "github_actions_secret" "crowdin_project_id_github_repo_secret" {
  repository      = var.github_repository
  secret_name     = "CROWDIN_PROJECT_ID"
  plaintext_value = var.crowdin_project_id
}

resource "github_actions_secret" "crowdin_api_token_github_repo_secret" {
  repository      = var.github_repository
  secret_name     = "CROWDIN_API_TOKEN"
  plaintext_value = var.crowdin_api_token
}
