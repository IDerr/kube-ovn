package util

import "strings"

// LabelSafeIP converts an IP address into a value usable as a Kubernetes label.
// IPv6 addresses contain ':' which is illegal in label values (and breaks
// OvnEip/OvnFip/OvnSnat creation for IPv6 external subnets). Replacing ':' with
// '-' yields a valid, unique, reversible label value. Must be applied consistently
// at every site that writes OR reads the eip_v6_ip label (CloudZero patch).
func LabelSafeIP(ip string) string {
	return strings.ReplaceAll(ip, ":", "-")
}
