myhttp Tool
===========

This is a simple command-line tool written in Go that makes HTTP requests and prints the address of the request along with the MD5 hash of the response. It can perform requests in parallel and limit the number of parallel requests to prevent exhausting local resources.

Usage
-----

To use the tool, run the following command:

`./myhttp [flags] <URL1> <URL2> ... <URLn>`

where `URL1`, `URL2`, ..., `URLn` are the URLs of the pages you want to retrieve. The tool will print the address of the request along with the MD5 hash of the response.

Building
--------

To build the `myhttp` tool, you'll need to have the Go installed on your system. Once you've installed Go, follow these steps:

1.  Clone the repository:

    `git clone https://github.com/xmoutaz/myhttp.git`

2.  Change to the `myhttp` directory:

    `cd myhttp`

3.  Build the tool:

    `go build`

    This will create an executable file named `myhttp` in the current directory.

### Flags

The following flags are available:

*   `-parallel`: Sets the maximum number of parallel requests. Default is 10.

Examples
--------

`./myhttp http://www.adjust.com http://google.com http://www.adjust.com http://google.com`

`./myhttp adjust.com http://adjust.com`

`./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com http://google.com http://adjust.com http://yandex.com http://facebook.com http://twitter.com http://reddit.com/r/funny http://reddit.com/r/notfunny http://yahoo.com http://baroquemusiclibrary.com`

Flags and Options
-----------------

The tool accepts the following flags:

*   `-parallel`: Sets the maximum number of parallel requests. Default is 10.

Unit Tests
----------

This tool has been tested using the standard Go testing package. You can run the tests by executing the following command:

`go test ./...`

Dependencies
------------

This tool has no external dependencies beyond Go's standard library.
