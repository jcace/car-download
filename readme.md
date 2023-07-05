
# Installation

You must have Aria setup and running first - https://aria2.github.io/


## Running
First you need to run aria as a daemon

```
aria2c --enable-rpc --rpc-listen-all
```


Then, run the downloader:

```bash
./downloader --out /path/to/download --aria-uri http://localhost:6800/jsonrpc --ddm-api https://delta-web.estuary.tech --ddm-token <token>
```