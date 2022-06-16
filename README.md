# Pong
![CodeGame Protocol Version](https://img.shields.io/badge/Protocol-v0.6-orange)
![CodeGame GameServer Version](https://img.shields.io/badge/GameServer-v0.1-yellow)
![CGE Version](https://img.shields.io/badge/CGE-v0.3-green)

An implementation of [Pong](https://en.wikipedia.org/wiki/Pong) for [CodeGame](https://code-game.org).

## Known instances

- `games.code-game.org/pong`

## Usage

```sh
# Run on default port 80
pong

# Specify a custom port
pong --port=8080

## Specify a custom port through an environment variable
CG_PORT=8080 pong
```

### Running with Docker

Prerequisites:
- [Docker](https://docker.com/)

```sh
# Download image
docker pull codegameproject/pong:0.2

# Run container
docker run -d -p <port-on-host-machine>:8080 --name pong codegameproject/pong:0.2
```

## License

Copyright (C) 2022 Julian Hofmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
