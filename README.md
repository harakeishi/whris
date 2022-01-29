![logo](logo.png)
# whereis
"whereis" is Displays management information for IPs associated with the domain.
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
$ go install github.com/harakeishi/whereis
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
 
# License
[MIT](LICENSE)

This software includes the work that is distributed in the Apache License 2.0.

このソフトウェアは、 Apache 2.0ライセンスで配布されている製作物が含まれています。