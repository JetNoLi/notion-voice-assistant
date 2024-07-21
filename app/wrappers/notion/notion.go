package notion

import "fmt"

//TODO: Make it completely Generic
// Move DB Logic here

type RequestBuilder[T any] struct {
	Req  *T
	errs []error
}

func (builder *RequestBuilder[any]) AddRelation(relation **PageCreateRelationProp, relationId string) {

	if *relation == nil {
		(*relation) = &PageCreateRelationProp{
			Relation: make([]Relation, 0),
		}
	}

	(*relation).Relation = append((*relation).Relation, Relation{ID: relationId})
}

func (builder *RequestBuilder[any]) AddMultiSelect(multiSelect **PageCreateMultiSelectProp, option string) {

	if *multiSelect == nil {
		*multiSelect = &PageCreateMultiSelectProp{
			MultiSelect: make([]MultiSelect, 0),
		}
	}

	(*multiSelect).MultiSelect = append((*multiSelect).MultiSelect, MultiSelect{Name: &option})
}

func (builder *RequestBuilder[any]) AddSelect(sel **PageCreateSelectProp, option string) {

	if *sel == nil {
		*sel = &PageCreateSelectProp{}
	}

	(*sel).Select = Select{
		Name: option,
	}
}

func (builder *RequestBuilder[any]) AddStatus(status **PageCreateStatusProp, option string) {

	if *status == nil {
		*status = &PageCreateStatusProp{}
	}

	(*status).Status = Status{
		Name: option,
	}
}

func (builder *RequestBuilder[any]) AddDate(sel **PageCreateDateProp, date string) {

	if *sel == nil {
		*sel = &PageCreateDateProp{}
	}

	(*sel).Date = DatePropValue{
		Start: date,
	}
}

func (builder *RequestBuilder[any]) AddTitle(name **PageCreateNameProp, title string) {
	if *name == nil {
		*name = &PageCreateNameProp{
			Title: make([]Text, 1),
		}
	}

	(*name).Title[0] = Text{
		Text: Content{
			Content: title,
		},
	}
}

type CreateTaskRequestBuilder struct {
	Builder *RequestBuilder[CreateTaskRequest]
}

func (nb *CreateTaskRequestBuilder) Add(option string, val string) {

	if nb.Builder == nil {
		nb.Builder = &RequestBuilder[CreateTaskRequest]{}
	}

	if nb.Builder.Req == nil {
		nb.Builder.Req = &CreateTaskRequest{}
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
			fmt.Println("error in builder", option, val)
			builder.errs = append(builder.errs, fmt.Errorf("invalid option type provided %s, only the supported types are allowed:\n 'db', 'categories', 'sub_category', 'status', 'project', 'priority', 'name', 'start_date', 'tags'", option))
		}
	}

}

func (nb *CreateTaskRequestBuilder) Error() error {
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

func (nb *CreateTaskRequestBuilder) Request() (*CreateTaskRequest, error) {
	err := nb.Error()

	if err != nil {
		return nil, err
	}

	return nb.Builder.Req, err
}
