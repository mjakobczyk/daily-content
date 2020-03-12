package newsapi

import "encoding/json"

// SourceDTO type.
type SourceDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type sourceRequestDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UnmarshalJSON provides custom unmarshaling for sourceRequestDTO type.
func (s *SourceDTO) UnmarshalJSON(data []byte) error {
	var sourceRequestDTO sourceRequestDTO

	if err := json.Unmarshal(data, &sourceRequestDTO); err != nil {
		return err
	}

	s.ID = sourceRequestDTO.ID
	s.Name = sourceRequestDTO.Name

	return nil
}
