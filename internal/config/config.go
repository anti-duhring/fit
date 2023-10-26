package config

type OptFunc func(*Opts)

type Opts struct {
	InitFolder string
}

func defaultOpts() Opts {
	return Opts{
		InitFolder: ".fit",
	}
}

func Init(opts []OptFunc) Opts {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	return o
}

func WithInitFolder(folder string) OptFunc {
	return func(o *Opts) {
		o.InitFolder = folder
	}
}
