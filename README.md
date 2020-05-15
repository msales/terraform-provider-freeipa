# terraform-provider-freeipa

## Build

Makefile has been made basically and will attempt to crosscompile for all distros. Due to kerberos and
using a non-native (no golang) library this requires gcc compiler which breaks very easy across the
distributions so your probably end up with failures.

TODO: separate make tasks to specify distro on build for local

```bash
make dist
```

## Usage Example

### Provider

For a test example you might not of configured your own certificates. If so
please download the self-signed certs from `<your-domain>/ipa/config/ca.crt` and
put save to `/etc/ipa/ca.crt`

```tf
provider "freeipa" {
  host     = var.freeipa_host
  username = var.freeipa_username
  password = var.freeipa_password
  base_dn  = var.freeipa_base_dn
}
```

Arguments:
* `host` - host name of IPA server, example *ipa.example.com* (string, required)
* `username` - username for authentication (string, required)
* `password` - password for authentication (string, required)
* `base_dn` - domain components (dc), example: *dc=ipa,dc=example,dc=com* (string, required)

### User

```tf
resource "freeipa_user" "user" {
  uid        = var.uid
  first_name = var.first_name
  last_name  = var.last_name
  email      = var.email
  groups     = var.groups
  uid_number = var.uid_number
  gid_number = var.gid_number
}
```

Arguments:
* `uid` - User name (string, required)
* `first_name` - First name (string, required)
* `last_name` - Last name (string, required)
* `email` - Email address (string, optional, computed)
* `groups` - Groups user will be added to (set/list, optional, computed)
* `uid_number` - User ID (string, optional, computed)
* `gid_number` - Group ID (string, optional, computed)

### Group

```tf
resource "freeipa_user" "user" {
  gid         = var.gid
  gid_number  = var.gid_number
  description = var.description
}
```

Arguments:
* `gid` - Group Name (string, required)
* `gid_number` - Group ID (string, optional, computed)
* `description` - Description for group (string, optional, default: "")

### DNS zone

```tf
resource "freeipa_dnszone" "zone" {
  idnsname = "zone."
  idnssoaretry = 900
  idnssoaminimum = 100
  idnsallowdynupdate = true
}
```

Arguments can be found in API browser in IPA server under `dnszone_add`

### DNS Record

Example A Record
```tf
resource "freeipa_dnsrecord" "arecord" {
  idnsname = "arecord"
  dnszoneidnsname = "zone."
  arecord = ["127.0.0.3", "127.0.0.4"]
}
```

Example MX Record
```tf
resource "freeipa_dnsrecord" "mxrecord" {
  idnsname = "mxrecord"
  dnszoneidnsname = "zone."
  mxrecord = ["0 testmx.pl", "1 testmx2.pl"]
}
```

Also you can reuse names of zone from main.tf
```tf
resource "freeipa_dnsrecord" "mxrecord" {
  idnsname = "mxrecord"
  dnszoneidnsname = freeipa_dnszone.zone.idnsname
  mxrecord = ["0 testmx.pl", "1 testmx2.pl"]
}
```

Useful docs:

https://www.freeipa.org/page/V2/DNS_Interface_Design#1._New_per-type_structured_API
https://www.terraform.io/docs/configuration-0-11/resources.html

Arguments can be found in API browser in IPA server under `dnsrecord_add` 

### Import

```bash
# terraform import <module path> <ipauniqueid for group>

terraform import freeipa_group.group <ipauniqueid for group>
terraform import freeipa_user.user <ipauniqueid for user>
```

If your using FreeIPA then your likely to find that the interface does not expose
the `ipauniqueid`. To get this you will need to log onto the IPA Server and run the
following command:

```bash
ipa user-show --all <username>
ipa group-show --all <groupname>
```

