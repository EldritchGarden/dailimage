# Dailimage
Get a random image with a simple API call

## Description
Dailimage is a simple lightweight web server designed to serve a random image. It's written in Go using Gin
and the docker image runs on Alpine. While it's designed to serve images, it actually makes no
real distinction between file types and can be used to serve any files.

## Usage and Endpoints
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
      <td>Gin logging level, defaults to 'relase', set to 'debug' for more logs</td>
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
      <td>Root dir for files to serve. Defaults to /media, only change this if you change the mount point in the container.</td>
    </tr>
    <tr>
      <td>TRUSTED_PROXIES</td>
      <td>Comma separated list of proxies to trust. If not set trusted proxies are disabled.</td>
    </tr>
  </tbody>
</table>

### API
<table>
  <thead>
    <tr>
      <th>Endpoint</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>/ping</td>
      <td>Returns a simple OK message</td>
    </tr>
    <tr>
      <td>/random</td>
      <td>Returns a random image from media dir</td>
    </tr>
    <tr>
      <td>/random/*subdir</td>
      <td>Returns a random image from a sub path under media dir. /random/art would return a file under /media/art.</td>
    </tr>
  </tbody>
</table>

## Setup
### Docker and Compose
Docker run:
`docker run -p 8080:8080 -v <media_path>:/media TODO`

Docker compose:
```yaml
services:
  dailimage:
    image: TODO
    ports:
      - 8080:8080
    volumes:
      - <media_path>:/media
    restart: unless-stopped
```
