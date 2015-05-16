package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type country struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *country) create() error {
	res, err := db.Exec("INSERT INTO countries (name) VALUES ($1)", c.Name)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New(funcName() + ": Failed to create country")
	}
	return nil
}

func (c *country) read() error {
	err := db.QueryRow("SELECT name FROM countries WHERE id = $1", &c.Id).Scan(&c.Name)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *country) update() error {
	res, err := db.Exec("UPDATE countries SET name = $2 WHERE id = $1", c.Id, c.Name)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("%s: Failed to update country: %+v", c)
	}
	return nil
}

func (c *country) delete() error {
	res, err := db.Exec("DELETE FROM countries WHERE id = $1", c.Id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("%s: Failed to update country: %+v", c)
	}
	return nil
}