resource "helm_release" "metrics-server" {
  name       = "metrics-server"
  repository = "https://kubernetes-charts.storage.googleapis.com"
  chart      = "metrics-server"
  version    = "2.11.1"
  namespace  = "kube-system"
  set {
    name  = "args[0]"
    value = "--kubelet-insecure-tls"
  }

  set {
    name  = "args[1]"
    value = "--kubelet-preferred-address-types=InternalIP"
  }
}

resource "helm_release" "external-dns" {
  name       = "external-dns"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "external-dns"
  version    = "2.20.4"
  namespace  = "kube-system"
  values = [
    "${file("helm-configs/external-dns.yaml")}"
  ]
  set_sensitive {
    name  = "cloudflare.apiKey"
    value = var.cloudflare_api_key
  }
  set_sensitive {
    name  = "cloudflare.email"
    value = var.cloudflare_email
  }
}
