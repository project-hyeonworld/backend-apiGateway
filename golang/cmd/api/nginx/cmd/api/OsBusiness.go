package nginx

import (
	"fmt"
	"os/exec"
)

type OsBusiness struct {
	application map[string]string
}

func NewOsBusiness() OsBusiness {
	return OsBusiness{}
}

func (b *OsBusiness) RestartNginx() error {
	// containerNmae, err := b.getContainerName(80)
	// if err != nil {
	// 	return err
	// }
	cmd := exec.Command("nginx", "-s", "reload")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to reload Nginx: %v, output: %s", err, output)
	}
	return nil
}

func (b *OsBusiness) getContainerName(port uint) (string, error) {
	cmd := exec.Command("docker", "ps", "-f", fmt.Sprintf("publish=%d", port), "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get container name: %v", err)
	}
	return string(output), nil
}
