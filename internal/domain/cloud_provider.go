package domain

type CloudProvider = AliasName

// Enterprise cloud providers
var (
	AWS = CloudProvider{
		Alias: "aws",
		Name:  "Amazon Web Services",
	}
	GCP = CloudProvider{
		Alias: "gcp",
		Name:  "Google Cloud Platform",
	}
	Azure = CloudProvider{
		Alias: "azure",
		Name:  "Microsoft Azure",
	}
	OracleCloud = CloudProvider{
		Alias: "oracle-cloud",
		Name:  "Oracle Cloud Infrastructure",
	}
	IBMCloud = CloudProvider{
		Alias: "ibm-cloud",
		Name:  "IBM Cloud",
	}
)

// Developer-friendly cloud providers
var (
	DigitalOcean = CloudProvider{
		Alias: "digital-ocean",
		Name:  "DigitalOcean",
	}
	Vultr = CloudProvider{
		Alias: "vultr",
		Name:  "Vultr",
	}
)

var (
	Cloudflare = CloudProvider{
		Alias: "cloudflare",
		Name:  "Cloudflare",
	}
)
