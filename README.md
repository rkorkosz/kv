# Persistent env key value store for multiple projects

I created this package to handle environment variables for many projects that I worked at.
Since env variables values shouldn't be kept in repo the need for other storage is big.
You can store env vars in seperate file that is listed in .gitignore, but experience teaches us,
that this approach can fail and sensitive data can end up in the repo.

## Installation

Linux users that uses amd64 architecture can download latest version from [releases page](https://github.com/rkorkosz/kv/releases)

If you have go installed you can always `go get` it

```
go get -u github.com/rkorkosz/kv
```

## Usage

1. To parse existing env file to kv database:

```
$ kv parse <filename>
```

2. To list all stored variables:

```
$ kv ls
```

3. To export all variables in shell:

```
$ export $(kv ls)
```

4. To store one variable:

```
$ kv set KEY=VAL
```

or
```
$ kv set KEY VAL
```

## Implementation

`kv` database by default is stored in `$HOME/.kv/kv.db`.
As of now single storage implementation is based on [bbolt](https://github.com/etcd-io/bbolt)

Suggestions and feature request are welcome :)
