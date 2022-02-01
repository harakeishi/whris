![logo](logo.png)


"whris" is Displays management information for IPs associated with the domain.

[![GitHub release](https://img.shields.io/github/release/harakeishi/whris.svg)](https://github.com/harakeishi/whris/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/harakeishi/whris)](https://goreportcard.com/report/github.com/harakeishi/whris) [![Test](https://github.com/harakeishi/whris/actions/workflows/test.yml/badge.svg)](https://github.com/harakeishi/whris/actions/workflows/test.yml)
# Table of Contents
- [Table of Contents](#table-of-contents)
- [DEMO](#demo)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)
- [Contributing](#contributing)
# DEMO
 
```bash
$ whris example.com      
Target domain:example.com
Target ip    :93.184.216.34

Network Admin:NETBLK-03-EU-93-184-216-0-24
Network name :EDGECAST-NETBLK-03
ip range     :93.184.216.0 - 93.184.216.255
country      :EU
```
 
 ```bash
$ whris yaserarenai.com 
Target domain:yaserarenai.com
Target ip    :163.44.185.212

Network Admin:GMO Pepabo, Inc.
Network name :LOLIPOP
ip range     :163.44.185.0 - 163.44.185.255
country      :JP
```
# Installation
 
```bash
$ go install github.com/harakeishi/whris@latest
```
 
# Usage
 
If you want to know the administrator of the IP associated with the domain, type `whris [target]`.
 
 You will then see the target domain and IP, and the administrator information for that IP (administrator name, network name, range of IPs to be managed, and country).
 
```bash
$ whris yaserarenai.com 
Target domain:yaserarenai.com
Target ip    :163.44.185.212

Network Admin:GMO Pepabo, Inc.
Network name :LOLIPOP
ip range     :163.44.185.0 - 163.44.185.255
country      :JP
```

If you want to see more details, use the `-v` option.

Then you will be able to see the higher level network administrator information.

```bash
$ whris yaserarenai.com -v
Target domain:yaserarenai.com
Target ip    :163.44.185.212

Network Admin:GMO Pepabo, Inc.
Network name :LOLIPOP
ip range     :163.44.185.0 - 163.44.185.255
country      :JP

=========Network Details=========
0:
 Network Admin:Administered by APNIC
 Network name :
 ip range     :163.0.0.0 - 163.255.255.255
 country      :
1:
 Network Admin:GMO Internet, Inc.
 Network name :interQ
 ip range     :163.44.64.0 - 163.44.191.255
 country      :JP
2:
 Network Admin:GMO Pepabo, Inc.
 Network name :LOLIPOP
 ip range     :163.44.185.0 - 163.44.185.255
 country      :JP
```
# License
[MIT](LICENSE)

This software includes the work that is distributed in the Apache License 2.0.

このソフトウェアは、 Apache 2.0ライセンスで配布されている製作物が含まれています。

# Contributing
Feel free to open issues
