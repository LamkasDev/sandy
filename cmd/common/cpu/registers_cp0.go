package cpu

var PlaystationCp0RegisterNames = map[PlaystationRegisterId]string{
	0x00: "index",
	0x01: "random",
	0x02: "entrylo0",
	0x03: "entrylo1",
	0x04: "context",
	0x05: "pagemask",
	0x06: "wired",
	0x07: "reserved",
	0x08: "badvaddr",
	0x09: "count",
	0x0a: "entryhi",
	0x0b: "compare",
	0x0c: "status",
	0x0d: "cause",
	0x0e: "epc",
	0x0f: "prid",
	0x10: "config",
	0x11: "lladdr",
	0x12: "watchlo",
	0x13: "watchhi",
	0x14: "reserved",
	0x15: "debug",
	0x16: "depc",
	0x17: "perfcnt",
	0x18: "errctl",
	0x19: "cacheerr",
	0x1a: "taglo",
	0x1b: "datalo",
	0x1c: "taghi",
	0x1d: "datahi",
	0x1e: "errorepc",
	0x1f: "desave",
}

const PlaystationCp0RegisterStatus = PlaystationRegisterId(0x0c)
