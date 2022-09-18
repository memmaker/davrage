
# davrage - Because nothing works

_davrage_ is a simple webdav server that provides the following features:

- Single binary that runs under Windows, Linux and OSX.
- Authentication via HTTP-Basic.
- TLS support - if needed.
- A simple user management which allows user-directory-jails.
- Fixes the Golang WebDAV implementation bugs where an unreadable file or named pipe inside a WebDAV share kills the server.
- It just logs to the damn stdout.
- It doesn't use any fancy configuration frameworks or file formats.

It perfectly fits if you would like to give some people the possibility to upload, download or share files with common tools like the OSX Finder, Windows Explorer or Nautilus under Linux ([or many other tools](https://en.wikipedia.org/wiki/Comparison_of_WebDAV_software#WebDAV_clients)).

## Usage

### Podman run

You'll have to provide at least a data directory.
Authentication is optional.

    podman run -d \
     -e DR_BIND_TO_IP="0.0.0.0" \
     -e DR_AUTH_FILE="/auth_user" \
     -p 8000:8000 \
     -v "$WEBDAV_DATA_DIR":/tmp \
     -v "$AUTHFILE":/auth_user \
     --name webdav \
     --network podnet \
     ghcr.io/memmaker/davrage:latest


### Configuration via Environment Variables

    Address: "DR_BIND_TO_IP",   Default: "127.0.0.1"
    Port:    "DR_BIND_TO_PORT", Default: "8000"
    Prefix:  "DR_URL_PREFIX",   Default: "/"
    Dir:     "DR_ROOT",         Default: "/tmp"
    Realm:   "DR_AUTH_REALM",   Default: "dav-rage"

### TLS files

Are also configured via environment variables, duh!

    Certificate File: "DR_TLS_CERT"
    Key File:         "DR_TLS_KEY"

### Users

Set this Environment Variable to a file containing the users and their passwords.

    DR_AUTH_FILE

The file must be in the following format:

    user1:password1
    user2:password2
    user3:password3

So, one user per line, username and password separated by a colon.
Passwords are stored in **bcrypt** format.

You can use [gobcrypt](https://github.com/memmaker/gobcrypt) to create them.

## Installation

Download the binary. Done.

## Connecting

You could simply connect to the webdav server with a http(s) connection and a tool that allows the webdav protocol.

For example: Under OSX you can use the default file management tool *Finder*. Press _CMD+K_, enter the server address (e.g. `http://localhost:8000`) and choose connect.

## History

An update and demolition of the work done here: https://github.com/micromata/dave
