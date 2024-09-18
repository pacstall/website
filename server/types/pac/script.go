package pac

import (
	"strings"
	"time"

	"github.com/pacstall/go-srcinfo"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/array"
)

type updateStatus struct {
	Unknown UpdateStatusValue
	Latest  UpdateStatusValue
	Minor   UpdateStatusValue
	Major   UpdateStatusValue
	Patch   UpdateStatusValue
}

var UpdateStatus = updateStatus{
	Major:   3,
	Minor:   2,
	Patch:   1,
	Latest:  0,
	Unknown: -1,
}

type UpdateStatusValue = int

type ArchDistroString struct {
	Arch   string `json:"arch,omitempty"`
	Distro string `json:"distro,omitempty"`
	Value  string `json:"value"`
}

func (a ArchDistroString) Equals(b types.Equaller) bool {
	other, ok := b.(ArchDistroString)
	if !ok {
		return false
	}

	return a.Arch == other.Arch && a.Distro == other.Distro && a.Value == other.Value
}

type Script struct {
	Architectures        []string           `json:"architectures"`
	PrettyName           string             `json:"prettyName"`
	Version              string             `json:"version"`
	LatestVersion        *string            `json:"latestVersion"`
	PackageName          string             `json:"packageName"`
	Maintainers          []string           `json:"maintainers"`
	Description          string             `json:"description"`
	Source               []ArchDistroString `json:"source"`
	RuntimeDependencies  []ArchDistroString `json:"runtimeDependencies"`
	BuildDependencies    []ArchDistroString `json:"buildDependencies"`
	OptionalDependencies []ArchDistroString `json:"optionalDependencies"`
	CheckDependencies    []ArchDistroString `json:"checkDependencies"`
	Conflicts            []ArchDistroString `json:"conflicts"`
	Breaks               []ArchDistroString `json:"breaks"`
	Gives                []ArchDistroString `json:"gives"`
	Replaces             []ArchDistroString `json:"replaces"`
	Sha1Sums             []ArchDistroString `json:"sha1sums"`
	Sha224Sums           []ArchDistroString `json:"sha224sums"`
	Sha256Sums           []ArchDistroString `json:"sha256sums"`
	Sha384Sums           []ArchDistroString `json:"sha384sums"`
	Sha512Sums           []ArchDistroString `json:"sha512sums"`
	Md5Sums              []ArchDistroString `json:"md5sums"`
	Priority             []ArchDistroString `json:"priority"`
	Recommends           []ArchDistroString `json:"recommends"`
	Suggests             []ArchDistroString `json:"suggests"`
	PacstallDependencies []ArchDistroString `json:"pacstallDependencies"`
	Enhances             []ArchDistroString `json:"enhances"`
	Repology             []string           `json:"repology"`
	RequiredBy           []string           `json:"requiredBy"`
	LastUpdatedAt        time.Time          `json:"lastUpdatedAt"`
	UpdateStatus         int                `json:"updateStatus"` // enum UpdateStatus
	Changelog            string             `json:"changelog"`
	Backup               []string           `json:"backup"`
	Compatible           []string           `json:"compatible"`
	Incompatible         []string           `json:"incompatible"`
	Epoch                string             `json:"epoch"`
	Install              string             `json:"install"`
	License              []string           `json:"license"`
	Mask                 []string           `json:"mask"`
	NoExtract            []string           `json:"noExtract"`
	ValidPGPKeys         []string           `json:"validPgpKeys"`
	Groups               []string           `json:"groups"`
}

func (p *Script) Type() types.PackageTypeName {
	for suffix, name := range types.PackageTypeSuffixToPackageTypeName {
		if strings.HasSuffix(p.PackageName, string(suffix)) {
			return name
		}
	}

	return types.PackageTypeSuffixToPackageTypeName["-git"]
}

func FromSrcInfo(info srcinfo.Srcinfo) *Script {
	return &Script{
		Version:              info.Version(),
		LatestVersion:        nil,
		PackageName:          info.Packages[0].Pkgname,
		Maintainers:          info.Maintainer,
		Description:          info.Pkgdesc,
		Source:               toArchDistroStrings(info.Source),
		RuntimeDependencies:  toArchDistroStrings(info.Depends),
		BuildDependencies:    toArchDistroStrings(info.MakeDepends),
		OptionalDependencies: toArchDistroStrings(info.OptDepends),
		Conflicts:            toArchDistroStrings(info.Conflicts),
		Breaks:               toArchDistroStrings(info.Breaks),
		Gives:                toArchDistroStrings(info.Gives),
		Replaces:             toArchDistroStrings(info.Replaces),
		PacstallDependencies: toArchDistroStrings(info.Pacdeps),
		Architectures:        info.Arch,
		Repology:             info.Repology,
		RequiredBy:           []string{},
		Sha1Sums:             toArchDistroStrings(info.SHA1Sums),
		Sha224Sums:           toArchDistroStrings(info.SHA224Sums),
		Sha256Sums:           toArchDistroStrings(info.SHA256Sums),
		Sha384Sums:           toArchDistroStrings(info.SHA384Sums),
		Sha512Sums:           toArchDistroStrings(info.SHA512Sums),
		PrettyName:           "",
		Changelog:            info.Changelog,
		Backup:               orEmptyArray(info.Backup),
		Compatible:           orEmptyArray(info.Compatible),
		Incompatible:         orEmptyArray(info.Incompatible),
		Epoch:                info.Epoch,
		Install:              info.Install,
		License:              orEmptyArray(info.License),
		Mask:                 orEmptyArray(info.Mask),
		NoExtract:            orEmptyArray(info.NoExtract),
		ValidPGPKeys:         orEmptyArray(info.ValidPGPKeys),
		Groups:               orEmptyArray(info.Groups),
		Enhances:             toArchDistroStrings(info.Enhances),
		CheckDependencies:    toArchDistroStrings(info.CheckDepends),
		Md5Sums:              toArchDistroStrings(info.MD5Sums),
		Priority:             toArchDistroStrings(info.Priority),
		Suggests:             toArchDistroStrings(info.Suggests),
		Recommends:           toArchDistroStrings(info.Recommends),
		UpdateStatus:         UpdateStatus.Unknown,
	}
}

func toArchDistroStrings(ads []srcinfo.ArchDistroString) []ArchDistroString {
	return array.SwitchMap(ads, func(it *array.Iterator[srcinfo.ArchDistroString]) ArchDistroString {
		return toArchDistroString(it.Value)
	})
}

func toArchDistroString(ads srcinfo.ArchDistroString) ArchDistroString {
	return ArchDistroString{
		Arch:   ads.Arch,
		Distro: ads.Distro,
		Value:  ads.Value,
	}
}

func orEmptyArray[T interface{}](items []T) []T {
	if items == nil {
		return []T{}
	}

	return items
}
