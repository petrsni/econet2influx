package econet2influx

type EconetData struct {
	SettingsVer               int            `json:"settingsVer"`
	EditableParamsVer         int            `json:"editableParamsVer"`
	SchedulesVer              int            `json:"schedulesVer"`
	RemoteMenuVer             map[string]int `json:"remoteMenuVer"`
	CurrentDataParamsEditsVer int            `json:"currentDataParamsEditsVer"`
	Curr                      Curr           `json:"curr"`
}

type Curr struct {
	Thermostat             int     `json:"thermostat"`
	PumpCOWorks            bool    `json:"pumpCOWorks"`
	Fan2Exhaust            bool    `json:"fan2Exhaust"`
	FeederWorks            bool    `json:"feederWorks"`
	Feeder                 bool    `json:"feeder"`
	MixerPumpWorks4        bool    `json:"mixerPumpWorks4"`
	LambdaSet              int     `json:"lambdaSet"`
	MixerPumpWorks1        bool    `json:"mixerPumpWorks1"`
	MixerPumpWorks2        bool    `json:"mixerPumpWorks2"`
	MixerPumpWorks3        bool    `json:"mixerPumpWorks3"`
	StatusCWU              int     `json:"statusCWU"`
	FuelLevel              int     `json:"fuelLevel"`
	TempOpticalSensor      float64 `json:"tempOpticalSensor"`
	FanPower               float64 `json:"fanPower"`
	MixerTemp1             *int    `json:"mixerTemp1"`
	MixerTemp3             *int    `json:"mixerTemp3"`
	MixerTemp2             *int    `json:"mixerTemp2"`
	MixerTemp4             *int    `json:"mixerTemp4"`
	BlowFan1Active         bool    `json:"blowFan1Active"`
	StatusCO               int     `json:"statusCO"`
	BoilerPower            int     `json:"boilerPower"`
	FeederOuter            bool    `json:"feederOuter"`
	PumpCWUWorks           bool    `json:"pumpCWUWorks"`
	PumpCWU                bool    `json:"pumpCWU"`
	AlarmOutput            bool    `json:"alarmOutput"`
	TempUpperBuffer        float64 `json:"tempUpperBuffer"`
	Fan                    bool    `json:"fan"`
	Lighter                bool    `json:"lighter"`
	LambdaStatus           int     `json:"lambdaStatus"`
	Transmission           int     `json:"transmission"`
	FuelStream             float64 `json:"fuelStream"`
	LighterWorks           bool    `json:"lighterWorks"`
	Mode                   int     `json:"mode"`
	AlarmOutputWorks       bool    `json:"alarmOutputWorks"`
	PumpSolar              bool    `json:"pumpSolar"`
	LambdaLevel            int     `json:"lambdaLevel"`
	ContactGZC             bool    `json:"contactGZC"`
	BlowFan1               bool    `json:"blowFan1"`
	BlowFan2               bool    `json:"blowFan2"`
	TempLowerBuffer        float64 `json:"tempLowerBuffer"`
	TempCO                 float64 `json:"tempCO"`
	PumpCO                 bool    `json:"pumpCO"`
	ContactGZCActive       bool    `json:"contactGZCActive"`
	PumpCirculation        bool    `json:"pumpCirculation"`
	OuterBoiler            bool    `json:"outerBoiler"`
	TempCOSet              int     `json:"tempCOSet"`
	OuterBoilerWorks       bool    `json:"outerBoilerWorks"`
	PumpFireplace          bool    `json:"pumpFireplace"`
	FeederOuterWorks       bool    `json:"feederOuterWorks"`
	MixerSetTemp4          int     `json:"mixerSetTemp4"`
	BoilerPowerKW          float64 `json:"boilerPowerKW"`
	Feeder2AdditionalWorks bool    `json:"feeder2AdditionalWorks"`
	PumpSolarWorks         bool    `json:"pumpSolarWorks"`
	MixerSetTemp1          int     `json:"mixerSetTemp1"`
	MixerSetTemp2          int     `json:"mixerSetTemp2"`
	MixerSetTemp3          int     `json:"mixerSetTemp3"`
	BlowFan2Active         bool    `json:"blowFan2Active"`
	TempCWUSet             int     `json:"tempCWUSet"`
	PumpCirculationWorks   bool    `json:"pumpCirculationWorks"`
	TempFlueGas            float64 `json:"tempFlueGas"`
	Fan2ExhaustWorks       bool    `json:"fan2ExhaustWorks"`
	PumpFireplaceWorks     bool    `json:"pumpFireplaceWorks"`
	TempFeeder             float64 `json:"tempFeeder"`
	FanWorks               bool    `json:"fanWorks"`
	Feeder2Additional      bool    `json:"feeder2Additional"`
}
