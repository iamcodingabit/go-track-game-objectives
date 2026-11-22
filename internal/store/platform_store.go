package store

import (
	"context"
	"database/sql"
)

type Platform struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type PlatformStore struct {
	db *sql.DB
}

func (ps *PlatformStore) GetAllPlatforms (ctx context.Context) ([]Platform, error) {
	var platforms []Platform

	rows, err := ps.db.QueryContext(ctx, `SELECT id, name FROM platform`)
	if err != nil {
		return platforms, err
	}	
	
	for rows.Next() {
		var platform Platform
		err := rows.Scan(&platform.ID, &platform.Name);
		if  err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}

	err = rows.Err();
	if err != nil {
		return nil, err
	}

	return platforms, nil
}

func(ps *PlatformStore) AddNewPlatform (ctx context.Context, newPlatform Platform) (Platform, error) {
	var platform Platform
	err := ps.db.QueryRowContext(ctx, `INSERT INTO platform(name) VALUES ($1) RETURNING id, name`, newPlatform.Name).Scan(&platform.ID, &platform.Name)
	if err != nil {
		return platform, err
	}

	return platform, nil
}