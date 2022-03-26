package option

type Setting struct {
	JSON             string
	EnvName          string
	Scopes           []string
	ValueInputOption string
}

type Default = Setting

type Option interface {
	apply(*Setting)
}

func Apply(options []Option, settings ...*Setting) *Setting {
	var setting *Setting
	if len(settings) > 0 {
		setting = settings[len(settings)-1]
	} else {
		setting = &Setting{}
	}

	for _, option := range options {
		option.apply(setting)
	}

	return setting
}

type FromJSON string

func (json FromJSON) apply(opt *Setting) {
	opt.JSON = string(json)
}

type FromEnv string

func (name FromEnv) apply(opt *Setting) {
	opt.EnvName = string(name)
}

type withScopes []string

func (scopes withScopes) apply(opt *Setting) {
	opt.Scopes = append(opt.Scopes, scopes...)
}

func WithScopes(scopes ...string) Option { //nolint:ireturn
	return withScopes(scopes)
}

type ValueInputOption string

func (vio ValueInputOption) apply(opt *Setting) {
	opt.ValueInputOption = string(vio)
}
