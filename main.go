package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"reflect"

	m "github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/types"
)

// type A struct {
// 	hdr         types.FileHeader

// 	[]any{
// 		&SegmentHeader{types.LC_SEGMENT_64, 0x48, "__PAGEZERO", 0x0, 0x100000000, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
// 		&SegmentHeader{types.LC_SEGMENT_64, 0x1d8, "__TEXT", 0x100000000, 0x1000, 0x0, 0x1000, 0x7, 0x5, 0x5, 0x0, 0},
// 		&SegmentHeader{LoadCmd: 0x19, Len: 0x138, Name: "__DATA", Addr: 0x100001000, Memsz: 0x1000, Offset: 0x1000, Filesz: 0x1000, Maxprot: 7, Prot: 3, Nsect: 0x3, Flag: 0x0, Firstsect: 0x5},
// 		&SegmentHeader{LoadCmd: 0x19, Len: 0x48, Name: "__LINKEDIT", Addr: 0x100002000, Memsz: 0x1000, Offset: 0x2000, Filesz: 0x140, Maxprot: 7, Prot: 1, Nsect: 0x0, Flag: 0x0, Firstsect: 0x8},
// 		nil, // LC_SYMTAB
// 		nil, // LC_DYSYMTAB
// 		nil, // LC_LOAD_DYLINKER
// 		nil, // LC_UUID
// 		nil, // LC_UNIXTHREAD
// 		----------- &LoadDylib{Dylib{LoadBytes: LoadBytes(nil), DylibCmd: types.DylibCmd{LoadCmd: 0xc, Len: 0x38, NameOffset: 0x18, Timestamp: 0x2, CurrentVersion: types.Version(0x10000), CompatVersion: types.Version(0x10000)}, Name: "/usr/lib/libgcc_s.1.dylib"}},
// 		&LoadDylib{Dylib{LoadBytes: LoadBytes(nil), DylibCmd: types.DylibCmd{LoadCmd: 0xc, Len: 0x38, NameOffset: 0x18, Timestamp: 0x2, CurrentVersion: types.Version(0x6f0104), CompatVersion: types.Version(0x10000)}, Name: "/usr/lib/libSystem.B.dylib"}},
// 	},
// 	loads       []any

// 	sections    []*types.SectionHeader
// 	[]*types.SectionHeader{
// 		{Name: "__text", Seg: "__TEXT", Addr: 0x100000f14, Size: 0x6d, Offset: 0xf14, Align: 0x2, Reloff: 0x0, Nreloc: 0x0, Flags: 0x80000400, Reserved1: 0x0, Reserved2: 0x0, Reserved3: 0x0, Type: 0x40},
// 		{Name: "__symbol_stub1", Seg: "__TEXT", Addr: 0x100000f81, Size: 0xc, Offset: 0xf81, Align: 0x0, Reloff: 0x0, Nreloc: 0x0, Flags: 0x80000408, Reserved1: 0x0, Reserved2: 0x6, Reserved3: 0x0, Type: 0x40},
// 		{Name: "__stub_helper", Seg: "__TEXT", Addr: 0x100000f90, Size: 0x18, Offset: 0xf90, Align: 0x2, Reloff: 0x0, Nreloc: 0x0, Flags: 0x0, Reserved1: 0, Reserved2: 0, Reserved3: 0, Type: 0x40},
// 		{Name: "__cstring", Seg: "__TEXT", Addr: 0x100000fa8, Size: 0xd, Offset: 0xfa8, Align: 0x0, Reloff: 0x0, Nreloc: 0x0, Flags: 0x2, Reserved1: 0x0, Reserved2: 0x0, Reserved3: 0x0, Type: 0x40},
// 		{Name: "__eh_frame", Seg: "__TEXT", Addr: 0x100000fb8, Size: 0x48, Offset: 0xfb8, Align: 0x3, Reloff: 0x0, Nreloc: 0x0, Flags: 0x6000000b, Reserved1: 0, Reserved2: 0, Reserved3: 0, Type: 0x40},
// 		{Name: "__data", Seg: "__DATA", Addr: 0x100001000, Size: 0x1c, Offset: 0x1000, Align: 0x3, Reloff: 0x0, Nreloc: 0x0, Flags: 0x0, Reserved1: 0, Reserved2: 0, Reserved3: 0, Type: 0x40},
// 		{Name: "__dyld", Seg: "__DATA", Addr: 0x100001020, Size: 0x38, Offset: 0x1020, Align: 0x3, Reloff: 0x0, Nreloc: 0x0, Flags: 0x0, Reserved1: 0x0, Reserved2: 0x0, Reserved3: 0x0, Type: 0x40},
// 		{Name: "__la_symbol_ptr", Seg: "__DATA", Addr: 0x100001058, Size: 0x10, Offset: 0x1058, Align: 0x2, Reloff: 0x0, Nreloc: 0x0, Flags: 0x7, Reserved1: 0x2, Reserved2: 0x0, Reserved3: 0x0, Type: 0x40}},
// }

// // 	nil,
// 	relocations map[string][]types.Reloc
// }

// const a = {
// 	"internal/testdata/gcc-amd64-darwin-exec.base64",
// 	types.FileHeader{Magic: 0xfeedfacf, CPU: types.CPUAmd64, SubCPU: 0x80000003, Type: 0x2, NCommands: 0xb, SizeCommands: 0x568, Flags: 0x85, Reserved: 0x0},

func main() {
	machoFile := m.File{
		// ByteOrder:
		// Loads:     []macho.Load{},
		// Sections:  []*macho.Section{},
		// Symtab:    &macho.Symtab{},
		// Dysymtab:  &macho.Dysymtab{},
	}
	machoFile.ByteOrder = binary.LittleEndian
	machoFile.FileHeader = types.FileHeader{Magic: 0xfeedfacf, CPU: types.CPUAmd64, SubCPU: 0x80000003, Type: 0x2, NCommands: 0xb, SizeCommands: 0x568, Flags: 0x85, Reserved: 0x0}

	// AddSegment contains AddLoad
	machoFile.AddSegment(&m.Segment{})

	// // Section depends on Segment
	// machoFile.AddSection(&types.Section{Relocs: []types.Reloc{}})

	machoFile.Save("macho")

	//fmt.Println(machoFile)


	//
}

func A(machoFile *m.File) {
	data, _ := base64.StdEncoding.DecodeString(Data)
	r := bytes.NewReader(data)
	f, _ := m.NewFile(r)
	if !reflect.DeepEqual(f.FileHeader, machoFile.FileHeader) {
		fmt.Println("not equal FileHeader, it is because file header depends on other parts of macho file: Segments, Section")
		// continue
	}
	fmt.Println(f.FileHeader)
	fmt.Println(machoFile.FileHeader)
}