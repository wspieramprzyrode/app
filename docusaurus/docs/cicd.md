---
id: cicd
title: GitHub Actions
---

## GitHub Actions

All build and deployment process using github actions

### Add taskfile and k8s tools to github workflow

Example workflow

```yaml
name: Install Tools
on: [push]
jobs:
  install-tools:
    runs-on: ubuntu-latest
    env:
      GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    steps:
      - uses: actions/checkout@v2
      - name: Set up k8s tools
        with:
          kubectl: '1.7.5'
          helm: '3.2.1'
          skaffold: '1.9.1'
        uses: daisaru11/setup-cd-tools@v1
      - name: Set up Taskfile
        uses: Arduino/actions/setup-taskfile@master
        with:
          version: '2.x'
      - name: List available tasks
        run: |
          task -l
```
