package cloudparser

type VirtualMachine struct {
	VMID string   `json:"vm_id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type FirewallRule struct {
	FWID      string `json:"fw_id"`
	SourceTag string `json:"source_tag"`
	DestTag   string `json:"dest_tag"`
}

type CloudEnvironment struct {
	VMs     []VirtualMachine `json:"vms"`
	FWRules []FirewallRule   `json:"fw_rules"`
}
