# netconf-tool-go
Work in progress to make a Go based version of [netconf-tool (Python)](https://github.com/BSpendlove/netconf-tool) as a project to learn more Go!

### Usage

Building:

```
git clone https://github.com/BSpendlove/netconf-tool-go
cd netconf-tool-go
go build netconf-tool/netconftool.go
```

Auto-completion:

- Bash
`netconf-tool completion bash > /etc/bash_completion.d/netconf-tool`
- For other shells, see `./netconftool completion bash --help` for supported auto-completion scripts that Cobra will generate for you

```
$ ./netconftool --help
netconf-tool is a Go based port of my original Python CLI tool netconf-tool
which uses the click module.

The idea of this tool is to provide some basic NETCONF functionality on demand so you can
interact with a NETCONF server on the fly with basic operations instead of having to write
temporary code for example to gather a NETCONF subscription and event data, or grab the running
configuration to then store and parse locally for offline development.

Usage:
  netconf-tool [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  operations   Perform various NETCONF Operation commands
  subscription Perform various NETCONF Subscription commands

Flags:
  -h, --help   help for netconf-tool

Use "netconf-tool [command] --help" for more information about a command.
```

## Feature Examples

### netconf-tool operations list-server-capabilities

Will list the NETCONF capabilities of the target server, you can export the URIs to a json file to use offline or just print them to stdout in string format, or in JSON format.

```
./netconftool operations list-server-capabilities --host 172.20.20.3 --username clab --password clab@123 --printcli --exportjson device_capabilities.json
(0) http://cisco.com/ns/yang/Cisco-IOS-XR-um-cdp-cfg?module=Cisco-IOS-XR-um-cdp-cfg&revision=2022-07-11
(...) ......
(730) http://cisco.com/ns/yang/cisco-xr-openconfig-network-instance-deviations?module=cisco-xr-openconfig-network-instance-deviations&revision=2021-08-20
2023/12/18 04:17:30 exported 731 NETCONF server capabilities to device_capabilities.json
```