package cmd_test

import (
	"testing"

	"github.com/dgethings/lsp-cisco-ios/scraper/cmd"
)

func TestNormaliseSection(t *testing.T) {
	tests := []struct {
		mode    string
		section string
		include bool
	}{
		{"", "", false},
		{"All configuration modes", "", true},
		{"Archive config mode; log config (configuration change logger) submode (config-archive-log-cfg)#", "config-archive-log-config", true},
		{"Archive configuration (config-archive)", "config-archive", true},
		{"Archive configuration mode, log config (configuration-change logger) submode (config-archive-log-cfg)#", "config-archive-log-config", true},
		{"Config-VLAN mode", "config-vlan", true},
		{"Configuration change logger configuration (config-archive-log-config)", "config-archive-log-config", true},
		{"Diagnostic (diag)", "", false},
		{"EXEC (>) Privileged EXEC (#) Diagnostic (diag)", "", false},
		{"EXEC mode", "", false},
		{"EXEC", "", false},
		{"Global Configuration", "config", true},
		{"Global configuration (config#)", "config", true},
		{"Global configuration (config)", "config", true},
		{"Global configuration", "config", true},
		{"Interface configuration (config-if)", "config-if", true},
		{"Interface configuration Frame Relay DLCI configuration Template configuration (config-template)", "config-template", true},
		{"Interface configuration mode", "config-if", true},
		{"Interface configuration", "config-if", true},
		{"Line configuration (config-line)", "config-line", true},
		{"Line configuration", "config-line", true},
		{"MST configuration (config-mst)", "config-mst", true},
		{"Main CPU submode", "", false},
		{"Privilege EXEC (#)", "", false},
		{"Privileged EXEC (#) Diagnostic (#)", "", false},
		{"Privileged EXEC (#) Diagnostic (diag)", "", false},
		{"Privileged EXEC (#) Diagnostic Mode (diag)", "", false},
		{"Privileged EXEC (#)", "", false},
		{"Privileged EXEC mode (#)", "", false},
		{"Privileged EXEC mode", "", false},
		{"Privileged EXEC on the Switch Processor", "", false},
		{"Privileged EXEC", "", false},
		{"Privileged EXEC(#)", "", false},
		{"Privileged Exec", "", false},
		{"ROM monitor mode", "", false},
		{"ROM monitor", "", false},
		{"Redundancy configuration (config-r) Main CPU redundancy configuration (config-r-mc)", "config-r", true},
		{"Redundancy configuration (config-red)", "config-red", true},
		{"Time-range configuration (config-time-range)", "config-time-range", true},
		{"User EXEC (>) Priviledged EXEC (#)", "", false},
		{"User EXEC (>) Privileged EXEC (#) Diagnostic (diag)", "", false},
		{"User EXEC (>) Privileged EXEC (#) Diagnostic (diag)--Cisco ASR 1000 Series Routers only", "", false},
		{"User EXEC (>) Privileged EXEC (#) Diagnostic Mode (diag)", "", false},
		{"User EXEC (>) Privileged EXEC (#)", "", false},
		{"User EXEC (>)", "", false},
		{"User EXEC (>)Privileged EXEC (#)", "", false},
		{"User EXEC Priviledged EXEC", "", false},
		{"User EXEC Privileged EXEC All configuration modes", "", false},
		{"User EXEC Privileged EXEC Diagnostic", "", false},
		{"User EXEC Privileged EXEC", "", false},
		{"User EXEC", "", false},
		{"User EXECPrivileged EXEC", "", false},
		{"privileged EXEC (#)", "", false},
		{"user EXEC privileged EXEC", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.mode, func(t *testing.T) {
			t.Parallel()
			sec, ok := cmd.ConfigSection(tt.mode)
			if tt.section != sec {
				t.Errorf("expected %s, actual %s", tt.section, sec)
			}
			if tt.include != ok {
				t.Errorf("expected %t, actual %t", tt.include, ok)
			}
		})
	}
}
