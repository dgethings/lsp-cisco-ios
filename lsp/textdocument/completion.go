package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var items = []protocol.CompletionItem{

		{
			Label: "activation-character ascii-number",
		},
		{
			Label: "no activation-character",
		},
		{
			Label: "alias mode command-alias original-command",
		},
		{
			Label: "no alias mode [command-alias]",
		},
		{
			Label: "archive",
		},
		{
			Label: "archive config",
		},
		{
			Label: "archive log config persistent save",
		},
		{
			Label: "archive tar {/create destination-urlflash:/ file-url | /table source-url | /xtract source-urlflash:/file-url [dir/file...]}",
		},
		{
			Label: "async-bootp tag [:hostname] data",
		},
		{
			Label: "no async-bootp",
		},
		{
			Label: "attach slot-number",
		},
		{
			Label: "attach module-number",
		},
		{
			Label: "autobaud",
		},
		{
			Label: "no autobaud",
		},
		{
			Label: "auto-sync {startup-config | config-register | bootvar | running-config | standard}",
		},
		{
			Label: "no auto-sync {startup-config | config-register | bootvar | standard}",
		},
		{
			Label: "autoupgrade disk-cleanup [crashinfo | core | image | irrecoverable]",
		},
		{
			Label: "no autoupgrade disk-cleanup [crashinfo | core | image | irrecoverable]",
		},
		{
			Label: "autoupgrade ida url url",
		},
		{
			Label: "no autoupgrade ida url url",
		},
		{
			Label: "autoupgrade status email [recipient [email-address] ] [smtp-server [smtp-server] ]",
		},
		{
			Label: "no autoupgrade status email [recipient [email-address] ] [smtp-server [smtp-server] ]",
		},
		{
			Label: "banner exec d message d",
		},
		{
			Label: "no banner exec",
		},
		{
			Label: "banner incoming d message d",
		},
		{
			Label: "no banner incoming",
		},
		{
			Label: "banner login d message d",
		},
		{
			Label: "no banner login",
		},
		{
			Label: "banner motd d message d",
		},
		{
			Label: "no banner motd",
		},
		{
			Label: "banner slip-ppp d message d",
		},
		{
			Label: "no banner slip-ppp",
		},
		{
			Label: "boot",
		},
		{
			Label: "boot file-url",
		},
		{
			Label: "boot filename [tftp-ip-address]",
		},
		{
			Label: "boot flash [flash-fs: ] [partition-number: ] [filename]",
		},
		{
			Label: "boot flash-fs: [filename]",
		},
		{
			Label: "boot [flash-fs: ] [partition-number: ] [filename]",
		},
		{
			Label: "boot usbflash0 [: filename]",
		},
		{
			Label: "boot bootldr file-url boot bootldr command",
		},
		{
			Label: "no boot bootldr",
		},
		{
			Label: "boot bootstrap file-url",
		},
		{
			Label: "no boot bootstrap file-url",
		},
		{
			Label: "boot bootstrap flash [filename]",
		},
		{
			Label: "no boot bootstrap flash [filename]",
		},
		{
			Label: "boot bootstrap [tftp] filename [ip-address]",
		},
		{
			Label: "no boot bootstrap [tftp] filename [ip-address]",
		},
		{
			Label: "boot bootstrap mop filename [interface-type interface-number]",
		},
		{
			Label: "no boot bootstrap mop filename [interface-type interface-number]",
		},
		{
			Label: "boot config file-system-prefix: [directory/] filename [nvbypass]",
		},
		{
			Label: "no boot config",
		},
		{
			Label: "boot config device:filename [nvbypass]",
		},
		{
			Label: "no boot config",
		},
		{
			Label: "boot host commandboot host remote-url",
		},
		{
			Label: "no boot host remote-url",
		},
		{
			Label: "boot network remote-url",
		},
		{
			Label: "no boot network remote-url",
		},
		{
			Label: "boot system {file-url | filename}",
		},
		{
			Label: "no boot system {file-url | filename}",
		},
		{
			Label: "boot system flash [flash-fs:] [partition-number:] [filename]",
		},
		{
			Label: "no boot system flash [flash-fs:] [partition-number:] [filename]",
		},
		{
			Label: "boot system mop filename [mac-address] [interface]",
		},
		{
			Label: "no boot system mop filename [mac-address] [interface]",
		},
		{
			Label: "boot system rom",
		},
		{
			Label: "no boot system rom",
		},
		{
			Label: "boot system {rcp | tftp | ftp} filename [ip-address]",
		},
		{
			Label: "no boot system {rcp | tftp | ftp} filename [ip-address]",
		},
		{
			Label: "cd [filesystem:] [directory]",
		},
		{
			Label: "clear archive log config [force | persistent]",
		},
		{
			Label: "clear catalyst6000 traffic-meter",
		},
		{
			Label: "clear configuration lock",
		},
		{
			Label: "clear diagnostic event-log {event-type {error | info | warning} | module {num | slot subslot | all}}",
		},
		{
			Label: "clear ip http client cache {all | session session-name | url complete-url}",
		},
		{
			Label: "clear logging [persistent [url filesystem:/directory]]",
		},
		{
			Label: "clear logging system [disk name]",
		},
		{
			Label: "clear logging xml",
		},
		{
			Label: "clear memory low-water-mark",
		},
		{
			Label: "clear mls statistics [module num]",
		},
		{
			Label: "clear parser cache",
		},
		{
			Label: "clear parser statistics",
		},
		{
			Label: "clear platform netint",
		},
		{
			Label: "clear processes interrupt mask detail",
		},
		{
			Label: "clear scp accounting",
		},
		{
			Label: "clear tcp {line line-number | local hostname port remote hostname port | tcb address}",
		},
		{
			Label: "clear vlan [vlan-id] counters",
		},
		{
			Label: "clock {auto | active [prefer] | passive [prefer]}",
		},
		{
			Label: "no clock",
		},
		{
			Label: "clock initialize nvram",
		},
		{
			Label: "no clock initialize nvram",
		},
		{
			Label: "config-register value",
		},
		{
			Label: "configure check syntax [source-location]",
		},
		{
			Label: "configuration mode exclusive {auto | manual} [expire seconds] [lock-show] [interleave] [terminate] [config-wait seconds] [retry-wait seconds]",
		},
		{
			Label: "no configuration mode exclusive",
		},
		{
			Label: "configure confirm",
		},
		{
			Label: "configure memory",
		},
		{
			Label: "configure replace target-url [nolock] list force ignorecase [revert trigger [error] [timer minutes] | time minutes]",
		},
		{
			Label: "configure revert {now | timer {minutes | idle minutes}}",
		},
		{
			Label: "configure terminal",
		},
		{
			Label: "configure terminal [lock]",
		},
		{
			Label: "configure terminal [revert {timer minutes | idle minutes}]",
		},
		{
			Label: "confreg [value]",
		},
		{
			Label: "continue",
		},
		{
			Label: "copy [/erase] [/ verify | /noverify] source-url destination-url",
		},
		{
			Label: "copy logging system target: filename",
		},
		{
			Label: "no copy logging system",
		},
		{
			Label: "copy xmodem: flashfilesystem:",
		},
		{
			Label: "copy ymodem: flashfilesystem:",
		},
		{
			Label: "copy /noverify source-url destination-url",
		},
		{
			Label: "databits {5 | 6 | 7 | 8}",
		},
		{
			Label: "no databits",
		},
		{
			Label: "data-character-bits {7 | 8}",
		},
		{
			Label: "no data-character-bits",
		},
		{
			Label: "default-value data-character-bits {7 | 8}",
		},
		{
			Label: "no default-value data-character-bits",
		},
		{
			Label: "default-value exec-character-bits {7 | 8}",
		},
		{
			Label: "no default-value exec-character-bits",
		},
		{
			Label: "default-value modem-interval milliseconds",
		},
		{
			Label: "no default-value modem-interval",
		},
		{
			Label: "default-value special-character-bits commanddefault-value special-character-bits {7 | 8}",
		},
		{
			Label: "no default-value special-character-bits",
		},
		{
			Label: "define interface-range macro-name interface-range",
		},
		{
			Label: "delete url [/force | /recursive]",
		},
		{
			Label: "diag command diag slot-number [halt | previous | post | verbose [wait] | wait]",
		},
		{
			Label: "no diag slot-number",
		},
		{
			Label: "diagnostic bootup level {minimal | complete}",
		},
		{
			Label: "no diagnostic bootup level",
		},
		{
			Label: "diagnostic cns {publish | subscribe} [subject]",
		},
		{
			Label: "no diagnostic cns {publish | subscribe} [subject]",
		},
		{
			Label: "diagnostic event-log size size",
		},
		{
			Label: "no diagnostic event-log size",
		},
		{
			Label: "diagnostic level {power-on | bypass}",
		},
		{
			Label: "diagnostic monitor intervalmodule number test {test-id | test-id-range | all} hh:mm:ss milliseconds days",
		},
		{
			Label: "diagnostic monitor syslog",
		},
		{
			Label: "diagnostic monitor module num test {test-id | test-id-range | all}",
		},
		{
			Label: "no diagnostic monitor {interval | syslog}",
		},
		{
			Label: "diagnostic monitor {bay slot/ bay | slot slot number | subslot slot/ subslot} test {test-id | test-id-range | all}",
		},
		{
			Label: "diagnostic monitor interval {bay slot/ bay | slot slot-no | subslot slot/ subslot} test {test-id | test-id-range | all} hh:mm:ss milliseconds days",
		},
		{
			Label: "diagnostic monitor syslog",
		},
		{
			Label: "diagnostic monitor threshold {bay slot/ bay | slot slot-no | subslot slot/ subslot} test {test-id | test-id-range | all} failure count failures [ {runs | days | hours | minutes | seconds | milliseconds} window_size]",
		},
		{
			Label: "diagnostic ondemand {iteration iteration-count | action-on-failure {continue error-count | stop}}",
		},
		{
			Label: "diagnostic schedule module {module-number | slot/subslot} test {test-id | all | complete | minimal | non-disruptive | [per-port [port {interface-port-number | port-number-list | all}]]} {on month dd yyyy hh:mm | daily hh:mm | weekly day-of-week hh:mm}",
		},
		{
			Label: "no diagnostic schedule module {module-number | slot/subslot} test {test-id | all | complete | minimal | non-disruptive | [per-port [port {interface-port-number | port-number-list | all}]]} {on month dd yyyy hh:mm | daily hh:mm | weekly day-of-week hh:mm}",
		},
		{
			Label: "diagnostic start module num test {test-id | test-id-range | minimal | complete | basic | per-port | non-disruptive | all} [port {num | port#-range | all}]",
		},
		{
			Label: "diagnostic start system test all",
		},
		{
			Label: "diagnostic start {bay slot/ bay | slot slot-no} test {test-id | test-id-range | all | complete | minimal | non-disruptive}",
		},
		{
			Label: "diagnostic start subslot slot/subslot test {test-id | test-id-range | all | complete | minimal | non-disruptive | [per-port [port {num | port#-range | all}]]}",
		},
		{
			Label: "diagnostic stop module num",
		},
		{
			Label: "diagnostic stop {bay slot/ bay | slot slot-no | subslot slot/ subslot}",
		},
		{
			Label: "dir [/ all] [/ recursive] [all-filesystems] [filesystem: ] [file-url]",
		},
		{
			Label: "disable [privilege-level]",
		},
		{
			Label: "disconnect-character ascii-number",
		},
		{
			Label: "no disconnect-character",
		},
		{
			Label: "dispatch-character ascii-number1 [ascii-number2. . . ]",
		},
		{
			Label: "no dispatch-character ascii-number1 [ascii-number2. . . ]",
		},
		{
			Label: "dispatch-machine name",
		},
		{
			Label: "no dispatch-machine",
		},
		{
			Label: "dispatch-timeout milliseconds",
		},
		{
			Label: "no dispatch-timeout",
		},
		{
			Label: "do command",
		},
		{
			Label: "downward-compatible-config version",
		},
		{
			Label: "no downward-compatible-config",
		},
		{
			Label: "editing",
		},
		{
			Label: "no editing",
		},
		{
			Label: "enable [privilege-level] [view [view-name] ]",
		},
		{
			Label: "no enable last-resort {password | succeed}",
		},
		{
			Label: "no enable last-resort",
		},
		{
			Label: "end",
		},
		{
			Label: "environment-monitor shutdown temperature [rommon | powerdown]",
		},
		{
			Label: "no environment-monitor shutdown temperature [rommon | powerdown]",
		},
		{
			Label: "environment temperature-controlled",
		},
		{
			Label: "no environment temperature-controlled",
		},
		{
			Label: "erase {/all nvram: | /no-squeeze-reserve-space filesystem: | filesystem: | startup-config}",
		},
		{
			Label: "erase {/all nvram: | filesystem: | startup-config}",
		},
		{
			Label: "errdisable detect cause {all | bpduguard | dtp-flap | l2ptguard | link-flap | packet-buffer-error | pagp-flap | rootguard | udld}",
		},
		{
			Label: "no errdisable detect cause {all | bpduguard | dtp-flap | l2ptguard | link-flap | pagp-flap | rootguard | udld}",
		},
		{
			Label: "errdisable recovery {cause {all | arp-inspection | bpduguard | channel-misconfig | dhcp-rate-limit | dtp-flap | gbic-invalid | l2ptguard | link-flap | pagp-flap | psecure-violation | security-violation | rootguard | udld | unicast-flood} | interval seconds}",
		},
		{
			Label: "no errdisable recovery {cause {all | arp-inspection | bpduguard | channel-misconfig | dhcp-rate-limit | dtp-flap | gbic-invalid | l2ptguard | link-flap | pagp-flap | psecure-violation | security-violation | rootguard | udld | unicast-flood} | interval seconds}",
		},
		{
			Label: "escape-character {break | char | default | none | soft}",
		},
		{
			Label: "no escape-character [soft]",
		},
		{
			Label: "default escape-character [soft]",
		},
		{
			Label: "exec",
		},
		{
			Label: "no exec",
		},
		{
			Label: "exec-banner",
		},
		{
			Label: "no exec-banner",
		},
		{
			Label: "exec-character-bits {7 | 8}",
		},
		{
			Label: "no exec-character-bits",
		},
		{
			Label: "exec-timeout minutes [seconds]",
		},
		{
			Label: "no exec-timeout",
		},
		{
			Label: "execute-on {slot slot-number | all | master} command",
		},
		{
			Label: "exit",
		},
		{
			Label: "exit",
		},
		{
			Label: "file privilege level level",
		},
		{
			Label: "no file privilege level level",
		},
		{
			Label: "file prompt prompt [alert | noisy | quiet]",
		},
		{
			Label: "file verify auto",
		},
		{
			Label: "no file verify auto",
		},
		{
			Label: "format filesystem1:",
		},
		{
			Label: "format [spare spare-number] filesystem1: [ [filesystem2:] [monlib-filename] ]",
		},
		{
			Label: "fsck [/ nocrc] [/ automatic] [/ all] [/ force] [filesystem:]",
		},
		{
			Label: "fsck [/ automatic] [/ all] [/ force] [filesystem:]",
		},
		{
			Label: "fsck [/ all] [/ force] [filesystem:]",
		},
		{
			Label: "full-help",
		},
		{
			Label: "help",
		},
		{
			Label: "hidekeys",
		},
		{
			Label: "no hidekeys",
		},
		{
			Label: "history",
		},
		{
			Label: "no history",
		},
		{
			Label: "history size number-of-lines",
		},
		{
			Label: "no history size",
		},
		{
			Label: "hold-character ascii-number",
		},
		{
			Label: "no hold-character",
		},
		{
			Label: "hostname name",
		},
		{
			Label: "hw-module module num reset",
		},
		{
			Label: "hw-module module num shutdown",
		},
		{
			Label: "insecure",
		},
		{
			Label: "no insecure",
		},
		{
			Label: "install {abort | activate | file {bootflash: | flash: | harddisk: | webui:} [auto-abort-timer timer timer prompt-level {all | none}] | add file {bootflash: | flash: | ftp: | harddisk: | http: | https: | pram: | rcp: | scp: | tftp: | webui:} [activate [auto-abort-timer timer prompt-level {all | none}commit]] | commit | auto-abort-timer stop | deactivate file {bootflash: | flash: | harddisk: | webui:} | label id {description description | label-name name} | remove {file {bootflash: | flash: | harddisk: | webui:} | inactive } | | rollback to {base | committed | id {install-ID } | label {label-name}}}",
		},
		{
			Label: "international",
		},
		{
			Label: "no international",
		},
		{
			Label: "ip bootp server",
		},
		{
			Label: "no ip bootp server",
		},
		{
			Label: "ip finger [rfc-compliant]",
		},
		{
			Label: "no ip finger",
		},
		{
			Label: "ip ftp passive",
		},
		{
			Label: "no ip ftp passive",
		},
		{
			Label: "ip ftp password [type] password",
		},
		{
			Label: "no ip ftp password",
		},
		{
			Label: "ip ftp source-interface interface-type interface-number",
		},
		{
			Label: "no ip ftp source-interface",
		},
		{
			Label: "ip ftp username username",
		},
		{
			Label: "no ip ftp username",
		},
		{
			Label: "ip rarp-server ip-address",
		},
		{
			Label: "no ip rarp-server ip-address",
		},
		{
			Label: "ip rcmd domain-lookup",
		},
		{
			Label: "no ip rcmd domain-lookup",
		},
		{
			Label: "ip rcmd rcp-enable",
		},
		{
			Label: "no ip rcmd rcp-enable",
		},
		{
			Label: "ip rcmd remote-host local-username {ip-address | host-name} remote-username [enable [level] ]",
		},
		{
			Label: "no ip rcmd remote-host local-username {ip-address | host-name} remote-username [enable [level] ]",
		},
		{
			Label: "ip rcmd remote-username username",
		},
		{
			Label: "no ip rcmd remote-username username",
		},
		{
			Label: "ip rcmd rsh-enable",
		},
		{
			Label: "no ip rcmd rsh-enable",
		},
		{
			Label: "ip rcmd source-interface interface-id",
		},
		{
			Label: "no ip rcmd source-interface interface-id",
		},
		{
			Label: "ip telnet source-interface interface",
		},
		{
			Label: "no ip telnet source-interface",
		},
		{
			Label: "ip tftp blocksize bytes",
		},
		{
			Label: "no ip tftp blocksize",
		},
		{
			Label: "ip tftp boot-interface type number",
		},
		{
			Label: "no ip tftp boot-interface",
		},
		{
			Label: "ip tftp min-timeout seconds",
		},
		{
			Label: "no ip tftp min-timeout",
		},
		{
			Label: "ip tftp source-interface interface-type interface-number",
		},
		{
			Label: "no ip tftp source-interface",
		},
		{
			Label: "ip wccp web-cache accelerated [ [group-address groupaddress] | [redirect-list access-list] | [group-list access-list] | [password password]]",
		},
		{
			Label: "no ip wccp web-cache accelerated",
		},
		{
			Label: "length screen-length",
		},
		{
			Label: "no length",
		},
		{
			Label: "load-interval seconds",
		},
		{
			Label: "no load-interval seconds",
		},
		{
			Label: "location text",
		},
		{
			Label: "no location",
		},
		{
			Label: "lock",
		},
		{
			Label: "lockable",
		},
		{
			Label: "no lockable",
		},
		{
			Label: "log config",
		},
		{
			Label: "logging buffered [discriminator discriminator-name] [buffer-size] [severity-level]",
		},
		{
			Label: "no logging buffered",
		},
		{
			Label: "default logging buffered",
		},
		{
			Label: "logging buginf",
		},
		{
			Label: "no logging buginf",
		},
		{
			Label: "logging enable",
		},
		{
			Label: "no logging enable",
		},
		{
			Label: "logging esm config",
		},
		{
			Label: "no logging esm config",
		},
		{
			Label: "logging event bundle-status",
		},
		{
			Label: "no logging event bundle-status",
		},
		{
			Label: "logging event link-status {default | boot}",
		},
		{
			Label: "no logging event link-status {default | boot}",
		},
		{
			Label: "logging event link-status [bchan | dchan | nfas]",
		},
		{
			Label: "no logging event link-status [bchan | dchan | nfas]",
		},
		{
			Label: "logging event subif-link-status",
		},
		{
			Label: "no logging event subif-link-status",
		},
		{
			Label: "logging event trunk-status",
		},
		{
			Label: "no logging event trunk-status",
		},
		{
			Label: "logging reload [message-limit number] [severity-level | alerts | critical | debugging | emergencies | errors | informational | notifications | warnings]",
		},
		{
			Label: "no logging reload",
		},
		{
			Label: "logging ip access-list cache {entries entries | interval seconds | rate-limit pps | threshold packets}",
		},
		{
			Label: "no logging ip access-list cache [entries | interval | rate-limit | threshold]",
		},
		{
			Label: "logging ip access-list cache [in | out]",
		},
		{
			Label: "no logging ip access-list cache",
		},
		{
			Label: "logging persistent {auto | manual}",
		},
		{
			Label: "no logging persistent {auto | manual}",
		},
		{
			Label: "logging persistent reload",
		},
		{
			Label: "no logging persistent reload",
		},
		{
			Label: "logging purge-log buffer days number-of-days [ time deletion-start-time ]",
		},
		{
			Label: "no logging purge-log buffer",
		},
		{
			Label: "logging size entries",
		},
		{
			Label: "no logging size",
		},
		{
			Label: "logging synchronous [level severity-level | all] [limit number-of-lines]",
		},
		{
			Label: "no logging synchronous [level severity-level | all] [limit number-of-lines]",
		},
		{
			Label: "logging system [disk name]",
		},
		{
			Label: "no logging system",
		},
		{
			Label: "logout",
		},
		{
			Label: "logout-warning [seconds]",
		},
		{
			Label: "logout-warning",
		},
		{
			Label: "macro {global {apply macro-name | description text | trace macro-name [keyword-to-value] value-first-keyword [keyword-to-value] value-second-keyword [keyword-to-value] value-third-keyword [keyword-to-value] } | name macro-name}",
		},
		{
			Label: "no macro {global {apply macro-name | description text | trace macro-name [keyword-to-value] value-first-keyword [keyword-to-value] value-second-keyword [keyword-to-value] value-third-keyword [keyword-to-value] } | name macro-name}",
		},
		{
			Label: "macro {apply macro-name | description text | trace macro-name [keyword-to-value] value-first-keyword [keyword-to-value] value-second-keyword [keyword-to-value] value-third-keyword [keyword-to-value] }",
		},
		{
			Label: "no macro {apply macro-name | description text | trace macro-name [keyword-to-value] value-first-keyword [keyword-to-value] value-second-keyword [keyword-to-value] value-third-keyword [keyword-to-value] }",
		},
		{
			Label: "maximum number",
		},
		{
			Label: "no maximum number",
		},
		{
			Label: "memory cache error-recovery {L1 | L2 | L3} {data | inst}",
		},
		{
			Label: "no memory cache error-recovery {L1 | L2 | L3} {data | inst}",
		},
		{
			Label: "memory cache error-recovery options {abort-if-same-content | blocking-mode | max-recoveries value | nvram-report | parity-check | window seconds}",
		},
		{
			Label: "no memory cache error-recovery options {abort-if-same-content | blocking-mode | max-recoveries value | nvram-report | parity-check | window seconds}",
		},
		{
			Label: "memory free low-watermark {processor threshold | io threshold}",
		},
		{
			Label: "no memory free low-watermark",
		},
		{
			Label: "memory lite",
		},
		{
			Label: "no memory lite",
		},
		{
			Label: "memory reserve {console size | critical [total-size] }",
		},
		{
			Label: "no memory reserve {console | critical}",
		},
		{
			Label: "memory reserve critical [total-size]",
		},
		{
			Label: "no memory reserve critical",
		},
		{
			Label: "memory reserve critical kilobytes",
		},
		{
			Label: "no memory reserve critical",
		},
		{
			Label: "memory sanity [buffer | queue | all]",
		},
		{
			Label: "no memory sanity",
		},
		{
			Label: "memory scan",
		},
		{
			Label: "no memory scan",
		},
		{
			Label: "memory-size iomem i/o-memory-percentage",
		},
		{
			Label: "no memory-size iomem",
		},
		{
			Label: "menu menu-name",
		},
		{
			Label: "menu menu-name single-space",
		},
		{
			Label: "menu clear-screen menu-name clear-screen",
		},
		{
			Label: "menu command menu menu-name command menu-item {command | menu-exit}",
		},
		{
			Label: "menu menu-name default menu-item",
		},
		{
			Label: "menu menu-name line-mode",
		},
		{
			Label: "menu menu-name options menu-item [login] [pause]",
		},
		{
			Label: "menu menu-name options menu-item {login | pause}",
		},
		{
			Label: "menu menu-name prompt d prompt d",
		},
		{
			Label: "menu menu-name status-line",
		},
		{
			Label: "menu menu-name text menu-item menu-text",
		},
		{
			Label: "menu menu-name title d menu-title d",
		},
		{
			Label: "microcode {oc12-atm | oc12-pos | oc3-pos4} {flash file-id [slot] | system [slot] }",
		},
		{
			Label: "no microcode {oc12-atm | oc12-pos | oc3-pos4} [flash file-id [slot] | system [slot] ]",
		},
		{
			Label: "microcode interface-type {flash-filesystem:filename [slot] | rom | system [slot] }",
		},
		{
			Label: "no microcode interface-type {flash-filesystem:filename [slot] | rom | system [slot] }",
		},
		{
			Label: "microcode {ecpa | pcpa} location",
		},
		{
			Label: "no microcode {ecpa | pcpa}",
		},
		{
			Label: "microcode reload [slot-number]",
		},
		{
			Label: "microcode reload [slot-number]",
		},
		{
			Label: "microcode reload {all | ecpa [slot slot-number] | pcpa [slot slot-number]}",
		},
		{
			Label: "mkdir directory",
		},
		{
			Label: "mkdir disk0:",
		},
		{
			Label: "mode {rpr | rpr-plus | sso}",
		},
		{
			Label: "mode {rpr | sso}",
		},
		{
			Label: "mode sso",
		},
		{
			Label: "monitor event-trace component {clear | continuous | destroy-buffer | disable | dump [pretty] | enable | one-shot}",
		},
		{
			Label: "monitor event-trace component {disable | dump | enable | size | stacktrace}",
		},
		{
			Label: "monitor event-trace all-traces {continuous [cancel] | dump [merged] [pretty]}",
		},
		{
			Label: "monitor event-trace l3 {clear | continuous [cancel] | disable | dump [pretty] | enable | interface type mod/port | one-shot}",
		},
		{
			Label: "monitor event-trace spa {clear | continuous [cancel] | disable | dump [pretty] | enable | one-shot}",
		},
		{
			Label: "monitor event-trace subsys {clear | continuous [cancel] | disable | dump [pretty] | enable | one-shot}",
		},
		{
			Label: "monitor event-trace component {disable | dump-file filename | enable | size number | stacktrace number} timestamps [datetime [localtime] [msec] [show-timezone] | uptime]",
		},
		{
			Label: "monitor event-trace component {disable | dump-file filename | enable | clear | continuous | one-shot}",
		},
		{
			Label: "monitor event-trace crypto pki { error | event | exceptions }",
		},
		{
			Label: "no monitor event-trace crypto pki { error | event | exceptions }",
		},
		{
			Label: "monitor event-trace crypto ipsec { error | event | exceptions }",
		},
		{
			Label: "no monitor event-trace crypto ipsec { error | event | exceptions }",
		},
		{
			Label: "monitor event-trace crypto ikev2 { error | event | exceptions }",
		},
		{
			Label: "no monitor event-trace crypto ikev2 { error | event | exceptions }",
		},
		{
			Label: "monitor event-trace dump-traces [pretty]",
		},
		{
			Label: "monitor event-trace dump-traces [pretty]",
		},
		{
			Label: "monitor event-trace dump-traces [pretty]",
		},
		{
			Label: "monitor event-trace dump-traces [pretty]",
		},
		{
			Label: "monitor event-trace crypto pki { error | event }",
		},
		{
			Label: "no monitor event-trace crypto pki { error | event }",
		},
		{
			Label: "monitor event-trace dump-traces [pretty]",
		},
		{
			Label: "monitor pcm-tracer capture-destination destination",
		},
		{
			Label: "no monitor pcm-tracer capture-destination",
		},
		{
			Label: "monitor pcm-tracer delayed-start seconds",
		},
		{
			Label: "no monitor pcm-tracer delayed-start",
		},
		{
			Label: "monitor pcm-tracer profile profile-number { {no} capture-tdm { [T1 | E1] | {analog-voice-port | bri-voice-port}port | ds0 | channel-numnumber}}",
		},
		{
			Label: "no monitor pcm-tracer profile profile-number",
		},
		{
			Label: "monitor permit-list",
		},
		{
			Label: "no monitor permit-list",
		},
		{
			Label: "monitor permit-list destination interface interface-type slot/port",
		},
		{
			Label: "no monitor permit-list destination interface interface-type slot/port",
		},
		{
			Label: "monitor permit-list destination interface interface-type slot/port-last-port",
		},
		{
			Label: "no monitor permit-list destination interface interface-type slot/port-last-port",
		},
		{
			Label: "monitor permit-list destination interface interface-type slot/port-last-port , [port-last-port]",
		},
		{
			Label: "no monitor permit-list destination interface interface-type slot/port-last-port , [port-last-port]",
		},
		{
			Label: "monitor session egress replication-mode centralized",
		},
		{
			Label: "no monitor session egress replication-mode centralized",
		},
		{
			Label: "monitor session egress replication-mode distributed",
		},
		{
			Label: "no monitor session egress replication-mode distributed",
		},
		{
			Label: "monitor session span-session-number type {erspan-destination | erspan-source | local | local-tx | rspan-destination | rspan-source}",
		},
		{
			Label: "no monitor session span-session-number type {erspan-destination | erspan-source | local | local-tx | rspan-destination | rspan-source}",
		},
		{
			Label: "mop device-code commandmop device-code {cisco | ds200}",
		},
		{
			Label: "no mop device-code {cisco | ds200}",
		},
		{
			Label: "mop retransmit-timer seconds",
		},
		{
			Label: "no mop retransmit-timer",
		},
		{
			Label: "mop retries count",
		},
		{
			Label: "no mop retries",
		},
		{
			Label: "more [/ascii | /binary | /compressed | /ebcdic] url",
		},
		{
			Label: "{more url | begin regular-expression}",
		},
		{
			Label: "{more url | exclude regular-expression}",
		},
		{
			Label: "{more url | include regular-expression}",
		},
		{
			Label: "more flh:logfile",
		},
		{
			Label: "motd-banner",
		},
		{
			Label: "no motd-banner",
		},
		{
			Label: "name-connection",
		},
		{
			Label: "nmsp enable",
		},
		{
			Label: "no nmsp enable",
		},
		{
			Label: "nmsp strong-cipher",
		},
		{
			Label: "no nmsp strong-cipher",
		},
		{
			Label: "no menu menu-name",
		},
		{
			Label: "notify",
		},
		{
			Label: "no notify",
		},
		{
			Label: "notify syslog [contenttype {plaintext | xml}]",
		},
		{
			Label: "no notify syslog [contenttype {plaintext | xml}]",
		},
		{
			Label: "padding ascii-number count",
		},
		{
			Label: "no padding ascii-number",
		},
		{
			Label: "parity {none | even | odd | space | mark}",
		},
		{
			Label: "no parity",
		},
		{
			Label: "parser cache",
		},
		{
			Label: "no parser cache",
		},
		{
			Label: "parser command serializer",
		},
		{
			Label: "no parser command serializer",
		},
		{
			Label: "parser config cache interface",
		},
		{
			Label: "no parser config cache interface",
		},
		{
			Label: "parser config partition",
		},
		{
			Label: "no parser config partition",
		},
		{
			Label: "parser maximum {latencylimit | utilizationlimit}",
		},
		{
			Label: "no parser maximum {latency | utilization}",
		},
		{
			Label: "partition flash-filesystem: [number-of-partitions] [partition-size]",
		},
		{
			Label: "no partition flash-filesystem:",
		},
		{
			Label: "partition flash partitions [size1 size2]",
		},
		{
			Label: "no partition flash",
		},
		{
			Label: "path url",
		},
		{
			Label: "no path url",
		},
		{
			Label: "periodic days-of-the-week hh:mm to [days-of-the-week] hh:mm",
		},
		{
			Label: "no periodic days-of-the-week hh:mm to [days-of-the-week] hh:mm",
		},
		{
			Label: "ping [ [protocol [tag]] {host-name | system-address}]",
		},
		{
			Label: "ping [hostname | system-address | [protocol | tag] {hostname | system-address}] [data [hex-data-pattern] | df-bit | repeat [repeat-count] | size [datagram-size] | source [source-address | async | bvi | ctunnel | dialer | ethernet | fastethernet | lex | loopback | multilink | null | port-channel | tunnel | vif | virtual-template | virtual-tokenring | xtagatm] | timeout [seconds] | validate]",
		},
		{
			Label: "ping ip {host-name | ip-address} [data [hex-data-pattern] | df-bit | repeat [repeat-count] | tos [service value] | size [datagram-size] source {source-address | source-interface}] [timeout seconds] [validate] [verbose]",
		},
		{
			Label: "ping srb name",
		},
		{
			Label: "ping vrf vrf-name [tag] [connection] target-address [connection-options]",
		},
		{
			Label: "platform shell",
		},
		{
			Label: "no platform shell",
		},
		{
			Label: "power enable module slot",
		},
		{
			Label: "no power enable module slot",
		},
		{
			Label: "power redundancy-mode {combined | redundant}",
		},
		{
			Label: "printer printer-name {line number | rotary number} [formfeed] [jobtimeout seconds] [newline-convert] [jobtypes type]",
		},
		{
			Label: "no printer printer-name",
		},
		{
			Label: "private",
		},
		{
			Label: "no private",
		},
		{
			Label: "process cpu statistics limit entry-percentage number [size seconds]",
		},
		{
			Label: "no process cpu statistics limit entry-percentage",
		},
		{
			Label: "process cpu threshold type {total | process | interrupt} rising percentage interval seconds [falling fall-percentage interval seconds]",
		},
		{
			Label: "no process cpu threshold type {total | process | interrupt}",
		},
		{
			Label: "process-max-time milliseconds",
		},
		{
			Label: "no process-max-time milliseconds",
		},
		{
			Label: "prompt string",
		},
		{
			Label: "no prompt [string]",
		},
		{
			Label: "prompt config hostname-length number",
		},
		{
			Label: "no prompt [config]",
		},
		{
			Label: "pwd",
		},
		{
			Label: "refuse-message d message d",
		},
		{
			Label: "no refuse-message",
		},
		{
			Label: "regexp optimize",
		},
		{
			Label: "no regexp optimize",
		},
		{
			Label: "reload [/verify | /noverify] [ [warm file] [line | in [hhh:mm | mmm [text] ] | at hh:mm [day month] [text] ] | reason [reason-string] | cancel]",
		},
		{
			Label: "remote command {module num | standby-rp | switch} command",
		},
		{
			Label: "remote login {module num | standby-rp | switch}",
		},
		{
			Label: "remote-span",
		},
		{
			Label: "no remote-span",
		},
		{
			Label: "rename url1 url2",
		},
		{
			Label: "request consent-token accept-response shell-access response-string",
		},
		{
			Label: "request consent-token generate-challenge shell-access auth-timeout time-validity-slot",
		},
		{
			Label: "request consent-token terminate-auth",
		},
		{
			Label: "request platform software package describe file URL [detail] [verbose]",
		},
		{
			Label: "request platform software package expand file source-URL [to destination-URL] [force] [verbose] [wipe]",
		},
		{
			Label: "request platform software package install rp rp-slot-number commit [verbose]",
		},
		{
			Label: "request platform software package install rp rp-slot-number file file-URL [auto-rollback minutes] [provisioning-file URL] [slot slot-number] [bay bay-number] [force] [on-reboot] [verbose]",
		},
		{
			Label: "request platform software package install rp rp-slot-number rollback [as-booted | provisioning-file provisioning-file-URL] [force] [on-reboot] [verbose]",
		},
		{
			Label: "request platform software package install rp rp-slot-number snapshot to URL [as snapshot-provisioning-filename] [force] [verbose] [wipe]",
		},
		{
			Label: "request platform software process release slot all",
		},
		{
			Label: "request platform software system shell [rp | esp | sip]",
		},
		{
			Label: "request platform software shell session output format format",
		},
		{
			Label: "request platform software snapshot slot {cancel | create | delete | restore} name",
		},
		{
			Label: "request platform software vty attach [permanent]",
		},
		{
			Label: "revision version",
		},
		{
			Label: "no revision",
		},
		{
			Label: "rmdir directory",
		},
		{
			Label: "rommon-pref [readonly | upgrade]",
		},
		{
			Label: "route-converge-interval seconds",
		},
		{
			Label: "no route-converge-interval",
		},
		{
			Label: "rsh {ip-address | host} [/user username] remote-command",
		},
		{
			Label: "scheduler allocate interrupt-time process-time",
		},
		{
			Label: "no scheduler allocate",
		},
		{
			Label: "scheduler heapcheck enable",
		},
		{
			Label: "no scheduler heapcheck enable",
		},
		{
			Label: "scheduler heapcheck poll",
		},
		{
			Label: "no scheduler heapcheck poll",
		},
		{
			Label: "scheduler heapcheck process [memory [fast] [io] [multibus] [pci] [processor] [checktype {all | data | magic | mlite-data | pointer | refcount | lite-chunks}]]",
		},
		{
			Label: "no scheduler heapcheck process",
		},
		{
			Label: "scheduler interrupt mask profile",
		},
		{
			Label: "no scheduler interrupt mask profile",
		},
		{
			Label: "scheduler interrupt mask size buffersize",
		},
		{
			Label: "no scheduler interrupt mask size",
		},
		{
			Label: "scheduler interrupt mask time threshold-time",
		},
		{
			Label: "no scheduler interrupt mask time",
		},
		{
			Label: "scheduler interval milliseconds",
		},
		{
			Label: "no scheduler interval",
		},
		{
			Label: "scheduler isr-watchdog",
		},
		{
			Label: "no scheduler isr-watchdog",
		},
		{
			Label: "scheduler max-sched-time milliseconds",
		},
		{
			Label: "no scheduler max-sched-time",
		},
		{
			Label: "scheduler process-watchdog {hang | normal | reload | terminate}",
		},
		{
			Label: "no scheduler process-watchdog",
		},
		{
			Label: "scheduler timercheck process pid",
		},
		{
			Label: "no scheduler timercheck process pid",
		},
		{
			Label: "scheduler timercheck system context",
		},
		{
			Label: "no scheduler timercheck system context",
		},
		{
			Label: "send {line-number | * | aux number | console number | log number [msg-ext] | tty number | vty number | xsm [client client-id] message text}",
		},
		{
			Label: "service compress-config",
		},
		{
			Label: "no service compress-config",
		},
		{
			Label: "service config",
		},
		{
			Label: "no service config",
		},
		{
			Label: "service counters max age seconds",
		},
		{
			Label: "no service counters max age",
		},
		{
			Label: "service decimal-tty",
		},
		{
			Label: "no service decimal-tty",
		},
		{
			Label: "service exec-wait",
		},
		{
			Label: "no service exec-wait",
		},
		{
			Label: "service hide-telnet-address",
		},
		{
			Label: "no service hide-telnet-address",
		},
		{
			Label: "service linenumber",
		},
		{
			Label: "no service linenumber",
		},
		{
			Label: "service nagle",
		},
		{
			Label: "no service nagle",
		},
		{
			Label: "service prompt config",
		},
		{
			Label: "no service prompt config",
		},
		{
			Label: "service sequence-numbers",
		},
		{
			Label: "no service sequence-numbers",
		},
		{
			Label: "service slave-log",
		},
		{
			Label: "no service slave-log",
		},
		{
			Label: "service tcp-keepalives-in",
		},
		{
			Label: "no service tcp-keepalives-in",
		},
		{
			Label: "service tcp-keepalives-out",
		},
		{
			Label: "no service tcp-keepalives-out",
		},
		{
			Label: "service tcp-small-servers [max-servers number | no-limit]",
		},
		{
			Label: "no service tcp-small-servers [max-servers number | no-limit]",
		},
		{
			Label: "service telnet-zero-idle",
		},
		{
			Label: "no service telnet-zeroidle",
		},
		{
			Label: "service timestamps [debug | log] [uptime | datetime [msec]] [localtime] [show-timezone] [year]",
		},
		{
			Label: "no service timestamps [debug | log]",
		},
		{
			Label: "service udp-small-servers [max-servers number | no-limit]",
		},
		{
			Label: "no service udp-small-servers [max-servers number | no-limit]",
		},
		{
			Label: "service-module apa traffic-management [monitor | inline]",
		},
		{
			Label: "service-module wlan-ap interface number bootimage [autonomous | unified]",
		},
		{
			Label: "service-module wlan-ap interface number reload",
		},
		{
			Label: "service-module wlan-ap interface number reset [bootloader | default-config]",
		},
		{
			Label: "service-module wlan-ap interface number session [clear | disconnect]",
		},
		{
			Label: "service-module wlan-ap interface number statistics",
		},
		{
			Label: "service-module wlan-ap interface number status",
		},
		{
			Label: "session slot mod processor processor-id",
		},
		{
			Label: "set memory debug incremental starting-time [none]",
		},
		{
			Label: "setup",
		},
		{
			Label: "{show command | append url}",
		},
		{
			Label: "{show command | begin regular-expression}",
		},
		{
			Label: "{show command | exclude regular-expression}",
		},
		{
			Label: "{show command | include regular-expression}",
		},
		{
			Label: "{show command | redirect url}",
		},
		{
			Label: "{show command | section [include | exclude] regular-expression}",
		},
		{
			Label: "{show command | tee [/ append] url}",
		},
		{
			Label: "slave auto-sync config",
		},
		{
			Label: "no slave auto-sync config",
		},
		{
			Label: "slave default-slot processor-slot-number",
		},
		{
			Label: "slave image {system | file-url}",
		},
		{
			Label: "slave reload",
		},
		{
			Label: "slave sync config",
		},
		{
			Label: "slave terminal",
		},
		{
			Label: "no slave terminal",
		},
		{
			Label: "software clean [file file url] [force] [switch nodes] [verbose]",
		},
		{
			Label: "software commit [switchnode] [verbose]",
		},
		{
			Label: "software expand {file source url | running} [todestination url] [switchnodes] [verbose]",
		},
		{
			Label: "software install file bundle url [switchnodes] [auto-rollbackminutes] [force] [on-reboot] [provisioning-fileprovisioning-file url] [force] [new] [verbose]",
		},
		{
			Label: "software install source switchnode [switchnode] [auto-rollbackminutes] [force] [on-reboot] [verbose] [new] [provisioning-fileprovisioning-file url]",
		},
		{
			Label: "software install source switch node [switch nodes] [auto-rollback minutes] [on-reboot] [provisioning-file provisioning-file url] [force] [verbose] [new]",
		},
		{
			Label: "software provision source {url bundle or package url | listlist-name} [packagepackage name or wildcard] [switchnode] [force] [verbose]",
		},
		{
			Label: "software repackage switchnode dest url and filename",
		},
		{
			Label: "software rollback [switchnode] [as-booted] [provisioning-fileprovisioning-file url] [on-reboot] [force] [verbose]",
		},
		{
			Label: "software source list list-name-string",
		},
		{
			Label: "no software source list list-name-string",
		},
		{
			Label: "software unistall bundle or package url [switchnode]",
		},
		{
			Label: "special-character-bits {7 | 8}",
		},
		{
			Label: "no special-character-bits",
		},
		{
			Label: "squeeze [/nolog] [/quiet] filesystem:",
		},
		{
			Label: "squeeze filesystem:",
		},
		{
			Label: "stack-mib portname portname",
		},
		{
			Label: "state-machine name state first-character last-character [next-state delay | transmit]",
		},
		{
			Label: "no state-machine name",
		},
		{
			Label: "stopbits {1 | 1. 5 | 2}",
		},
		{
			Label: "no stopbits",
		},
		{
			Label: "storm-control {broadcast | multicast | unicast} level level [. level]",
		},
		{
			Label: "no storm-control {broadcast | multicast | unicast} level",
		},
		{
			Label: "sync-restart-delay timer",
		},
		{
			Label: "no sync-restart-delay timer",
		},
		{
			Label: "systat all",
		},
		{
			Label: "[default] system flowcontrol bus {auto | on}",
		},
		{
			Label: "no system flowcontrol bus",
		},
		{
			Label: "system jumbomtu mtu-size",
		},
		{
			Label: "no system jumbomtu",
		},
		{
			Label: "tdm clock priority priority-number {slot/ds1-port | slot/ds3-port:ds1-port | external | freerun}",
		},
		{
			Label: "no tdm clock priority priority-number {slot/ds1-port | slot/ds3-port:ds1-port | external | freerun}",
		},
		{
			Label: "terminal databits {5 | 6 | 7 | 8}",
		},
		{
			Label: "terminal data-character-bits {7 | 8}",
		},
		{
			Label: "terminal dispatch-character ascii-number [ascii-number2...]",
		},
		{
			Label: "terminal dispatch-timeout milliseconds",
		},
		{
			Label: "terminal download",
		},
		{
			Label: "terminal editing",
		},
		{
			Label: "terminal no editing",
		},
		{
			Label: "terminal escape-character ascii-number",
		},
		{
			Label: "terminal exec-character-bits {7 | 8}",
		},
		{
			Label: "terminal flowcontrol {none | software [in | out] | hardware}",
		},
		{
			Label: "terminal full-help",
		},
		{
			Label: "terminal history",
		},
		{
			Label: "terminal no history",
		},
		{
			Label: "terminal history size number-of-lines",
		},
		{
			Label: "terminal no history size",
		},
		{
			Label: "terminal hold-character ascii-number",
		},
		{
			Label: "terminal no hold-character",
		},
		{
			Label: "terminal international",
		},
		{
			Label: "no terminal international",
		},
		{
			Label: "terminal keymap-type keymap-name",
		},
		{
			Label: "terminal length screen-length",
		},
		{
			Label: "terminal monitor",
		},
		{
			Label: "terminal notify",
		},
		{
			Label: "terminal no notify",
		},
		{
			Label: "terminal padding ascii-number count",
		},
		{
			Label: "terminal parity {none | even | odd | space | mark}",
		},
		{
			Label: "terminal rxspeed bps",
		},
		{
			Label: "terminal special-character-bits 7 | 8",
		},
		{
			Label: "terminal speed bps",
		},
		{
			Label: "terminal start-character ascii-number",
		},
		{
			Label: "terminal stopbits {1 | 1.5 | 2}",
		},
		{
			Label: "terminal stop-character ascii-number",
		},
		{
			Label: "terminal telnet break-on-ip",
		},
		{
			Label: "terminal telnet refuse-negotiations",
		},
		{
			Label: "terminal telnet speed default-speed maximum-speed",
		},
		{
			Label: "terminal telnet sync-on-break",
		},
		{
			Label: "terminal telnet transparent",
		},
		{
			Label: "terminal terminal-type terminal-type",
		},
		{
			Label: "terminal txspeed bps",
		},
		{
			Label: "terminal width characters",
		},
		{
			Label: "terminal-queue entry-retry-interval seconds",
		},
		{
			Label: "no terminal-queue entry-retry-interval",
		},
		{
			Label: "terminal-type {terminal-name | terminal-type}",
		},
		{
			Label: "no terminal-type",
		},
		{
			Label: "test cable-diagnostics tdr interface type number",
		},
		{
			Label: "test flash",
		},
		{
			Label: "test interfaces",
		},
		{
			Label: "test memory",
		},
		{
			Label: "test memory destroy [chunk | mgd-chunk | force-chunk | dangling-reference] chunk-id",
		},
		{
			Label: "test platform police get",
		},
		{
			Label: "test platform police set rate",
		},
		{
			Label: "tftp-server flash [partition-number :] filename1 [alias filename2] [access-list-number]",
		},
		{
			Label: "tftp-server rom alias filename1 [access-list-number]",
		},
		{
			Label: "no tftp-server {flash [partition-number :] filename1 | rom alias filename2}",
		},
		{
			Label: "tftp-server flash [device :] [partition-number :] filename",
		},
		{
			Label: "no tftp-server flash [device :] [partition-number :] filename",
		},
		{
			Label: "tftp-server flash device : filename",
		},
		{
			Label: "no tftp-server flash device : filename",
		},
		{
			Label: "time-period minutes",
		},
		{
			Label: "no time-period minutes",
		},
		{
			Label: "trace [protocol] [destination]",
		},
		{
			Label: "trace [protocol] [destination]",
		},
		{
			Label: "traceroute [vrf vrf-name | topology topology-name] [protocol] destination",
		},
		{
			Label: "traceroute mac source-mac-address {destination-mac-address | interface type interface-number destination-mac-address} [vlan vlan-id] [detail]",
		},
		{
			Label: "traceroute mac interface type interface-number source-mac-address {destination-mac-address | interface type interface-number destination-mac-address} [vlan vlan-id] [detail]",
		},
		{
			Label: "traceroute mac ip {source-ip-address | source-hostname} {destination-ip-address | destination-hostname} [detail]",
		},
		{
			Label: "undelete index [filesystem :]",
		},
		{
			Label: "unprofile {process {process-ID | process-name}} {start-address end-address increment | all | task}",
		},
		{
			Label: "upgrade automatic abortversion",
		},
		{
			Label: "no upgrade automatic abortversion",
		},
		{
			Label: "upgrade automatic getversion {cisco username username password password image image | url} { [at hh:mm] | now | [in hh:mm]} [disk-management {auto | confirm | no}]",
		},
		{
			Label: "upgrade automatic runversion {at hh:mm | now | in hh:mm}",
		},
		{
			Label: "upgrade filesystem monlib {disk0 | disk1}",
		},
		{
			Label: "upgrade rom-monitor slot num {sp | rp} file filename",
		},
		{
			Label: "upgrade rom-monitor slot num {sp | rp} {invalidate | preference} {region1 | region2}",
		},
		{
			Label: "upgrade rom-monitor filename URL slot",
		},
		{
			Label: "upgrade filesystem monlib {disk0 | disk1}",
		},
		{
			Label: "upgrade rom-monitor preference [readonly | upgrade]",
		},
		{
			Label: "vacant-message [d message d]",
		},
		{
			Label: "no vacant-message",
		},
		{
			Label: "verify [/md5 [md5-value] ] filesystem: [file-url]",
		},
		{
			Label: "verify {/md5 flash-filesystem [expected-md5-signature] | /ios flash-filesystem | flash-filesystem}",
		},
		{
			Label: "vtp {domain domain-name | file filename | interface interface-name [only] | mode {client | off | server | transparent} {vlan | mst | unknown} | password password-value [hidden | secret] | pruning | version {1 | 2 | 3}}",
		},
		{
			Label: "no vtp",
		},
		{
			Label: "warm-reboot [count number] [uptime minutes]",
		},
		{
			Label: "no warm-reboot count number uptime minutes",
		},
		{
			Label: "where",
		},
		{
			Label: "width characters",
		},
		{
			Label: "no width",
		},
		{
			Label: "write core [hostname [LINE] | destination-address [LINE]]",
		},
		{
			Label: "write memory",
		},
		{
			Label: "write mib-data",
		},
		{
			Label: "write network [host-file-address]",
		},
		{
			Label: "xmodem [-c] [-y] [-e] [-f] [-r] [-x] [ [-s] data-rate] [filename]",
		},
  }
  return items, nil
}

