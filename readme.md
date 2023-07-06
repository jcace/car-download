
# Installation

You must have Aria setup and running first - https://aria2.github.io/


## Running
First you need to run aria as a daemon

```
aria2c --enable-rpc --rpc-listen-all
```

In this aria window, you can see the download progress as the downloader script runs. 


Then, run the downloader:

```bash
./downloader --out /path/to/download --aria-uri ws://localhost:6800/jsonrpc --ddm-api https://delta-web.estuary.tech --ddm-token <token>
```