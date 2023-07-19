# The Pacstall Website

## Stats

<p align="center"><img alt="Repobeats analytics image" src="https://repobeats.axiom.co/api/embed/eda50f5088638b97b7f64e29c0f2d65c7fb89568.svg" /></p>

## How to Run

First, clone this repo and go to its directory.

```sh
git clone https://github.com/pacstall/website
cd website
```

Then, choose an option below depending on if you're planning to make changes to the website or to just run it.

### Run it for Development

```sh
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --force-recreate --build
```

### Run it in Production

```sh
docker-compose pull && docker-compose down && docker-compose up -d
```