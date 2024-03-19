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
    --host STRING               Host address to listen (default: '0.0.0.0:4936')

Examples:
    gotest
    gotest --host 0.0.0.0:4936

See more about GoTest at https://github.com/duruyao/gotest

```

Start gotest in the foreground.

```bash
gotest --host=0.0.0.0:4936
```

Start gotest in the background.

```bash
gotest --host=0.0.0.0:4936 &
```

Start gotest in the background via **Nohup**.

```bash
nohup gotest --host=0.0.0.0:4936 </dev/null 1>/$PWD/gotest.out 2>&1 &
```
