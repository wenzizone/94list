package build

import "expvar"

var (
	// revision denotes the git commit revision at which the binary is built. The default value
	// provided here is father of all git revisions, e.g., an empty git tree. It will be replaced to
	// HEAD commit during build time.
	revision = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"

	// version denotes the version of Jaeger which the binary is tagged with. The default value
	// provided here is a placeholder and should not appear. It will be replaced with the
	// build time version.
	version = "v0.0.0"
)

// ExposeBuildInfo method exposes the build information such as revision and version via the expvar
// package.
func ExposeBuildInfo() {
	expvar.NewString("rev").Set(revision)
	expvar.NewString("ver").Set(version)
}

// Revision returns the current server revision which is just some sort of uniquely identifiable hash.
func Revision() string {
	return revision
}

// Version returns the current server version which is a string given at build time through
// a script or job (run by Jenkins for example).
func Version() string {
	return version
}
