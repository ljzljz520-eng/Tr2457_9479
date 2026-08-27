package archive

import (
	"encoding/json"
	"frontend_go/model"
)

func Marshal(rs []model.Record) ([]byte, error) { return json.Marshal(Archived(rs)) }
func Unmarshal(data []byte) ([]model.Record, error) {
	var rs []model.Record
	err := json.Unmarshal(data, &rs)
	return rs, err
}
func RestoreAll(m *Manager, rs []model.Record) error {
	for _, r := range rs {
		if r.IsArchived() {
			if err := m.db.PutRecord(r); err != nil {
				return err
			}
		}
	}
	return nil
}
