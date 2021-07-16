# ownCloud Infinite Scale: Jupyter Notebooks

**This project is under heavy development, it's not in a working state yet!**

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.13. For the frontend it's also required to have [NodeJS](https://nodejs.org/en/download/package-manager/) and [Yarn](https://yarnpkg.com/lang/en/docs/install/) installed.

> Follow these [instructions](docs/configuration-with-ocis.md) and configure ocis to run ocis-jupyter.  

```console
# clone and cd into ocis-jupyter

yarn install
yarn build

make generate build

./bin/ocis-jupyter -h
```

## Security

If you find a security issue please contact security@owncloud.com first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2019 ownCloud GmbH <https://owncloud.com>
```
