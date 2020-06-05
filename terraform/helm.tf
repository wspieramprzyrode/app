resource "helm_release" "metrics-server" {
  name        = "metrics-server"
  repository  = "https://kubernetes-charts.storage.googleapis.com"
  chart       = "metrics-server"
  version     = "2.11.1"
  max_history = 10
  namespace   = "kube-system"
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
  name        = "external-dns"
  repository  = "https://charts.bitnami.com/bitnami"
  chart       = "external-dns"
  version     = "2.20.4"
  max_history = 10
  namespace   = "kube-system"
  values = [
    "${file("helm/configs/external-dns.yaml")}"
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

resource "helm_release" "ambassador" {
  name             = "ambassador"
  repository       = "https://www.getambassador.io"
  chart            = "ambassador"
  max_history      = 10
  version          = "6.4.1"
  create_namespace = true
  namespace        = "ambassador"
  values = [
    "${file("helm/configs/ambassador.yaml")}"
  ]
  set_sensitive {
    name  = "licenseKey.value"
    value = var.ambassador_licence
  }
}

resource "helm_release" "cert-manager" {
  name             = "cert-manager"
  repository       = "https://charts.jetstack.io"
  chart            = "cert-manager"
  max_history      = 10
  version          = "v0.15.1"
  create_namespace = true
  namespace        = "cert-manager"
  values = [
    "${file("helm/configs/cert-manager.yaml")}"
  ]
}

resource "kubernetes_namespace" "k8s_namespace_dev" {
  metadata {
    name = "dev"
  }
}

resource "helm_release" "cluster-resources" {
  depends_on = [
    helm_release.cert-manager,
    google_service_account_key.wpieramprzyrode_api_dev_key,
    kubernetes_namespace.k8s_namespace_dev
  ]
  name             = "cluster-common-resources"
  create_namespace = true
  namespace        = "app-system"
  chart            = "./helm/chart"
  values = [
    "${file("helm/configs/cluster-resources.yaml")}"
  ]
  set_sensitive {
    name  = "cloudflare.apiToken"
    value = var.cloudflare_api_key
  }
  set_sensitive {
    name  = "cloudflare.email"
    value = var.cloudflare_email
  }
  set_sensitive {
    name  = "letsencrypt.email"
    value = var.letsencrypt_email
  }
  set_sensitive {
    name  = "gceCredentialsFileDev"
    value = google_service_account_key.wpieramprzyrode_api_dev_key.private_key
  }
}
