package gojenkins

type Queue struct {
	Items []Item `json:"items"`
}

type Item struct {
	Actions                    []Action `json:"actions"`
	Blocked                    bool     `json:"blocked"`
	Buildable                  bool     `json:"buildable"`
	Id                         int      `json:"id"`
	InQueueSince               int64    `json:"inQueueSince"`
	Params                     string   `json:"params"`
	Stuck                      bool     `json:"stuck"`
	Task                       Task     `json:"task"`
	URL                        string   `json:"url"`
	Why                        string   `json:"why"`
	BuildableStartMilliseconds int64    `json:"buildableStartMilliseconds"`
	Pending                    bool     `json:"pending"`
}

type Action struct {
	Causes               []Cause               `json:"causes"`
	ParameterDefinitions []ParameterDefinition `json:"parameterDefinitions"`
	Parameters           []Parameter           `json:"parameters"`
	LastBuiltRevision    Revision              `json:"lastBuiltRevision"`
	RemoteUrls           []string              `json:"remoteUrls"`
}

type Cause struct {
	ShortDescription string `json:"shortDescription"`
	UserId           string `json:"userId"`
	UserName         string `json:"userName"`
	UpstreamCause
}

type ParameterDefinition struct {
	Name string `json:"name"`
}

type Parameter struct {
	Class string `json:"_class"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Revision struct {
	SHA1   string `json:"SHA1"`
	Branches []Branch `json:"branch"`
}

type Branch struct {
	SHA1 string `json:"SHA1"`
	Name string `json:"name"`
}

type Task struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}
