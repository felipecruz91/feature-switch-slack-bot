package main

import (
	"fmt"
	"log"
	"regexp"
)

// ParseExperimentName parses the experiment name from the Slack text message
func ParseExperimentName(textMessage string) (experimentName string, err error) {
	var result string
	r, err := regexp.Compile(`check (.*)`)
	match := r.MatchString(textMessage)

	if match {
		index := r.FindStringIndex(textMessage)
		result = textMessage[index[0]+6 : index[1]]
		return result, nil
	}

	parseError := fmt.Errorf("Unable to parse the Experiment name from '%s'", textMessage)
	log.Print(err)
	return result, parseError
}
