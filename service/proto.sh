#!/usr/bin/env bash

set -e -u -o pipefail # Fail on error

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $dir

"$dir/../scripts/proto.sh" $dir

keysapp="$dir/../../app"

cp *.proto "$keysapp/src/main/rpc/"
cp *.d.ts "$keysapp/src/main/rpc/"
cp *.d.ts "$keysapp/src/renderer/rpc/"
rm *.d.ts
mv *.ts "$keysapp/src/renderer/rpc/"