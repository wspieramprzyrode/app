---
id: inventory
title: Inventory Service
---
## Basic informations

Component directory in github repo structure: ```inventory```.

We using golang to create this service.

It use Google Firestore as data storage created by terraform - defined in file ```terraform/google.tf```.

### CI

Build process is defined in file ```.gihub/workflows/inventory_service.yaml```

We build and publish application as docker image ```wspieramprzyrode/inventory```.

### Deployment

Deployment od service in created as helm chart and located in ```chart/charts/inventory```.
