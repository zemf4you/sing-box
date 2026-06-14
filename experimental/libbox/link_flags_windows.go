//go:build windows

package libbox

import "net"

// Pacific seam: upstream only implements linkFlags on unix (link_flags_unix.go)
// and panics on every other GOOS (link_flags_stub.go), which crashes the
// platform NetworkInterfaces() path on Windows. Windows has no syscall.IFF_*
// raw flags — Go's net.Interface.Flags is already a net.Flags bitmask — so the
// desktop platform hands those bits through as-is and this conversion is the
// identity.
func linkFlags(rawFlags uint32) net.Flags {
	return net.Flags(rawFlags)
}
