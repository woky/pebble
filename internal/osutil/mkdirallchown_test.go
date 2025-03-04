// Copyright (c) 2014-2020 Canonical Ltd
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License version 3 as
// published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package osutil_test

import (
	"strings"

	"gopkg.in/check.v1"

	"github.com/canonical/pebble/internal/osutil"
)

type mkdacSuite struct{}

var _ = check.Suite(&mkdacSuite{})

func (mkdacSuite) TestSlashySlashy(c *check.C) {
	for _, dir := range []string{
		// these must start with "/" (because d doesn't end in /, and we
		// are _not_ using filepath.Join, on purpose)
		"/foo/bar",
		"/foo/bar/",
	} {
		d := c.MkDir()
		// just in case
		c.Assert(strings.HasSuffix(d, "/"), check.Equals, false)
		err := osutil.MkdirAllChown(d+dir, 0755, osutil.NoChown, osutil.NoChown)
		c.Assert(err, check.IsNil, check.Commentf("%q", dir))
	}
}

// Add some very basic tests of the functionality (chown requires root, so use
// NoChown for these). Permissions and user/group are tested in other places.
func (mkdacSuite) TestMkdirAllChown(c *check.C) {
	tmpDir := c.MkDir()

	err := osutil.MkdirAllChown(tmpDir+"/foo/bar", 0o755, osutil.NoChown, osutil.NoChown)
	c.Assert(err, check.IsNil)
	c.Assert(osutil.IsDir(tmpDir+"/foo"), check.Equals, true)
	c.Assert(osutil.IsDir(tmpDir+"/foo/bar"), check.Equals, true)

	err = osutil.MkdirChown(tmpDir+"/foo/bar", 0o755, osutil.NoChown, osutil.NoChown)
	c.Assert(err, check.ErrorMatches, `.*: file exists`)
}

func (mkdacSuite) TestMkdirChown(c *check.C) {
	tmpDir := c.MkDir()

	err := osutil.MkdirChown(tmpDir+"/foo", 0o755, osutil.NoChown, osutil.NoChown)
	c.Assert(err, check.IsNil)
	c.Assert(osutil.IsDir(tmpDir+"/foo"), check.Equals, true)

	err = osutil.MkdirChown(tmpDir+"/foo", 0o755, osutil.NoChown, osutil.NoChown)
	c.Assert(err, check.ErrorMatches, `.*: file exists`)
}
