package configuration_secret

type Value struct {
	session string
}

func (v *Value) Init() {
	v.session = "{MY_SESSION_APPLICATION_NAME}"
}
