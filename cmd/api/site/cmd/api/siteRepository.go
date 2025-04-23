package site

import (
	"path/filepath"
	"way-manager/api/shared/common/model"
)

type Repository struct {
	availableSiteFilePath string
}

func NewRepository() *Repository {
	return &Repository{availableSiteFilePath: filepath.Join("nginx", "site-available") + "site-service.conf"}
}

func (r *Repository) Add(proxyServer *model.ProxyServer) error {
	// err := os.WriteFile(r.availableSiteFilePath, []byte(), 0644)
	// if err != nil {
	// 	fmt.Errorf("failed to write Nginx config: %v", err)
	// }
	panic("")
}
