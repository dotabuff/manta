# Manta

[![Build Status](https://travis-ci.org/dotabuff/manta.svg?branch=master)](https://travis-ci.org/dotabuff/manta)

Manta is a Dota 2 replay parser written in [Go](https://golang.org). It targets the Source 2 (Dota 2 Reborn) game engine. To parse Source 1 (original Dota 2) replays, take a look at [Yasha](https://github.com/dotabuff/yasha).

**Project Status:**

- Dota 2 Reborn (Source 2) is now the main Dota 2 client. Yay!
- Manta is structurally feature complete and in use serving production workloads at Dotabuff.
- Manta currently handles nearly all packets, including the packet entities.

## Getting Started

Manta is a low-level replay parser, meaning that it will provide you access to the raw data in the replay, but doesn't provide any opinion on how that data should be structured for your use case. You'll need to create callback functions, inspect the raw data, and decide how you're going to use it.

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

You cannot import Manta and Yasha in the same binary due to protocol buffer conflicts.

## License

MIT, see the LICENSE file.

## Help

If you have any questions, you can find us in the #dota2replay channel on QuakeNet.

## Authors and Acknowledgements

Manta is maintained and development is sponsored by [Dotabuff](http://www.dotabuff.com), a leading Dota 2 community website with an emphasis on statistics. Manta wouldn't exist without the efforts of a number of people:

* [Michael Fellinger](https://github.com/manveru) built Dotabuff's Source 1 parser [yasha](https://github.com/dotabuff/yasha).
* [Robin Dietrich](https://github.com/invokr) built the C++ parser [Alice](https://github.com/AliceStats/Alice).
* [Martin Schrodt](https://github.com/spheenik) built the Java parser [clarity](https://github.com/skadistats/clarity).
* [Drew Schleck](https://github.com/dschleck), built an original C++ parser [edith](https://github.com/dschleck/edith).
