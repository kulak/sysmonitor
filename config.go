package main

import "os"

type Feature string

type SrcDst struct {
	UseSudo bool   `yaml:"use-sudo"`
	Src     string `yaml:"src"`
	Dst     string `yaml:"dst"`
}

type Config struct {
	Features  []Feature `yaml:"features"`
	OutputDir string    `yaml:"output-dir"`

	RsyncDirs    []SrcDst `yaml:"rsync-dirs"`
	RsyncArgs    []string `yaml:"rsync-args"`
	BtrfsDevices []string `yaml:"btrfs-devices"`

	// private state
	out *os.File
}

const BtrfsFeature Feature = "btrfs"
const JournalFeature Feature = "journal"
const ZfsFeature Feature = "zfs"
const RsyncFeature Feature = "rsync"

var SupportedFeatures = []Feature{
	BtrfsFeature,
	JournalFeature,
	ZfsFeature,
	RsyncFeature,
}

func (c *Config) FeatureEnabled(feature Feature) bool {
	for _, v := range c.Features {
		if v == feature {
			return true
		}
	}
	return false
}
