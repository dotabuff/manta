# Manta

[![Build Status](https://travis-ci.org/dotabuff/manta.svg?branch=master)](https://travis-ci.org/dotabuff/manta) [![GoDoc](https://godoc.org/github.com/dotabuff/manta?status.svg)](https://godoc.org/github.com/dotabuff/manta)

Manta is a fully functional Dota 2 replay parser written in [Go](https://golang.org), targeting the Source 2 (Dota 2 Reborn) game engine.

## Getting Started

Manta is a low-level replay parser, meaning that it will provide you access to the raw data in the replay, but doesn't provide any opinion on how that data should be structured for your use case. You'll need to create callback functions, inspect the raw data, and decide how you're going to use it.

## Usage

Get the code:

    go get github.com/dotabuff/manta

Use it to parse a replay:

```go
import (
  "log"
  "os"

  "github.com/dotabuff/manta"
  "github.com/dotabuff/manta/dota"
)

func main() {
  // Create a new parser instance from a file. Alternatively see NewParser([]byte)
  f, err := os.Open("my_replay.dem")
  if err != nil {
    log.Fatalf("unable to open file: %s", err)
  }
  defer f.Close()

  p, err := manta.NewStreamParser(f)
  if err != nil {
    log.Fatalf("unable to create parser: %s", err)
  }

  // Register a callback, this time for the OnCUserMessageSayText2 event.
  p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
    log.Printf("%s said: %s\n", m.GetParam1(), m.GetParam2())
    return nil
  })

  // Start parsing the replay!
  p.Start()

  log.Printf("Parse Complete!\n")
}
```

## Developing

To run `make update` you will need the latest version of the `protobuf` package:

`go get -u github.com/golang/protobuf/...`

You will also need GNU sed. To install GNU sed on Mac OS X:

```
# Install GNU sed
brew install gnu-sed
# Create a symlink in /usr/local/bin
ln -s /usr/local/bin/gsed /usr/local/bin/sed
# Ensure that /usr/local/bin is foremost in your PATH
export PATH="/usr/local/bin:$PATH"
```


## License

Manta is distributed under the [MIT license](https://github.com/dotabuff/manta/blob/master/LICENSE).

## Code of Conduct

Manta has adopted the [Contributor Covenant Code of Conduct](https://github.com/dotabuff/manta/blob/master/CONDUCT.md).

## Getting Help

The best place to ask questions about Dota 2 replay parsing is the #dota2replay channel on QuakeNet, where we're happy to answer any questions you may have. Please only open Github issues for actual bugs in manta, not questions about usage.

Looking to parse Source 1 (original Dota 2) replays? Take a look at [Yasha](https://github.com/dotabuff/yasha).

## Authors and Acknowledgements

Manta is maintained and development is sponsored by [Dotabuff](http://www.dotabuff.com), a leading Dota 2 community website with an emphasis on statistics. Manta wouldn't exist without the efforts of a number of people:

* [Michael Fellinger](https://github.com/manveru) built Dotabuff's Source 1 parser [yasha](https://github.com/dotabuff/yasha).
* [Robin Dietrich](https://github.com/invokr) built the C++ parser [Alice](https://github.com/AliceStats/Alice).
* [Martin Schrodt](https://github.com/spheenik) built the Java parser [clarity](https://github.com/skadistats/clarity).
* [Drew Schleck](https://github.com/dschleck), built an original C++ parser [edith](https://github.com/dschleck/edith).
