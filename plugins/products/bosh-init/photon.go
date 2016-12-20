package boshinit

import (
	"fmt"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/plugins/products/bosh-init/enaml-gen/blobstore"
	"github.com/enaml-ops/omg-cli/plugins/products/bosh-init/enaml-gen/photoncpi"
)

const (
	PhotonCPIReleaseName     = "bosh-photon-cpi"
	PhotonCPIJobName         = "cpi"
	PhotonPersistentDiskSize = 32768

	PhotonCPIURL = "https://s3.amazonaws.com/concourse-photon/bosh-photon-cpi-1.0.0.tgz"
	PhotonCPISHA = "71626961a8505447fa34ca569f97f8c70a0ef39a"

	PhotonBoshReleaseURL = "https://bosh.io/d/github.com/cloudfoundry/bosh?v=257.3"
	PhotonBoshReleaseSHA = "e4442afcc64123e11f2b33cc2be799a0b59207d0"

	PhotonStemcellURL = "https://bosh.io/d/stemcells/bosh-vsphere-esxi-ubuntu-trusty-go_agent?v=3232.1"
	PhotonStemcellSHA = "169df93e3e344cd84ac6ef16d76dd0276e321a25"
)

type PhotonBoshInitConfig struct {
	photoncpi.Photon
	MachineType string
	NetworkName string
}

type PhotonBosh struct {
	BoshInitConfig *PhotonBoshInitConfig
	Base           *BoshBase
}

// NewPhotonBoshBase creates a BoshBase initialized with Photon defaults.
func NewPhotonBoshBase() *BoshBase {
	return &BoshBase{
		CPIJobName:         PhotonCPIJobName,
		CPIReleaseURL:      PhotonCPIURL,
		CPIReleaseSHA:      PhotonCPISHA,
		BoshReleaseURL:     PhotonBoshReleaseURL,
		BoshReleaseSHA:     PhotonBoshReleaseSHA,
		GOAgentReleaseURL:  PhotonStemcellURL,
		GOAgentSHA:         PhotonStemcellSHA,
		PersistentDiskSize: PhotonPersistentDiskSize,
	}
}

func NewPhotonIaaSProvider(cfg *PhotonBoshInitConfig, boshBase *BoshBase) IAASManifestProvider {
	boshBase.CPIJobName = PhotonCPIJobName
	return &PhotonBosh{
		BoshInitConfig: cfg,
		Base:           boshBase,
	}
}

func (g *PhotonBosh) CreateCPIRelease() enaml.Release {
	return enaml.Release{
		Name: PhotonCPIReleaseName,
		URL:  g.Base.CPIReleaseURL,
		SHA1: g.Base.CPIReleaseSHA,
	}
}

func (g *PhotonBosh) CreateCPITemplate() enaml.Template {
	return enaml.Template{
		Name:    PhotonCPIJobName,
		Release: PhotonCPIReleaseName,
	}
}

func (g *PhotonBosh) CreateDiskPool() enaml.DiskPool {
	return enaml.DiskPool{
		Name:     "disks",
		DiskSize: g.Base.PersistentDiskSize,
		CloudProperties: map[string]interface{}{
			"disk_flavor": "core-200",
		},
	}
}

func (g *PhotonBosh) resourcePoolCloudProperties() interface{} {
	return map[string]interface{}{
		"vm_flavor":   g.BoshInitConfig.MachineType,
		"disk_flavor": "core-200",
	}
}

func (g *PhotonBosh) CreateResourcePool() (*enaml.ResourcePool, error) {
	return g.Base.CreateResourcePool(g.resourcePoolCloudProperties)
}

func (g *PhotonBosh) CreateManualNetwork() enaml.ManualNetwork {
	net := enaml.NewManualNetwork("private")
	net.AddSubnet(enaml.Subnet{
		Range:   g.Base.NetworkCIDR,
		Gateway: g.Base.NetworkGateway,
		DNS:     g.Base.NetworkDNS,
		Static: []string{
			g.Base.PrivateIP,
		},
		CloudProperties: map[string]interface{}{
			"network_id": g.BoshInitConfig.NetworkName,
		},
	})
	return net
}

func (g *PhotonBosh) CreateVIPNetwork() enaml.VIPNetwork {
	return enaml.VIPNetwork{
		Name: "vip",
		Type: "vip",
	}
}

func (g *PhotonBosh) CreateJobNetwork() *enaml.Network {
	// photon just needs the default private network provided by boshbase
	return nil
}

func (g *PhotonBosh) CreateCloudProvider() enaml.CloudProvider {
	return enaml.CloudProvider{
		Template: g.CreateCPITemplate(),
		MBus:     fmt.Sprintf("https://mbus:%s@%s:6868", g.Base.MBusPassword, g.Base.PrivateIP),
		Properties: &photoncpi.PhotoncpiJob{
			Photon: &g.BoshInitConfig.Photon,
			Agent: &photoncpi.Agent{
				Mbus: fmt.Sprintf("https://mbus:%s@0.0.0.0:6868", g.Base.MBusPassword),
			},
			Blobstore: &photoncpi.Blobstore{
				Provider: "local",
				Options: map[string]interface{}{
					"blobstore_path": "/var/vcap/micro_bosh/data/cache",
				},
			},
			Ntp: g.createNTP(),
		},
	}
}

func (g *PhotonBosh) createNTP() interface{} {
	return g.Base.NtpServers
}

func (g *PhotonBosh) CreateCPIJobProperties() map[string]interface{} {
	return map[string]interface{}{
		"photon": g.BoshInitConfig.Photon,
		"agent": &photoncpi.Agent{
			Mbus: fmt.Sprintf("nats://nats:%s@%s:4222", g.Base.NatsPassword, g.Base.PrivateIP),
		},
	}
}

func (g *PhotonBosh) CreateDeploymentManifest() (*enaml.DeploymentManifest, error) {
	var manifest = g.Base.CreateDeploymentManifest()
	manifest.AddRelease(g.CreateCPIRelease())
	if rp, err := g.CreateResourcePool(); err != nil {
		return nil, err
	} else {
		manifest.AddResourcePool(*rp)
	}
	manifest.AddDiskPool(g.CreateDiskPool())
	manifest.AddNetwork(g.CreateManualNetwork())
	manifest.AddNetwork(g.CreateVIPNetwork())
	boshJob := manifest.Jobs[0]
	boshJob.AddTemplate(g.CreateCPITemplate())
	n := g.CreateJobNetwork()
	if n != nil {
		boshJob.AddNetwork(*n)
	}
	for name, val := range g.CreateCPIJobProperties() {
		boshJob.AddProperty(name, val)
	}
	boshJob.AddProperty("blobstore", g.getJobPropertyBlobstore())
	manifest.Jobs[0] = boshJob
	manifest.SetCloudProvider(g.CreateCloudProvider())
	return manifest, nil
}

func (g *PhotonBosh) getJobPropertyBlobstore() map[string]interface{} {
	return map[string]interface{}{
		"address": g.Base.PrivateIP,
		"port":    25250,
		"agent": blobstore.Agent{
			User:     "agent",
			Password: g.Base.NatsPassword,
		},
		"director": blobstore.Director{
			User:     "director",
			Password: g.Base.DirectorPassword,
		},
		"provider": "dav",
		"options": map[string]interface{}{
			"endpoint": fmt.Sprintf("http://%v:25250", g.Base.PrivateIP),
			"user":     "agent",
			"password": g.Base.NatsPassword,
		},
	}
}
