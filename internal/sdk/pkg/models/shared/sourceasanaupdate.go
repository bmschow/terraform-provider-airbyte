// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle - PAT Credentials
type SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle string

const (
	SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitlePatCredentials SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle = "PAT Credentials"
)

func (e SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle) ToPointer() *SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle {
	return &e
}

func (e *SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PAT Credentials":
		*e = SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle: %v", v)
	}
}

// SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken - Choose how to authenticate to Github
type SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken struct {
	// PAT Credentials
	OptionTitle *SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle `json:"option_title,omitempty"`
	// Asana Personal Access Token (generate yours <a href="https://app.asana.com/0/developer-console">here</a>).
	PersonalAccessToken string `json:"personal_access_token"`
}

// SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle - OAuth Credentials
type SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle string

const (
	SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitleOAuthCredentials SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle = "OAuth Credentials"
)

func (e SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle) ToPointer() *SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle {
	return &e
}

func (e *SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth Credentials":
		*e = SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle: %v", v)
	}
}

// SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth - Choose how to authenticate to Github
type SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// OAuth Credentials
	OptionTitle  *SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle `json:"option_title,omitempty"`
	RefreshToken string                                                                             `json:"refresh_token"`
}

type SourceAsanaUpdateAuthenticationMechanismType string

const (
	SourceAsanaUpdateAuthenticationMechanismTypeSourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth           SourceAsanaUpdateAuthenticationMechanismType = "source-asana-update_Authentication mechanism_Authenticate via Asana (Oauth)"
	SourceAsanaUpdateAuthenticationMechanismTypeSourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceAsanaUpdateAuthenticationMechanismType = "source-asana-update_Authentication mechanism_Authenticate with Personal Access Token"
)

type SourceAsanaUpdateAuthenticationMechanism struct {
	SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth           *SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth
	SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken *SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken

	Type SourceAsanaUpdateAuthenticationMechanismType
}

func CreateSourceAsanaUpdateAuthenticationMechanismSourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth(sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth) SourceAsanaUpdateAuthenticationMechanism {
	typ := SourceAsanaUpdateAuthenticationMechanismTypeSourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth

	return SourceAsanaUpdateAuthenticationMechanism{
		SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth: &sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth,
		Type: typ,
	}
}

func CreateSourceAsanaUpdateAuthenticationMechanismSourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken(sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken) SourceAsanaUpdateAuthenticationMechanism {
	typ := SourceAsanaUpdateAuthenticationMechanismTypeSourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken

	return SourceAsanaUpdateAuthenticationMechanism{
		SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken: &sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceAsanaUpdateAuthenticationMechanism) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken := new(SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken); err == nil {
		u.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken = sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken
		u.Type = SourceAsanaUpdateAuthenticationMechanismTypeSourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken
		return nil
	}

	sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth := new(SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth); err == nil {
		u.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth = sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth
		u.Type = SourceAsanaUpdateAuthenticationMechanismTypeSourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAsanaUpdateAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		return json.Marshal(u.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken)
	}

	if u.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth != nil {
		return json.Marshal(u.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth)
	}

	return nil, nil
}

type SourceAsanaUpdate struct {
	// Choose how to authenticate to Github
	Credentials *SourceAsanaUpdateAuthenticationMechanism `json:"credentials,omitempty"`
}
