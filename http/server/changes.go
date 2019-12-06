package server

import (
	"strconv"

	"github.com/keys-pub/keys"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type changes struct {
	docs        []*keys.Document
	version     int64
	versionNext int64
	badRequest  error
}

func (s *Server) changes(c echo.Context, path string) (*changes, error) {
	request := c.Request()
	ctx := request.Context()

	var version keys.TimeMs
	if f := c.QueryParam("version"); f != "" {
		i, err := strconv.Atoi(f)
		if err != nil {
			return &changes{badRequest: ErrBadRequest(c, errors.Wrapf(err, "invalid version"))}, nil
		}
		version = keys.TimeMs(i)
	}
	plimit := c.QueryParam("limit")
	if plimit == "" {
		plimit = "100"
	}
	limit, err := strconv.Atoi(plimit)
	if err != nil {
		return &changes{badRequest: ErrBadRequest(c, errors.Wrapf(err, "invalid limit"))}, nil
	}
	if limit > 100 {
		return &changes{badRequest: ErrBadRequest(c, errors.Wrapf(err, "invalid limit, too large"))}, nil
	}

	logger.Infof(ctx, "Changes %s", path)
	chngs, to, err := s.fi.Changes(ctx, path, keys.TimeFromMillis(version), limit)
	if err != nil {
		return nil, err
	}

	logger.Infof(ctx, "Changes %s, found ", path, len(chngs))
	paths := make([]string, 0, len(chngs))
	for _, a := range chngs {
		paths = append(paths, a.Path)
	}
	docs, err := s.fi.GetAll(ctx, paths)
	if err != nil {
		return nil, err
	}

	versionNext := keys.TimeMs(0)
	if to.IsZero() {
		versionNext = version
	} else {
		versionNext = keys.TimeToMillis(to)
	}

	logger.Infof(ctx, "Changes %s, version next: %d", path, versionNext)

	return &changes{
		docs:        docs,
		version:     int64(version),
		versionNext: int64(versionNext),
	}, nil
}