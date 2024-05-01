package service

type GitVersionResolver interface {
	ParseGitPackageVersion(sources []string) (string, error)
}
