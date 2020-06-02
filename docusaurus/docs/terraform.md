---
id: terraform
title: Terraform
sidebar_label: Terraform
---

## Basic informations

[Code repository](https://github.com/wspieramprzyrode/app)

Infrastructure is defined in terraform.

Data is stored in ```terraform``` directory

Infrastructure state is stored in [Terraform Cloud](https://app.terraform.io) and it is described in file ```terraform/init.tf```

## CI

Every pull request with changes in ```terraform``` is formatted, validated and planned by github workflow defined in file ```.gihub/workflows/ui_test.yaml```

## Providers

All Providers are defined in file ```terraform/provider.tf```

### DigitalOcean

All resources are defined in file ```terraform/digitalocean.tf```

We are using kubernetes cluster located in fra1 region

### CloudFlare

All resources are defined in file ```terraform/cloudflare.tf```

We are using following domains for our services:

- wspieramprzyrode.pl

To manage records created inside kubernetes cluster we are using [external-dns](https://github.com/kubernetes-sigs/external-dns) deployed as helm chart

### Google and Firebase

All resources are defined in file ```terraform/google.tf```

Prepare google credentials file to serve from [Terraform Cloud](https://app.terraform.io) variable:

- create google_account.json file as one line file
- paste it as ```GOOGLE_CREDENTIALS``` variable in [Terraform Cloud](https://app.terraform.io) workflow site
