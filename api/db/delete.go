package db

import (
	"context"

	"github.com/spaceuptech/space-api-go/api/config"
	"github.com/spaceuptech/space-api-go/api/model"
	"github.com/spaceuptech/space-api-go/api/utils"
)

// Delete contains the methods for the delete operation
type Delete struct {
	ctx    context.Context
	meta   *model.Meta
	op     string
	find   utils.M
	config *config.Config
}

func initDelete(ctx context.Context, db, col, op string, config *config.Config) *Delete {
	return &Delete{
		ctx:    ctx,
		meta:   &model.Meta{DB: db, Col: col, Token: config.Token, Project: config.Project},
		op:     op,
		find:   make(utils.M),
		config: config}
}

// Where sets the where clause for the request
func (d *Delete) Where(conds ...utils.M) *Delete {
	if len(conds) == 1 {
		d.find = utils.GenerateFind(conds[0])
	} else {
		d.find = utils.GenerateFind(utils.And(conds...))
	}
	return d
}

// Apply executes the operation and returns the result
func (d *Delete) Apply() (*model.Response, error) {
	return d.config.Transport.Delete(d.ctx, d.meta, d.op, d.find)
}

func (d *Delete) getProject() string {
	return d.config.Project
}
func (d *Delete) getDb() string {
	return d.meta.DB
}
func (d *Delete) getToken() string {
	return d.config.Token
}
func (d *Delete) getCollection() string {
	return d.meta.Col
}
func (d *Delete) getOperation() string {
	return d.op
}
func (d *Delete) getFind() utils.M {
	return d.find
}
