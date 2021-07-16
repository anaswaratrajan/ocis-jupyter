---
title: "Running"
date: 2018-05-02T00:00:00+00:00
weight: 50
geekdocRepo: https://github.com/owncloud/ocis-hello
geekdocEditPath: edit/master/docs
geekdocFilePath: configuration-with-ocis.md
---

### Configuring ocis-jupyter with ocis
We will need various services to run ocis
#### Running ocis
In order to run this extension we will need to configure ocis first. For that clone and build the ocis single binary from the github repo `https://github.com/owncloud/ocis`. 

Update the pheonix config file `web-config.json` with the following contents. We're registering the client of `ocis-jupyter` extension to pheonix. 

```json
{
    "server": "https://localhost:9200",
    "theme": "owncloud",
    "version": "0.1.0",
    "openIdConnect": {
        "metadata_url": "https://localhost:9200/.well-known/openid-configuration",
        "authority": "https://localhost:9200",
        "client_id": "phoenix",
        "response_type": "code",
        "scope": "openid profile email"
    }, 
    "apps": [
        "files",
        "draw-io",
        "pdf-viewer",
        "markdown-editor",
        "media-viewer"
    ], 
    "external_apps": [
        {
            "id": "ocis-jupyter",
            "path": "/ocis-jupyter.js"
        }
    ]   
}

```
Here we can add the url for the js file from where the ocis-jupyter app will be loaded.

After that we will need a configuration file for ocis where we can specify the path for the ocis-jupyter app in the backend. For this you can use the existing `proxy-example.json` file from the [ocis-proxy](https://github.com/owncloud/ocis-proxy/blob/master/config/proxy-example.json) repo. Just add these two endpoints at the end for the ocis-jupyter app. Now ocis-proxy knows where to route the requests when the client hits these endpoints. 
```json
        {
          "endpoint": "/api/v0/convert",
          "backend": "http://localhost:9105"
        },
        {
          "endpoint": "/ocis-jupyter.js",
          "backend": "http://localhost:9105"
        },
```

With all this in place we can finally start ocis. 

Start ocis-server like so declaring the path to web-config file path explicitly

```
WEB_UI_CONFIG=/../ocis/web/config/web-config.json ./bin/ocis server
```

Then kill ocis-proxy service and start the proxy seperately by declaring proxy-config file path. 

```
./bin/ocis kill proxy

PROXY_CONFIG_FILE=/../ocis/proxy/config/proxy-example.json ./bin/proxy server

``` 

After this we will need to start the ocis-jupyter service.
For that just build ocis-jupyter binary.

```
cd ocis-jupyter 
make generate build
```
And Run the service
```
bin/ocis-jupyter server
```
