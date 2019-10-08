# Synergy
### A simple command line tool to fetch details of resources created on GCP from terraform state files in remote location.

Synergy relies on __*carbon-status.json*__ file which contains the details of resources and the location to their terraform state files.

## Build locally
Requires _*go*_ to build.
```
$ git clone https://tools.lowes.com/stash/scm/e-dpiac/synergy.git
$ cd synergy
$ go mod vendor
$ make build
```
## Pre-requisites:
* GOOGLE_APPLICATION_CREDENTIALS env variable must be set to access the GCP project.

## Usage
Listing environments within cluster resources:
* `synergy --bucket=<STATUS_BUCKET> --project=<STATUS_PROJECT> --path=<STATUS_JSON_PATH> state read --resource=admin_cluster --environ=npe`
* `synergy --bucket=<STATUS_BUCKET> --project=<STATUS_PROJECT> --path=<STATUS_JSON_PATH> status list-environs --resource=admin_cluster`

## Output
TBD
