package configmodify

import (
	"bufio"
	"os"
	"regexp"
)

// SelinuxMod : Disable selinux
func SelinuxMod() {
	err := IsFileExist("/etc/sysconfig/selinux")
	PrintErr("File not found", err)
	if err != nil {
		return
	}
	f, err := os.OpenFile("/etc/sysconfig/selinux", os.O_RDWR, 0600)
	PrintErr("Couldn't open file", err)
	if err != nil {
		return
	}
	rd := bufio.NewReader(f)
	buf := make([]byte, 1024*1024)
	//wt := bufio.NewWriter(f)
	n, err1 := rd.Read(buf)
	PrintErr("Couldn't read selinux", err1)
	if err1 != nil {
		return
	}
	// Just ensuring there is some data in selinux file.
	var m string
	if n > 10 {
		re := regexp.MustCompile(`[^\S]SELINUX=\S.*\n`)
		m = re.ReplaceAllString(string(buf[:n]), "\nSELINUX=disabled\n")
	}

	out, err := os.OpenFile("/etc/sysconfig/selinux", os.O_WRONLY, 0600)

	PrintErr("Write: /etc/sysconfig/selinux", err)
	defer out.Close()
	w := bufio.NewWriter(out)
	w.WriteString(m)
	w.Flush()
}
