# arhat-proto

[![CI](https://github.com/arhat-dev/arhat-proto/workflows/CI/badge.svg)](https://github.com/arhat-dev/arhat-proto/actions?query=workflow%3ACI)
[![PkgGoDev](https://pkg.go.dev/badge/arhat.dev/arhat-proto)](https://pkg.go.dev/arhat.dev/arhat-proto)
[![GoReportCard](https://goreportcard.com/badge/arhat.dev/arhat-proto)](https://goreportcard.com/report/arhat.dev/arhat-proto)
[![codecov](https://codecov.io/gh/arhat-dev/arhat-proto/branch/master/graph/badge.svg)](https://codecov.io/gh/arhat-dev/arhat-proto)

Extension protocol for [`arhat`](https://github.com/arhat-dev/arhat)

## Workflow

1. Install proto tools with `make install.proto.all`
1. Update `*.proto` in `src` directory
1. Generate protobuf files with `make gen.proto.all`
1. Test with `make test.all`

## Usage

- Go

  ```bash
  go get -u arhat.dev/arhat-proto/arhatgopb
  ```

## License

```txt
Copyright 2020 The arhat.dev Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
