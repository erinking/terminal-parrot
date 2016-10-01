# :parrot: for your terminal

![demo](http://dropit.velvetcache.org.s3.amazonaws.com/jmhobbs/NzczFOYq4g/termbox-parrot-color.gif)

## Using

Either grab a build on the [releases page](https://github.com/jmhobbs/terminal-parrot/releases) or clone and run...

    go get -u github.com/nsf/termbox-go
    make

You can also build a docker image and run it in a container with...

    make docker
    docker run -it --rm partyparrot (-args)

### Quitting

Hit the escape key to quit.

### -loops

You can limit your parrots enthusiasm with the `-loops` flag.

### :fastparrot:

Set the frame delay with the `-delay` flag (defaults to 75, use 25 for :fastparrot:)

### :aussieparrot:

Use `-orientation aussie` for our friends down under.

## Thanks

Idea from seeing [this tweet from @rachsmithtweets](https://twitter.com/rachsmithtweets/status/742785722290212868)

Thanks to [termbox-go](https://github.com/nsf/termbox-go) for making it easy.

Thanks to [jp2a](https://csl.name/jp2a/) for nice ASCII art conversion.
