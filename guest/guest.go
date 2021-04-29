package guest

import (
	"fmt"
	"path/filepath"

	"gitlab.alibaba-inc.com/seadent/pkg/common"
	"gitlab.alibaba-inc.com/seadent/pkg/image/utils"
	v1 "gitlab.alibaba-inc.com/seadent/pkg/types/api/v1"
	ssh2 "gitlab.alibaba-inc.com/seadent/pkg/utils/ssh"
)

type Interface interface {
	Apply(cluster *v1.Cluster) error
	Delete(cluster *v1.Cluster) error
}

type Default struct{}

func NewGuestManager() Interface {
	return &Default{}
}

func (d *Default) Apply(cluster *v1.Cluster) error {
	ssh := ssh2.NewSSHByCluster(cluster)
	image, err := utils.GetImage(cluster.Spec.Image)
	if err != nil {
		return fmt.Errorf("get cluster image failed, %s", err)
	}
	masters := cluster.Spec.Masters.IPList
	if len(masters) == 0 {
		return fmt.Errorf("failed to found master")
	}
	clusterRootfs := filepath.Join(common.DefaultClusterRootfsDir, cluster.Name)
	for i := range image.Spec.Layers {
		if image.Spec.Layers[i].Type != common.CMDCOMMAND {
			continue
		}
		if err := ssh.CmdAsync(masters[0], fmt.Sprintf(common.CdAndExecCmd, clusterRootfs, image.Spec.Layers[i].Value)); err != nil {
			return err
		}
	}
	return nil
}

func (d Default) Delete(cluster *v1.Cluster) error {
	panic("implement me")
}