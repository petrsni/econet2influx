package main

import (
	"fmt"
	"log/slog"

	"github.com/kardianos/service"
	"github.com/petrsni/econet2influx/internal/app/econet2influx"
)

type CLI struct {
	Install InstallCmd `cmd:"" help:"install as svc"`
	Run     RunCmd     `default:"1" cmd:"" help:"run mode"`
}

type InstallCmd struct {
	//Conf string `arg:"" name:"conf" help:"install as svc" type:"path"`
	CliParams
}

type RunCmd struct {
	CliParams
}

type CliParams struct {
	Influx       string `name:"influx" help:"influxdb url" type:"string" ENV:"INFLUX_URL"`
	InfluxOrg    string `name:"influxorg" help:"influxdb org" type:"string" ENV:"INFLUX_ORG"`
	InfluxBucket string `name:"influxbucket" help:"influxdb bucket" type:"string" ENV:"INFLUX_BUCKET"`
	InfluxToken  string `name:"influxtoken" help:"influxdb token" type:"string" ENV:"INFLUX_TOKEN"`
	EconetIP     string `name:"econetip" help:"econet url" type:"string" ENV:"ECONET_URL"`
	EconetUser   string `name:"econetuser" help:"econet user" type:"string" ENV:"ECONET_USER"`
	EconetPass   string `name:"econetpass" help:"econet pass" type:"string" ENV:"ECONET_PASS"`
}

func (p *CliParams) Print() {
	slog.Info(fmt.Sprintf("influxdb url: %s", p.Influx))
	slog.Info(fmt.Sprintf("influxdb org: %s", p.InfluxOrg))
	slog.Info(fmt.Sprintf("influxdb bucket: %s", p.InfluxBucket))
	slog.Info(fmt.Sprintf("influxdb token: %s", p.InfluxToken))
	slog.Info(fmt.Sprintf("econet url: %s", p.EconetIP))
	slog.Info(fmt.Sprintf("econet user: %s", p.EconetUser))
	slog.Info(fmt.Sprintf("econet pass: %s", p.EconetPass))
}

func (c *RunCmd) Run(appctx *econet2influx.AppCtx) error {
	// create a service
	//create a service, then install it
	svconf := NewSvcConf(&c.CliParams)
	prg := &program{params: &c.CliParams, appctx: appctx}
	svc, err := service.New(prg, svconf)
	if err != nil {
		return err
	}
	err = svc.Run()
	if err != nil {
		return err
	}
	return nil
}

func (c *InstallCmd) Run(appctx *econet2influx.AppCtx) error {
	//create a service, then install it
	svconf := NewSvcConf(&c.CliParams)
	prg := &program{params: &c.CliParams, appctx: appctx}
	svc, err := service.New(prg, svconf)
	if err != nil {
		return err
	}
	err = svc.Install()

	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf("service registered as '%s' with '%s' ", svconf.Name, svc.Platform()))
	return nil

}

type program struct {
	params *CliParams
	appctx *econet2influx.AppCtx
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.work()
	return nil
}
func (p *program) work() {
	p.params.Print()
	econet2influx.Run(p.params.Influx, p.params.InfluxOrg, p.params.InfluxBucket, p.params.InfluxToken, p.params.EconetIP, p.params.EconetUser, p.params.EconetPass, *p.appctx)

}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func NewSvcConf(p *CliParams) *service.Config {
	return &service.Config{
		Name:        "econet2influx",
		DisplayName: "econet2influx",
		Description: "delivers ecoNET data to influxdb",
		Arguments:   []string{"run", "--influx", p.Influx, "--influxorg", p.InfluxOrg, "--influxbucket", p.InfluxBucket, "--influxtoken", p.InfluxToken, "--econetip", p.EconetIP, "--econetuser", p.EconetUser, "--econetpass", p.EconetPass},
	}
}
