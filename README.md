# sotu(疎通)

ネットワークの疎通確認をするためのプログラムです

# Usage

## HTTP疎通確認

Run server

```
$ sotu server
HTTP Server is running at 0.0.0.0:8888
```

Run client

```
$ sotu client
HTTP/1.0 200 OK
Connection: close

HTTP connection was successful
```

## TCP疎通確認

Run server

```
$ sotu server -protocol tcp
TCP Server is running at 0.0.0.0:8888
```

Run client

```
$ sotu client -protocol tcp
reply from server= Hello
```