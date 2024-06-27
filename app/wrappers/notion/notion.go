package notion

import "fmt"

//TODO: Make it completely Generic
// Move DB Logic here

type NotionRequestBuilder[T any] struct {
	Req  *T
	errs []error
}

func (builder *NotionRequestBuilder[any]) AddRelation(relation **NotionPageCreateRelationProp, relationId string) {

	if *relation == nil {
		(*relation) = &NotionPageCreateRelationProp{
			Relation: make([]NotionRelation, 0),
		}
	}

	(*relation).Relation = append((*relation).Relation, NotionRelation{ID: relationId})
	fmt.Println("add rleation", (*relation).Relation)
}

func (builder *NotionRequestBuilder[any]) AddMultiSelect(multiSelect **NotionPageCreateMultiSelectProp, option string) {

	if *multiSelect == nil {
		*multiSelect = &NotionPageCreateMultiSelectProp{
			MultiSelect: make([]NotionMultiSelect, 0),
		}
	}

	(*multiSelect).MultiSelect = append((*multiSelect).MultiSelect, NotionMultiSelect{Name: &option})
}

func (builder *NotionRequestBuilder[any]) AddSelect(sel **NotionPageCreateSelectProp, option string) {

	if *sel == nil {
		*sel = &NotionPageCreateSelectProp{}
	}

	(*sel).Select = NotionSelect{
		Name: option,
	}
}

func (builder *NotionRequestBuilder[any]) AddStatus(status **NotionPageCreateStatusProp, option string) {

	if *status == nil {
		*status = &NotionPageCreateStatusProp{}
	}

	(*status).Status = NotionStatus{
		Name: option,
	}
}

func (builder *NotionRequestBuilder[any]) AddDate(sel **NotionPageCreateDateProp, date string) {

	if *sel == nil {
		*sel = &NotionPageCreateDateProp{}
	}

	(*sel).Date = NotionDatePropValue{
		Start: date,
	}
}

func (builder *NotionRequestBuilder[any]) AddTitle(name **NotionPageCreateNameProp, title string) {
	if *name == nil {
		*name = &NotionPageCreateNameProp{
			Title: make([]NotionText, 1),
		}
	}

	(*name).Title[0] = NotionText{
		Text: NotionContent{
			Content: title,
		},
	}
}

type NotionCreateTaskRequestBuilder struct {
	Builder *NotionRequestBuilder[NotionCreateTaskRequest]
}

func (nb *NotionCreateTaskRequestBuilder) Add(option string, val string) {

	if nb.Builder == nil {
		nb.Builder = &NotionRequestBuilder[NotionCreateTaskRequest]{}
	}

	if nb.Builder.Req == nil {
		nb.Builder.Req = &NotionCreateTaskRequest{}
	}

	builder := nb.Builder
	args := builder.Req

	switch option {
	case "db":
		{
			args.Parent.DatabaseID = val
		}
	case "categories":
		{
			builder.AddRelation(&args.Properties.Categories, val)
		}
	case "sub_category":
		{
			builder.AddRelation(&args.Properties.SubCategory, val)
		}
	case "status":
		{
			builder.AddStatus(&args.Properties.Status, val)
		}

	case "project":
		{
			builder.AddRelation(&args.Properties.Project, val)
		}
	case "priority":
		{
			builder.AddSelect(&args.Properties.Priority, val)
		}
	case "name":
		{
			builder.AddTitle(&args.Properties.Name, val)
		}
	case "start_date":
		{
			builder.AddDate(&args.Properties.StartDate, val)
		}
	case "tags":
		{
			builder.AddMultiSelect(&args.Properties.Tags, val)
		}
	// TODO: Add In Resource Case
	// case "resource": {

	// }
	case "default":
		{
			builder.errs = append(builder.errs, fmt.Errorf("invalid option type provided %s, only the supported types are allowed:\n 'db', 'categories', 'sub_category', 'status', 'project', 'priority', 'name', 'start_date', 'tags'", option))
		}
	}

}

func (nb *NotionCreateTaskRequestBuilder) Error() error {
	if len(nb.Builder.errs) == 0 {
		return nil
	}

	errMsg := ""

	for _, err := range nb.Builder.errs {
		if errMsg == "" {
			errMsg = err.Error()
			continue
		}
		errMsg = errMsg + ", " + err.Error()
	}

	return fmt.Errorf(errMsg)
}

func (nb *NotionCreateTaskRequestBuilder) Request() (*NotionCreateTaskRequest, error) {
	err := nb.Error()

	if err != nil {
		return nil, err
	}

	return nb.Builder.Req, err
}
