package gofetch

import (
	"fmt"
	"os"
	"strings"

	"github.com/OrlandoRomo/gofetch/command"
	"github.com/shirou/gopsutil/mem"
)

type macos struct{}

func NewMacOS() command.OSInformer {
	return &macos{}
}

// GetName returns the current user name
func (mac *macos) GetName() (string, error) {
	return command.ExecuteCommand("whoami")
}

// GetOSVersion returns the name of the current OS, version and kernel version
func (mac *macos) GetOSVersion() (string, error) {
	return command.ExecuteCommand("uname", "-srm")
}

// GetHostname returns the hostname of the machine
func (mac *macos) GetHostname() (string, error) {
	return os.Hostname()
}

// GetUptime returns the up time of the current OS
func (mac *macos) GetUptime() (string, error) {
	uptime, err := command.ExecuteCommand("uptime")
	if err != nil {
		return "", err
	}
	uptime = strings.Replace(uptime, "\r\n", "", -1)
	uptimes := strings.Split(uptime, " ")
	return uptimes[4], nil
}

// GetNumberPackages return the number of packages install by homebrew
func (mac *macos) GetNumberPackages() (string, error) {
	cmd := "ls /usr/local/Cellar/ | wc -l"
	packages, err := command.ExecuteCommand("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(packages), nil
}

// GetShellInformation return the used shell and its version
func (mac *macos) GetShellInformation() (string, error) {
	shell, err := command.ExecuteCommand(os.ExpandEnv("$SHELL"), "--version")
	if err != nil {
		return "", err
	}
	return shell, nil
}

// GetResolution returns the resolution of thee current monitor
func (mac *macos) GetResolution() (string, error) {
	cmd := "system_profiler SPDisplaysDataType  | grep Resolution"
	resolution, err := command.ExecuteCommand("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	resolutions := strings.Split(resolution, "Resolution: ")
	resolution = strings.TrimSpace(resolutions[1])
	return resolution, nil
}

// GetDesktopEnvironment returns the resolution of thee current monitor
func (mac *macos) GetDesktopEnvironment() (string, error) {
	return "Aqua", nil
}

// GetTerminalInfo get the current terminal name
func (mac *macos) GetTerminalInfo() (string, error) {
	terminal, err := command.ExecuteCommand("echo", os.ExpandEnv("$TERM_PROGRAM"))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(terminal), nil
}

// GetCPU returns the name of th CPU
func (mac *macos) GetCPU() (string, error) {
	cmd := "sysctl -a | grep machdep.cpu.brand_string"
	cpuInfo, err := command.ExecuteCommand("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	cpu := strings.Split(cpuInfo, ": ")
	cpuInfo = strings.Replace(cpu[1], "\n\r", "", -1)
	cpuInfo = strings.TrimSpace(cpuInfo)
	return cpuInfo, nil
}

// GetGPU returns the name of the GPU
func (mac *macos) GetGPU() (string, error) {
	cmd := "system_profiler SPDisplaysDataType | grep 'Chipset Model'"
	gpu, err := command.ExecuteCommand("bash", "-c", cmd)
	if err != nil {
		return "", err
	}
	gpus := strings.Split(gpu, "Chipset Model: ")
	gpu = strings.TrimSpace(gpus[1])
	return gpu, nil
}

// GetMemoryUsage returns the memory usage of the computer
func (mac *macos) GetMemoryUsage() (string, error) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	total := memStat.Total / (1024 * 1024)
	used := memStat.Used / (1024 * 1024)
	return fmt.Sprintf("%v MB / %v MB", used, total), nil
}