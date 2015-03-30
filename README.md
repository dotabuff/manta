# Manta: Dota 2 (Source 1) replay parser written in Go

This project is an evolution from [Yasha](https://github.com/dotabuff/yasha),
which is the Source 1 parser.

The aim of Manta is to make the parser easier to understand and allow much more
performance.

In particular, Yasha wasn't able to parallelize parsing and
processing, and couldn't skip any packets, with this new implementation we're
planning to change that.

This will make porting projects that currently use Yasha difficult, so beware.

## Installation

Simple as:

    go get github.com/dotabuff/manta

And in your code:

    import "github.com/dotabuff/manta"

Please be aware that you _can't import_ Manta and Yasha in the same binary!
The Protocol Buffer definitions conflict, and will panic.

## License

MIT, see the LICENSE file.

## Help

If you have any questions, just ask manveru in the #dota2replay channel on QuakeNet.

## Acknowledgements

* [Robin Dietrich](https://github.com/invokr) worked out how the new CDemoPackets are
  encoded and helped a lot in other areas as well.
