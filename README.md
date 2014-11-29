# AWS IPs

Client utility (`awsips`) to query IP addresses used by Amazon Web Services.

The tool has to be used in one of several modes (`-m` flag):

* `regions`: returns a JSON array over all supported regions
* `services`: returns a JSON array over all supported services
* `list-files`: generates JSON files (in `$PWD`) for every region and service, containing a JSON array of IP addresses in CIDR notation
* `list-tree`:returns a JSON object mapping regions to services, and services to an array of IP addresses in CIDR notation

When using a list mode, the resulting output can be filtered by:

* region or multiple, comma-seperated regions (`-rf` flag)
* service or multiple, comma-seperated services (`-sf` flag)

The JSON (including the IP addresses) emitted by `awsip` is always sorted lexicographically to make diffing and checksum generation easier. 

## Example

```
$ awsips -m list-tree -sf CLOUDFRONT,ROUTE53_HEALTHCHECKS

{
    "CLOUDFRONT": {
        "GLOBAL": [
            "204.246.164.0/22",
            "204.246.168.0/22",
            "204.246.174.0/23",
            "204.246.176.0/20",
            "205.251.192.0/19",
            "205.251.249.0/24",
            "205.251.250.0/23",
            "205.251.252.0/23",
            "205.251.254.0/24",
            "216.137.32.0/19",
            "54.182.0.0/16",
            "54.192.0.0/16",
            "54.230.0.0/16",
            "54.239.128.0/18",
            "54.239.192.0/19",
            "54.240.128.0/18"
        ]
    },
    "ROUTE53_HEALTHCHECKS": {
        "ap-northeast-1": [
            "54.248.220.0/26",
            "54.250.253.192/26"
        ],
        "ap-southeast-1": [
            "54.251.31.128/26",
            "54.255.254.192/26"
        ],
        "ap-southeast-2": [
            "54.252.254.192/26",
            "54.252.79.128/26"
        ],
        "eu-west-1": [
            "176.34.159.192/26",
            "54.228.16.0/26"
        ],
        "sa-east-1": [
            "177.71.207.128/26",
            "54.232.40.64/26"
        ],
        "us-east-1": [
            "107.23.255.0/26",
            "54.243.31.192/26"
        ],
        "us-west-1": [
            "54.183.255.128/26",
            "54.241.32.64/26"
        ],
        "us-west-2": [
            "54.244.52.192/26",
            "54.245.168.0/26"
        ]
    }
}
```
