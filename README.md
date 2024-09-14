# WASA Photo

**Copyright © 2023 Alessio Borgi**

**PROJECT SCOPE**: Build a complete Web Architecture for an Instagram-Like Platform, from the API Documentation, passing through the Backend and the Frontend, and ending with the Containerisation.

**PROJECT RESULTS**:
- OpenAPI Documentation.
- Go Backend.
- JavaScript+VueJS+Axios+HTML+CSS Frontend.
- Docker Containerizzzation. 
- Agile Development and Iterative Behaviour on the site development based on User Feedbacks. 
- 25 API Developed.
	- Possibility to Authenticate the user using a Bearer Token.
	- Possibility to Create, Modify, Retrieve, Delete an User Account. 	 
	- Possibility to Follow, UnFollow another User Account.
	- Possibility to view Followers, Following and Banned Accounts.
	- Possibility to Ban, UnBan another User Account.
	- Possibility to Upload, Delete a Photo onto a User Account.
	- Possibility to Like, UnLike, Comment, UnComment a User Account's Photo.
	- Possibility to Scroll all the Stream of photos related to users we follow and possiblity to interact with them (Comment, Like, etc...)

**PROJECT REPOSITORY**: https://github.com/alessioborgi/WASA_Photo

## Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an example of an API server
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is an example of a web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding

Other project files include:
* `open-npm.sh` starts a new (temporary) container using `node:lts` image for safe web frontend development (you don't want to use `npm` in your system, do you?)

## Go vendoring

This project uses [Go Vendoring](https://go.dev/ref/mod#vendoring). You must use `go mod vendor` after changing some dependency (`go get` or `go mod tidy`) and add all files under `vendor/` directory in your commit.

For more information about vendoring:

* https://go.dev/ref/mod#vendoring
* https://www.ardanlabs.com/blog/2020/04/modules-06-vendoring.html

## Node/NPM vendoring

This repository contains the `webui/node_modules` directory with all dependencies for Vue.JS. You should commit the content of that directory and both `package.json` and `package-lock.json`.

## How to set up a new project from this template

You need to:

* Change the Go module path to your module path in `go.mod`, `go.sum`, and in `*.go` files around the project
* Rewrite the API documentation `doc/api.yaml`
* If no web frontend is expected, remove `webui` and `cmd/webapi/register-webui.go`
* If no cronjobs or health checks are needed, remove them from `cmd/`
* Update top/package comment inside `cmd/webapi/main.go` to reflect the actual project usage, goal, and general info
* Update the code in `run()` function (`cmd/webapi/main.go`) to connect to databases or external resources
* Write API code inside `service/api`, and create any further package inside `service/` (or subdirectories)

## How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## License

See [LICENSE](LICENSE).
