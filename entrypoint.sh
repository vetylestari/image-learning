#!/bin/sh

make init-schema
make migrate-up

chmod +x /app/bin/go-starter-template
/app/bin/go-starter-template
