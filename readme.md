[![Go Report Card](https://goreportcard.com/badge/github.com/gtrafimenkov/hetzctl)](https://goreportcard.com/report/github.com/gtrafimenkov/hetzctl)

# Hetzner Server Management Tool

hetzctl is a CLI tool for managing Hetzner servers.

## Usage

```bash

$ hetzctl server list
+-------------+--------+-----------------+------------+-----------+---------+----------+--------+-----------+-----------+------------+
| SERVER NAME | NUMBER |       IP        |  PRODUCT   |    DC     | TRAFFIC | FLATRATE | STATUS | THROTTLED | CANCELLED | PAID UNTIL |
+-------------+--------+-----------------+------------+-----------+---------+----------+--------+-----------+-----------+------------+
| s1          | 111111 | 144.44.444.114  | PX70       | FSN1-DC11 | 30 TB   | true     | ready  | false     | false     | 2018-01-14 |
| s2          | 111112 | 138.138.138.138 | PX121-SSD  | FSN1-DC8  | 50 TB   | true     | ready  | false     | false     | 2018-01-10 |
+-------------+--------+-----------------+------------+-----------+---------+----------+--------+-----------+-----------+------------+
```

## Configuration

You need to provide username and password for accessing [Hetzner API](https://robot.your-server.de/doc/webservice/en.html).

There are few methods to do so:
  - using configuration file `~/.hetzctl.yaml`
  - using environment variables `HETZCTL_USERNAME` and `HETZCTL_PASSWORD`
  - using command line arguments `--username` and `--password`

### Configuration file format

```
username: foo
password: bar
```

## Changelog

[changelog.md](changelog.md)

## License

MIT
