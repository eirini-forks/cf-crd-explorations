package handlers

type CFAPIAppResource struct {
	GUID          string                  `json:"guid"`
	Name          string                  `json:"name"`
	State         string                  `json:"state"`
	CreatedAt     string                  `json:"created_at"`
	UpdatedAt     string                  `json:"updated_at"`
	Lifecycle     CFAPIAppLifecycle       `json:"lifecycle,omitempty"`
	Relationships CFAPIAppRelationships   `json:"relationships"`
	Links         map[string]CFAPIAppLink `json:"links"`
	Metadata      CFAPIMetadata           `json:"metadata"`
}

type CFAPIAppResourceWithEnvVars struct {
	CFAPIAppResource
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
}

type CFAPIAppLifecycle struct {
	Type string                `json:"type"`
	Data CFAPIAppLifecycleData `json:"data"`
}

type CFAPIAppLifecycleData struct {
	Buildpacks []string `json:"buildpacks,omitempty"`
	Stack      string   `json:"stack,omitempty"`
}

type CFAPIAppRelationships struct {
	Space CFAPIAppRelationshipsSpace `json:"space"`
}

type CFAPIAppRelationshipsSpace struct {
	Data CFAPIAppRelationshipsSpaceData `json:"data"`
}

type CFAPIAppRelationshipsSpaceData struct {
	GUID string `json:"guid"`
}

type CFAPIAppLink struct {
	Href   string `json:"href"`
	Method string `json:"method,omitempty"`
}

type CFAPIMetadata struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type CFAPIPackageResource struct {
	Type          string                       `json:"type"`
	Relationships CFAPIPackageAppRelationships `json:"relationships"`
	Data          CFAPIPackageData             `json:"data"`
}

type CFAPIPackageAppRelationships struct {
	App CFAPIPackageAppRelationshipsApp `json:"app"`
}

type CFAPIPackageAppRelationshipsApp struct {
	Data CFAPIPackageAppRelationshipsAppData `json:"data"`
}

type CFAPIPackageAppRelationshipsAppData struct {
	GUID string `json:"guid"`
}

type CFAPIPackageData struct {
	Image    string `json:"image"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CFAPIErrorData struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Code   string `json:"code"`
}
