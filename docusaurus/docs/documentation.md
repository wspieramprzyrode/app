---
id: documentation
title: Dokumentacja
---

## Documentation generation process

Project documentation is manage by [Docusaurus](https://docusaurus.io/).

Files are defined in ```docusaurus``` directory.

Build process is defined in file ```.gihub/workflows/ui_test.yaml```.

Master branch documentation is released as github pages.

DNS record for docs subdomain is defined in ```.terraform/cloudflare.tf``` file and manage by terrafrom scripts.

Build process is defined in file ```.gihub/workflows/documentation.yaml```

Application is build and publish as docker image ```wspieramprzyrode/docs```

## API documentation

Api doc is described as OpenApi 3.0 format and defined in this [file](https://raw.githubusercontent.com/wspieramprzyrode/app/master/oas3.yaml).

Every changes in master branch are deploy automatically to [Apiary](https://wspieramprzyrode.docs.apiary.io/)
