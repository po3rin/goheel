# Goheel

this CLI tail file's contents

## Feature

coloring abnormal log (error , warn ,etc...) and watching file change.

## Quickstart

Set this folder on 
```
GOPATH/sec/github.com/po3rin/
```
and exec command in project folder

```
$ make dep
$ make build
$ ./goheel test.log
```

## Options

### --num, -n

dicide how many lines show.

```
$ ./goheel -n=5 test.log
```

inspect a line and coloring Abnormal log.

### --color, -c

inspect a line and coloring Abnormal log.

```
$ ./goheel -c test.log
```

### --watch -w

watching file's change.

```
$ ./goheel -w test.log
```