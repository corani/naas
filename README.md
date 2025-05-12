# 👎 no-as-a-service in Go

"Inspired" by <https://github.com/hotheadhacker/no-as-a-service>, from where I copied the list of reasons as well.

Not intended for productive usage, though requests are logged and limited by IP to 100/minute.

## Usage

Build using `./build.sh -b`

Run as a cli with `./bin/naas`

Run as a service with `NAAS_PORT=3939 ./bin/naas`

Run using docker-compose: `docker-compose -f docker/docker-compose.yml up`

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
