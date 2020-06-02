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

Every pull request and master branches are checked automatically in the mobile and app repositories.

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

Master branch documentation is released as github pages.

DNS record for docs subdomain is defined in ```.terraform/cloudflare.tf``` file and manage by terrafrom scripts.

Build process is defined in file ```.gihub/workflows/documentation.yaml```

Application is build and publish as docker image ```wspieramprzyrode/docs```
