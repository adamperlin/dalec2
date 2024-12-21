package azlinux

import (
	"github.com/Azure/dalec/frontend/rpm/distro"
)

const (
	Mariner2TargetKey     = "mariner2"
	tdnfCacheNameMariner2 = "mariner2-tdnf-cache"

	Mariner2Ref               = "mcr.microsoft.com/cbl-mariner/base/core:2.0"
	Mariner2FullName          = "CBL-Mariner 2"
	Mariner2WorkerContextName = "dalec-mariner2-worker"
)

var Mariner2Config = &distro.Config{
	ImageRef:   "mcr.microsoft.com/cbl-mariner/base/core:2.0",
	ContextRef: Mariner2WorkerContextName,

	ReleaseVer:         "2.0",
	BuilderPackages:    builderPackages,
	BasePackages:       []string{"distroless-packages-minimal", "prebuilt-ca-certificates"},
	RepoPlatformConfig: &defaultAzlinuxRepoPlatform,
	InstallFunc:        distro.TdnfInstall,
}

// func (mariner2) tdnfCacheMount(root string) llb.RunOption {
// 	return llb.AddMount(filepath.Join(root, tdnfCacheDir), llb.Scratch(), llb.AsPersistentCacheDir(tdnfCacheNameMariner2, llb.CacheMountLocked))
// }
