package postgres

import "github.com/DionisiyGri/sample-go/internal/model"

// GetAll used to get all entities. It is just a plug.
func (s *SomeService) GetAll() ([]model.Some, error) {
	var items []model.Some
	rows, err := s.db.Query("SELECT * from some_table")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := model.Some{}
		if err := rows.Scan(
			&item.ID,
			&item.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// PostSome used to save singe entity. It is just a plug.
func (s *SomeService) PostSome(some *model.Some) (*model.Some, error) {
	var item model.Some
	stmt := `INSERT INTO some_table(name) VALUES ($1) RETURNING id, name`
	err := s.db.QueryRow(stmt, some.Name).Scan(&item.ID, &item.Name)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
