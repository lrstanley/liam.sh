Title: Makefiles for Go projects
Show: true
Time: 19 Aug 17 22:16 UTC

I've started releasing enough Go projects that making a user-distributed tool comes close in line with Makefiles to simplify my workflow when:

   * Pushing code
   * Generating clean, properly named binaries for multiple distributions and architectures
   * Managing my workspace (cleaning up temporary files)
   * Compressing binaries (see [Shrinking the size of go binaries](https://liam.sh/post/shrinking-the-size-of-go-golang-binaries) post)
   * Updating and managing dependencies

### An example

The below example is pulled from my recently released [nagios-notify-irc](https://github.com/lrstanley/nagios-notify-irc) project. It has entries for releasing, cleaning, and managing the dependencies of the project.

```
.DEFAULT_GOAL := build

GOPATH := $(shell go env | grep GOPATH | sed 's/GOPATH="\(.*\)"/\1/')
PATH := $(GOPATH)/bin:$(PATH)
export $(PATH)

BINARY=notify-irc
LD_FLAGS += -s -w

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

release: clean fetch ## Generate a release, but don't publish to GitHub.
	$(GOPATH)/bin/goreleaser --skip-validate --skip-publish

publish: clean fetch ## Generate a release, and publish to GitHub.
	$(GOPATH)/bin/goreleaser

snapshot: clean fetch ## Generate a snapshot release.
	$(GOPATH)/bin/goreleaser --snapshot --skip-validate --skip-publish

update-deps: fetch ## Updates all dependencies to the latest available versions.
	$(GOPATH)/bin/govendor add +external
	$(GOPATH)/bin/govendor remove +unused
	$(GOPATH)/bin/govendor update +external

fetch: ## Fetches the necessary dependencies to build.
	test -f $(GOPATH)/bin/govendor || go get -u -v github.com/kardianos/govendor
	test -f $(GOPATH)/bin/goreleaser || go get -u -v github.com/goreleaser/goreleaser
	$(GOPATH)/bin/govendor sync

clean: ## Cleans up generated files/folders from the build.
	/bin/rm -rfv "dist/" "${BINARY}"

build: fetch ## Builds the application.
	go build -ldflags "${LD_FLAGS}" -i -v -o ${BINARY}
```

The benefit of this task-style workflow is that you can combine specific tasks with other dependant tasks, like the `release` task which runs `clean` and `fetch` before it starts itself.

In addition, for Go projects, a useful combination is a `Makefile` with [GoReleaser](https://github.com/goreleaser/goreleaser). GoReleaser is a utility written in Go which allows streamlining making archives of binaries with special treatment for specific operating systems (e.g. making a Windows binary contained in a **.zip** file), and allows uploading them to Github using Github's release system. It also supports creating **rpm** and **deb** archives (amongst others) by taking advantage of [FPM](https://github.com/jordansissel/fpm). GoReleaser also has `pre` and `post` hooks, which let you run commands before the binaries are placed into archives.

### Compressing using GoReleaser

Below is an example of a Makefile task (`$ make compress`) that you could call as a GoReleaser `hook.post`, which will compress all supported binaries using the [UPX](https://upx.github.io/) binary compression utility. This allows for streamlined, small, compressed archives which can easily be distributed for many systems. Below is an excerpt which you could use:

```
compress:
	(which /usr/bin/upx > /dev/null && find dist/*/* | xargs -I{} -n1 -P 4 /usr/bin/upx --best "{}") || echo "not using upx for binary compression"
```

This will ensure that if `upx` isn't found/isn't installed, it won't fail outright. In addition, with the combination of `xargs`, this will compress 4 binaries in parallel. May want to adjust if you have less than 4 cores. Adjust `--best` to something less intense to improve the speed of how long UPX takes (but trading for less of a compression size`). See my [Shrinking the size of go binaries](https://liam.sh/post/shrinking-the-size-of-go-golang-binaries) post for more info on the useful arguments to UPX.

### Help documentation

Lastly, if you look over the Makefile at the top of this post, you will notice a **help** task (`$ make help`). This will list all of the tasks in the Makefile, and will pull the double hash (`##`) documentation lines and print them out, since this isn't a native feature of `make`. An example:

```
$ make help
release                        Generate a release, but don't publish to GitHub.
publish                        Generate a release, and publish to GitHub.
snapshot                       Generate a snapshot release.
update-deps                    Updates all dependencies to the latest available versions.
fetch                          Fetches the necessary dependencies to build.
clean                          Cleans up generated files/folders from the build.
build                          Builds the application.
```

My use of Makefiles will likely continue to evolve over time for Go projects, since the utility/tooling space in the Go environment and community is ever changing and improving. However, for now, this is already tremendously useful over the previous way my workflow for Python projects worked.
