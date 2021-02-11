package decisiontree

import (
	"sugar-level-client/models"
)

type Tree struct {
	nodes          map[SugarClassification]*Tree
	name           string
	classification SugarClassification
	lowest         float32
	highest        float32
}

type SugarClassification int

const (
	Unknown SugarClassification    = 0
	Low                            = 1
	Normal                         = 2
	High                           = 3
)

func (c SugarClassification) String() string {
	switch c {
	case Unknown:
		return "Unknown"
	case Low:
		return "Low"
	case Normal:
		return "Normal"
	case High:
		return "High"
	default:
		return "Unknown"
	}
}

var _tree *Tree

func Init() {
	var morning = Tree{name: "morning", lowest: 0.50, highest: 0.78}
	morning.nodes = make(map[SugarClassification]*Tree)
	morning.nodes[Low] = &Tree{
		name:    "low-afternoon",
		lowest:  0.15,
		highest: 0.24,
		nodes: map[SugarClassification]*Tree{
			Low: {name: "low-evening", classification: Low},
			Normal: {
				name:    "normal-evening",
				lowest:  0.12,
				highest: 0.21,
				nodes: map[SugarClassification]*Tree{
					Low:    {name: "result-low", classification: Low},
					Normal: {name: "result-normal", classification: Normal},
					High:   {name: "result-high", classification: Normal},
				},
			},
			High: {
				name:    "high-evening",
				lowest:  0.12,
				highest: 0.21,
				nodes: map[SugarClassification]*Tree{
					Low:    {name: "result-low", classification: Low},
					Normal: {name: "result-normal", classification: Normal},
					High:   {name: "result-high", classification: High},
				},
			},
		},
	}
	morning.nodes[Normal] = &Tree{
		name:    "normal-afternoon",
		lowest:  0.15,
		highest: 0.24,
		nodes: map[SugarClassification]*Tree{
			Low: {
				name:    "low-evening",
				lowest:  0.12,
				highest: 0.21,
				nodes: map[SugarClassification]*Tree{
					Low:    {name: "result-low", classification: Low},
					Normal: {name: "result-normal", classification: Normal},
					High:   {name: "result-high", classification: Normal},
				},
			},
			Normal: {name: "evening-normal", classification: Normal},
			High: {
				name:    "high-evening",
				lowest:  0.12,
				highest: 0.21,
				nodes: map[SugarClassification]*Tree{
					Low:    {name: "result-low", classification: Normal},
					Normal: {name: "result-normal", classification: Normal},
					High:   {name: "result-high", classification: High},
				},
			},
		},
	}
	morning.nodes[High] = &Tree{
		name:    "high-afternoon",
		lowest:  0.15,
		highest: 0.24,
		nodes: map[SugarClassification]*Tree{
			Low: {
				name:    "evening-low",
				lowest:  0.12,
				highest: 0.21,
				nodes: map[SugarClassification]*Tree{
					Low:    {name: "result-low", classification: Low},
					Normal: {name: "result-normal", classification: Normal},
					High:   {name: "result-high", classification: High},
				},
			},
			Normal: {
				name:    "evening-normal",
				lowest:  0.12,
				highest: 0.21,
				nodes: map[SugarClassification]*Tree{
					Low:    {name: "result-low", classification: Normal},
					Normal: {name: "result-normal", classification: Normal},
					High:   {name: "result-high", classification: High},
				},
			},
			High: {name: "evening-high", classification: High},
		},
	}

	_tree = &morning
}

func CheckUserLevel(user *models.User) SugarClassification {
	return classifyUser(&user.Samples, _tree)
}

func classifyUser(samples *[]models.Sample, head *Tree) SugarClassification {
	var node = head
	var result SugarClassification

	for _, sample := range *samples {
		classification := node.classify(sample.Value)
		_, ok := node.nodes[classification]
		if ok == true {
			node = node.nodes[classification]
			if node.classification != Unknown {
				result = node.classification
				break
			}
		}
	}

	return result
}

func (node *Tree) classify(sample float32) SugarClassification {
	if sample < node.lowest {
		return Low
	} else if sample >= node.lowest && sample <= node.highest {
		return Normal
	} else { //then sample is great than node.highest
		return High
	}
}
