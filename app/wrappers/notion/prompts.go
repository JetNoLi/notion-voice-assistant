package notion

import (
	"fmt"

	"github.com/jetnoli/notion-voice-assistant/utils"
)

const IntroPrompt = `I am creating a notion database with the following fields ; `

// $FIELDS

const OptionsPrompt = `for all the fields which require a select or multi select here are my options ; `

// $OPTIONS
const DBRelationsPrompt = `for all the db relations, the key is the name of the field and then its type ; here are my relations`

// $DB_RELATION
const StruturePrompt = `; suggest a selection for each and possibly multiple for multi select options. If no property appears to match properly, omit the key from the response. Here is the prompt - `

// $PROMPT
const ReturnPrompt = ` ; To form the response create a json object, with 3 primary keys, a name for the task, options and relations. For options please return the name of the selected option or in the case of multi selects a string array of the selected and for relations please return the Name and PageID as well as the DBID, return relations as arrays,  with a max of 2 items, if there are none omit the key, ; return only json in the response as it will be parsed to json for another request. Prioritize returning valid json within the max token count, fitting this format, remember omit enitrely if the key has no value - `

type CreatePromptArgs struct {
	Prompt       string
	DBFields     any
	ReturnStruct any
	DBRelations  any
	Options      any
}

// prompt string, fields string, dbRelations string, options string
func CreatePrompt(args CreatePromptArgs) string {

	fields := utils.PrintStructType(args.DBFields)
	options := fmt.Sprintf("%#v", args.Options)
	dbRelations := fmt.Sprintf("%#v", args.DBRelations)
	returnStruct := utils.PrintStructType(args.ReturnStruct)

	prompt := IntroPrompt + fields + " " + OptionsPrompt + " OPTIONS: " + options + " " + DBRelationsPrompt + " DB RELATIONS: " + dbRelations + StruturePrompt + args.Prompt + ReturnPrompt + returnStruct

	return prompt
}
