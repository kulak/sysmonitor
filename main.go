package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

//go:embed res/*
var res embed.FS

type Config struct {
	OutputDir    string   `yaml:"output-dir"`
	BtrfsDevices []string `yaml:"btrfs-devices"`

	// private state
	out *os.File
}

func main() {
	var err error
	var report ReportContext
	report.Now = time.Now()
	report.Groups = make([]Group, 0)
	var help bool
	var confFileName string
	// var dryRun bool
	flag.BoolVar(&help, "h", false, "Print Help and Exit")
	// flag.BoolVar(&dryRun, "n", false, "Dry Run")
	flag.StringVar(&confFileName, "cf", "sysmonitor.yaml", "Configuration file name")
	flag.Parse()

	var conf = &Config{}
	if help {
		printHelp(conf)
		return
	}

	if err = loadConfig(confFileName, conf); err != nil {
		fmt.Printf("failed to load config file %s: %v", confFileName, err)
		return
	}

	var templ *template.Template
	if templ, err = template.New("report").ParseFS(res, "res/report.html"); err != nil {
		fmt.Printf("failed to load template file res/report.html: %v", err)
		return
	}

	report.Groups = append(report.Groups, btrfsReport(conf)...)
	report.Groups = append(report.Groups, journalReport())

	var writer io.Writer
	if writer, err = conf.OutputWriter(); err != nil {
		fmt.Printf("failed to get output writer: %v", err)
		return
	}
	defer conf.CloseOutputWriter()
	if err = templ.ExecuteTemplate(writer, "report", &report); err != nil {
		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println("finished with error")
		return
	}
}

func execReport(app string, cmdArgs []string, title string) Group {
	var msgs []Message
	cmd := exec.Command(app, cmdArgs...)
	var out []byte
	var err error
	if out, err = cmd.CombinedOutput(); err != nil {
		msgs = append(msgs, Msg(err.Error(), errorLvl, p))
		if len(out) > 0 {
			msgs = append(msgs, Msg(string(out), errorLvl, code))
		}
	} else {
		msgs = append(msgs, Msg(string(out), infoLvl, code))
	}
	return Group{
		Title:       title,
		Description: fmt.Sprintf("%s %s", app, strings.Join(quote(cmd.Args), " ")),
		Msgs:        msgs,
	}
}

func quote(strs []string) []string {
	var rv []string
	for _, each := range strs {
		rv = append(rv, fmt.Sprintf("'%s'", each))
	}
	return rv
}

func printHelp(conf *Config) {
	var err error
	fmt.Printf("Usage: %s [OPTIONS]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("Config File Example:")
	var content []byte
	if content, err = yaml.Marshal(conf); err != nil {
		panic(err)
	}
	os.Stdout.Write(content)
	fmt.Println()
}

func loadConfig(fileName string, target *Config) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, target)
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
