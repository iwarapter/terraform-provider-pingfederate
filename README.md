PingFederate Terraform Provider
==================

- Website: https://registry.terraform.io/providers/iwarapter/pingfederate/latest
- ![Status: Experimental](https://img.shields.io/badge/status-experimental-EAAA32) [![Gitter](https://badges.gitter.im/iwarapter/terraform-provider-pingfederate.svg)](https://gitter.im/iwarapter/terraform-provider-pingfederate?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
  [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=github.com.iwarapter.terraform-provider-pingfederate&metric=coverage)](https://sonarcloud.io/dashboard?id=github.com.iwarapter.terraform-provider-pingfederate)
  [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=github.com.iwarapter.terraform-provider-pingfederate&metric=alert_status)](https://sonarcloud.io/dashboard?id=github.com.iwarapter.terraform-provider-pingfederate)
  ![ci](https://github.com/iwarapter/terraform-provider-pingfederate/workflows/ci/badge.svg)
  ![GitHub release (latest by date)](https://img.shields.io/github/v/release/iwarapter/terraform-provider-pingfederate)
  [![Github All Releases](https://img.shields.io/github/downloads/iwarapter/terraform-provider-pingfederate/total.svg)]()

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)

Using the Provider
----------------------

Please see the terraform registry docs for detailed usage documentation:
https://registry.terraform.io/providers/iwarapter/pingfederate/latest/docs

The provider is current tested against the following versions of PingFederate

| PingFederate | Status                                                                                                                                                                                                  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ~~9.3.3~~    | PingIdentity no longer supply working containers for this version                                                                                                                                       |
| ~~10.0.6~~   | PingIdentity no longer supply working containers for this version                                                                                                                                       |
| 10.1.5       | [![ci](https://github.com/iwarapter/terraform-provider-pingfederate/actions/workflows/build.yaml/badge.svg)](https://github.com/iwarapter/terraform-provider-pingfederate/actions/workflows/build.yaml) |
| 10.2.4       | [![ci](https://github.com/iwarapter/terraform-provider-pingfederate/actions/workflows/build.yaml/badge.svg)](https://github.com/iwarapter/terraform-provider-pingfederate/actions/workflows/build.yaml) |
| 10.3.1       | [![ci](https://github.com/iwarapter/terraform-provider-pingfederate/actions/workflows/build.yaml/badge.svg)](https://github.com/iwarapter/terraform-provider-pingfederate/actions/workflows/build.yaml) |

Whilst 9.3.3 and 10.0.x should both still be compatible, PingIdentity no longer supply **working** docker images and so have been removed from the regression test pack.

Developing the Provider
---------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](https://github.com/iwarapter/terraform-provider-pingfederate#requirements) before proceeding).

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/development/terraform-providers/`).

Clone repository to: `$HOME/development/terraform-providers/`

```sh
$ git clone git@github.com:iwarapter/terraform-provider-pingfederate.git
...
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the local directory.

```sh
$ make build
...
$ terraform-provider-pingfederate
...
```

Using the Provider
----------------------

Please see the terraform registry docs for detailed usage documentation:
https://registry.terraform.io/providers/iwarapter/pingfederate/latest/docs

Testing the Provider
---------------------------

In order to test the provider, you can run `make sweep test`.

```sh
$ make sweep test
```

This will run the acceptance tests by initializing a local docker container to execute the functional tests against.
