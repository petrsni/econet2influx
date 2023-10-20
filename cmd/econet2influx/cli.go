package main

import (
	"github.com/petrsni/econet2influx/internal/app/econet2influx"
)

type CLI struct {
	Install InstallCmd `cmd:"" help:"install as svc"`
	Run     RunCmd     `cmd:"" help:"run mode"`
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

func (c *RunCmd) Run(appctx *econet2influx.AppCtx) error {
	err := econet2influx.Run(c.Influx, c.InfluxOrg, c.InfluxBucket, c.InfluxToken, c.EconetIP, c.EconetUser, c.EconetPass, *appctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *InstallCmd) Run(appctx *econet2influx.AppCtx) error {
	/*
		svcConf := service.Config{
			Name:        "econet2influx",
			DisplayName: "econet2influx",
			Description: "delivers ecoNET data to influxdb",
			EnvVars: map[string]string{
				"INFLUX_URL":    c.Influx,
				"INFLUX_ORG":    c.InfluxOrg,
				"INFLUX_BUCKET": c.InfluxBucket,
				"INFLUX_TOKEN":  c.InfluxToken,
				"ECONET_URL":    c.EconetIP,
				"ECONET_USER":   c.EconetUser,
				"ECONET_PASS":   c.EconetPass,
			},
		}
	*/
	return nil
}
