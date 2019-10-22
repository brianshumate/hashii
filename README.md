# hashii

[![No Maintenance Intended](http://unmaintained.tech/badge.svg)](http://unmaintained.tech/)

## What?

ANSI colorized ASCII versions of the HashiCorp logo emitted to your terminal:

![](https://github.com/brianshumate/hashii/blob/master/share/screenshot.png)

## Why?

I am casually learning the Go programming language for various values of great good!

> NOTE: this is just a personal learning example and so it is not open to collaboration; issues are intentionally not enabled for this repository.

## How?

If you have an ANSI color capable terminal and a Go environment, you can download, compile, and run the `hashii` command like this:

### Install hashii
```
$ go get -u github.com/brianshumate/hashii
```

You can also clone this repository, and use `make`, but you will still need the Go language environment installed:

```
$ git clone https://github.com/brianshumate/hashii
$ cd hashii
$ make
```

### Run hashii

```
$ hashii
```

or

```
$ hashii -size=medium -color=random
```

or

```
$ hashii -size=large -color=mix
```

The available colors are standard ANSI foreground colors: _blue_, _cyan_, _green_, _magenta_, _red_, and _yellow_; there are also two special color modes: _mix_ and _random_.

### BONUS ✨

Need to send a beacon — a signal of your HashiJoy? Well neighbor, try:

𝗗𝗮𝘇𝘇𝗹𝗲 𝗠𝗼𝗱𝗲

```
$ hashii -dazzle -size=large
```

🎈🎉  *Great fun at parties*!

**Enjoy**!

## License?

BSD 2 Clause

## Copyright

The HashiCorp 'H' logo is a registered trademark of HashiCorp and is used with permission.

## Who?

[Brian Shumate](https://github.com/brianshumate)
