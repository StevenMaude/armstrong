# armstrong

## About

Retrieve EPO.BIN GPS data from Garmin's servers on any platform that you
can compile a Go binary for, rather than only those that support Garmin
Connect. Since Garmin's software doesn't run natively on Linux, it's
useful there.

EPO.BIN is GPS Extended Prediction Orbit (EPO) satellite data valid for
7 days, used to help speed up GPS locking on Garmin devices that support
this.

For Forerunner devices, the file should be copied to
`GARMIN/REMOTESW/EPO.BIN` on the watch. 

## Alternatives

There are other existing options, e.g. running `curl` and then a short
Ruby script as detailed in this [blog
post](https://www.kluenter.de/garmin-ephemeris-files-and-linux/), or
using [postrunner](https://github.com/scrapper/postrunner), a more
complete application which also uses Ruby. However, this code can be
compiled into a standalone binary, for different platforms, and then run
without any dependencies.

## Usage

### Build

With Go installed:

* `go build armstrong.go`

Or, use the `Dockerfile` via `make`:

* `make build`

### Run

Just run the compiled `armstrong` binary. `EPO.BIN` should be generated
in the same directory.

## Credits

Code uses the POST data and idea from the [blog post
here](https://www.kluenter.de/garmin-ephemeris-files-and-linux/) and
is also based on the GPL v2 licensed code in
[postrunner](https://github.com/scrapper/postrunner), specifically
[`EPO_Downloader.rb`](https://github.com/scrapper/postrunner/blob/93b5fc82a4d8ef9c6587abdd43b2e91ea829cfda/lib/postrunner/EPO_Downloader.rb). Go code is my contribution.
