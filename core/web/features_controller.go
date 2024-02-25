package web

import (
	"github.com/gin-gonic/gin"

	"github.com/DeAI-Artist/MintAI/core/services/chainlink"
	"github.com/DeAI-Artist/MintAI/core/web/presenters"
)

// FeaturesController manages the feature flags
type FeaturesController struct {
	App chainlink.Application
}

const (
	FeatureKeyCSA          string = "csa"
	FeatureKeyFeedsManager string = "feeds_manager"
)

// Index retrieves the features
// Example:
// "GET <application>/features"
func (fc *FeaturesController) Index(c *gin.Context) {
	resources := []presenters.FeatureResource{
		*presenters.NewFeatureResource(FeatureKeyCSA, fc.App.GetConfig().Feature().UICSAKeys()),
		*presenters.NewFeatureResource(FeatureKeyFeedsManager, fc.App.GetConfig().Feature().FeedsManager()),
	}

	jsonAPIResponse(c, resources, "features")
}
