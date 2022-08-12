package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"time"

	"gitlab.com/nest-machine/sysmonitor/core"
	"gitlab.com/nest-machine/sysmonitor/features"
	"gopkg.in/yaml.v3"
)

//go:embed res/*
var res embed.FS

func main() {
	var err error
	var report core.ReportContext
	report.Now = time.Now()
	report.Groups = make([]core.Group, 0)
	var help bool
	var confFileName string
	// var dryRun bool
	flag.BoolVar(&help, "h", false, "Print Help and Exit")
	// flag.BoolVar(&dryRun, "n", false, "Dry Run")
	flag.StringVar(&confFileName, "cf", "sysmonitor.yaml", "Configuration file name")
	flag.Parse()

	var conf = &core.Config{}
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

	if conf.FeatureEnabled(core.BtrfsFeature) {
		report.Groups = append(report.Groups, features.BtrfsReport(conf)...)
	}
	if conf.FeatureEnabled(core.ZfsFeature) {
		report.Groups = append(report.Groups, features.ZfsReport())
	}
	if conf.FeatureEnabled(core.JournalFeature) {
		report.Groups = append(report.Groups, features.JournalReport())
	}
	if conf.FeatureEnabled(core.RsyncFeature) {
		report.Groups = append(report.Groups, features.RsyncReport(conf)...)
	}
	if conf.FeatureEnabled(core.AppsFeature) {
		report.Groups = append(report.Groups, features.AppsReport(conf)...)
	}

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

func printHelp(conf *core.Config) {
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

func loadConfig(fileName string, target *core.Config) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, target)
}
