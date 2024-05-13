syslog-to-journald
==============

forward syslog from network (514/UDP) to systemd-journald

Requirements
------------

To compile and run `syslog-to-journald` you need:

* [systemd](https://www.github.com/systemd/systemd)
* [go](https://golang.org)
* escalated privileges

Build and install
-----------------

Building and installing is very easy. Just run:

> make && make install

This will place an executable at `/usr/local/bin/syslog-to-journald`.
Additionally, a systemd unit file is installed to
`/etc/systemd/system/syslog-to-journald.service`.

Usage
-----

Just run `/usr/local/bin/syslog-to-journald` or start a systemd unit with
`systemctl enable --now syslog-to-journald.service`. Make sure UDP port 514 is not blocked
in your firewall or used by any service.

Use `journalctl` to view the logs:

    $ journalctl -u syslog-to-journald
    Jun 07 08:15:22 server 10.0.0.1[548]: dhcp,info mikrotik1: intern assigned 10.0.0.50 to 00:11:22:33:44:55
    Jun 07 09:16:59 server 10.0.0.1[548]: interface,info mikrotik1: en7 link down
    Jun 07 09:17:17 server 10.0.0.1[548]: interface,info mikrotik1: en7 link up (speed 100M, full duplex)
    Jun 07 10:07:16 server 10.1.1.1[548]: wireless,info mikrotik2: 00:11:22:33:44:66@wl2-guest: connected, signal strength -36
    Jun 07 10:07:21 server 10.1.1.1[548]: dhcp,info mikrotik2: guest assigned 192.168.1.50 to 00:11:22:33:44:66

Filtering is available with matching `SYSLOG_IDENTIFIER` the ip address:

    $ journalctl -u syslog-to-journald SYSLOG_IDENTIFIER=10.1.1.1
    Jun 07 10:07:16 server 10.1.1.1[548]: wireless,info mikrotik2: 00:11:22:33:44:66@wl2-guest: connected, signal strength -36
    Jun 07 10:07:21 server 10.1.1.1[548]: dhcp,info mikrotik2: guest assigned 192.168.1.50 to 00:11:22:33:44:66

License and warranty
--------------------

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
[GNU General Public License](LICENSE) for more details.

Special thanks
--------------

[udp514-journal](https://git.eworm.de/cgit.cgi/udp514-journal/) project for almost everything.

Source code
-----------

[gitlab](https://gitlab.com/rsipos/syslog-to-journald)
[github mirror](https://github.com/rsipos/syslog-to-journald)
