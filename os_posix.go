// +build !windows

package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func create_sock_flag(name, desc string) *string {
	return flag.String(name, "unix", desc)
}

// Full path of the current executable
func get_executable_filename() string {
	// try readlink first
	path, err := os.Readlink("/proc/self/exe")
	if err == nil {
		return path
	}
	// use argv[0]
	path = os.Args[0]
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	if file_exists(path) {
		return path
	}
	// Fallback : use "gocode" and assume we are in the PATH...
	path, err = exec.LookPath("gocode")
	if err == nil {
		return path
	}
	return ""
}

// config location

func config_dir() string {
	return filepath.Join(xdg_home_dir(), "gocode")
}

func config_file() string {
	return filepath.Join(xdg_home_dir(), "gocode", "config.json")
}

func get_local_filename(filename string) string {
	newname := strings.Replace(filename, PATH_CLIENT, PATH_SERVER, -1)
	return filepath.ToSlash(newname)
}
