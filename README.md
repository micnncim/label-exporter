# label-exporter

Export GitHub labels in the form of YAML and JSON.

## Installation

```
$ go get github.com/micnncim/label-exporter/cmd/label-exporter
```

## Usage

```
$ label-exporter --help
usage: label-exporter [<flags>] <owner> <repo>

Flags:
      --help  Show context-sensitive help (also try --help-long and --help-man).
  -y, --yaml  Use the YAML format.
  -j, --json  Use the JSON format.

Args:
  <owner>  Owner of the repository.
  <repo>   Repository whose wanted labels.
```

## Example

```
$ label-exporter micnncim lable-exporter --yaml
- color: d73a4a
  description: Something isn't working
  name: bug
- color: 0075ca
  description: Improvements or additions to documentation
  name: documentation
- color: cfd3d7
  description: This issue or pull request already exists
  name: duplicate
```

```
$ label-exporter micnncim lable-exporter --json | jq
[
  {
    "name": "bug",
    "description": "Something isn't working",
    "color": "d73a4a"
  },
  {
    "name": "documentation",
    "description": "Improvements or additions to documentation",
    "color": "0075ca"
  },
  {
    "name": "duplicate",
    "description": "This issue or pull request already exists",
    "color": "cfd3d7"
  },
]
```
