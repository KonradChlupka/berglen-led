# Berglen LED

@jamesjarvis, @KonradChlupka

Somewhere in an east london flat, there is an LED strip running this code.

You can too!

## How to run this code

If you already have a raspberry pi with a programmable LED strip wired up (TODO: write up how to do this), then all you need to do is either grab the latest binary from the [releases](https://github.com/KonradChlupka/berglen-led/releases), or you can build from source.

If you have the latest binary installed on the RPI, simply run it with:

```bash
sudo ./led-lights
```

Or if you want to build from source, you can build it with:

```bash
./build.sh
```

And then scp the generated file over to your PI.

## Command options

There are a few CLI options for the binary, you can discover these by running

```bash
sudo ./led-lights --help
```

## HTTP endpoints

You can even modify the behaviour while it is running, simply make a HTTP GET request to the following endpoints:

- `curl http://pi.local:8888/wipe` - Pauses the RGB background program and runs a blue wipe across the LEDs, before resuming the background.
- ... more coming soon™️!
