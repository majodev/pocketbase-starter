# github.com/majodev/pocketbase-starter

Template for extending [Pocketbase](https://pocketbase.io/) with [Go](https://pocketbase.io/docs/go-overview/) in a containerized environment via [VSCode remote containers](https://code.visualstudio.com/docs/remote/containers).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/majodev/pocketbase-starter/blob/master/LICENSE)
[![Build and Test](https://github.com/majodev/pocketbase-starter/actions/workflows/build-test-publish.yml/badge.svg)](https://github.com/majodev/pocketbase-starter/actions)

- [github.com/majodev/pocketbase-starter](#githubcommajodevpocketbase-starter)
    - [Requirements](#requirements)
    - [Quickstart](#quickstart)
    - [Visual Studio Code](#visual-studio-code)
    - [Production build](#production-build)
    - [Kubernetes deployment](#kubernetes-deployment)
    - [Uninstall](#uninstall)
  - [Maintainers](#maintainers)
  - [License](#license)

### Requirements

Requires the following local setup for development:

- [Docker CE](https://docs.docker.com/install/) (19.03 or above)
- [Docker Compose](https://docs.docker.com/compose/install/) (1.25 or above)
- [VSCode Extension: Remote - Containers](https://code.visualstudio.com/docs/remote/containers) (`ms-vscode-remote.remote-containers`)

This project makes use of the [Remote - Containers extension](https://code.visualstudio.com/docs/remote/containers) provided by [Visual Studio Code](https://code.visualstudio.com/). A local installation of Go/Pocketbase tool-chain is **no longer required** when using this setup.

### Quickstart

Create a new git repository through the GitHub template repository feature ([use this template](https://github.com/majodev/pocketbase-starter/generate)). You will then start with a **single initial commit** in your own repository. 

```bash
# First replace all occurances of 'github.com/majodev/pocketbase-starter' and then 'pocketbase-starter' with your own repository URI and project-name

# Then easily start the docker-compose dev environment through our helper
./docker-helper.sh --up
```

You should be inside the 'service' docker container with a bash shell.

```bash
development@94242c61cf2b:/app$ # inside your container...

# Shortcut for make init, make build, make info and make test
make all

# Print all available make targets
make help

# Start the pocketbase dev server pipeline and watch for changes to *.go files
make watch

# Execute all tests
make test

# Rebuild only after changes to files
make
```

### Visual Studio Code

Run `CMD+SHIFT+P` `Go: Install/Update Tools` **after** attaching to the container with VSCode to auto-install all golang related vscode extensions.

### Production build

```bash
docker build . -t pocketbase-starter
docker run -p 8090:8090 pocketbase-starter
```

### Kubernetes deployment

See the sample static manifests in `deploy/` for a simple deployment to a Kubernetes cluster.

### Uninstall

Simply run `./docker-helper --destroy` in your working directory (on your host machine) to wipe all docker related traces of this project (and its volumes!).

## Maintainers

- [Mario Ranftl - @majodev](https://github.com/majodev)


## License

[MIT](LICENSE) © Mario Ranftl