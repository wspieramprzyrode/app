terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "wspieramprzyrode"

    workspaces {
      name = "app"
    }
  }
}