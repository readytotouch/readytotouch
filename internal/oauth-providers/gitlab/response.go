package gitlab

type UserResponse struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	//State                          string        `json:"state"`
	//AvatarURL                      string        `json:"avatar_url"`
	//WebURL                         string        `json:"web_url"`
	//CreatedAt                      time.Time     `json:"created_at"`
	//Bio                            string        `json:"bio"`
	//BioHTML                        string        `json:"bio_html"`
	//Location                       string        `json:"location"`
	//PublicEmail                    string        `json:"public_email"`
	//Skype                          string        `json:"skype"`
	//Linkedin                       string        `json:"linkedin"`
	//Twitter                        string        `json:"twitter"`
	//WebsiteURL                     string        `json:"website_url"`
	//Organization                   string        `json:"organization"`
	//JobTitle                       string        `json:"job_title"`
	//WorkInformation                interface{}   `json:"work_information"`
	//LastSignInAt                   time.Time     `json:"last_sign_in_at"`
	//ConfirmedAt                    time.Time     `json:"confirmed_at"`
	//LastActivityOn                 string        `json:"last_activity_on"`
	//ThemeID                        int           `json:"theme_id"`
	//ColorSchemeID                  int           `json:"color_scheme_id"`
	//ProjectsLimit                  int           `json:"projects_limit"`
	//CurrentSignInAt                time.Time     `json:"current_sign_in_at"`
	//Identities                     []interface{} `json:"identities"`
	//CanCreateGroup                 bool          `json:"can_create_group"`
	//CanCreateProject               bool          `json:"can_create_project"`
	//TwoFactorEnabled               bool          `json:"two_factor_enabled"`
	//External                       bool          `json:"external"`
	//PrivateProfile                 bool          `json:"private_profile"`
	//SharedRunnersMinutesLimit      interface{}   `json:"shared_runners_minutes_limit"`
	//ExtraSharedRunnersMinutesLimit interface{}   `json:"extra_shared_runners_minutes_limit"`
}
