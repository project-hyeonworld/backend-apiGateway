package nginx

type Controller struct {
	svc IService
}

func NewController(svc IService) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Restart() error {
	if err := c.svc.Restart(); err != nil {
		return err
	}
	return nil
}
