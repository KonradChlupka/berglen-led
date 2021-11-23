package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KonradChlupka/berglen-led/colourutils"
	"github.com/KonradChlupka/berglen-led/engine"
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
	"github.com/urfave/cli/v2"
)

const (
	brightness = 100
	numLEDs    = 240

	flagColourWipe = "colour_wipe_colour"
)

func main() {
	opt := ws281x.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = numLEDs

	ws, err := ws281x.MakeWS2811(&opt)
	if err != nil {
		log.Fatalf("Failed to initiate the dumbass LEDs: %s", err)
	}

	leds := engine.NewLightstrip(ws)

	err = leds.Init()
	if err != nil {
		log.Fatalf("Failed to initiate the light strip: %s", err)
	}
	defer leds.Close()

	app := &cli.App{
		Name:  "Berglen LED",
		Usage: "Shine bright like an east london flat",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagColourWipe,
				Aliases: []string{"c"},
				Usage:   "Colour wipe colour",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("Hello world")

			colourString := ctx.String(flagColourWipe)
			colour := colourutils.WHITE
			switch colourString {
			case "RED":
				colour = colourutils.RED
			default:
				colour = colourutils.WHITE
			}

			return leds.ColourWipe(colour)
		},
	}

	app.EnableBashCompletion = true

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
