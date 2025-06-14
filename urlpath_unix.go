//go:build unix

package urlpath

import (
	"net/url"
	"path"
)

// fromString must not call path.Clean on URL.Path!
func fromString(pth string) (*url.URL, error) {
	return url.Parse(pth)
}

func isAbsURL(u *url.URL) bool {
	return u.Scheme != "" || path.IsAbs(u.Path)
}

func toString(u *url.URL) string {
	if u.Scheme != "" {
		return u.Scheme + "://" + u.Host + path.Join("/", u.Path)
	}
	return path.Clean(u.Path)
}

func normalize(u, wd *url.URL) string {
	if u.Scheme != "" || u.Host != "" {
		u.Path = path.Clean(u.Path)
	} else {
		u.Scheme = wd.Scheme
		u.Host = wd.Host
		if !path.IsAbs(u.Path) {
			u.Path = path.Join(wd.Path, u.Path)
		} else {
			u.Path = path.Clean(u.Path)
		}
	}
	return toString(u)
}
