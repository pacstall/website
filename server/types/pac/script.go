package pac

import (
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
	Arch   string `json:"arch"`
	Distro string `json:"distro"`
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
	Conflicts            []ArchDistroString `json:"conflicts"`
	Breaks               []ArchDistroString `json:"breaks"`
	Gives                []ArchDistroString `json:"gives"`
	Replaces             []ArchDistroString `json:"replaces"`
	Hash                 *string            `json:"hash"`
	PacstallDependencies []ArchDistroString `json:"pacstallDependencies"`
	Repology             []string           `json:"repology"`
	RequiredBy           []string           `json:"requiredBy"`
	LastUpdatedAt        time.Time          `json:"lastUpdatedAt"`
	UpdateStatus         int                `json:"updateStatus"` // enum UpdateStatus
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
		PrettyName:           "",
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
