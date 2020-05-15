package freeipa

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	dnsSchemaZoneName                 = "idnsname"
	dnsSchemaReverseZone              = "name_from_ip"
	dnsSchemaSkipOverlapCheck         = "skip_overlap_check"
	dnsSchemaSkipNameserverCheck      = "skip_nameserver_check"
	dnsSchemaZoneForwarders           = "idnsforwarders"
	dnsSchemaForwardPolicy            = "idnsforwardpolicy"
	dnsSchemaAuthoritativeNameserver  = "idnssoamname"
	dnsSchemaAdministratorEmail       = "idnssoarname"
	dnsSchemaSOASerial                = "idnssoaserial"
	dnsSchemaSOARefresh               = "idnssoarefresh"
	dnsSchemaSOARetry                 = "idnssoaretry"
	dnsSchemaSOAExpire                = "idnssoaexpire"
	dnsSchemaSOAMinimum               = "idnssoaminimum"
	dnsSchemaTTL                      = "dnsttl"
	dnsSchemaDefaultTTL               = "dnsdefaultttl"
	dnsSchemaDNSClass                 = "dnsclass"
	dnsSchemaBINDUpdatePolicy         = "idnsupdatepolicy"
	dnsSchemaDynamicUpdate            = "idnsallowdynupdate"
	dnsSchemaAllowQuery               = "idnsallowquery"
	dnsSchemaAllowTransfer            = "idnsallowtransfer"
	dnsSchemaAllowPTRSync             = "idnsallowsyncptr"
	dnsSchemaAllowInLineDNSSECSigning = "idnssecinlinesigning"
	dnsSchemaNSEC3ParamRecord         = "nsec3paramrecord"
)

var dnsSchemaKeys = []string{
	dnsSchemaZoneName,
	dnsSchemaReverseZone,
	dnsSchemaSkipOverlapCheck,
	dnsSchemaSkipNameserverCheck,
	dnsSchemaZoneForwarders,
	dnsSchemaForwardPolicy,
	dnsSchemaAuthoritativeNameserver,
	dnsSchemaAdministratorEmail,
	dnsSchemaSOASerial,
	dnsSchemaSOARefresh,
	dnsSchemaSOARetry,
	dnsSchemaSOAExpire,
	dnsSchemaSOAMinimum,
	dnsSchemaTTL,
	dnsSchemaDefaultTTL,
	dnsSchemaDNSClass,
	dnsSchemaBINDUpdatePolicy,
	dnsSchemaDynamicUpdate,
	dnsSchemaAllowQuery,
	dnsSchemaAllowTransfer,
	dnsSchemaAllowPTRSync,
	dnsSchemaAllowInLineDNSSECSigning,
	dnsSchemaNSEC3ParamRecord,
}

func resourceDNSZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSZoneCreate,
		Read:   resourceDNSZoneRead,
		Update: resourceDNSZoneUpdate,
		Delete: resourceDNSZoneDelete,
		Exists: resourceDNSZoneExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			dnsSchemaZoneName: {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			dnsSchemaReverseZone: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaSkipOverlapCheck: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaSkipNameserverCheck: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaZoneForwarders: {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dnsSchemaForwardPolicy: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAuthoritativeNameserver: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAdministratorEmail: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaSOASerial: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSOARefresh: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSOARetry: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSOAExpire: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSOAMinimum: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaTTL: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDefaultTTL: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDNSClass: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaBINDUpdatePolicy: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaDynamicUpdate: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaAllowQuery: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAllowTransfer: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAllowPTRSync: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaAllowInLineDNSSECSigning: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaNSEC3ParamRecord: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDNSZoneCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSZoneCreate - %s", d.Id())

	c := m.(*Connection)

	options := map[string]interface{}{}

	for _, key := range dnsSchemaKeys {
		if val, ok := d.GetOk(key); ok {
			options[key] = val
		}
	}

	log.Print("[TRACE] creating DNS zone")

	dns, err := c.client.CreateDNSZone(options)
	if err != nil {
		log.Printf("[ERROR] Error creating DNS Zone %v", err)
		return err
	}

	d.SetId(dns.Name.String())

	return resourceDNSZoneRead(d, m)
}

func resourceDNSZoneRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSZoneRead - %s", d.Id())

	c := m.(*Connection)

	dns, err := c.client.GetDNSZone(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading DNS Zone %s - %s", dns, err)
		return err
	}

	return nil
}

func resourceDNSZoneUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSZoneUpdate - %s", d.Id())

	c := m.(*Connection)

	dns, err := c.ldapClient.GetDNSZone(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading DNS Zone %s - %s", dns, err)
		return err
	}
	d.Partial(true)

	for _, key := range dnsSchemaKeys {
		if d.HasChange(key) {
			_, newVal := d.GetChange(key)
			c.client.DNSZoneMod(*dns, key, newVal)
			d.SetPartial(key)
			log.Printf("[TRACE] resource resourceDNSZoneUpdate dns %s key %s updated", *dns, key)
		}
	}

	d.Partial(false)

	return resourceDNSZoneRead(d, m)
}

func resourceDNSZoneDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSZoneDelete - %s", d.Id())

	c := m.(*Connection)

	ns, err := c.ldapClient.GetDNSZone(d.Id())

	if err != nil {
		log.Printf("[ERROR] Error reading DNS Zone %s - %s", ns, err)
		return err
	}

	return c.client.DeleteDNSZone(*ns)
}

func resourceDNSZoneExists(d *schema.ResourceData, m interface{}) (bool, error) {
	log.Printf("[TRACE] resourceDNSZoneDelete - %s", d.Id())

	c := m.(*Connection)

	exists, err := c.ldapClient.DNSZoneExists(d.Id())

	if err != nil {
		return false, err
	}

	return exists, nil
}
