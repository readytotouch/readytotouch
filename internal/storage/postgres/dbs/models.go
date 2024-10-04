// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbs

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type FeatureWait string

const (
	FeatureWaitOrganizerGolangCompanies  FeatureWait = "organizer_golang_companies"
	FeatureWaitOrganizerGolangVacancies  FeatureWait = "organizer_golang_vacancies"
	FeatureWaitOrganizerRustCompanies    FeatureWait = "organizer_rust_companies"
	FeatureWaitOrganizerRustVacancies    FeatureWait = "organizer_rust_vacancies"
	FeatureWaitOrganizerZigCompanies     FeatureWait = "organizer_zig_companies"
	FeatureWaitOrganizerZigVacancies     FeatureWait = "organizer_zig_vacancies"
	FeatureWaitOrganizerScalaCompanies   FeatureWait = "organizer_scala_companies"
	FeatureWaitOrganizerScalaVacancies   FeatureWait = "organizer_scala_vacancies"
	FeatureWaitOrganizerElixirCompanies  FeatureWait = "organizer_elixir_companies"
	FeatureWaitOrganizerElixirVacancies  FeatureWait = "organizer_elixir_vacancies"
	FeatureWaitOrganizerClojureCompanies FeatureWait = "organizer_clojure_companies"
	FeatureWaitOrganizerClojureVacancies FeatureWait = "organizer_clojure_vacancies"
)

func (e *FeatureWait) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FeatureWait(s)
	case string:
		*e = FeatureWait(s)
	default:
		return fmt.Errorf("unsupported scan type for FeatureWait: %T", src)
	}
	return nil
}

type NullFeatureWait struct {
	FeatureWait FeatureWait
	Valid       bool // Valid is true if FeatureWait is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFeatureWait) Scan(value interface{}) error {
	if value == nil {
		ns.FeatureWait, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FeatureWait.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFeatureWait) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FeatureWait), nil
}

type SocialProvider string

const (
	SocialProviderGithub    SocialProvider = "github"
	SocialProviderGitlab    SocialProvider = "gitlab"
	SocialProviderBitbucket SocialProvider = "bitbucket"
	SocialProviderFigma     SocialProvider = "figma"
	SocialProviderDribbble  SocialProvider = "dribbble"
	SocialProviderBehance   SocialProvider = "behance"
	SocialProviderLinkedin  SocialProvider = "linkedin"
	SocialProviderXing      SocialProvider = "xing"
	SocialProviderGoogle    SocialProvider = "google"
)

func (e *SocialProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SocialProvider(s)
	case string:
		*e = SocialProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for SocialProvider: %T", src)
	}
	return nil
}

type NullSocialProvider struct {
	SocialProvider SocialProvider
	Valid          bool // Valid is true if SocialProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSocialProvider) Scan(value interface{}) error {
	if value == nil {
		ns.SocialProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SocialProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSocialProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.SocialProvider), nil
}

type Company struct {
	ID        int64
	CreatedAt time.Time
	CreatedBy int64
}

type CompanyViewDailyStat struct {
	CompanyID int64
	CreatedAt time.Time
	ViewCount int64
}

type FeatureViewDailyStat struct {
	Feature   FeatureWait
	CreatedAt time.Time
	ViewCount int64
}

type FeatureViewStat struct {
	Feature   FeatureWait
	ViewCount int64
}

type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type UserFavoriteCompany struct {
	UserID    int64
	CompanyID int64
	Favorite  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFeatureWaitlist struct {
	UserID    int64
	Feature   FeatureWait
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFeatureWaitlistStat struct {
	Feature   FeatureWait
	UserCount int64
}

type UserOnlineDailyCountStat struct {
	CreatedAt time.Time
	UserCount int64
}

type UserOnlineDailyStat struct {
	CreatedAt time.Time
	UserID    int64
}

type UserOnlineHourlyStat struct {
	UserID    int64
	CreatedAt time.Time
}

type UserSocialProfile struct {
	ID                   int64
	UserID               int64
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Email                string
	Username             string
	Name                 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            sql.NullTime
}

type UserSocialProfileChangeHistory struct {
	ID                  int64
	UserID              int64
	UserSocialProfileID int64
	Email               string
	Username            string
	Name                string
	CreatedAt           time.Time
	DeletedAt           sql.NullTime
}

type WipLinkedinCompany struct {
	ID         int64
	VanityName string
	Name       string
	CreatedAt  time.Time
	CreatedBy  int64
}

type WipLinkedinCompanyRequestHistory struct {
	ID                int64
	VanityName        string
	LinkedinCompanyID sql.NullInt64
	ResponsePayload   json.RawMessage
	CreatedAt         time.Time
	CreatedBy         int64
}

type WipUserToLinkedinCompany struct {
	UserID            int64
	LinkedinCompanyID int64
	Active            bool
	CreatedAt         time.Time
	CreatedBy         int64
	UpdatedAt         time.Time
	UpdatedBy         int64
}
