package organizers

import "github.com/readytotouch/readytotouch/internal/domain"

// https://chatgpt.com/share/66e88b74-6470-8007-825c-e26def03e5f6
var (
	CompanyStartupMap = map[string]bool{
		"readytotouch":    true,
		"dochq":           true,
		"bereal-app":      true,
		"victoriametrics": true,
		"ferretdb":        true,
		"salesforgeai":    true,
		"weaviate-io":     true,
		"yassir":          true,
		"vodeno":          true,
		"woolsocks":       true,
		"xataio":          true,
		"firebolt":        true,
		"isovalent":       true,
		"zep-ai":          true,
		"dusty-robotics":  true,
		"fincompare":      true,
		"goflink":         true,
		"asset-reality":   true,
		"mycafu":          true,
		"rollee":          true,
		"treecardapp":     true,
		"voltus-inc.":     true,
		"toggle-ai":       true,
		"hearx-group":     true,
		"nuro-inc.":       true,
		"cloudwalk-inc":   true,
		"checkout":        true,
	}
)

func ToCompanyType(alias string) domain.CompanyType {
	if CompanyStartupMap[alias] {
		return domain.CompanyTypeStartup
	}

	return domain.CompanyTypeProduct
}
