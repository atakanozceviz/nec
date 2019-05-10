package csproj

type Csproj struct {
	Name       string
	FilePath   string
	ItemGroups []ItemGroup `xml:"ItemGroup,omitempty"`
}

type ItemGroup struct {
	ProjectReferences []ProjectReference `xml:"ProjectReference,omitempty"`
	PackageReferences []PackageReference `xml:"PackageReference"`
}

type ProjectReference struct {
	Include string `xml:"Include,attr,omitempty"`
}
type PackageReference struct {
	Text    string `xml:",chardata"`
	Include string `xml:"Include,attr"`
	Version string `xml:"Version,attr"`
}

// IsTest return true if project is a test project, if not return false.
func (c *Csproj) IsTest() bool {
	for _, ig := range c.ItemGroups {
		for _, pr := range ig.PackageReferences {
			if pr.Include == "Microsoft.NET.Test.Sdk" {
				return true
			}
		}
	}
	return false
}
