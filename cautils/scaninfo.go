package cautils

import (
	"path/filepath"

	"github.com/armosec/kubescape/cautils/getter"
	"github.com/armosec/kubescape/cautils/opapolicy"
)

type ScanInfo struct {
	PolicyGetter       getter.IPolicyGetter
	PolicyIdentifier   opapolicy.PolicyIdentifier
	Format             string
	Output             string
	ExcludedNamespaces string
	InputPatterns      []string
	Silent             bool
}

func (scanInfo *ScanInfo) Init() {
	// scanInfo.setSilentMode()
	scanInfo.setOutputFile()
	scanInfo.setGetter()

}
func (scanInfo *ScanInfo) setGetter() {
	scanInfo.PolicyGetter = getter.NewArmoAPI()
}
func (scanInfo *ScanInfo) setSilentMode() {
	if scanInfo.Format == "json" || scanInfo.Format == "junit" {
		scanInfo.Silent = true
	}
	if scanInfo.Output != "" {
		scanInfo.Silent = true
	}
}

func (scanInfo *ScanInfo) setOutputFile() {
	if scanInfo.Output == "" {
		return
	}
	if scanInfo.Format == "json" {
		if filepath.Ext(scanInfo.Output) != "json" {
			scanInfo.Output += ".json"
		}
	}
	if scanInfo.Format == "junit" {
		if filepath.Ext(scanInfo.Output) != "xml" {
			scanInfo.Output += ".xml"
		}
	}
}