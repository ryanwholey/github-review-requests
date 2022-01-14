# github-review-requests

A CLI to send desktop notifications when new GitHub review requests are created.

## Install

```sh
brew install ryanwholey/formulas/github-review-requests
```

## Usage

```sh
grr --username <user> --interval 5m
```

The CLI expects `GH_TOKEN` to be set in the environment.

### Flags

| Flag | Description | Required | Default |
|---|---|---|---|
| `--username\|-u` | GitHub username | `true` | `GH_USERNAME` env var, else `""` |
| `--storage\|-s` | Path to a storage file | `false` | `~/.github-review-request-storage`|
| `--interval\|-i` | Interval to poll for new requests (e.g. `30s`, `5m`, etc.). If zero, run once and exit. | `false` | `0` |
| `--clean\|-c` | Remove storage file on initialization | `false` | `false` |

## Release

```sh
git tag v0.0.0
git push --tags
```
