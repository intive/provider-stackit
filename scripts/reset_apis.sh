#!/bin/bash

find apis -name 'zz_generated.managed.go' -delete
find apis -name 'zz_*generated*.go' -delete
find apis -name '*_types.go' -delete
