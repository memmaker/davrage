
# davrage - Because nothing works

_davrage_ is a simple webdav server that provides the following features:

- Single binary that runs under Windows, Linux and OSX.
- Authentication via HTTP-Basic.
- TLS support - if needed.
- A simple user management which allows user-directory-jails.
- Fixes the Golang WebDAV implementation bugs where an unreadable file or named pipe inside a WebDAV share kills the server.

It perfectly fits if you would like to give some people the possibility to upload, download or share files with common tools like the OSX Finder, Windows Explorer or Nautilus under Linux ([or many other tools](https://en.wikipedia.org/wiki/Comparison_of_WebDAV_software#WebDAV_clients)).

## Usage

TODO

## Installation

TODO

## Connecting

You could simply connect to the webdav server with a http(s) connection and a tool that allows the webdav protocol.

For example: Under OSX you can use the default file management tool *Finder*. Press _CMD+K_, enter the server address (e.g. `http://localhost:8000`) and choose connect.

## History

An update and demolition of the work done here: https://github.com/micromata/dave