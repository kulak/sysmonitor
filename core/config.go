package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Feature string

type SrcDst struct {
	UseSudo bool   `yaml:"use-sudo"`
	Src     string `yaml:"src"`
	Dst     string `yaml:"dst"`
}

type Application struct {
	Description string   `yaml:"description"`
	App         string   `yaml:"app"`
	Args        []string `yaml:"args"`
}

type Config struct {
	Features  []Feature `yaml:"features"`
	OutputDir string    `yaml:"output-dir"`

	RsyncDirs    []SrcDst      `yaml:"rsync-dirs"`
	RsyncArgs    []string      `yaml:"rsync-args"`
	BtrfsDevices []string      `yaml:"btrfs-devices"`
	Apps         []Application `yaml:"apps"`

	// private state
	out *os.File
}

const BtrfsFeature Feature = "btrfs"
const JournalFeature Feature = "journal"
const ZfsFeature Feature = "zfs"
const RsyncFeature Feature = "rsync"
const AppsFeature Feature = "apps"

var SupportedFeatures = []Feature{
	BtrfsFeature,
	JournalFeature,
	ZfsFeature,
	RsyncFeature,
	AppsFeature,
}

func (c *Config) FeatureEnabled(feature Feature) bool {
	for _, v := range c.Features {
		if v == feature {
			return true
		}
	}
	return false
}

func (c *Config) OutputWriter() (io.Writer, error) {
	if c.OutputDir == "" {
		return os.Stdout, nil
	}
	if c.out == nil {
		fileName := filepath.Join(c.OutputDir, fmt.Sprintf("report-%03d.html", time.Now().YearDay()))
		f, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		c.out = f
	}
	return c.out, nil
}

func (c *Config) CloseOutputWriter() {
	if c.out == nil {
		return
	}
	c.out.Close()
	c.out = nil
}
