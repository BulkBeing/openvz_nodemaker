package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"openvz_nodemaker/configmodify"
	"os"
	"os/exec"
	"strings"
)

// For human readable conversion
const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
)

func main() {
	err := downloadFile("/etc/yum.repos.d/openvz.repo", "http://download.openvz.org/openvz.repo")
	configmodify.FatalErr("Error in downloading file:", err)
	out := runCommand("rpm", "--import", "http://download.openvz.org/RPM-GPG-Key-OpenVZ")
	fmt.Println(out)
	out = runCommand("yum", "install", "vzkernel", "vzctl", "vzquota", "-y")
	fmt.Println(out)
	configmodify.SysctlModify()
	out = runCommand("sysctl", "-p")
	fmt.Println(out)
	configmodify.SelinuxMod()
	fmt.Println("Reboot now to use the new Kernel")
}

// Download file and save it as "filepath"
func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if (err != nil) || (resp.StatusCode != 200) {
		if err == nil {
			err = fmt.Errorf(resp.Status)
		}
		return err
	}
	defer resp.Body.Close()

	d, err1 := io.Copy(out, resp.Body)
	if err1 != nil {
		return err1
	}
	fmt.Println("Downloaded:", filepath, " size:", convertFileSize(d))
	return nil
}

// filesize in bytes
func convertFileSize(filesize int64) string {
	switch {
	case filesize > GB:
		return fmt.Sprintf("%v GB", filesize/GB)
	case filesize > MB:
		return fmt.Sprintf("%v MB", (filesize / MB))
	case filesize > KB:
		return fmt.Sprintf("%v KB", (filesize / KB))
	default:
		return fmt.Sprintf("%v B", filesize)
	}
}

func runCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	log.Println("Executing:", command, strings.Join(args, " "))
	out, err := cmd.CombinedOutput()
	configmodify.FatalErr(fmt.Sprintf("Error executing command: %s", command), err)
	return string(out)
}
