# minflow_serve_proto

This repo contains compiled protobufs from tensorflow library. 

Run the following command with docker to recompile.

```
docker build . -t goflow_proto:latest
docker run -it -v $GOPATH:/go goflow_proto:latest bash
```

You should then be able to run, `run.sh` to recompile protobufs
