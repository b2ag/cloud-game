package github

import (
	"testing"

	"github.com/giongto35/cloud-game/v2/pkg/emulator/libretro/core"
)

func TestBuildbotRepo(t *testing.T) {
	testAddress := "http://test.me"
	tests := []struct {
		file        string
		compression string
		arch        core.ArchInfo
		resultUrl   string
	}{
		{
			file: "uber_core",
			arch: core.ArchInfo{
				Os:     "linux",
				Arch:   "x86_64",
				LibExt: ".so",
			},
			resultUrl: testAddress + "/" + "linux/x86_64/latest/uber_core.so?raw=true",
		},
		{
			file:        "uber_core",
			compression: "zip",
			arch: core.ArchInfo{
				Os:     "linux",
				Arch:   "x86_64",
				LibExt: ".so",
			},
			resultUrl: testAddress + "/" + "linux/x86_64/latest/uber_core.so.zip?raw=true",
		},
		{
			file: "uber_core",
			arch: core.ArchInfo{
				Os:     "osx",
				Arch:   "x86_64",
				Vendor: "apple",
				LibExt: ".dylib",
			},
			resultUrl: testAddress + "/" + "apple/osx/x86_64/latest/uber_core.dylib?raw=true",
		},
	}

	for _, test := range tests {
		repo := NewGithubRepo(testAddress, test.compression)
		data := repo.GetCoreData(test.file, test.arch)
		if data.Url != test.resultUrl {
			t.Errorf("seems that expected link address is incorrect (%v) for file %s %+v",
				data.Url, test.file, test.arch)
		}
	}
}
