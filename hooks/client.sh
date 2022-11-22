#!/bin/sh

(cd client && npm run lint:fix)
(cd server && make fmt && make test)