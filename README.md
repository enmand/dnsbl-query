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
mage -v build:binary
./bin/dnsbl-query <configuration>
```

### Deploying dnsbl-query

You can use the [Helm](https://helm.sh) chart to configure manifests, or deploy
directly into your cluster.

TODO

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information on how to contribute
