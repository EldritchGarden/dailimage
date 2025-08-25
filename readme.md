- [Dailimage](#dailimage)
  - [Description](#description)
  - [Usage and Endpoints](#usage-and-endpoints)
    - [Configuration](#configuration)
    - [API](#api)
      - [Examples](#examples)
  - [Setup](#setup)
    - [Docker and Compose](#docker-and-compose)

# Dailimage
Get a random image with a simple API call

## Description
Dailimage is a simple lightweight web server designed to serve a random image.
It's written in Go using Gin and the docker image runs on Alpine.
While it's designed to serve images, it actually makes no real distinction
between file types and can be used to serve any files.

## Usage and Endpoints
Check out the [examples](docs/example/) for scripts, snippets, etc. for using
dailimage.

### Configuration
All env vars are optional.
<table>
  <thead>
    <tr>
      <th>ENV Var</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>GIN_MODE</td>
      <td>Gin logging level, defaults to 'relase', set to 'debug'
      for more logs</td>
    </tr>
    <tr>
      <td>BIND_ADDR</td>
      <td>Address to listen on. Defaults to 0.0.0.0</td>
    </tr>
    <tr>
      <td>BIND_PORT</td>
      <td>Port to listen on. Defaults to 8080</td>
    </tr>
    <tr>
      <td>MEDIA_ROOT</td>
      <td>Root dir for files to serve. Defaults to /media, only change this if
      you change the mount point in the container.</td>
    </tr>
    <tr>
      <td>TRUSTED_PROXIES</td>
      <td>Comma separated list of proxies to trust. If not set trusted proxies
      are disabled.</td>
    </tr>
  </tbody>
</table>

### API
- `/ping` : Return a simple '200 OK' message
- `/random` : Return a random image from the media library
- `/random/*subpath` : Return a random image from `subpath/`
- `/slideshow[?...]` : Web page with auto-refresh for browser use
  - `interval=30` : Seconds between refreshes
  - `subpath=somedir` : Use images from the specified sub folder
  - `mode=(full | frame)` : View mode
    - `full` uses the entire window and crops excess
    - `frame` fits the image into the window with a small border

#### Examples
With a media folder structure like:
```
+ media
| + family
| | + vacation
| + art
```

`/random` will return an image from anywhere in 'media'

`/random/family` will return an image from 'family' or 'vacation' but not 'art'

`/slideshow?interval=60&subpath=art&mode=full` will show a new fullscreen image under 'art' every 60 seconds

## Setup
### Docker and Compose
[Docker Hub](https://hub.docker.com/r/eldritchgarden/dailimage)

Docker run:
`docker run -p 8080:8080 -u 1000:1000 -v <media_path>:/media eldritchgarden/dailimage:latest`

Docker compose:
```yaml
services:
  dailimage:
    image: eldritchgarden/dailimage:latest
    user: 1000:1000 # optional
    ports:
      - 8080:8080
    volumes:
      - <media_path>:/media
    restart: unless-stopped
```

*Note: I recommend you pin the version, but latest will track the latest
release version*
