// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceApifyDatasetUpdate struct {
	// If set to true, only clean items will be downloaded from the dataset. See description of what clean means in <a href="https://docs.apify.com/api/v2#/reference/datasets/item-collection/get-items">Apify API docs</a>. If not sure, set clean to false.
	Clean *bool `json:"clean,omitempty"`
	// ID of the dataset you would like to load to Airbyte.
	DatasetID *string `json:"datasetId,omitempty"`
	// Your application's Client Secret. You can find this value on the <a href="https://console.apify.com/account/integrations">console integrations tab</a> after you login.
	Token string `json:"token"`
}
