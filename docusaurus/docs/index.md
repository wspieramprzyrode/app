---
id: index
title: Getting started
---

## Wspieram Przyrodę developer documentation

[Web application repo](https://github.com/wspieramprzyrode/app)

[Mobile application repo](https://github.com/wspieramprzyrode/mobile)

## Tools

### SonarCloud

[Wspieram Przyrodę SonarCloud Page](https://sonarcloud.io/organizations/wspieramprzyrode)

Every pull request and branch master are automated checked at app and mobile repositories.

## Translations

Project translations are managed by [Cronwdin](https://crowdin.com/) service.

[Translations](https://crowdin.com/project/wspieramprzyrode/)

Configuration is defined in file ```.crowdin.yaml```

Secrets for github actions are managed byc terraform and defined in ```.terraform/github.tf``` file.

Values with credentials are defined in [Terraform Cloud](https://app.terraform.io).

## Documentation generation process

Project documentation is manage by [Docusaurus](https://docusaurus.io/).

Files are defined in ```docusaurus``` directory.

Build process is defined in file ```.gihub/workflows/ui_test.yaml```.

Documentation from branch master are deployed as github pages.

DNS record for docs subdomain is defined in ```.terraform/cloudflare.tf``` file and manage by terrafrom scripts.

Build process is defined in file ```.gihub/workflows/documentation.yaml```

We build and publish application as docker image ```wspieramprzyrode/docs```
