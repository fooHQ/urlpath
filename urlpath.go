package urlpath

import (
	"net/url"
	"path"
)

// Abs resolves a path to an absolute URL string relative to a working directory.
// It parses the input path and working directory as URLs, normalizes the path
// based on the working directory, and returns the resulting absolute URL string.
// If the working directory is empty, it defaults to "/".
// On Unix, an absolute URL is determined by a non-empty scheme or host with an
// absolute path (or a file URL with an absolute path).
// On Windows, an absolute URL may also include a volume name (e.g., "C:") or a
// path starting with a forward slash. Windows paths with backslashes are converted
// to forward slashes for consistency.
// Returns an error if either the path or working directory fails to parse.
func Abs(pth, wd string) (string, error) {
	if wd == "" {
		wd = "/"
	}
	u, err := fromString(pth)
	if err != nil {
		return "", err
	}
	wdu, err := fromString(wd)
	if err != nil {
		return "", err
	}
	return normalize(u, wdu), nil
}

// Base returns the last element of the path component of the input URL string.
// It parses the input as a URL and extracts the base name of the path using
// path.Base. The result is the final component of the path, which may be a file
// or directory name.
// Returns an error if the input path fails to parse.
func Base(pth string) (string, error) {
	u, err := fromString(pth)
	if err != nil {
		return "", err
	}
	return path.Base(u.Path), nil
}

// Dir returns the directory component of the input URL string by removing the
// last path element.
// It parses the input as a URL, extracts the directory using path.Dir, and
// returns the resulting URL string.
// Returns an error if the input path fails to parse.
func Dir(pth string) (string, error) {
	u, err := fromString(pth)
	if err != nil {
		return "", err
	}
	u.Path = path.Dir(u.Path)
	return toString(u), nil
}

// Ext returns the file extension of the path component of the input URL string.
// It parses the input as a URL and extracts the extension using path.Ext.
// The extension includes the leading dot (e.g., ".txt") or is empty if no
// extension is present.
// Returns an error if the input path fails to parse.
func Ext(pth string) (string, error) {
	u, err := fromString(pth)
	if err != nil {
		return "", err
	}
	return path.Ext(u.Path), nil
}

// Clean cleans the path component of the input URL string, removing unnecessary
// elements like "." and "..".
// It parses the input as a URL, applies path.Clean to the path component, and
// returns the resulting URL string.
// Returns an error if the input path fails to parse.
func Clean(pth string) (string, error) {
	u, err := fromString(pth)
	if err != nil {
		return "", err
	}
	u.Path = path.Clean(u.Path)
	return toString(u), nil
}

// IsAbs determines if the input URL string represents an absolute URL.
// On Unix, a URL is absolute if it has a non-empty scheme or host and an absolute
// path (or is a file URL with an absolute path).
// On Windows, a URL is absolute if it has a non-empty scheme, host, volume name
// (e.g., "C:"), or an absolute path starting with a forward slash.
// Returns a boolean indicating if the URL is absolute and an error if the input
// path fails to parse.
func IsAbs(pth string) (bool, error) {
	u, err := fromString(pth)
	if err != nil {
		return false, err
	}
	return isAbsURL(u), nil
}

// Join combines multiple URL path strings into a single URL string.
// It parses each input string as a URL, joins their path components using
// path.Join, and returns the resulting URL string.
// The first element sets the scheme and host, and subsequent elements contribute
// to the path. If the first element has a scheme or host, it is preserved in the
// output.
// Returns an error if any input path fails to parse.
func Join(elem ...string) (string, error) {
	var uf *url.URL
	for _, e := range elem {
		u, err := fromString(e)
		if err != nil {
			return "", err
		}
		if uf == nil {
			uf = u
			continue
		}
		uf.Path = path.Join(uf.Path, u.Path)
	}
	return toString(uf), nil
}

// Split splits the path component of the input URL string into its directory and
// file components.
// It parses the input as a URL, uses path.Split to separate the path into
// directory and file parts, and returns the directory as a URL string and the
// file as a string.
// Returns an error if the input path fails to parse.
func Split(pth string) (string, string, error) {
	u, err := fromString(pth)
	if err != nil {
		return "", "", err
	}
	pthDir, pthFile := path.Split(u.Path)
	u.Path = pthDir
	return toString(u), pthFile, nil
}

// Match checks if the path component of the input URL string matches the given
// pattern.
// It parses the input as a URL, converts it to a string, and uses path.Match to
// check if the pattern matches the resulting path.
// The pattern follows the syntax of path.Match, supporting wildcards like "*" and
// "?".
// Returns a boolean indicating if the pattern matches and an error if the input
// path fails to parse.
func Match(pattern, name string) (bool, error) {
	u, err := fromString(name)
	if err != nil {
		return false, err
	}
	name = toString(u)
	return path.Match(pattern, name)
}

// Path extracts the path component of the input URL string.
// It parses the input as a URL, cleans the path component to remove unnecessary
// elements like "." and "..", and returns the resulting path as a string.
// On Windows, if the path includes a volume name (e.g., "C:"), the leading slash
// is removed from the output. For file URLs with a non-empty host, the host is
// included in the output path (e.g., "//host/path").
// Returns an error if the input path fails to parse.
func Path(pth string) (string, error) {
	u, err := fromString(pth)
	if err != nil {
		return "", err
	}
	return pathURL(u), nil
}

func isFileURL(u *url.URL) bool {
	return u.Scheme == "file" || u.Scheme == ""
}
