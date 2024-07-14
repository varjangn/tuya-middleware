# Tuya-middleware
- Rest API middleware for Tuya Cloud APIs.
- This middleware provides SDK like wrapper for Tuya Cloud API which is easily extensible.
- App exposes Rest API endpoint for Tuya Cloud Functionality with auto refresing tokens.

# Setup
- Install the go dependencies.
```bash
go mod download
```

- Ensure configs are created in `config` directory for your environment. Name of the config file should be `config/config-<env>.yml`.

Sample Config
```yml
server:
  Port: :5000
  Mode: Development
  Debug: false

logger:
  Development: true
  DisableCaller: false
  DisableStacktrace: false
  Encoding: json
  Level: info

tuya:
  Host: https://openapi.tuyain.com # ensure host url is as per your data center
  ClientId: <<tuya-client-id>>
  Secret: <<tuya-client-secret>>
```

And then set the environment variable for your config. default is local if no environment variable is found.

```bash
export config="development" # for the config file config/config-development.yml
```

# Running locally
After installing the dependencies and setting configs. We can start using `make` for building and running application binary.

- Running the application locally
```bash
make run
```
or use below command to build the binary in `bin` directory
```bash
make build
```

