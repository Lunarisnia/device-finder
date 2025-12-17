package finder

import (
	"github.com/Lunarisnia/device-finder/internal/pingy"
	"github.com/Lunarisnia/device-finder/internal/tinycli"
)

func Run(ctx tinycli.Context) error {
	deviceIp := ctx.Argument("--ip")
	received, err := pingy.Ping(deviceIp)
	if err != nil {
		return err
	}
	if received > 0 {
		// TODO: Call the discord bot API
	}
	return nil
}
