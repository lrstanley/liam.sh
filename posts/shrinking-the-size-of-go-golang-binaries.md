Title: Shrinking the size of Go (Golang) Binaries
Show: true
Time: 06 Dec 16 19:18 EST

I recently started using Go a little over 4 months ago. It was an amazing change from things like Python, NodeJS, and other misc. languages in that it was very opinionated. The syntax was sound, it performed very well, it had static typing, and it compiled down into a single binary.

Single binaries are amazing, especially in the way that Go handles them. This meant that I could easily distribute my code without a bunch of shared libraries and complex dependencies, it was as simple as dropping the application or utility onto the server and that was it. It allowed a very flexible distribution release system that made pushing out updates **way** easier. The only downside, was that the binary size was quite large. Ever since (I believe) Go 1.4, the binary size has been slowly increasing, however with the release of Go 1.7, just recompiling the same exact program with the new Go release, the binary size was much smaller.

To give a rough estimate, a small program with a few linked Go packages, a Go program may compiled down into a single binary that may be 5mb in size, or 15mb in size. This was still very manageable, however I wanted smaller, _**much smaller**_.


## How to shrink the size of your Go binary?

First off, I would recommend ensuring you are compiling with the latest Go release (at the time of writing this, Go 1.7.4). Secondly, we must first understand what is within the Go binary. You have your application code that is compiled down to ASM/machine code, but what else is there? Well, there are things like DWARF and symbol tables. These are a visual mapping between the original code, and the sections of code that were translated into ASM/machine code. These _symbol tables_ basically allow you to debug the program using standard GNU Unix tools like GDB, or "GNU Debugger". In most cases these symbol tables aren't really necessary, especially considering if your code is fairly clean.

How do we get rid of them? With many packages submitted to things like the Debian system repositories, they recommend stripping the DWARF/symbol tables from the executables when they are submitted, to ensure that they are much smaller. We can do this during the build process of the binary. Below is an example of how this may work:

```
$ cd $GOPATH/src/github.com/username/package
$ go build -ldflags "-s -w"
```

Ok, so what kind of decrease/margin do we get from just this? (running this against my [Automated Site Testing Utility, Marill](https://github.com/lrstanley/marill)):

Before:

```
$ go build
$ ls -lah marill
-rwxrwxr-x 1 liam liam 9.2M Dec  6 18:22 marill
```

After:

```
$ go build -ldflags "-s -w"
$ ls -lah marill
-rwxrwxr-x 1 liam liam 5.8M Dec  6 18:23 marill
```

That's 63% of the original size! awesome! But, wait.. what _ELSE_ can we do to decrease the size?

## More awesomeness of shrinkage you say?

The next thing we can do, because the binary is now mainly just machine code, is to try and compress it. What we are going to use next is called [Ultimate Packer for eXecutables](https://upx.github.io/) (or UPX for short). What this does is compresses parts of the binary that would give a considerable decrease in size using many different passes, and techniques. So, when your application is executed, it will decompress it at runtime (this only adds a minor amount of delay during application startup, usually less than 100ms for my applications, commonly under 50ms). Let's start with installing it.

For Ubuntu 16.04, the package within the Ubuntu repositories is named `upx-ucl`, however it may differ in your package manager/repos.

```
$ apt install upx-ucl
```

Now ensure that you have built the binary with the above commands. We will then want to decide what level of shrinking we want. All the way from very fast (`-1`, fast, but binary will be larger), or best (`-best`, smaller binary, but will take longer to complete). `--brute` also exists, which will attempt quite a few more times to make it even smaller. You can use `-#` for an in between from fast to best as well (e.g. `-4`, `-6`, etc).

Let's start off with fast. What size margins do we get from this? (remember the size before was **5.8M**):

```
$ upx -1 -q marill
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2013
UPX 3.91        Markus Oberhumer, Laszlo Molnar & John Reiser   Sep 30th 2013

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   6073248 ->   2503732   41.23%  linux/ElfAMD   marill

Packed 1 file.
```

**41.23%** (2.4M), awesome! That means that now we're at a total combined percentage of **26%!** But, ok, come on... let's push it farther. How low can we go?

Let's try `--best` with `--brute`...:

```
$ upx --best --brute marill
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2013
UPX 3.91        Markus Oberhumer, Laszlo Molnar & John Reiser   Sep 30th 2013

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   6073248 ->   1574868   25.93%  linux/ElfAMD   marill

Packed 1 file.
```

**25.93%, which is 1.6M!** AWESOME! That's **17%** of the original size! 1.6M is very manageable, and makes distributing your application/utility much faster.

## What to do from here?

If you are following the above guide, you are likely making a user-facing application, and not a library. Meaning you probably have a Makefile or some form of script to build it for you, if your application requires external files, you are building things into a **.tar.gz** or **.zip** file, etc. How I currently tackle the above processes in my applications and utilities is by using a `Makefile`. See a very long example of one of these [here](https://github.com/lrstanley/marill/blob/master/Makefile) or maybe [here](https://github.com/lrstanley/links.ml/blob/master/Makefile).

Below is a good starting point for a Go project's Makefile:

```Makefile
.DEFAULT_GOAL := build

BINARY=your-package-name

LD_FLAGS += -s -w

compress:
	(which upx > /dev/null && upx -9 -q ${BINARY} > /dev/null) || echo "UPX not installed"

build:
	rm -vf ${BINARY}
	go build -ldflags "${LD_FLAGS}" -v -o ${BINARY}
```

Simply run the above with:

```
$ make build compress
```

Or without compression:

```
$ make build # (or just 'make' will select build by default)
```

## Caveats?

Using the above methods, there are a few caveats and downsides that should be noted, as shown below.

   * If you are using Go pre-1.5, I would not recommend using the first method to shrink the size of your binaries. Simply because users have experienced very weird crashes and issues with those flags pre-1.5.
   * If you are using Go pre-1.5, you will also likely need to change the syntax of the ld-flags functionality, because this has changed in 1.5 to use the method I used above.
   * UPX does not work with all binary types. You can run `upx --help` to get a full list of what is supported, and what isn't. Most x86 and x64/amd binary types are supported, however you may have issues compressing things on arm, and freebsd (and other less popular architectures/binary types).
   * UPX using the `--brute` method and `--best` method does take a bit of time (1-2 minutes on my i7 laptop, with 12gb of ram). However, just dropping `--brute` will make it much faster with still a considerable size decrease.
   * As mentioned above with using the `-s` and `-w` `LD_FLAG` options, this may make using things like GDB much more difficult, though this should not be an issue in most cases (and, you should be using Delve!)


Now _**Go**_ have some fun with your small binaries! (see what I did there!)
