# Release process for veidemann-api

## Prerequisites

Make sure the repository secret named API_COMMIT_PAT is set with a valid
[fine grained personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token)
with the following rights:

| Service         | Access         |
|-----------------|----------------|
| Actions         | Read-only      |
| Commit statuses | Read and write |
| Contents        | Read and write |
| Metadata        | Read-only      |
| Pull requests   | Read and write |
| Secrets         | Read-only      |
| Variables       | Read and write |

## Release

Cut a release using [GitHubs release feature](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository):

* Tag name and release title should be  `v<major>.<minor>.<patch>[-<pre-release>]` where `<pre-release>` is optional.
* Use the auto generated release notes as a starting point for the release notes.

This will trigger the release workflow which will generate language specific stubs. 
The generated stubs for Go will be pushed to the `master` branch and tagged with the release tag prefixed with `go/`.