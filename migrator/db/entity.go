/**
 * This file is part of the raoptimus/db-migrator.go library
 *
 * @copyright Copyright (c) Evgeniy Urvantsev <resmus@gmail.com>
 * @license https://github.com/raoptimus/db-migrator.go/blob/master/LICENSE.md
 * @link https://github.com/raoptimus/db-migrator.go
 */
package db

import (
	"sort"
	"time"
)

type (
	MigrationEntity struct {
		Version   string
		ApplyTime int
	}
	MigrationEntityList []MigrationEntity
)

func (s MigrationEntityList) Len() int {
	return len(s)
}

func (s MigrationEntityList) Less(i, j int) bool {
	return s[i].Version < s[j].Version
}

func (s MigrationEntityList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s MigrationEntityList) SortByVersion() {
	sort.Sort(s)
}

func (s MigrationEntity) ApplyTimeFormat() string {
	return time.Unix(int64(s.ApplyTime), 0).Format("2006-01-02 15:04:05")
}
