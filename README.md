
# Bucks

A simple commandline currency converter written in Go.

## Demo

Insert gif or link to demo

  
## Documentation

[Documentation](https://bucks.docs.danielms.site)

  
## Features

- Convert currencies easily
- Powered by third party APIs
- Super fast!

  
## Installation

`bucks` requires the use of a configuration file. The standard location for this file is
`$HOME/.bucks.yaml` and is a `yaml` file.

```yaml
# bucks example configuration yaml
providers:
  currencyconverterapi:
    # API key from provider, empty means it will not be used even in fallback
    key: ""
    # set to true if this should be the preferred API.
    default: false
```
If `bucks` does not find this file it will fail to run. 

```bash
# todo
#docker run -it --rm danielmichaels/bucks:latest <int> <currency> <currency>
```

```bash
go get github.com/danielmichaels/bucks
```
## Usage/Examples

The prefered option is to run as container.

```shell
docker run -it --rm danielmichaels/bucks:latest 100 usd eur
# you can alias this command for more brevity
```

`bucks` can also be run as a binary when installed from github. See [Installation](#installation) for more details.


```shell
bucks 100 usd eur
```

  
## License

[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)
  
## Related

Here are some related projects


  
## Running Tests

To run tests, run the following command
# todo
```bash
go test
```

  
## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.

