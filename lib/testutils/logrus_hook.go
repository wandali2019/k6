/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2019 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package testutils

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

// SimpleLogrusHook implements the logrus.Hook interface and could be used to check
// if log messages were outputted
type SimpleLogrusHook struct {
	HookedLevels []log.Level
	mutex        sync.Mutex
	messageCache []log.Entry
}

// Levels just returns whatever was stored in the HookedLevels slice
func (smh *SimpleLogrusHook) Levels() []log.Level {
	return smh.HookedLevels
}

// Fire saves whatever message the logrus library passed in the cache
func (smh *SimpleLogrusHook) Fire(e *log.Entry) error {
	smh.mutex.Lock()
	defer smh.mutex.Unlock()
	smh.messageCache = append(smh.messageCache, *e)
	return nil
}

// Drain returns the currently stored messages and deletes them from the cache
func (smh *SimpleLogrusHook) Drain() []log.Entry {
	smh.mutex.Lock()
	defer smh.mutex.Unlock()
	res := smh.messageCache
	smh.messageCache = []log.Entry{}
	return res
}

var _ log.Hook = &SimpleLogrusHook{}
