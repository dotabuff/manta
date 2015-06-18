# Manta

[![Build Status](https://travis-ci.org/dotabuff/manta.svg?branch=master)](https://travis-ci.org/dotabuff/manta)

Manta parses Dota 2 replays based on the *currently unreleased* Source 2 game engine.

If you're interested in parsing current Dota 2 replays (which are played on the Source 1 game engine), take a look at [Yasha](https://github.com/dotabuff/yasha).

*Source 2 Status:* The Dota 2 Source 2 client has been released as a beta

*Manta Status:* Manta is currently incomplete and unusable.

## Improvements from Yasha

Compared to Yasha (the Dotabuff Source 1 parser), Manta will be easier to understand and allow much more performance. In particular, Yasha wasn't able to parallelize parsing and processing, and couldn't skip any packets. With this new implementation we're planning to change that.

This will make porting projects that currently use Yasha difficult, so beware.

**Warning:** Please be aware that you *cannot import Manta and Yasha in the same binary*! The Protocol Buffer definitions conflict, and will panic.

## Usage

Get the code:

    go get github.com/dotabuff/manta

Use it to parse a replay:

```go
import (
  "github.com/dotabuff/manta"
  "github.com/dotabuff/manta/dota"
)

func main() {
  // Create a new parser instance from a file. Alternatively see NewParser([]byte)
  p, _ := manta.NewParserFromFile("my_replay.dem")

  // Register a callback, this time for the OnCUserMessageSayText2 event.
  p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
    fmt.Printf("%s said: %s", m.GetParam1(), m.GetParam2())
  })

  // Start parsing the replay!
  p.Start()
}
```

## License

MIT, see the LICENSE file.

## Help

If you have any questions, you can find us in the #dota2replay channel on QuakeNet.

## Acknowledgements

* [Robin Dietrich](https://github.com/invokr) worked out how the new CDemoPackets are encoded and helped a lot in other areas as well. See his [alice2](https://github.com/invokr/alice2) project for a C++ implementation (work in progress).
