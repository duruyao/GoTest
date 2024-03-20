# GoTest

GoTest parses test results (CSV files) and creates interactive visualizations in the web.

## Install

### Download executable

Download the compiled **[releases](https://github.com/duruyao/gotest/releases)** for your platform.

## Usage

The following operation takes the Linux platform as an example.

Type `gotest -h` to show help manual.

```text

   _____    _______        _
  / ____|  |__   __|      | |
 | |  __  ___ | | ___  ___| |_
 | | |_ |/ _ \| |/ _ \/ __| __|
 | |__| | (_) | |  __/\__ \ |_
  \_____|\___/|_|\___||___/\__|


Usage: gotest [OPTIONS]

GoTest parses test results and creates interactive visualizations in the web

Options:
    -h, --help                  Display this help message
    -v, --version               Print version information and quit
    --dir  STRING               Directory of dataset (default: './dataset')
    --host STRING               Host address to listen (default: '0.0.0.0:4936')

Examples:
    gotest
    gotest --dir ./dataset
    gotest --dir ./dataset --host 0.0.0.0:4936

```

```shell
tar -zxf dataset.tar.gz
nohup ${PWD}/gotest --dir ./dataset --host 0.0.0.0:4936 </dev/null 1>/${PWD}/gotest.out 2>&1 &

if [ -n "$(command -v open)" ]; then OPEN_EXE="open"; else OPEN_EXE="xdg-open"; fi
${OPEN_EXE} "http://0.0.0.0:4936/history?project=vc0728&test_type=similarity&branch=dev&test_case_id=OC9BbGV4TmV0&commit_short_sha=a25708c7"
${OPEN_EXE} "http://0.0.0.0:4936/history?project=vc0768&test_type=similarity&branch=dev&test_case_id=MTYvQWxleE5ldA==&commit_short_sha=c5dc5db1"
${OPEN_EXE} "http://0.0.0.0:4936/history?project=vc0728&test_type=accuracy&branch=dev&test_case_id=QWxleE5ldG9sZDhTdGF0aXN0aWNzU3RhdGlzdGljcw==&commit_short_sha=a25708c7"
${OPEN_EXE} "http://0.0.0.0:4936/history?project=vc0768&test_type=accuracy&branch=dev&test_case_id=QWxleE5ldG5ldzBPdXRsaWVyX1JlbW92ZU91dGxpZXJfUmVtb3ZlRXVjbGlkZWFu&commit_short_sha=c5dc5db1"

```
