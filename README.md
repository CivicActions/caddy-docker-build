# caddy-docker-build

Simple Caddy Docker build, using go modules for plugins.

The existing [caddy-builder](https://github.com/abiosoft/caddy-docker/blob/master/BUILDER.md) project is good for a quick start, but can be non deterministic as plugins and dependencies are not strictly versioned, leading to flakey builds.

This is a simpler more direct approach - to use the image, simply use this as your FROM in your web projects Dockerfile, copy in your Caddyfile(s) to /etc, your web files to /srv and adjust the CMD with caddy args as needed.

To customize your own Caddy build:
 * Fork this repository
 * Alter the modules needed in main.go and the versions/mappings required in go.mod
 * Run `go get -u` to update go.sum as needed
 * Adjust the `apk add` command for any other system dependencies needed
 * Run `docker build .` to test
 * Commit and add a CI or Docker Hub build as you prefer
