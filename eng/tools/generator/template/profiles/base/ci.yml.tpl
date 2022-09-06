# NOTE: Please refer to https://aka.ms/azsdk/engsys/ci-yaml before editing this file.
trigger:
  branches:
    include:
      - main
      - feature/*
      - hotfix/*
      - release/*
  paths:
    include:
    - sdk/profiles/{{packageName}}/

pr:
  branches:
    include:
      - main
      - feature/*
      - hotfix/*
      - release/*
  paths:
    include:
    - sdk/profiles/{{packageName}}/

stages:
- template: /eng/pipelines/templates/jobs/archetype-sdk-client.yml
  parameters:
    IncludeRelease: true
    ServiceDirectory: 'profiles/{{packageName}}'
