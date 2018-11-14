# hashii

## What?

ANSI colorized ASCII versions of the HashiCorp logo emitted to your terminal:

![](https://github.com/brianshumate/hashii/blob/master/share/screenshot.png)

## Why?

I am casually learning the Go programming language for various values of great good!

> NOTE: this is just a personal learning example and so it is not open to collaboration; issues are intentionally not enabled for this repository.

## How?

If you have an ANSI color capable terminal and a Go environment, you can install the `hashii` command like this:

```
$ go get -u github.com/brianshumate/hashii
```

Then, run it:

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

You can also clone this repository, and use `make`:

```
$ git clone https://github.com/brianshumate/hashii
$ cd hashii
$ make
$ ./hashii
```

### Available Colors

- blue
- cyan
- green
- magenta
- mix
- random
- red
- yellow

### BONUS ✨

Need to send a beacon — a signal of your HashiJoy? Well neighbor, try:

𝗗𝗮𝘇𝘇𝗹𝗲 𝗠𝗼𝗱𝗲

```
$ hashii -dazzle -size=large
```

🎈🎉  *Great fun at parties*!

**Enjoy**!

## License?

BSD

## Copyright

The HashiCorp logo is Copyright HashiCorp and is used with permission.

## Who?

[Brian Shumate](https://github.com/brianshumate)
