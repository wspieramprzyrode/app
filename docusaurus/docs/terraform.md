---
id: terraform
title: Terraform
sidebar_label: Terraform
---

## Basic informations

[Code repository](https://github.com/wspieramprzyrode/app)

Infrastructure are defined in terraform.

Data are stored in ```terraform``` directory

State of infrastructure is store at [Terraform Cloud](https://app.terraform.io) is described in file ```terraform/init.tf```

## CI

Every pull request with changes in ```terraform``` is formatted, validated and planned by github workflow defined in file ```.gihub/workflows/ui_test.yaml```

## Providers

All Providers are defined  in file ```terraform/provider.tf```

### DigitalOcean

All resources are defined in file ```terraform/digitalocean.tf```

We using kubernetes cluster located in fra1 region

### CloudFlare

All resources are defined in file ```terraform/cloudflare.tf```

We using dns services to our domains:

- wspieramprzyrode.pl

To manage records created inside kubernetes cluster we using [external-dns](https://github.com/kubernetes-sigs/external-dns) deployed as helm chart

### Google and Firebase

All resources are defined in file ```terraform/google.tf```

Prepare google credentials file to serve from [Terraform Cloud](https://app.terraform.io) variable:

- create google_account.json file as one line file
- paste as ```GOOGLE_CREDENTIALS``` variable in [Terraform Cloud](https://app.terraform.io) workflow site
