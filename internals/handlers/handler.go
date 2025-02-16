package handlers

import "database/sql"

type Handler struct {
	Database *sql.DB
}

func NewHandler(database *sql.DB) *Handler {
	return &Handler{Database: database}
}
