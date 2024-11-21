package models

import "time"

type Todo struct {
    Title string
    Deadline time.Time
}
