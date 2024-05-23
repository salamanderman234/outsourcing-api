package configs

type fileConfig struct {
	MaxSize       uint
	BasePath      string
	AcceptedMimes []string
}

type Resource struct {
	Config fileConfig
	Path   string
}

type resourceConfig struct {
	resourcesConfig map[string]Resource
}

var ResourceConfig resourceConfig

func (r *resourceConfig) setResourceConfig() {
	// file configs
	imageConfig := fileConfig{
		MaxSize:       2000,
		BasePath:      "/images",
		AcceptedMimes: []string{"jpg", "png", "jpeg", "webp"},
	}

	r.resourcesConfig = map[string]Resource{
		"category.icon": {
			Config: imageConfig,
			Path:   "/category",
		},
		"transaction.mou": {
			Config: imageConfig,
			Path:   "/transaction/mou",
		},
	}
}

func (r *resourceConfig) GetResourceConfig(key string) Resource {
	resource, ok := r.resourcesConfig[key]
	if !ok {
		return Resource{
			Config: fileConfig{
				MaxSize:       5000,
				BasePath:      "/etc",
				AcceptedMimes: []string{"*"},
			},
			Path: "",
		}
	}
	return resource
}
