package freeipa

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	dnsSchemaRecordName              = "idnsname"
	dnsSchemaRecordZoneName          = "dnszoneidnsname"
	dnsSchemaRecordTTL               = "dnsttl"
	dnsSchemaRecordClass             = "dnsclass"
	dnsSchemaForce                   = "force"
	dnsSchemaARecord                 = "arecord"
	dnsSchemaAIPAddress              = "a_part_ip_address"
	dnsSchemaACreateReverse          = "a_extra_create_reverse"
	dnsSchemaAAAARecord              = "aaaarecord"
	dnsSchemaAAAAIPAddress           = "aaaa_part_ip_address"
	dnsSchemaAAAACreateReverse       = "aaaa_extra_create_reverse"
	dnsSchemaA6Record                = "a6record"
	dnsSchemaA6Data                  = "a6_part_data"
	dnsSchemaAFSDBRecord             = "afsdbrecord"
	dnsSchemaAFSDBSubtype            = "afsdb_part_subtype"
	dnsSchemaAFSDBHostname           = "afsdb_part_hostname"
	dnsSchemaAPLRecord               = "aplrecord"
	dnsSchemaCERTRecord              = "certrecord"
	dnsSchemaCERTType                = "cert_part_type"
	dnsSchemaCERTKeyTag              = "cert_part_key_tag"
	dnsSchemaCERTAlgorithm           = "cert_part_algorithm"
	dnsSchemaCERTCertOrCRL           = "cert_part_certificate_or_crl"
	dnsSchemaCNAMERecord             = "cnamerecord"
	dnsSchemaCNAMEHostname           = "cname_part_hostname"
	dnsSchemaDHCIDRecord             = "dhcidrecord"
	dnsSchemaDLVRecord               = "dlvrecord"
	dnsSchemaDLVKeyTag               = "dlv_part_key_tag"
	dnsSchemaDLVAlgorithm            = "dlv_part_algorithm"
	dnsSchemaDLVDigestType           = "dlv_part_digest_type"
	dnsSchemaDLVDigest               = "dlv_part_digest"
	dnsSchemaDNAMERecord             = "dnamerecord"
	dnsSchemaDNAMETarget             = "dname_part_target"
	dnsSchemaDSRecord                = "dsrecord"
	dnsSchemaDSKeyTag                = "ds_part_key_tag"
	dnsSchemaDSAlgorithm             = "ds_part_algorithm"
	dnsSchemaDSDigestType            = "ds_part_digest_type"
	dnsSchemaDSDigest                = "ds_part_digest"
	dnsSchemaHIPRecord               = "hiprecord"
	dnsSchemaIPSECRecord             = "ipseckeyrecord"
	dnsSchemaKeyRecord               = "keyrecord"
	dnsSchemaKXRecord                = "kxrecord"
	dnsSchemaKXPreference            = "kx_part_preference"
	dnsSchemaKXExchanger             = "kx_part_exchanger"
	dnsSchemaLOCRecord               = "locrecord"
	dnsSchemaLOCLatDeg               = "loc_part_lat_deg"
	dnsSchemaLOCLatMin               = "loc_part_lat_min"
	dnsSchemaLOCLatSec               = "loc_part_lat_sec"
	dnsSchemaLOCLatDir               = "loc_part_lat_dir"
	dnsSchemaLOCLonDeg               = "loc_part_lon_deg"
	dnsSchemaLOCLonMin               = "loc_part_lon_min"
	dnsSchemaLOCLonSec               = "loc_part_lon_sec"
	dnsSchemaLOCLonDir               = "loc_part_lon_dir"
	dnsSchemaLOCAltitude             = "loc_part_altitude"
	dnsSchemaLOCSize                 = "loc_part_size"
	dnsSchemaLOCHorizontalPrecision  = "loc_part_h_precision"
	dnsSchemaLOCVerticalPrecision    = "loc_part_v_precision"
	dnsSchemaMXRecord                = "mxrecord"
	dnsSchemaMXPreference            = "mx_part_preference"
	dnsSchemaMXExchanger             = "mx_part_exchanger"
	dnsSchemaNAPTRRecord             = "naptrrecord"
	dnsSchemaNAPTROrder              = "naptr_part_order"
	dnsSchemaNAPTRPreference         = "naptr_part_preference"
	dnsSchemaNAPTRFlags              = "naptr_part_flags"
	dnsSchemaNAPTRService            = "naptr_part_service"
	dnsSchemaNAPTRRegexp             = "naptr_part_regexp"
	dnsSchemaNAPTRReplacement        = "naptr_part_replacement"
	dnsSchemaNSRecord                = "nsrecord"
	dnsSchemaNSHostname              = "ns_part_hostname"
	dnsSchemaNSECRecord              = "nsecrecord"
	dnsSchemaPTRRecord               = "ptrrecord"
	dnsSchemaPTRHostname             = "ptr_part_hostname"
	dnsSchemaRRSIGRecord             = "rrsigrecord"
	dnsSchemaRPRecord                = "rprecord"
	dnsSchemaSIGRecord               = "sigrecord"
	dnsSchemaSPFRecord               = "spfrecord"
	dnsSchemaSRVRecord               = "srvrecord"
	dnsSchemaSRVPriority             = "srv_part_priority"
	dnsSchemaSRVWeight               = "srv_part_weight"
	dnsSchemaSRVPort                 = "srv_part_port"
	dnsSchemaSRVTarget               = "srv_part_target"
	dnsSchemaSSHFPRecord             = "sshfprecord"
	dnsSchemaSSHFPAlgorithm          = "sshfp_part_algorithm"
	dnsSchemaSSHFPFingerprintType    = "sshfp_part_fp_type"
	dnsSchemaSSHFPFingerPrint        = "sshfp_part_fingerprint"
	dnsSchemaTLSARecord              = "tlsarecord"
	dnsSchemaTLSACertUsage           = "tlsa_part_cert_usage"
	dnsSchemaTLSASelector            = "tlsa_part_selector"
	dnsSchemaTLSAMatchingType        = "tlsa_part_matching_type"
	dnsSchemaTLSACertAssociationData = "tlsa_part_cert_association_data"
	dnsSchemaTXTRecord               = "txtrecord"
	dnsSchemaTXTData                 = "txt_part_data"
	dnsSchemaURIRecord               = "urirecord"
	dnsSchemaURIPriority             = "uri_part_priority"
	dnsSchemaURIWeight               = "uri_part_weight"
	dnsSchemaURITarget               = "uri_part_target"
)

var dnsSchemaRecordKeys = []string{
	dnsSchemaRecordName,
	dnsSchemaRecordZoneName,
	dnsSchemaRecordTTL,
	dnsSchemaRecordClass,
	dnsSchemaForce,
	dnsSchemaARecord,
	dnsSchemaAIPAddress,
	dnsSchemaACreateReverse,
	dnsSchemaAAAARecord,
	dnsSchemaAAAAIPAddress,
	dnsSchemaAAAACreateReverse,
	dnsSchemaA6Record,
	dnsSchemaA6Data,
	dnsSchemaAFSDBRecord,
	dnsSchemaAFSDBSubtype,
	dnsSchemaAFSDBHostname,
	dnsSchemaAPLRecord,
	dnsSchemaCERTRecord,
	dnsSchemaCERTType,
	dnsSchemaCERTKeyTag,
	dnsSchemaCERTAlgorithm,
	dnsSchemaCERTCertOrCRL,
	dnsSchemaCNAMERecord,
	dnsSchemaCNAMEHostname,
	dnsSchemaDHCIDRecord,
	dnsSchemaDLVRecord,
	dnsSchemaDLVKeyTag,
	dnsSchemaDLVAlgorithm,
	dnsSchemaDLVDigestType,
	dnsSchemaDLVDigest,
	dnsSchemaDNAMERecord,
	dnsSchemaDNAMETarget,
	dnsSchemaDSRecord,
	dnsSchemaDSKeyTag,
	dnsSchemaDSAlgorithm,
	dnsSchemaDSDigestType,
	dnsSchemaDSDigest,
	dnsSchemaHIPRecord,
	dnsSchemaIPSECRecord,
	dnsSchemaKeyRecord,
	dnsSchemaKXRecord,
	dnsSchemaKXPreference,
	dnsSchemaKXExchanger,
	dnsSchemaLOCRecord,
	dnsSchemaLOCLatDeg,
	dnsSchemaLOCLatMin,
	dnsSchemaLOCLatSec,
	dnsSchemaLOCLatDir,
	dnsSchemaLOCLonDeg,
	dnsSchemaLOCLonMin,
	dnsSchemaLOCLonSec,
	dnsSchemaLOCLonDir,
	dnsSchemaLOCAltitude,
	dnsSchemaLOCSize,
	dnsSchemaLOCHorizontalPrecision,
	dnsSchemaLOCVerticalPrecision,
	dnsSchemaMXRecord,
	dnsSchemaMXPreference,
	dnsSchemaMXExchanger,
	dnsSchemaNAPTRRecord,
	dnsSchemaNAPTROrder,
	dnsSchemaNAPTRPreference,
	dnsSchemaNAPTRFlags,
	dnsSchemaNAPTRService,
	dnsSchemaNAPTRRegexp,
	dnsSchemaNAPTRReplacement,
	dnsSchemaNSRecord,
	dnsSchemaNSHostname,
	dnsSchemaNSECRecord,
	dnsSchemaPTRRecord,
	dnsSchemaPTRHostname,
	dnsSchemaRRSIGRecord,
	dnsSchemaRPRecord,
	dnsSchemaSIGRecord,
	dnsSchemaSPFRecord,
	dnsSchemaSRVRecord,
	dnsSchemaSRVPriority,
	dnsSchemaSRVWeight,
	dnsSchemaSRVPort,
	dnsSchemaSRVTarget,
	dnsSchemaSSHFPRecord,
	dnsSchemaSSHFPAlgorithm,
	dnsSchemaSSHFPFingerprintType,
	dnsSchemaSSHFPFingerPrint,
	dnsSchemaTLSARecord,
	dnsSchemaTLSACertUsage,
	dnsSchemaTLSASelector,
	dnsSchemaTLSAMatchingType,
	dnsSchemaTLSACertAssociationData,
	dnsSchemaTXTRecord,
	dnsSchemaTXTData,
	dnsSchemaURIRecord,
	dnsSchemaURIPriority,
	dnsSchemaURIWeight,
	dnsSchemaURITarget,
}

func resourceDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSRecordCreate,
		Read:   resourceDNSRecordRead,
		Update: resourceDNSRecordUpdate,
		Delete: resourceDNSRecordDelete,
		Exists: resourceDNSRecordExists,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			dnsSchemaRecordName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dnsSchemaRecordZoneName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			dnsSchemaRecordTTL: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaRecordClass: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaForce: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaARecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaAIPAddress: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaACreateReverse: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaAAAARecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaAAAAIPAddress: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAAAACreateReverse: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			dnsSchemaA6Record: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaA6Data: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAFSDBRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaAFSDBSubtype: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaAFSDBHostname: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaAPLRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaCERTRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaCERTType: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaCERTKeyTag: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaCERTAlgorithm: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaCERTCertOrCRL: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaCNAMERecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaCNAMEHostname: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaDHCIDRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaDLVRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaDLVKeyTag: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDLVAlgorithm: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDLVDigestType: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDLVDigest: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaDNAMERecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaDNAMETarget: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaDSRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaDSKeyTag: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDSAlgorithm: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDSDigestType: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaDSDigest: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaHIPRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaIPSECRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaKeyRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaKXRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaKXPreference: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaKXExchanger: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaLOCRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaLOCLatDeg: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaLOCLatMin: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaLOCLatSec: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			dnsSchemaLOCLatDir: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaLOCLonDeg: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaLOCLonMin: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaLOCLonSec: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			dnsSchemaLOCLonDir: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaLOCAltitude: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			dnsSchemaLOCSize: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			dnsSchemaLOCHorizontalPrecision: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			dnsSchemaLOCVerticalPrecision: {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			dnsSchemaMXRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaMXPreference: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaMXExchanger: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaNAPTRRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaNAPTROrder: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaNAPTRPreference: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaNAPTRFlags: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaNAPTRService: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaNAPTRRegexp: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaNAPTRReplacement: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaNSRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaNSHostname: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaNSECRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaPTRRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaPTRHostname: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaRRSIGRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaRPRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaSIGRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaSPFRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaSRVRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaSRVPriority: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSRVWeight: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSRVPort: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSRVTarget: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaSSHFPRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaSSHFPAlgorithm: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSSHFPFingerprintType: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaSSHFPFingerPrint: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaTLSARecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaTLSACertUsage: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaTLSASelector: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaTLSAMatchingType: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaTLSACertAssociationData: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaTXTRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaTXTData: {
				Type:     schema.TypeString,
				Optional: true,
			},
			dnsSchemaURIRecord: {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			dnsSchemaURIPriority: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaURIWeight: {
				Type:     schema.TypeInt,
				Optional: true,
			},
			dnsSchemaURITarget: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSRecordCreate - %s - zone %s", d.Id(), d.Get(dnsSchemaRecordZoneName).(string))

	c := m.(*Connection)

	options := map[string]interface{}{}

	for _, key := range dnsSchemaRecordKeys {
		if val, ok := d.GetOk(key); ok {
			options[key] = val
		}
	}

	log.Print("[TRACE] creating DNS Record")

	dns, err := c.client.CreateDNSRecord(options)

	if err != nil {
		log.Printf("[ERROR] Error creating DNS Record %v", err)
		return err
	}

	d.SetId(dns.Name.String())

	return resourceDNSRecordRead(d, m)
}

func resourceDNSRecordRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSRecordRead - %s - zone %s", d.Id(), d.Get(dnsSchemaRecordZoneName).(string))

	c := m.(*Connection)

	zone := d.Get(dnsSchemaRecordZoneName).(string)
	_, err := c.client.GetDNSRecord(d.Id(), zone)

	if err != nil {
		log.Printf("[ERROR] Error reading DNS Record %s for zone %s - %s", d.Id(), zone, err)
		return err
	}

	return nil
}

func resourceDNSRecordUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSRecordUpdate - %s - zone %s", d.Id(), d.Get(dnsSchemaRecordZoneName).(string))

	c := m.(*Connection)

	zone := d.Get(dnsSchemaRecordZoneName).(string)
	rec, err := c.ldapClient.GetDNSRecord(d.Id(), zone)

	if err != nil {
		log.Printf("[ERROR] Error reading DNS Record %s for zone %s - %s", d.Id(), zone, err)
		return err
	}
	d.Partial(true)

	for _, key := range dnsSchemaRecordKeys {
		if !d.HasChange(key) {
			continue
		}

		_, newVal := d.GetChange(key)
		c.client.DNSRecordMod(*rec, zone, key, newVal)
		d.SetPartial(key)
		log.Printf("[TRACE] resource resourceDNSRecordUpdate record %s key %s updated", *rec, key)
	}

	d.Partial(false)

	return resourceDNSRecordRead(d, m)
}

func resourceDNSRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] resourceDNSRecordDelete - %s - zone %s", d.Id(), d.Get(dnsSchemaRecordZoneName).(string))

	c := m.(*Connection)

	zone := d.Get(dnsSchemaRecordZoneName).(string)
	rec, err := c.ldapClient.GetDNSRecord(d.Id(), zone)

	if err != nil {
		log.Printf("[ERROR] Error reading DNS Record %s for zone %s - %s", rec, zone, err)
		return err
	}

	return c.client.DeleteDNSRecord(*rec, zone)
}

func resourceDNSRecordExists(d *schema.ResourceData, m interface{}) (bool, error) {
	log.Printf("[TRACE] resourceDNSRecordExists - %s - zone %s", d.Id(), d.Get(dnsSchemaRecordZoneName).(string))

	c := m.(*Connection)

	exists, err := c.ldapClient.DNSRecordExists(d.Id(), d.Get(dnsSchemaRecordZoneName).(string))

	if err != nil {
		log.Printf("[ERROR] Error checking existence dns Record %s - %s", d.Id(), err)
		return false, err
	}

	return exists, nil
}
