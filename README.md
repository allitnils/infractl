# infractl

Azure infrastructure auditing CLI — drift detection, cost analysis, tagging compliance, and policy validation.

## Overview

`infractl` is a command-line tool for auditing Azure infrastructure. It surfaces compliance gaps, cost anomalies, configuration drift from ARM templates, and Azure Policy violations — all from your terminal.

Built for platform and infrastructure engineers who need fast, scriptable visibility into their Azure estate without leaving the shell.

## Installation

```bash
go install github.com/allitnils/infractl@latest
```

Or download a pre-built binary from [Releases](https://github.com/allitnils/infractl/releases).

## Authentication

Set your subscription ID:

```bash
export AZURE_SUBSCRIPTION_ID=<your-subscription-id>
# or use the flag on any command
infractl audit --subscription <your-subscription-id>
```

In production, infractl uses `DefaultAzureCredential` — it will pick up environment variables, managed identity, or `az login` credentials automatically.

## Commands

### audit

Run compliance checks against Azure resources.

```bash
infractl audit                              # all checks
infractl audit --checks tagging,encryption  # specific checks
infractl audit --output json                # JSON output for piping
infractl audit --subscription <id>          # explicit subscription
```

Available checks: `tagging`, `encryption`, `networking`, `identity`, `backup`, `all`

### cost

Analyse cost trends and surface anomalies.

```bash
infractl cost --days 30
infractl cost --days 7 --output json
```

### drift

Compare live Azure resource state against an ARM template.

```bash
infractl drift --template ./azuredeploy.json
infractl drift --template ./azuredeploy.json --resource-group my-rg
infractl drift --template ./azuredeploy.json --output json
```

### policy

Evaluate Azure Policy compliance state across the subscription.

```bash
infractl policy
infractl policy --output json
```

### version

```bash
infractl version
```

## Configuration

Create `~/.infractl.yaml` to set defaults:

```yaml
subscription: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
output: table
```

## Output formats

| Format | Flag | Use case |
|--------|------|----------|
| `table` | `-o table` | Human-readable terminal output with lipgloss colouring |
| `json` | `-o json` | Machine-readable; pipe to `jq`, feed into dashboards or alerts |

## Audit checks

| Check | Severity | What it validates |
|-------|----------|-------------------|
| `tagging` | MEDIUM | Required tags present: `Environment`, `Owner`, `CostCenter`, `Project` |
| `encryption` | HIGH | Disk encryption on VMs; HTTPS-only traffic on storage accounts |
| `networking` | CRITICAL | NSGs with unrestricted inbound rules (allow-all) |
| `identity` | MEDIUM | Managed identity enabled — reduces credential sprawl |
| `backup` | HIGH | Backup policy assigned to virtual machines |

## Architecture

```
infractl/
├── cmd/           Cobra subcommands (audit, cost, drift, policy, version)
├── internal/
│   ├── azure/     AzureClient interface + MockClient (swap for SDK in production)
│   ├── audit/     Runner orchestrates checks; checks/ contains one file per check
│   ├── drift/     ARM template parser + live-state comparator
│   ├── output/    TablePrinter (lipgloss) and JSONPrinter
│   └── config/    Viper-based config loading
└── tests/         Integration tests using MockClient
```

The `AzureClient` interface in `internal/azure/client.go` is the seam for the real Azure SDK. The `MockClient` used in tests and demo mode returns deterministic data — swap it for an SDK-backed implementation pointing at your subscription for production use.

## Examples

```bash
# Audit all resources, output JSON, filter only failures with jq
infractl audit --output json | jq '[.[] | select(.Passed == false)]'

# Check drift against a template, alert if any changes found
infractl drift --template prod.json --output json | jq 'if length > 0 then error("drift detected") else empty end'

# Cost summary piped to a log
infractl cost --days 7 --output json >> /var/log/infracosts.json
```

## Contributing

1. Fork the repo
2. Create a feature branch: `git checkout -b feat/my-check`
3. Add your check in `internal/audit/checks/`
4. Implement the `Check` interface: `Evaluate(r azure.Resource) CheckResult`
5. Register it in `internal/audit/runner.go`
6. Run tests: `go test ./...`
7. Submit a PR
