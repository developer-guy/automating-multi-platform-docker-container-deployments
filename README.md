# automating-multi-platform-docker-container-deployments
Automating Multi-Platform Docker Container Deployments in Kubernetes with ArgoCD presentation at Docker Istanbul Community Meetup Group


### How to build and push multi-arch images to GitHub Container Registrt (ghcr.io)

There is an amazing command called `docker init` that can be used to initialize a Dockerfile with a template.

```bash
docker init
```

Then, you should do some modifications on the Dockerfile to make it multi-arch but instead of using `QEMU` to build multi-arch images, we will use Go's native cross-compilation support in which we'll be using `GOOS` and `GOARCH` variables.

```bash
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
...


ARG TARGETOS TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/server .
...
```

this modification will make your build process a lot faster!!! 🚀

Then, you should build and push your image to GitHub Container Registry (ghcr.io) with the following command:

> ⚠️ You should be running your own builder instance to be able to build multi-arch images since the Docker Daemon is not supporting multi-arch builds yet.
> docker buildx create --name my-own-builder --use --bootstrap

> Also do not forget to replace the `developerguy` with your own GitHub username.

```bash
docker buildx build -t developerguy/hello-go-server:latest --platform linux/amd64,linux/arm64 . --push
```
