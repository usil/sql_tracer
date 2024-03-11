# Sql Tracer

<p float="left">
  <img src="./.coverage/branches.svg">
  <img src="./.coverage/functions.svg">
  <img src="./.coverage/lines.svg">
  <img src="./.coverage/statements.svg">
</p>


<p align="center">
  <img src="./.assets/logo.png" width=200 ></img>
</p>

A tiny tool to get how fast id resolving your database.

## Settings

|name| sample value |description|
|:--|:--|:--|
|h| 10.10.20.30 | database host like localhost, 127.0.0.1, an ip, etc|

## requirements

- Go

## Dependencies

```
go get github.com/sijms/go-ora/v2
```

## Run

```
go run src/main/go/sql_tracer.go -p 1521 -h 192.168.0.10 -u system -ps admin123 -s xe
```

## Build

**Linux/Mac**

```
go build -o build/sql_tracer src/main/go/sql_tracer.go
```

**Windows**

```
go build -o build/sql_tracer.exe src/main/go/sql_tracer.go
```



## Advanced settings

For notitifations, templates, plugins, etc go to the [wiki](https://github.com/usil/sql_tracer/wiki)

## Acknowledgments

- https://easydrawingguides.com/how-to-draw-bob-the-minion/
- https://www.textstudio.com/logo/minions-411

## Contributors

<table>
  <tbody>    
    <td>
      <img src="https://avatars0.githubusercontent.com/u/3322836?s=460&v=4" width="100px;"/>
      <br />
      <label><a href="http://jrichardsz.github.io/">JRichardsz</a></label>
      <br />
    </td>
  </tbody>
</table>