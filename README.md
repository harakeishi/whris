![logo](logo.png)


"whereis" is Displays management information for IPs associated with the domain.

[![GitHub release](https://img.shields.io/github/release/harakeishi/whereis.svg)](https://github.com/harakeishi/whereis/releases)[![Go Report Card](https://goreportcard.com/badge/github.com/harakeishi/whereis)](https://goreportcard.com/report/github.com/harakeishi/whereis)
# Table of Contents
- [DEMO](#demo)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)
- [Contributing](#contributing)
# DEMO
 
```bash
$ whereis exmple.com      
Target domain:exmple.com
Target ip    :67.210.233.131

Network Admin:Globalcon.net, LLC (GLOBA-72)
Network name :GLOBALCON-WEBHOST-NET1
ip range     :67.210.233.128 - 67.210.233.255
country      :
```
 
 ```bash
$ whereis yaserarenai.com 
Target domain:yaserarenai.com
Target ip    :163.44.185.212

Network Admin:GMO Pepabo, Inc.
Network name :LOLIPOP
ip range     :163.44.185.0 - 163.44.185.255
country      :JP
```
# Installation
 
```bash
$ go install github.com/harakeishi/whereis@latest
```
 
# Usage
 
If you want to know the administrator of the IP associated with the domain, type `whereis [target]`.
 
 You will then see the target domain and IP, and the administrator information for that IP (administrator name, network name, range of IPs to be managed, and country).
 
```bash
$ whereis yaserarenai.com 
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
$ whereis yaserarenai.com -v
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
