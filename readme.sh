#!/usr/bin/env bash

set -e

C="dist/caddy_darwin_amd64/caddy"

{
 echo "# My Own Caddy"
 echo "## Version"
 $C version | awk '{print "**"$1"** `"$2"`"}'

 echo "## Modules"
 $C list-modules | awk '{print "- ["$1"](https://caddyserver.com/docs/modules/"$1")"}'
 echo "## Info"
 echo '```text'
 $C build-info
 echo '```'

} > README.md
