package api

import (
	"log"

	shell "github.com/vaslabs/pi-web-agent/internal"
)

type Os_Info_Response struct {
	Id               string ""
	Version_Id       string ""
	Version_Codename string ""
}

func OS_Info() Os_Info_Response {
	return OS_Info_From("/etc", "os-release")
}

func OS_Info_From(path string, filename string) Os_Info_Response {
	config, err := shell.Read_Env_Config(path, filename)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
		return Os_Info_Response{}
	} else {
		return Os_Info_Response{
			config.GetString("ID"),
			config.GetString("VERSION_ID"),
			config.GetString("VERSION_CODENAME"),
		}
	}
}

type Temperature_Response struct {
	Temp string
}

func Measure_Temperature() Temperature_Response {
	return Measure_Temperature_From("/opt/vc/bin/vcgencmd", "measure_temp")
}

func Measure_Temperature_From(command string, args ...string) Temperature_Response {
	output, error := shell.RunSingle(command, args...)
	if error != nil {
		log.Fatalf("Error while measuring temperature: %s", error.Error())
		return Temperature_Response{"unknown"}
	} else {
		config, err := shell.ReadConfigFromString(output)
		if err != nil {
			log.Fatalf("Error reading temperature output %s", err.Error())
			return Temperature_Response{"unknown"}
		} else {
			return Temperature_Response{config.GetString("temp")}
		}
	}
}