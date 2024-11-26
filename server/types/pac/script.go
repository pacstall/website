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
	PackageName          string             `json:"packageName"`
	PrettyName           string             `json:"prettyName"`
	Description          string             `json:"description"`
	Version              string             `json:"version"`
	Release              string             `json:"release"`
	Epoch                string             `json:"epoch"`
	LatestVersion        *string            `json:"latestVersion"`
	Homepage             string             `json:"homepage"`
	Priority             string             `json:"priority"`
	Architectures        []string           `json:"architectures"`
	License              []string           `json:"license"`
	Gives                []ArchDistroString `json:"gives"`
	RuntimeDependencies  []ArchDistroString `json:"runtimeDependencies"`
	CheckDependencies    []ArchDistroString `json:"checkDependencies"`
	BuildDependencies    []ArchDistroString `json:"buildDependencies"`
	OptionalDependencies []ArchDistroString `json:"optionalDependencies"`
	PacstallDependencies []ArchDistroString `json:"pacstallDependencies"`
	CheckConflicts       []ArchDistroString `json:"checkConflicts"`
	BuildConflicts       []ArchDistroString `json:"buildConflicts"`
	Conflicts            []ArchDistroString `json:"conflicts"`
	Provides             []ArchDistroString `json:"provides"`
	Breaks               []ArchDistroString `json:"breaks"`
	Replaces             []ArchDistroString `json:"replaces"`
	Enhances             []ArchDistroString `json:"enhances"`
	Recommends           []ArchDistroString `json:"recommends"`
	Suggests             []ArchDistroString `json:"suggests"`
	Mask                 []string           `json:"mask"`
	Compatible           []string           `json:"compatible"`
	Incompatible         []string           `json:"incompatible"`
	Maintainers          []string           `json:"maintainers"`
	Source               []ArchDistroString `json:"source"`
	NoExtract            []string           `json:"noExtract"`
	NoSubmodules         []string           `json:"NoSubmodules"`
	Md5Sums              []ArchDistroString `json:"md5sums"`
	Sha1Sums             []ArchDistroString `json:"sha1sums"`
	Sha224Sums           []ArchDistroString `json:"sha224sums"`
	Sha256Sums           []ArchDistroString `json:"sha256sums"`
	Sha384Sums           []ArchDistroString `json:"sha384sums"`
	Sha512Sums           []ArchDistroString `json:"sha512sums"`
	Backup               []string           `json:"backup"`
	Repology             []string           `json:"repology"`
	RequiredBy           []string           `json:"requiredBy"`
	UpdateStatus         int                `json:"updateStatus"` // enum UpdateStatus
	LastUpdatedAt        time.Time          `json:"lastUpdatedAt"`
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
		PackageName:          info.Packages[0].Pkgname, // needs to be looped for every pkgname within pkgbase
		PrettyName:           "",
		Description:          info.Pkgdesc,
		Version:              info.Version(),
		Release:              info.Pkgrel,
		Epoch:                info.Epoch,
		LatestVersion:        nil,
		Homepage:             info.URL,
		Priority:             info.Priority,
		Architectures:        info.Arch,
		License:              orEmptyArray(info.License),
		Gives:                toArchDistroStrings(info.Gives),
		RuntimeDependencies:  toArchDistroStrings(info.Depends),
		CheckDependencies:    toArchDistroStrings(info.CheckDepends),
		BuildDependencies:    toArchDistroStrings(info.MakeDepends),
		OptionalDependencies: toArchDistroStrings(info.OptDepends),
		PacstallDependencies: toArchDistroStrings(info.Pacdeps),
		CheckConflicts:       toArchDistroStrings(info.CheckConflicts),
		BuildConflicts:       toArchDistroStrings(info.MakeConflicts),
		Conflicts:            toArchDistroStrings(info.Conflicts),
		Provides:             toArchDistroStrings(info.Provides),
		Breaks:               toArchDistroStrings(info.Breaks),
		Replaces:             toArchDistroStrings(info.Replaces),
		Enhances:             toArchDistroStrings(info.Enhances),
		Recommends:           toArchDistroStrings(info.Recommends),
		Suggests:             toArchDistroStrings(info.Suggests),
		Mask:                 orEmptyArray(info.Mask),
		Compatible:           orEmptyArray(info.Compatible),
		Incompatible:         orEmptyArray(info.Incompatible),
		Maintainers:          info.Maintainer,
		Source:               toSourceStrings(info.Source),
		NoExtract:            orEmptyArray(info.NoExtract),
		NoSubmodules:         orEmptyArray(info.NoSubmodules),
		Md5Sums:              toArchDistroStrings(info.MD5Sums),
		Sha1Sums:             toArchDistroStrings(info.SHA1Sums),
		Sha224Sums:           toArchDistroStrings(info.SHA224Sums),
		Sha256Sums:           toArchDistroStrings(info.SHA256Sums),
		Sha384Sums:           toArchDistroStrings(info.SHA384Sums),
		Sha512Sums:           toArchDistroStrings(info.SHA512Sums),
		Backup:               orEmptyArray(info.Backup),
		Repology:             info.Repology,
		RequiredBy:           []string{},
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

func toSourceStrings(ads []srcinfo.ArchDistroString) []ArchDistroString {
	return array.SwitchMap(ads, func(it *array.Iterator[srcinfo.ArchDistroString]) ArchDistroString {
		return ArchDistroString{
			Arch:   it.Value.Arch,
			Distro: it.Value.Distro,
			Value:  extractSourceUrl(it.Value.Value),
		}
	})
}

func orEmptyArray[T interface{}](items []T) []T {
	if items == nil {
		return []T{}
	}

	return items
}

func extractSourceUrl(source string) string {
	parts := strings.Split(source, "::")
	if len(parts) > 1 {
		return parts[1]
	}
	return parts[0]
}
