---
id: infrastructure
title: Infrastructure
sidebar_label: Infrastructure
---


## Terraform

Infrastructure are described in terraform. Data are stored in ```terraform``` directory
State of infrastructure is store at <https://app.terraform.io>

Prepare google credentials file to serve from app.terraform.io variable: create google_account.json file as one liner file and paste as GOOGLE_CREDENTIALS variable in app.terraform.io site

## DNS and CDN

We using cloudflare to manage dns services for domain ```wspieramprzyrode.pl```

## Kubernetes

We using kubernetes deployed to digitalocean cloud with:

- metrics-server
- external-dns
- ambassador

all deployments are described in terraform manifests.
