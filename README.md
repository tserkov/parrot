<p align="center"><img src="https://github.com/tserkov/parrot/blob/master/app/src/assets/parrot.svg" width="50%"></p>

[![GoDoc](https://godoc.org/github.com/tserkov/parrot?status.svg)](https://godoc.org/github.com/tserkov/parrot)

# Parrot
:bird: Parrot is a syslog server with a real-time web dashboard.

:zap: Parrot supports RFC3164, RFC5424, and RFC6587 over TCP, UDP, and/or Unix Sockets.

:hammer: Support for forwarding received logs to another syslog server is planned in a future version.

## Installation
:arrow_double_down: Download the [latest release](https://github.com/mcuadros/go-syslog/releases) and extract the appropriate binary to wherever you'd like (renaming it to `parrot` and placing it in `/usr/local/bin` seems nice for *nix systems).

## Usage
:computer: Just call the binary (`parrot`?) with the appropriate arguments.

### Arguments
- `-forward` _[tcp | udp | unix]://[host:port | path]_
  - forward received logs to the specified syslog server ___(not yet implemented)___
- `-listen` _[tcp | udp | unix]://[host:port | path]_
  - listen for syslog messages on the specified location _(required)_
  - specify multiple times for multiple listeners!
- `-silent` _(defaults to `false`)_
  - specify to disable info-level logging
- `-web` _ip:port_
  - a `host:port` for the dashboard webserver to serve from _(defaults to `127.0.0.1:8080`)_

## Examples
- Listen for syslog messages on the same server sent to standard port 514 over UDP:
  - `parrot -listen udp://127.0.0.1:514`
- Listen for syslog messages from any interface sent to standard port 514 over TCP:
  - `parrot -listen tcp://0.0.0.0:514`
- Listen for syslog messages from a Unix socket:
  - `parrot -listen unix:///var/tmp/parrot_syslog.sock`
- Listen for syslog messages from any interface sent to TCP port 514 and UDP port 515:
  - `parrot -listen tcp://0.0.0.0:514 -listen udp://0.0.0.0:515`
- Listen for syslog messages on the same server sent to standard port 514 over UDP, while serving the web dashboard over any interface on port 8080:
  - `parrot -listen tcp://127.0.0.1:514 -web 0.0.0.0:8080`

## Technology
Parrot is written in Golang 1.11 for the backend syslog listener, forward, and webserver.  VueJS and Bulma are used for the dashboard, and are embedded into the executables.

## TODO
- Tests
- Support for forwarding
- OSX support (untested; built for linux-amd64 only at the moment)

## License
Parrot is licensed GNU GPL v3.

Parrot icon by [Freepik](http://www.freepik.com "Freepik") from [flaticon.com](https://www.flaticon.com/ "Flaticon") is licensed by [CC BY 3.0](http://creativecommons.org/licenses/by/3.0/ "Creative Commons BY 3.0")