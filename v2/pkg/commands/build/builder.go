package build

import (
	"github.com/josStorer/wails/v2/internal/project"
	"github.com/josStorer/wails/v2/pkg/clilogger"
)

// Builder defines a builder that can build Wails applications
type Builder interface {
	SetProjectData(projectData *project.Project)
	BuildFrontend(logger *clilogger.CLILogger) error
	CompileProject(options *Options) error
	OutputFilename(options *Options) string
	CleanUp()
}
