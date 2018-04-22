package configmodify

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// SysctlModify : Appends necessay settings to sysctl.conf if they are not already there.
func SysctlModify() {
	config := map[string]int{
		"net.ipv4.ip_forward":                  1,
		"net.ipv4.conf.default.proxy_arp":      0,
		"net.ipv4.conf.all.rp_filter":          1,
		"kernel.sysrq":                         1,
		"net.ipv4.conf.default.send_redirects": 1,
		"net.ipv4.conf.all.send_redirects":     0,
		"net.ipv4.icmp_echo_ignore_broadcasts": 1,
		"net.ipv4.conf.default.forwarding":     1,
	}
	// Map for storing the settings that are not present and needs to be appended in sysctl.conf
	toadd := make(map[string]int)
	data := make([]byte, 1024*1024)
	f, err := os.OpenFile("/etc/sysctl.conf", os.O_APPEND|os.O_RDWR, 0600)
	FatalErr("Error in reading file /etc/sysctl.conf", err)
	defer f.Close()
	_, err1 := f.Read(data)
	FatalErr("Error reading from file: /etc/sysctl.conf", err1)
	for key, val := range config {
		pat := fmt.Sprintf("[^\\S]%s\\s?=\\s?%d", key, val)
		re := regexp.MustCompile(pat)
		m := re.FindString(string(data))
		if m == "" {
			toadd[key] = val
		}
	}
	w := bufio.NewWriter(f)
	for key, val := range toadd {
		setting := fmt.Sprintf("%s = %d\n", key, val)
		fmt.Println("writing to /etc/sysctl.conf", setting)
		_, err2 := w.WriteString(setting)
		FatalErr("Error writing /etc/sysctl.conf", err2)
	}
	w.Flush()
}
