#!/bin/bash

set -eu
shopt -s nullglob

TAG=${1:-}
if [ -z $TAG ]; then
  echo 'ERROR: tag is not set'
  exit 1
fi

echo " [*] Clean old files"
rm -rf tmp
rm -rf envoy

echo " [*] Clone repos"
mkdir tmp
cd tmp
git clone git@github.com:envoyproxy/envoy.git
git clone git@github.com:lyft/protoc-gen-validate.git
git clone git@github.com:gogo/protobuf.git
git clone git@github.com:istio/gogo-genproto.git

echo " [*] Checkout tag"
cd envoy
git checkout tags/$TAG

echo " [*] Copy files"
cp -r api/envoy ../..
cd ../..
find envoy -name BUILD -exec rm {} \;

echo " [*] Generate protobufs"

arg=''

mappings=(
  'google/api/annotations.proto=github.com/gogo/googleapis/google/api'
  'google/api/http.proto=github.com/gogo/googleapis/google/api'
  'google/rpc/code.proto=github.com/gogo/googleapis/google/rpc'
  'google/rpc/error_details.proto=github.com/gogo/googleapis/google/rpc'
  'google/rpc/status.proto=github.com/gogo/googleapis/google/rpc'

  'google/protobuf/any.proto=github.com/gogo/protobuf/types'
  'google/protobuf/duration.proto=github.com/gogo/protobuf/types'
  'google/protobuf/empty.proto=github.com/gogo/protobuf/types'
  'google/protobuf/struct.proto=github.com/gogo/protobuf/types'
  'google/protobuf/timestamp.proto=github.com/gogo/protobuf/types'
  'google/protobuf/wrappers.proto=github.com/gogo/protobuf/types'

  'gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto'
  
  'trace.proto=istio.io/gogo-genproto/opencensus/proto/trace'
  'metrics.proto=istio.io/gogo-genproto/prometheus'
)
for mapping in "${mappings[@]}"
do
  arg+=",M$mapping"
done

for path in $(find envoy -type d)
do
  path_protos=(${path}/*.proto)
  if [[ ${#path_protos[@]} > 0 ]]; then
    for path_proto in "${path_protos[@]}"
    do
      mapping=${path_proto}=github.com/altipla-consulting/envoy-api/${path}
      arg+=",M$mapping"
    done
  fi
done

for path in $(find envoy -type d)
do
  path_protos=(${path}/*.proto)
  if [[ ${#path_protos[@]} > 0 ]]
  then
    echo "> generate protos ./${path}"
    actools protoc \
    -I. \
    -Itmp/protoc-gen-validate \
    -Itmp/gogo-genproto/prometheus \
    -Itmp/gogo-genproto/googleapis \
    -Itmp/gogo-genproto/opencensus/proto/trace \
    -Itmp/protobuf \
    -I/opt/googleapis \
    --go_out=plugins=grpc,paths=source_relative${arg}:. \
    ${path}/*.proto
  fi
done

echo " [*] Test generated files"
go install ./envoy/...
