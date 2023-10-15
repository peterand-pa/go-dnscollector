
# Logger: Syslog

Syslog logger to local syslog system or remote one.

* local or remote server
* custom text format
* supported format: text, json
* tls support

Options:

* `facility`: (string) Set the syslog logging facility
* `transport`: (string) Transport to use to a remote log daemon or local one. local|tcp|udp|unix|tcp+tls
* `remote-address`: (string) Remote address host:port
* `retry-interval`: (integer) interval in second between retry reconnect
* `mode`: (string) output format: text, json, or flat-json
* `text-format`: (string) output text format, please refer to the default text format to see all available directives, use this parameter if you want a specific format
* `tls-insecure`: (boolean) insecure skip verify
* `tls-min-version`: (string) min tls version, default to 1.2
* `format`: (string) Set syslog formatter between `unix` (default), [`rfc3164`](https://www.rfc-editor.org/rfc/)rfc3164 or [`rfc5424`](https://www.rfc-editor.org/rfc/rfc5424) or [`rfc5425`](https://www.rfc-editor.org/rfc/rfc5425)
* `chan-buffer-size`: (integer) channel buffer size used on incoming dns message, number of messages before to drop it.
* `tag`: (string) syslog tag or MSGID

Default values:

```yaml
syslog:
  severity: INFO
  facility: DAEMON
  transport: local
  remote-address: ""
  retry-interval: 10
  text-format: ""
  mode: text
  tls-insecure: false
  tls-min-version: 1.2
  format: ""
  chan-buffer-size: 65535
  tag: ""
```