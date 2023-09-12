// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationPineconePinecone string

const (
	DestinationPineconePineconePinecone DestinationPineconePinecone = "pinecone"
)

func (e DestinationPineconePinecone) ToPointer() *DestinationPineconePinecone {
	return &e
}

func (e *DestinationPineconePinecone) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinecone":
		*e = DestinationPineconePinecone(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconePinecone: %v", v)
	}
}

type DestinationPineconeEmbeddingFakeMode string

const (
	DestinationPineconeEmbeddingFakeModeFake DestinationPineconeEmbeddingFakeMode = "fake"
)

func (e DestinationPineconeEmbeddingFakeMode) ToPointer() *DestinationPineconeEmbeddingFakeMode {
	return &e
}

func (e *DestinationPineconeEmbeddingFakeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationPineconeEmbeddingFakeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeEmbeddingFakeMode: %v", v)
	}
}

// DestinationPineconeEmbeddingFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationPineconeEmbeddingFake struct {
	Mode *DestinationPineconeEmbeddingFakeMode `json:"mode,omitempty"`
}

type DestinationPineconeEmbeddingCohereMode string

const (
	DestinationPineconeEmbeddingCohereModeCohere DestinationPineconeEmbeddingCohereMode = "cohere"
)

func (e DestinationPineconeEmbeddingCohereMode) ToPointer() *DestinationPineconeEmbeddingCohereMode {
	return &e
}

func (e *DestinationPineconeEmbeddingCohereMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationPineconeEmbeddingCohereMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeEmbeddingCohereMode: %v", v)
	}
}

// DestinationPineconeEmbeddingCohere - Use the Cohere API to embed text.
type DestinationPineconeEmbeddingCohere struct {
	CohereKey string                                  `json:"cohere_key"`
	Mode      *DestinationPineconeEmbeddingCohereMode `json:"mode,omitempty"`
}

type DestinationPineconeEmbeddingOpenAIMode string

const (
	DestinationPineconeEmbeddingOpenAIModeOpenai DestinationPineconeEmbeddingOpenAIMode = "openai"
)

func (e DestinationPineconeEmbeddingOpenAIMode) ToPointer() *DestinationPineconeEmbeddingOpenAIMode {
	return &e
}

func (e *DestinationPineconeEmbeddingOpenAIMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationPineconeEmbeddingOpenAIMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeEmbeddingOpenAIMode: %v", v)
	}
}

// DestinationPineconeEmbeddingOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPineconeEmbeddingOpenAI struct {
	Mode      *DestinationPineconeEmbeddingOpenAIMode `json:"mode,omitempty"`
	OpenaiKey string                                  `json:"openai_key"`
}

type DestinationPineconeEmbeddingType string

const (
	DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingOpenAI DestinationPineconeEmbeddingType = "destination-pinecone_Embedding_OpenAI"
	DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingCohere DestinationPineconeEmbeddingType = "destination-pinecone_Embedding_Cohere"
	DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingFake   DestinationPineconeEmbeddingType = "destination-pinecone_Embedding_Fake"
)

type DestinationPineconeEmbedding struct {
	DestinationPineconeEmbeddingOpenAI *DestinationPineconeEmbeddingOpenAI
	DestinationPineconeEmbeddingCohere *DestinationPineconeEmbeddingCohere
	DestinationPineconeEmbeddingFake   *DestinationPineconeEmbeddingFake

	Type DestinationPineconeEmbeddingType
}

func CreateDestinationPineconeEmbeddingDestinationPineconeEmbeddingOpenAI(destinationPineconeEmbeddingOpenAI DestinationPineconeEmbeddingOpenAI) DestinationPineconeEmbedding {
	typ := DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingOpenAI

	return DestinationPineconeEmbedding{
		DestinationPineconeEmbeddingOpenAI: &destinationPineconeEmbeddingOpenAI,
		Type:                               typ,
	}
}

func CreateDestinationPineconeEmbeddingDestinationPineconeEmbeddingCohere(destinationPineconeEmbeddingCohere DestinationPineconeEmbeddingCohere) DestinationPineconeEmbedding {
	typ := DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingCohere

	return DestinationPineconeEmbedding{
		DestinationPineconeEmbeddingCohere: &destinationPineconeEmbeddingCohere,
		Type:                               typ,
	}
}

func CreateDestinationPineconeEmbeddingDestinationPineconeEmbeddingFake(destinationPineconeEmbeddingFake DestinationPineconeEmbeddingFake) DestinationPineconeEmbedding {
	typ := DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingFake

	return DestinationPineconeEmbedding{
		DestinationPineconeEmbeddingFake: &destinationPineconeEmbeddingFake,
		Type:                             typ,
	}
}

func (u *DestinationPineconeEmbedding) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPineconeEmbeddingFake := new(DestinationPineconeEmbeddingFake)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPineconeEmbeddingFake); err == nil {
		u.DestinationPineconeEmbeddingFake = destinationPineconeEmbeddingFake
		u.Type = DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingFake
		return nil
	}

	destinationPineconeEmbeddingOpenAI := new(DestinationPineconeEmbeddingOpenAI)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPineconeEmbeddingOpenAI); err == nil {
		u.DestinationPineconeEmbeddingOpenAI = destinationPineconeEmbeddingOpenAI
		u.Type = DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingOpenAI
		return nil
	}

	destinationPineconeEmbeddingCohere := new(DestinationPineconeEmbeddingCohere)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPineconeEmbeddingCohere); err == nil {
		u.DestinationPineconeEmbeddingCohere = destinationPineconeEmbeddingCohere
		u.Type = DestinationPineconeEmbeddingTypeDestinationPineconeEmbeddingCohere
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPineconeEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationPineconeEmbeddingFake != nil {
		return json.Marshal(u.DestinationPineconeEmbeddingFake)
	}

	if u.DestinationPineconeEmbeddingOpenAI != nil {
		return json.Marshal(u.DestinationPineconeEmbeddingOpenAI)
	}

	if u.DestinationPineconeEmbeddingCohere != nil {
		return json.Marshal(u.DestinationPineconeEmbeddingCohere)
	}

	return nil, nil
}

// DestinationPineconeIndexing - Pinecone is a popular vector store that can be used to store and retrieve embeddings.
type DestinationPineconeIndexing struct {
	// Pinecone index to use
	Index string `json:"index"`
	// Pinecone environment to use
	PineconeEnvironment string `json:"pinecone_environment"`
	PineconeKey         string `json:"pinecone_key"`
}

type DestinationPineconeProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `json:"chunk_overlap,omitempty"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
}

type DestinationPinecone struct {
	DestinationType DestinationPineconePinecone `json:"destinationType"`
	// Embedding configuration
	Embedding DestinationPineconeEmbedding `json:"embedding"`
	// Pinecone is a popular vector store that can be used to store and retrieve embeddings.
	Indexing   DestinationPineconeIndexing              `json:"indexing"`
	Processing DestinationPineconeProcessingConfigModel `json:"processing"`
}
