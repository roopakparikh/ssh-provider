# SSH Machine Controller
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Froopakparikh%2Fssh-provider.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Froopakparikh%2Fssh-provider?ref=badge_shield)


A machine controller for the [Cluster API](http://sigs.k8s.io/cluster-api) reference implementation.


## Testing

### Machine Actuator

The machine actuator tests will run against a mock SSH server, but until then they run against an actual host, and require valid SSH credentials. Generate them before running tests:

    $ ./machine/testdata/generate-ssh-credentials.sh username /path/to/private-ssh-key > ./machine/testdata/ssh-credentials.yaml

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Froopakparikh%2Fssh-provider.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Froopakparikh%2Fssh-provider?ref=badge_large)