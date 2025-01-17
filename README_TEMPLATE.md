# Conduit Connector for <resource>

[Conduit](https://conduit.io) connector for <resource>.

## How to build?

Run `make build` to build the connector.

## Source

A source connector pulls data from an external resource and pushes it to
downstream resources via Conduit.

### Configuration

| name                  | description                           | required | default value |
|-----------------------|---------------------------------------|----------|---------------|
| `source_config_param` | Description of `source_config_param`. | true     | 1000          |

## Destination

A destination connector pushes data from upstream resources to an external
resource via Conduit.

### Configuration

| name                       | description                                | required | default value |
|----------------------------|--------------------------------------------|----------|---------------|
| `destination_config_param` | Description of `destination_config_param`. | true     | 1000          |

## Testing

Run `make test` to run all the unit tests. Run `make test-integration` to run
the integration tests.

The Docker compose file at `test/docker-compose.yml` can be used to run the
required resource locally.

## How to release?

The release is done in two steps:

- Bump the version in [connector.yaml](/connector.yaml). This can be done
  with [bump_version.sh](/scripts/bump_version.sh) script, e.g.
  `scripts/bump_version.sh 2.3.4` (`2.3.4` is the new version and needs to be a
  valid semantic version). This will also automatically create a PR for the
  change.
- Tag the connector, which will kick off a release. This can be done
  with [tag.sh](/scripts/tag.sh).

## Known Issues & Limitations

- Known issue A
- Limitation A

## Planned work

- [ ] Item A
- [ ] Item B
