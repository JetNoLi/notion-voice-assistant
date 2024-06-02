package serve

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jetnoli/notion-voice-assistant/utils"
)

func Html(path string) (htmlData []byte, err error) {
	absPath, err := filepath.Abs(path)

	if err != nil {
		return htmlData, err
	}

	htmlData, err = os.ReadFile(absPath)

	return htmlData, err
}

// __keyname__ in html gets replaced with value in data map -> [keyname]: value
// if any value in data, not injected or any value in html, not injected, will throw error
func AndInjectHtml(path string, data map[string]string) (htmlData []byte, err error) {
	rawHtmlData, err := Html(path)

	if err != nil {
		return htmlData, err
	}

	dataStr := string(rawHtmlData)

	numDataKeys := len(data)
	numHtmlKeys := utils.GetStringOccurrences("__[^_]+__", dataStr)

	replacedKeyCount := 0

	for key, htmlValue := range data {
		htmlKey := "__" + key + "__"
		tempPrevStr := dataStr

		dataStr = strings.ReplaceAll(dataStr, htmlKey, htmlValue)

		if dataStr != tempPrevStr {
			replacedKeyCount += 1
		}
	}

	if replacedKeyCount != numDataKeys {
		return htmlData, fmt.Errorf(`number of keys to replace in data: %d, does not match keys replaced: %d`, numDataKeys, replacedKeyCount)
	}

	if replacedKeyCount != numHtmlKeys {
		return htmlData, fmt.Errorf(`number of keys to replace in html: %d, does not match keys replaced: %d`, numHtmlKeys, replacedKeyCount)
	}

	return []byte(dataStr), nil
}
