# xray-lite

A lean fork of [Xray-core](https://github.com/XTLS/Xray-core) that keeps only the
**VLESS** family (plus SOCKS) and strips everything else. The release binary is
about 40% smaller than stock Xray-core (19.8 MB vs 32.9 MB, `-ldflags="-s -w"`).

This is an unofficial fork. It is not affiliated with Project X. All credit for the
protocol and the code goes to the upstream authors; see [LICENSE](LICENSE) (MPL-2.0).

## What is included

| Area | Kept |
|---|---|
| Proxies | VLESS inbound/outbound (Vision flow, post-quantum `encryption`, XUDP, reverse), SOCKS (inbound also speaks HTTP, as upstream), freedom, blackhole |
| Transports | RAW (TCP), XHTTP (including H3), UDP |
| Security | TLS (uTLS, ECH), REALITY |
| Core apps | dispatcher, proxyman, router (random / round-robin balancers), DNS, policy, stats, log, geodata |
| Config | JSON only |
| CLI | `run`, `version`, `uuid`, `x25519`, `mldsa65`, `mlkem768`, `vlessenc`, `tls` |

## What was removed

VMess, Trojan, Shadowsocks (+2022), WireGuard, TUN, MASQUE, Hysteria, Dokodemo, HTTP
proxy as a standalone protocol, loopback and DNS outbounds; the mKCP, gRPC, WebSocket,
HTTPUpgrade, XDrive transports; the browser dialer; finalmask packet masks (only
`quicParams` remains); observatory, the least-ping / least-load balancers, metrics,
the gRPC command API; TOML and YAML configs, FakeDNS, and the API/convert/wg commands.

Configs that reference a removed protocol or transport fail at load time.

## Build

```sh
go build -trimpath -ldflags="-s -w" -o xray ./main
./xray run -c config.json
```

Integration tests in `testing/scenarios` use `dokodemo-door` as a local entry point and
build the binary with `-tags xraytest`; that tag is not part of the shipped binary.
