#!/bin/bash
docker rm -f go-extract-files
docker run -v $PWD/output:/output --name go-extract-files go-extract-files
