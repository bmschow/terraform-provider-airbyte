// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType string

const (
	SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthTypeToken SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType = "Token"
)

func (e SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType) ToPointer() *SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType {
	return &e
}

func (e *SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType: %v", v)
	}
}

// SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken - Choose how to authenticate to Harvest.
type SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken struct {
	// Log into Harvest and then create new <a href="https://id.getharvest.com/developers"> personal access token</a>.
	APIToken string                                                                           `json:"api_token"`
	AuthType *SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType `json:"auth_type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken

func (c *SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken) UnmarshalJSON(bs []byte) error {
	data := _SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "api_token")
	delete(additionalFields, "auth_type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType string

const (
	SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthTypeClient SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType = "Client"
)

func (e SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType) ToPointer() *SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType {
	return &e
}

func (e *SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType: %v", v)
	}
}

// SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth - Choose how to authenticate to Harvest.
type SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth struct {
	AuthType *SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType `json:"auth_type,omitempty"`
	// The Client ID of your Harvest developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Harvest developer application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth

func (c *SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth) UnmarshalJSON(bs []byte) error {
	data := _SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "auth_type")
	delete(additionalFields, "client_id")
	delete(additionalFields, "client_secret")
	delete(additionalFields, "refresh_token")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type SourceHarvestAuthenticationMechanismType string

const (
	SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth         SourceHarvestAuthenticationMechanismType = "source-harvest_Authentication mechanism_Authenticate via Harvest (OAuth)"
	SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceHarvestAuthenticationMechanismType = "source-harvest_Authentication mechanism_Authenticate with Personal Access Token"
)

type SourceHarvestAuthenticationMechanism struct {
	SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth         *SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth
	SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken *SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken

	Type SourceHarvestAuthenticationMechanismType
}

func CreateSourceHarvestAuthenticationMechanismSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth(sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth) SourceHarvestAuthenticationMechanism {
	typ := SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth

	return SourceHarvestAuthenticationMechanism{
		SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth: &sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth,
		Type: typ,
	}
}

func CreateSourceHarvestAuthenticationMechanismSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken(sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken) SourceHarvestAuthenticationMechanism {
	typ := SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken

	return SourceHarvestAuthenticationMechanism{
		SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken: &sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceHarvestAuthenticationMechanism) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken := new(SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken); err == nil {
		u.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken = sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken
		u.Type = SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken
		return nil
	}

	sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth := new(SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth); err == nil {
		u.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth = sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth
		u.Type = SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceHarvestAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		return json.Marshal(u.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken)
	}

	if u.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth != nil {
		return json.Marshal(u.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth)
	}

	return nil, nil
}

type SourceHarvestHarvest string

const (
	SourceHarvestHarvestHarvest SourceHarvestHarvest = "harvest"
)

func (e SourceHarvestHarvest) ToPointer() *SourceHarvestHarvest {
	return &e
}

func (e *SourceHarvestHarvest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "harvest":
		*e = SourceHarvestHarvest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestHarvest: %v", v)
	}
}

type SourceHarvest struct {
	// Harvest account ID. Required for all Harvest requests in pair with Personal Access Token
	AccountID string `json:"account_id"`
	// Choose how to authenticate to Harvest.
	Credentials *SourceHarvestAuthenticationMechanism `json:"credentials,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
	ReplicationEndDate *time.Time `json:"replication_end_date,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	ReplicationStartDate time.Time            `json:"replication_start_date"`
	SourceType           SourceHarvestHarvest `json:"sourceType"`
}
