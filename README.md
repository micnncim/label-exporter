# label-exporter

Export GitHub labels in the form of YAML, JSON and table.

## Installation

Create a personal access token with `public_repo` scope at 
https://github.com/settings/tokens/new and:

```console
$ export GITHUB_TOKEN=enter_the_token_here
```

Then either download a binary release from 
https://github.com/micnncim/label-exporter/releases or:

```console
$ go get github.com/micnncim/label-exporter/cmd/label-exporter
```

## Usage

```console
$ label-exporter --help
usage: label-exporter [<flags>] <owner> <repo>

Flags:
      --help   Show context-sensitive help (also try --help-long and --help-man).
  -y, --yaml   Use the YAML format.
  -j, --json   Use the JSON format.
  -t, --table  Use the table format.

Args:
  <owner>  Owner of the repository.
  <repo>   Repository whose wanted labels.
```

## Example

```console
$ label-exporter micnncim label-exporter --yaml
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

```console
$ label-exporter micnncim label-exporter --json | jq
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
  }
]
```

```console
$ label-exporter micnncim label-exporter --table
+------------------+--------------------------------+--------+
|       NAME       |          DESCRIPTION           | COLOR  |
+------------------+--------------------------------+--------+
| bug              | Something isn't working        | d73a4a |
| documentation    | Improvements or additions to   | 0075ca |
|                  | documentation                  |        |
| duplicate        | This issue or pull request     | cfd3d7 |
|                  | already exists                 |        |
+------------------+--------------------------------+--------+
```
