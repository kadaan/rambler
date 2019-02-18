## Setup your machine

`rambler` is written in [Go](https://golang.org/).

Prerequisites:

- `make`
- [Go 1.11+](https://golang.org/doc/install)

Clone `rambler` anywhere:

```sh
$ git clone git@github.com:kadaan/rambler.git
```

Install the build and lint dependencies:

```console
$ make setup
```

A good way of making sure everything is all right is running the test suite:

```console
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```console
$ make build
```

When you are satisfied with the changes, we suggest you run:

```console
$ make ci
```

Which runs all the linters and tests.

## Create a commit

Commit messages should be well formatted, and to make that "standarized", we
are using Convetional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `rambler` fork and open a pull request against the
master branch.