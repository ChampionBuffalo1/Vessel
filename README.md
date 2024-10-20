# Vessel
A docker alternative to containerd with better Terminal experience

## Supported Commands
[Docker Cheatsheet](https://docs.docker.com/get-started/docker_cheatsheet.pdf)

- `image`
    - `ls` List local images
    - `rm` Delete an image (Supports bulk delete)
- `run` Create & run a container from image name
- `container`
    - `start` Starts a existing container
    - `stop` Stops a existing container
    - `rm` Removes a stopped container
- `pull` Pull an image from the container registry

### Note
`containerd` doesn't build any images it only pulls & runs containers. 
For building images (e.g `docker build`) use [`moby/buildkit`](https://pkg.go.dev/github.com/moby/buildkit/client)

> [!NOTE]
> Use `log/slog` for strucutred logging to a file rather than `fmt.Println` in `pkg` 
