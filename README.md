##Description
Demo application use for the Cloud foundry Meetup [https://github.com/shinji62/cf-meetup-tokyo-2016-03-01--1]Tokyo Cloud Foundry #1 Meetup
It's a simple reverse proxy aim to be use with the CloudFoundry Route Service

Basically this will transform incoming body into ASCII BANNER

##Usage

### Build (tested with go 1.5.3)
```
$godep go build
```

###Just push to cloudfoundry
```
$cf push 
```

### Create your service (User provided Service)
```bash
$cf cups my-proxy -r https://route-service.local.pcfdev.io
```
*Using PcfDev that's why the domain is local.pcfdev.io

### Just bind to the domain you want to intercept request  / response

```bash
$cf bind-route-service local.pcfdev.io my-proxy -n php-simple
```
*Here, I want to forward everyrequest php-simple.local.pcfdev.io to my-proxy

