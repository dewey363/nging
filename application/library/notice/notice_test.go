/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package notice

import (
	"testing"
	"fmt"

	"github.com/webx-top/com"
)

func TestOpenMessage(t *testing.T) {
	OpenMessage(`testUser`, `testType`)
	user := Default().User[`testUser`]
	if len(user.Notice.Types) != 1 {
		t.Errorf(`Size of types != %v`, 1)
	}
	if user.Notice.Types[`testType`] != true {
		t.Error(`Type of testType != true`)
	}
	
	clientID := OpenClient(`testUser`)
	if user.Clients != 1 {
		t.Error(`Number of clients (`+fmt.Sprint(user.Clients)+`) != 1`)
	}
	
	CloseClient(`testUser`, clientID)
	if user.Clients != 0 {
		t.Error(`Number of clients (`+fmt.Sprint(user.Clients)+`) != 0`)
	}
	
	com.Dump(Default())
}

func TestSend(t *testing.T) {
	go func() {
		t.Log(string(RecvJSON(`testUser`, 0)))
	}()
	Send(`testUser`, NewMessageWithValue(`testType`, `testTitle`, `testContent`))
	com.Dump(Default())
}
