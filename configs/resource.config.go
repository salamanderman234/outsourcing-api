package configs

import "github.com/salamanderman234/outsourcing-api/app/types"

type resourceConfig struct {
	FileConfigs map[types.FileConfigKey]types.FileConfig
	BasePath    string
}

func (r *resourceConfig) setResourceConfig() {
	// file configs
	imageConfig := types.FileConfig{
		MaxSize:       2000,
		BasePath:      "/images",
		AcceptedMimes: []string{"image/jpg", "image/png", "image/jpeg", "image/webp"},
	}

	pdfConfig := types.FileConfig{
		MaxSize:       20000,
		BasePath:      "/documents",
		AcceptedMimes: []string{"application/pdf"},
	}

	r.FileConfigs[types.ImageConfig] = imageConfig
	r.FileConfigs[types.PDFConfig] = pdfConfig
}

func (r *resourceConfig) GetFileConfig(key types.FileConfigKey) types.FileConfig {
	fileConfig, ok := r.FileConfigs[key]
	if !ok {
		return types.FileConfig{
			MaxSize:       5000,
			BasePath:      "/etc",
			AcceptedMimes: []string{},
		}
	}
	return fileConfig
}

var ResourceConfig = resourceConfig{
	BasePath:    "./storage",
	FileConfigs: map[types.FileConfigKey]types.FileConfig{},
}
