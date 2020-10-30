# DNSBL-query

dnsbl-query is a service that can be used to check
[DNSBL](https://www.spamhaus.org/whitepapers/dnsbl_function/) records for an IP
address. dnsbl-query currently supports IPv4 addresses.

## Getting Started

dnsbl-query offers a GraphQL API for checking, and enqueing checks on IP addresses
against a DNSBL service. By default, this uses the
[Spamhaus DNSBL service](https://www.spamhaus.org/faq/section/DNSBL%20Usage%23200)
but it can be configured to point to any supported DNSBL service.

### Running dnsbl-query

To run dnsbl-query locally, you can use [Docker Compose](https://docs.docker.com/compose/)
to get started with limited configuration time.

You can also run the binary locally using the following flow:

```zsh
mage -v go:build
./bin/dnsbl-query
```

Configuration can be provided either by environment variables, or CLI flags. You
can get a full list of configuration using `-h` on the bianry.

#### Release

The CI/CD tooling will automatically take care of pushing the latest version of
the built image to Dockerhub. To run this manually, you can use

```zsh
mage -v docker:push
```

#### Deploying dnsbl-query

You can use the [Helm](https://helm.sh) chart to configure manifests, or deploy
directly into your cluster.

The Helm chart is automatically linted and tested on PR.

### Contributing

Pull requests are welcome to make changes. All pull requests must complete their
GitHub workflows successfully before they will be merged.

Many packages that contains some non-generated code has a `doc.go` which explains
the package level documentation. A quick run down of the internal packages:

- `internal/auth` -- authentication tooling and middlewares for the GraphQL service
- `internal/cmd` -- entry points and sub-commands
- `internal/dnsbl` -- DNSBL query service
- `internal/ent` -- [ent](github.com/facebook/ent) custom compiler with gqlgen support
- `internal/ent/schema` -- entity schema definitions
- `internal/ent/schema/mixin` -- entity mixins for additional common fields
- `internal/flags` -- flag definitions for entry points
- `internal/graphql` -- gqlgen-based GraphQL generation and server
- `internal/graphql/schema` -- *.graphql schema definitions
- `internal/graphql/internal/model` -- non-ent models for gqlgen autobinding
- `internal/graphql/internal/resolvers` -- gqlgen resolvers
- `internal/worker` -- background worker library
