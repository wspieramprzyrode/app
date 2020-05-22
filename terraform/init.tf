terraform {
  required_version = "~> 0.12.25"
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "wspieramprzyrode"

    workspaces {
      name = "app"
    }
  }
}