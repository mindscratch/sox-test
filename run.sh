#!/bin/bash

docker run --name sox-test --net=host -v $(pwd):/opt/go/src/github.com/mindscratch/sox-test -it gosox bash
