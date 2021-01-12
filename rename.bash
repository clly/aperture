#!/usr/bin/bash

package=$1
receiver=$2

if [[ -z $package ]]; then
    echo "Usage: $0 package receiver"
    exit 1
fi

if [[ -z $receiver ]]; then
    echo "Usage: $0 package receiver"
    exit 1
fi

ack -l apterture | xargs sed -i "s~GOTEMPLATE~${package}~"
ack -l Aperture | xargs sed -i "s~GODAEMONTEMPLATE~${receiver}~"
