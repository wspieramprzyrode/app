---
id: kubernetes
title: Kubernetes
sidebar_label: Kubernetes
---

We using [kubernetes](https://kubernetes.io) deployed to [DigitalOcean](https://www.digitalocean.com/) cloud with:

- [metrics-server](https://github.com/kubernetes-sigs/metrics-server)
- [external-dns](https://github.com/kubernetes-sigs/external-dns)
- [ambassador](https://www.getambassador.io/)

All deployments are defined in terraform file ```terraform/helm.tf``` (for helm charts deployments).
