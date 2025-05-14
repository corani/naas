# 👎 no-as-a-service in Go

"Inspired" by <https://github.com/hotheadhacker/no-as-a-service>, from where I copied the list of reasons as well.

Not intended for productive usage, though requests are logged and limited by IP to 100/minute.

## Usage


Build using `./build.sh -b`

Run as a CLI (prints a random reason):

```sh
./bin/naas
```

Run as a service (API and web UI):

```sh
NAAS_PORT=3939 ./bin/naas -serve
```

Run using docker-compose:

```sh
docker-compose -f docker/docker-compose.yml up
```

This will start the service with the `-serve` flag automatically.

## Web UI

The service includes a simple web UI for exploring the reasons interactively. Open your browser and navigate to [http://localhost:3939](http://localhost:3939) (or the port you configured) after starting the service. The UI is served from the root path (`/`) and allows you to get random reasons with a button click.

## Endpoints

- `/no` returns a JSON object with a random reason.
- `/ping` returns status 200 (for health check).
- `/version` returns a JSON object with version info.
- `/debug` (when started with environment `NAAS_DEBUG=true`) exposes the `pprof` endpoints.

## Example

Example Request:

```http
GET /no
```

Example response:

```json
{
  "reason": "I've never been so sure of anything as I am of saying no." 
}
```
