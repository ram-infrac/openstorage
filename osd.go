package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"

	"github.com/libopenstorage/openstorage/apiserver"
	osdcli "github.com/libopenstorage/openstorage/cli"
	"github.com/libopenstorage/openstorage/drivers/aws"
	"github.com/libopenstorage/openstorage/drivers/nfs"
	"github.com/libopenstorage/openstorage/volume"
)

const (
	version = "0.3"
)

var (
	drivers = []string{aws.Name, nfs.Name}
)

type osd struct {
	// Drivers map[string][]volume.DriverParams
	Drivers map[string]volume.DriverParams
}

type Config struct {
	Osd osd
}

func start(c *cli.Context) {
	cfg := Config{}

	file := c.String("file")
	if file != "" {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(b, &cfg)
		if err != nil {
			panic(err)
		}

	}

	if !osdcli.DaemonMode(c) {
		cli.ShowAppHelp(c)
	}

	// Start the drivers.
	for d, v := range cfg.Osd.Drivers {
		fmt.Println("Starting driver: ", d)
		_, err := volume.New(d, v)
		if err != nil {
			panic(err)
		}

		// Create a unique path for a UNIX socket that the driver will listen on.
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			panic(err)
		}
		uuid := string(out)
		uuid = strings.TrimSuffix(uuid, "\n")

		sock := "/tmp/" + uuid
		err = apiserver.StartDriver(d, 0, sock)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "osd"
	app.Usage = "Open Storage CLI"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json,j",
			Usage: "output in json",
		},
		cli.BoolFlag{
			Name:  osdcli.DaemonAlias,
			Usage: "Start OSD in daemon mode",
		},
		cli.StringSliceFlag{
			Name:  "provider, p",
			Usage: "provider name and options: name=btrfs,root_vol=/var/openstorage/btrfs",
			Value: new(cli.StringSlice),
		},
		cli.StringFlag{
			Name:  "file,f",
			Usage: "file to read the OSD configuration from.",
			Value: "",
		},
	}
	app.Action = start
	app.Commands = []cli.Command{
		{
			Name:        "volume",
			Aliases:     []string{"v"},
			Usage:       "Manage volumes",
			Subcommands: osdcli.VolumeCommands(),
		},
		{
			Name:        "driver",
			Aliases:     []string{"d"},
			Usage:       "Manage drivers",
			Subcommands: osdcli.DriverCommands(),
		},
	}
	app.Run(os.Args)
}

func init() {
}