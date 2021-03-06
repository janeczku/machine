package egoscale

// ListResourceDetails lists the resource tag(s) (but different from listTags...)
type ListResourceDetails struct {
	ResourceType string `json:"resourcetype" doc:"list by resource type"`
	ForDisplay   bool   `json:"fordisplay,omitempty" doc:"if set to true, only details marked with display=true, are returned. False by default"`
	Key          string `json:"key,omitempty" doc:"list by key"`
	Keyword      string `json:"keyword,omitempty" doc:"List by keyword"`
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"pagesize,omitempty"`
	ResourceID   *UUID  `json:"resourceid,omitempty" doc:"list by resource id"`
	Value        string `json:"value,omitempty" doc:"list by key, value. Needs to be passed only along with key"`
	_            bool   `name:"listResourceDetails" description:"List resource detail(s)"`
}

// ListResourceDetailsResponse represents a list of resource details
type ListResourceDetailsResponse struct {
	Count          int           `json:"count"`
	ResourceDetail []ResourceTag `json:"resourcedetail"`
}

func (*ListResourceDetails) response() interface{} {
	return new(ListResourceDetailsResponse)
}
