package constants

import "time"

const (
	// ObfuscationString is the string is used to hide sensitive data such as token in logs
	ObfuscationString = "*************"

	// TimestampFormat is the common timestamp format ACR Builder uses
	TimestampFormat = time.RFC3339

	// SourceNameWebArchive is the name of the web archive source
	SourceNameWebArchive = "web archive"

	// SourceNameGit is the name of the git source
	SourceNameGit = "git repository"

	// SourceNameLocal is the name of local source
	SourceNameLocal = "local directory"

	// DependencyTypeBuild denotes build time dependency
	DependencyTypeBuild = "build"

	// DependencyTypeRuntime denotes runtime dependency
	DependencyTypeRuntime = "runtime"

	// FromStdin denotes when context or dockerfile is to be read from stdin
	FromStdin = "-"

	// NoBaseImageSpecifierLatest is the empty base image
	// Note that :latest is not valid in the FROM clause, but we're
	// always appending :latest to tags during processing.
	NoBaseImageSpecifierLatest = "scratch:latest"
)
