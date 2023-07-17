
# Installation

You must have Aria setup and running first - https://aria2.github.io/


## Build
```bash
make build
```

This will output the executable `car-downloader`

You can install it (copy to /usr/local/bin) with:

```bash
make install
```


## Running
First you need to run aria as a daemon

```
aria2c --enable-rpc --rpc-listen-all
```

In this aria window, you can see the download progress as the downloader script runs. 


Then, run the downloader:

```bash
./car-downloader --out /path/to/download --aria-uri ws://localhost:6800/jsonrpc --ddm-api https://delta-dm.estuary.tech --ddm-token <token>
```

### Options
```
NAME:
   CAR Downloader - a simple utility to download carfiles from a content data source

USAGE:
   CAR Downloader [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --out value, -o value        output directory for downloaded files (default: "./")
   --aria-uri value, -a value   aria2c RPC URI (default: "ws://localhost:6800/jsonrpc")
   --ddm-api value, -d value    DDM URI to fetch contents from
   --ddm-token value, -t value  DDM self-service token
   --interval value, -i value   interval (in minutes) to check for new content to download (default: 10)
   --help, -h                   show help
Required flags "ddm-api, ddm-token, interval" not set
```


