package econet2influx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type AppCtx struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Logger *slog.Logger
}

func Run(influx, org, bucket, token, econet, econetusr, econetpw string, appctx AppCtx) error {
	appctx.Logger.Info("starting...")

	client := influxdb2.NewClient(influx, token)

	writeAPI := client.WriteAPIBlocking(org, bucket)
	defer client.Close()
	go readAndWrite(influx, org, bucket, token, econet, econetusr, econetpw, writeAPI, appctx.Ctx)
	<-appctx.Ctx.Done()
	appctx.Logger.Info("stopping...")
	return nil
}

func readAndWrite(influx, org, bucket, token, econet, econetusr, econetpw string, ewapi api.WriteAPIBlocking, ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	cl := EconetHttpClient()
	for {
		select {
		case <-ticker.C:
			// read from econet
			d, err := econetRead(cl, econet, econetusr, econetpw)
			if err != nil {
				slog.Error("could not read from econet", err)
				ticker.Reset(1 * time.Minute)
				continue
			}
			slog.Info("read from econet")
			// write to influx
			writePoint(ewapi, d, ctx)
			ticker.Reset(1 * time.Minute)
		case <-ctx.Done():
			return

		}
	}

}

func writePoint(ewapi api.WriteAPIBlocking, d EconetData, ctx context.Context) {
	// write all fields from curr struct to influx w/ current timestamp
	p := influxdb2.NewPoint("thermostat",
		map[string]string{"device": "boiler"},
		map[string]interface{}{
			"thermostat":             d.Curr.Thermostat,
			"pumpCOWorks":            d.Curr.PumpCOWorks,
			"fan2Exhaust":            d.Curr.Fan2Exhaust,
			"feederWorks":            d.Curr.FeederWorks,
			"feeder":                 d.Curr.Feeder,
			"mixerPumpWorks4":        d.Curr.MixerPumpWorks4,
			"lambdaSet":              d.Curr.LambdaSet,
			"mixerPumpWorks1":        d.Curr.MixerPumpWorks1,
			"mixerPumpWorks2":        d.Curr.MixerPumpWorks2,
			"mixerPumpWorks3":        d.Curr.MixerPumpWorks3,
			"statusCWU":              d.Curr.StatusCWU,
			"fuelLevel":              d.Curr.FuelLevel,
			"tempOpticalSensor":      d.Curr.TempOpticalSensor,
			"fanPower":               d.Curr.FanPower,
			"mixerTemp1":             d.Curr.MixerTemp1,
			"mixerTemp3":             d.Curr.MixerTemp3,
			"mixerTemp2":             d.Curr.MixerTemp2,
			"mixerTemp4":             d.Curr.MixerTemp4,
			"blowFan1Active":         d.Curr.BlowFan1Active,
			"statusCO":               d.Curr.StatusCO,
			"boilerPower":            d.Curr.BoilerPower,
			"feederOuter":            d.Curr.FeederOuter,
			"pumpCWUWorks":           d.Curr.PumpCWUWorks,
			"pumpCWU":                d.Curr.PumpCWU,
			"alarmOutput":            d.Curr.AlarmOutput,
			"tempUpperBuffer":        d.Curr.TempUpperBuffer,
			"fan":                    d.Curr.Fan,
			"lighter":                d.Curr.Lighter,
			"lambdaStatus":           d.Curr.LambdaStatus,
			"transmission":           d.Curr.Transmission,
			"fuelStream":             d.Curr.FuelStream,
			"lighterWorks":           d.Curr.LighterWorks,
			"mode":                   d.Curr.Mode,
			"alarmOutputWorks":       d.Curr.AlarmOutputWorks,
			"pumpSolar":              d.Curr.PumpSolar,
			"lambdaLevel":            d.Curr.LambdaLevel,
			"contactGZC":             d.Curr.ContactGZC,
			"blowFan1":               d.Curr.BlowFan1,
			"blowFan2":               d.Curr.BlowFan2,
			"tempLowerBuffer":        d.Curr.TempLowerBuffer,
			"tempCO":                 d.Curr.TempCO,
			"pumpCO":                 d.Curr.PumpCO,
			"contactGZCActive":       d.Curr.ContactGZCActive,
			"pumpCirculation":        d.Curr.PumpCirculation,
			"outerBoiler":            d.Curr.OuterBoiler,
			"tempCOSet":              d.Curr.TempCOSet,
			"outerBoilerWorks":       d.Curr.OuterBoilerWorks,
			"pumpFireplace":          d.Curr.PumpFireplace,
			"feederOuterWorks":       d.Curr.FeederOuterWorks,
			"mixerSetTemp4":          d.Curr.MixerSetTemp4,
			"boilerPowerKW":          d.Curr.BoilerPowerKW,
			"feeder2AdditionalWorks": d.Curr.Feeder2AdditionalWorks,
			"pumpSolarWorks":         d.Curr.PumpSolarWorks,
			"mixerSetTemp1":          d.Curr.MixerSetTemp1,
			"mixerSetTemp2":          d.Curr.MixerSetTemp2,
			"mixerSetTemp3":          d.Curr.MixerSetTemp3,
			"blowFan2Active":         d.Curr.BlowFan2Active,
			"tempCWUSet":             d.Curr.TempCWUSet,
			"pumpCirculationWorks":   d.Curr.PumpCirculationWorks,
			"tempFlueGas":            d.Curr.TempFlueGas,
			"fan2ExhaustWorks":       d.Curr.Fan2ExhaustWorks,
			"pumpFireplaceWorks":     d.Curr.PumpFireplaceWorks,
			"tempFeeder":             d.Curr.TempFeeder,
			"fanWorks":               d.Curr.FanWorks,
			"feeder2Additional":      d.Curr.Feeder2Additional,
		},
		time.Now())
	writeCtx, _ := context.WithTimeout(ctx, 5*time.Second)
	err := ewapi.WritePoint(writeCtx, p)
	if err != nil {
		slog.Error("could not write to influx", err)
	}
	// Flush writes
	err = ewapi.Flush(writeCtx)
	if err != nil {
		slog.Error("could not flush to influx", err)
	}

}

func econetRead(client *http.Client, ip, user, pass string) (EconetData, error) {
	url := fmt.Sprintf("http://%s/econet/regParams", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EconetData{}, err
	}
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		return EconetData{}, err
	}
	if resp.StatusCode != 200 {
		return EconetData{}, fmt.Errorf("econet returned status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EconetData{}, err
	}
	ret := EconetData{}
	err = json.Unmarshal(body, &ret)
	return ret, nil
}

func EconetHttpClient() *http.Client {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	return httpClient
}

/*
for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}
*/
