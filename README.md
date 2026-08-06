# fmtcheck

[![Release](https://img.shields.io/github/v/release/libnudget/fmtcheck?logo=github&label=latest)](https://github.com/libnudget/fmtcheck/releases)

A GitHub Action that enforces consistent formatting across your CI
pipeline.

fmtcheck runs your language's canonical formatter in check mode and fails
the build on any diff. Formatting debates stay out of code review because
the machine settles them in CI.

## Usage

```yaml
- uses: libnudget/fmtcheck@main
```

Checks the repository root. To check a subdirectory:

```yaml
- uses: libnudget/fmtcheck@main
  with:
    path: src
```

## What it checks

- Go source files (`*.go`) must match `gofmt` output.
- `.git` and `vendor` directories are skipped.

## Inputs

| Name | Required | Default | Description |
| --- | --- | --- | --- |
| `path` | no | `.` | Directory to check. |

## Local development

The check can be run without Docker:

```sh
go run . --path .
```

## Development

```sh
go vet ./...
go test ./...
gofmt -l .
```

## License

MIT
