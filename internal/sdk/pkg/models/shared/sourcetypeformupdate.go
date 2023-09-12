// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType string

const (
	SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthTypeAccessToken SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType = "access_token"
)

func (e SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType) ToPointer() *SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType {
	return &e
}

func (e *SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType: %v", v)
	}
}

type SourceTypeformUpdateAuthorizationMethodPrivateToken struct {
	// Log into your Typeform account and then generate a personal Access Token.
	AccessToken string                                                       `json:"access_token"`
	AuthType    *SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType `json:"auth_type,omitempty"`
}

type SourceTypeformUpdateAuthorizationMethodOAuth20AuthType string

const (
	SourceTypeformUpdateAuthorizationMethodOAuth20AuthTypeOauth20 SourceTypeformUpdateAuthorizationMethodOAuth20AuthType = "oauth2.0"
)

func (e SourceTypeformUpdateAuthorizationMethodOAuth20AuthType) ToPointer() *SourceTypeformUpdateAuthorizationMethodOAuth20AuthType {
	return &e
}

func (e *SourceTypeformUpdateAuthorizationMethodOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceTypeformUpdateAuthorizationMethodOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTypeformUpdateAuthorizationMethodOAuth20AuthType: %v", v)
	}
}

type SourceTypeformUpdateAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                                                  `json:"access_token"`
	AuthType    *SourceTypeformUpdateAuthorizationMethodOAuth20AuthType `json:"auth_type,omitempty"`
	// The Client ID of the Typeform developer application.
	ClientID string `json:"client_id"`
	// The Client Secret the Typeform developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access_token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceTypeformUpdateAuthorizationMethodType string

const (
	SourceTypeformUpdateAuthorizationMethodTypeSourceTypeformUpdateAuthorizationMethodOAuth20      SourceTypeformUpdateAuthorizationMethodType = "source-typeform-update_Authorization Method_OAuth2.0"
	SourceTypeformUpdateAuthorizationMethodTypeSourceTypeformUpdateAuthorizationMethodPrivateToken SourceTypeformUpdateAuthorizationMethodType = "source-typeform-update_Authorization Method_Private Token"
)

type SourceTypeformUpdateAuthorizationMethod struct {
	SourceTypeformUpdateAuthorizationMethodOAuth20      *SourceTypeformUpdateAuthorizationMethodOAuth20
	SourceTypeformUpdateAuthorizationMethodPrivateToken *SourceTypeformUpdateAuthorizationMethodPrivateToken

	Type SourceTypeformUpdateAuthorizationMethodType
}

func CreateSourceTypeformUpdateAuthorizationMethodSourceTypeformUpdateAuthorizationMethodOAuth20(sourceTypeformUpdateAuthorizationMethodOAuth20 SourceTypeformUpdateAuthorizationMethodOAuth20) SourceTypeformUpdateAuthorizationMethod {
	typ := SourceTypeformUpdateAuthorizationMethodTypeSourceTypeformUpdateAuthorizationMethodOAuth20

	return SourceTypeformUpdateAuthorizationMethod{
		SourceTypeformUpdateAuthorizationMethodOAuth20: &sourceTypeformUpdateAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceTypeformUpdateAuthorizationMethodSourceTypeformUpdateAuthorizationMethodPrivateToken(sourceTypeformUpdateAuthorizationMethodPrivateToken SourceTypeformUpdateAuthorizationMethodPrivateToken) SourceTypeformUpdateAuthorizationMethod {
	typ := SourceTypeformUpdateAuthorizationMethodTypeSourceTypeformUpdateAuthorizationMethodPrivateToken

	return SourceTypeformUpdateAuthorizationMethod{
		SourceTypeformUpdateAuthorizationMethodPrivateToken: &sourceTypeformUpdateAuthorizationMethodPrivateToken,
		Type: typ,
	}
}

func (u *SourceTypeformUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceTypeformUpdateAuthorizationMethodPrivateToken := new(SourceTypeformUpdateAuthorizationMethodPrivateToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceTypeformUpdateAuthorizationMethodPrivateToken); err == nil {
		u.SourceTypeformUpdateAuthorizationMethodPrivateToken = sourceTypeformUpdateAuthorizationMethodPrivateToken
		u.Type = SourceTypeformUpdateAuthorizationMethodTypeSourceTypeformUpdateAuthorizationMethodPrivateToken
		return nil
	}

	sourceTypeformUpdateAuthorizationMethodOAuth20 := new(SourceTypeformUpdateAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceTypeformUpdateAuthorizationMethodOAuth20); err == nil {
		u.SourceTypeformUpdateAuthorizationMethodOAuth20 = sourceTypeformUpdateAuthorizationMethodOAuth20
		u.Type = SourceTypeformUpdateAuthorizationMethodTypeSourceTypeformUpdateAuthorizationMethodOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceTypeformUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceTypeformUpdateAuthorizationMethodPrivateToken != nil {
		return json.Marshal(u.SourceTypeformUpdateAuthorizationMethodPrivateToken)
	}

	if u.SourceTypeformUpdateAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceTypeformUpdateAuthorizationMethodOAuth20)
	}

	return nil, nil
}

type SourceTypeformUpdate struct {
	Credentials SourceTypeformUpdateAuthorizationMethod `json:"credentials"`
	// When this parameter is set, the connector will replicate data only from the input forms. Otherwise, all forms in your Typeform account will be replicated. You can find form IDs in your form URLs. For example, in the URL "https://mysite.typeform.com/to/u6nXL7" the form_id is u6nXL7. You can find form URLs on Share panel
	FormIds []string `json:"form_ids,omitempty"`
	// The date from which you'd like to replicate data for Typeform API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}
