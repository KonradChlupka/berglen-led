package main

import (
	"fmt"
	"os"

	"github.com/KonradChlupka/berglen-led/engine"
	"github.com/KonradChlupka/berglen-led/server"
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
	"github.com/urfave/cli/v2"
)

const (
	defaultBrightness = 255
	numLEDs           = 240

	flagBrightness = "brightness"
	flagColourWipe = "colour_wipe_colour"
	flagRainbow    = "rainbow"
)

func main() {
	app := &cli.App{
		Name:  "Berglen LED",
		Usage: "Shine bright like an east london flat",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagBrightness,
				Aliases: []string{"b"},
				Usage:   "Brightness [0-255]",
			},
			&cli.StringFlag{
				Name:    flagColourWipe,
				Aliases: []string{"c"},
				Usage:   "Colour wipe colour",
			},
			&cli.BoolFlag{
				Name:    flagRainbow,
				Aliases: []string{"r"},
				Usage:   "Rainbow setting",
			},
		},
		Action: func(ctx *cli.Context) error {
			brightness := ctx.Int(flagBrightness)
			if brightness < 0 || brightness > 255 {
				return fmt.Errorf("brightness requires a value of 0>=b<=255, not %d", brightness)
			}
			if !ctx.IsSet(flagBrightness) {
				brightness = defaultBrightness
			}

			// Setup light strip connection.
			opt := ws281x.DefaultOptions
			opt.Channels[0].Brightness = brightness
			opt.Channels[0].LedCount = numLEDs

			ws, err := ws281x.MakeWS2811(&opt)
			if err != nil {
				return fmt.Errorf("failed to initiate the dumbass LEDs: %w", err)
			}

			leds := engine.NewLightstrip(ws)

			err = leds.Init()
			if err != nil {
				return fmt.Errorf("failed to initiate the light strip: %w", err)
			}
			defer leds.Close()

			server := server.NewServer(leds)
			return server.Serve()

			// // If rainbow option applied, run that.
			// isRainbow := ctx.Bool(flagRainbow)
			// if isRainbow {
			// 	fmt.Println("Starting rainbow")
			// 	return leds.RainbowRGB(ctx.Context)
			// }

			// // Otherwise run colour wipe.
			// colourString := ctx.String(flagColourWipe)
			// colour := colourutils.OFF
			// switch colourString {
			// case "OFF":
			// 	colour = colourutils.OFF
			// case "RED":
			// 	colour = colourutils.RED
			// case "BLUE":
			// 	colour = colourutils.BLUE
			// case "GREEN":
			// 	colour = colourutils.GREEN
			// default:
			// 	colour = colourutils.WHITE
			// }

			// fmt.Printf("Starting colour wipe")
			// return leds.ColourWipe(colour)
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
