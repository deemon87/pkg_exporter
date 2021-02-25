package info

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

// Pkg type for save information about packages
type Pkg struct {
	Name          string
	Version       string
	DesiredAction string
	PackageStatus string
	ErrorFlags    string
}

// GetPkgInfo get information about installed packages
func GetPkgInfo() *[]Pkg {
	var out bytes.Buffer

	cmd := exec.Command("dpkg-query", "-f", "${Package}=${source:Version}=${db:Status-Want}=${db:Status-Status}=${db:Status-Eflag}\n", "-W")
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	data := strings.Fields(out.String())
	var packages []Pkg
	for _, pack := range data {
		line := strings.Split(pack, "=")
		packages = append(packages, Pkg{
			Name:          line[0],
			Version:       line[1],
			DesiredAction: line[2],
			PackageStatus: line[3],
			ErrorFlags:    line[4],
		})
	}

	return &packages
}
