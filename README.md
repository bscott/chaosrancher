# chaosrancher
[![Docker Repository on Quay](https://quay.io/repository/bscott/chaosrancher/status "Docker Repository on Quay")](https://quay.io/repository/bscott/chaosrancher)

## Why

Test how your system behaves under arbitrary service failures.

## Setup

ChaosRancher reads from enviroment varibles to authenticate to Rancher API (v2-beta is only supported at the moment, v1 is not)

```console
# Set the url that Rancher is on
$ export RANCHER_URL=http://<server_ip>:8080
# Set the access key, i.e. username
$ export RANCHER_ACCESS_KEY=<accessKey_of_account_api_key>
# Set the secret key, i.e. password
$ export RANCHER_SECRET_KEY=<secretKey_of_account_api_key>
```

## Example

Running it will kill a service in any stack every 10 minutes by default.

```console
$ chaosrancher
```

or Deploy ChaosRancher in Rancher's Sandbox environment:

[Try Rancher](https://try.rancher.com/login)

## Filtering Services 

To select a specific environment that should be targeted:

```console
# Set the environment to use, you can use either environment ID or environment name
$ export RANCHER_ENVIRONMENT=<environment_id>
```

## Contributing

Feel free to create issues or submit pull requests.