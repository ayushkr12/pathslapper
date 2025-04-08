# pathslapper
Tool to append a path to each URL from stdin. Useful for quick enumeration.

## Installation:

Using go:

```sh
go install github.com/ayushkr12/pathslapper@latest
```

## Usage:

Basic usage:

```sh
$ cat urls.txt
https://example.com
https://test.io/

$ cat urls.txt | pathslapper "/admin"
https://example.com/admin
https://test.io/admin
```

## Workflow:

Example workflow for appending paths and probing:

```sh
cat hosts.txt | pathslapper "/website-backup.zip" | httpx -silent
```

- see https://github.com/projectdiscovery/httpx for more info about httpx
