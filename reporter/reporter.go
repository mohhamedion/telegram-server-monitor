package reporter

type Reporter struct {
	layers []ReportLayer
}

type ReportLayer interface {
	Inform(message string)
}

func (reporter *Reporter) Inform(message string) {
	for _, layer := range reporter.layers {
		layer.Inform(message)
	}
}

func (reporter *Reporter) AddLayer(layer ReportLayer) {
	// reporter.layers
	reporter.layers = append(reporter.layers, layer)
}

func GetInstence() Reporter {
	return Reporter{}
}
