#!/bin/sh

# Copyright 2020 The arhat.dev Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex

GOPATH=$(go env GOPATH)
export GOPATH

PROTO_SOURCE="./src/*.proto"

fix_pb_gen_json_name() {
  pb_files="$*"
  cmd_sed="sed"

  command -v gsed && cmd_sed="gsed"

  for f in ${pb_files}; do
    ${cmd_sed} -i -r 's/,json=(\w+?),omitempty,proto3" json:"\w+(,omitempty)?"/,json=\1,omitempty,proto3" json:"\1,omitempty"/g' "${f}"
    ${cmd_sed} -i -r 's/,json=(\w+?),proto3" json:"\w+(,omitempty)?"/,json=\1,proto3" json:"\1"/g' "${f}"
  done
}

_do_gen_proto_go() {
  rm arhatgopb/*.pb.go || true

  # shellcheck disable=SC2086
  protoc \
    -I"${GOPATH}/src" \
    -I"${GOPATH}/src/github.com/gogo/protobuf/protobuf" \
    -I"./src" \
    --gogoslick_out "./arhatgopb" \
    --gogoslick_opt "paths=source_relative" \
    ${PROTO_SOURCE}

  # fix_pb_gen_json_name ./arhatgopb/*.pb.go
}

_do_gen_proto_python() {
  rm arhatpythonpb/*_pb*.py || true

  # shellcheck disable=SC2086
  pipenv run \
  python -m grpc_tools.protoc \
    -I"${GOPATH}/src" \
    -I"${GOPATH}/src/github.com/gogo/protobuf/protobuf" \
    -I"./src" \
    --python_out "./arhatpythonpb" \
    ${PROTO_SOURCE}
}

_do_gen_proto_c() {
  rm arhatnanopb/*.pb.c arhatnanopb/*.pb.h || true

  # shellcheck disable=SC2086
  pipenv run \
  python build/nanopb/generator/nanopb_generator.py \
    --no-timestamp \
    -x github.com/gogo/protobuf/gogoproto/gogo.proto \
    --output-dir ./arhatnanopb \
    -I"${GOPATH}/src" \
    -I"${GOPATH}/src/github.com/gogo/protobuf/protobuf" \
    -I"./src" \
    ${PROTO_SOURCE}
}

_do_gen_proto_rust() {
  rm arhatrustpb/*.pb.rs || true

  # shellcheck disable=SC2086
  pipenv run \
  protoc \
    -I ./src \
    -I "${GOPATH}/src/github.com/gogo/protobuf/protobuf" \
    -I "${GOPATH}/src" \
    --plugin "protoc-gen-rust=$(pwd)/build/pb-jelly/pb-jelly-gen/codegen/codegen.py" \
    --rust_out ./arhatrustpb \
    ${PROTO_SOURCE}
}

CODE_LANG=$(printf "%s" "$@" | cut -d. -f3)

"_do_gen_proto_${CODE_LANG}"
