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
	PackageBase          string             `json:"packageBase"`
	BaseIndex            int                `json:"baseIndex"`
	BaseTotal            int                `json:"baseTotal"`
	BaseChildren         []string           `json:"baseChildren"`
	Description          string             `json:"description"`
	Version              string             `json:"version"`
	SourceVersion        string             `json:"sourceVersion"`
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

func FromSrcInfo(info srcinfo.Srcinfo) []*Script {
	var scripts []*Script
	if len(info.Packages) > 1 {
		children := make([]string, len(info.Packages))
		for i, pkg := range info.Packages {
			children[i] = pkg.Pkgname
		}
		scripts = append(scripts, &Script{
			PackageName:          info.Pkgbase,
			PrettyName:           "",
			PackageBase:          info.Pkgbase,
			BaseIndex:            0,
			BaseTotal:            len(info.Packages),
			BaseChildren:         children,
			Description:          info.Pkgdesc,
			Version:              info.Version(),
			SourceVersion:		  info.Pkgver,
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
			Repology:             orEmptyArray(info.Repology),
			RequiredBy:           []string{},
			UpdateStatus:         UpdateStatus.Unknown,
		})
	}
	for i, pkg := range info.Packages {
		scripts = append(scripts, &Script{
			PackageName:          pkg.Pkgname,
			PrettyName:           "",
			PackageBase:          info.Pkgbase,
			BaseIndex:            i + 1,
			BaseTotal:            len(info.Packages),
			BaseChildren:         nil,
			Description:          fallback[string, string](pkg.Pkgdesc, info.Pkgdesc, nil),
			Version:              info.Version(),
			SourceVersion:		  info.Pkgver,
			Release:              info.Pkgrel,
			Epoch:                info.Epoch,
			LatestVersion:        nil,
			Homepage:             info.URL,
			Priority:             fallback[string, string](pkg.Priority, info.Priority, nil),
			Architectures:        info.Arch,
			License:              fallback(pkg.License, info.License, orEmptyArray),
			Gives:                fallback(pkg.Gives, info.Gives, toArchDistroStrings),
			RuntimeDependencies:  fallback(pkg.Depends, info.Depends, toArchDistroStrings),
			CheckDependencies:    fallback(pkg.CheckDepends, info.CheckDepends, toArchDistroStrings),
			BuildDependencies:    toArchDistroStrings(info.MakeDepends),
			OptionalDependencies: fallback(pkg.OptDepends, info.OptDepends, toArchDistroStrings),
			PacstallDependencies: fallback(pkg.Pacdeps, info.Pacdeps, toArchDistroStrings),
			CheckConflicts:       fallback(pkg.CheckConflicts, info.CheckConflicts, toArchDistroStrings),
			BuildConflicts:       toArchDistroStrings(info.MakeConflicts),
			Conflicts:            fallback(pkg.Conflicts, info.Conflicts, toArchDistroStrings),
			Provides:             fallback(pkg.Provides, info.Provides, toArchDistroStrings),
			Breaks:               fallback(pkg.Breaks, info.Breaks, toArchDistroStrings),
			Replaces:             fallback(pkg.Replaces, info.Replaces, toArchDistroStrings),
			Enhances:             fallback(pkg.Enhances, info.Enhances, toArchDistroStrings),
			Recommends:           fallback(pkg.Recommends, info.Recommends, toArchDistroStrings),
			Suggests:             fallback(pkg.Suggests, info.Suggests, toArchDistroStrings),
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
			Backup:               fallback(pkg.Backup, info.Backup, orEmptyArray),
			Repology:             fallback(pkg.Repology, info.Repology, orEmptyArray),
			RequiredBy:           []string{},
			UpdateStatus:         UpdateStatus.Unknown,
		})
	}

	return scripts
}

func fallback[T any, R any](pkgValue, infoValue T, transform func(T) R) R {
	if isNonEmpty(pkgValue) {
		if transform != nil {
			return transform(pkgValue)
		}
		return any(pkgValue).(R)
	}
	if transform != nil {
		return transform(infoValue)
	}
	return any(infoValue).(R)
}

func isNonEmpty[T any](value T) bool {
	switch v := any(value).(type) {
	case string:
		return v != ""
	case []string:
		return len(v) > 0
	case []srcinfo.ArchDistroString:
		return len(v) > 0
	default:
		return true
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
