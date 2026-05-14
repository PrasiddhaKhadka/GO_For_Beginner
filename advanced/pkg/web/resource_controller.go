package web

import "errors"

type ResourceController interface {
	Index(ctx *Ctx) (*Result, error)
	Create(ctx *Ctx) (*Result, error)
	Store(ctx *Ctx) (*Result, error)
	Show(ctx *Ctx) (*Result, error)
	Edit(ctx *Ctx) (*Result, error)
	Update(ctx *Ctx) (*Result, error)
	Delete(ctx *Ctx) (*Result, error)
}

type UnimplementedResourceConrtoller struct{}

// "Get /campaign"
func (c UnimplementedResourceConrtoller) Index(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}

// "GET /campaign/create"
func (c UnimplementedResourceConrtoller) Create(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}

// "POST /campaign"
func (c UnimplementedResourceConrtoller) Store(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}

// "GET /campaign/{id}"
func (c UnimplementedResourceConrtoller) Show(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}

// "PUT /campaign/{id}"
func (c UnimplementedResourceConrtoller) Update(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}

// "GET /campaign/edit/{campaign_id}"
func (c UnimplementedResourceConrtoller) Edit(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}

// "DELETE /campaign/{campaign_id}"
func (c UnimplementedResourceConrtoller) Delete(ctx *Ctx) (*Result, error) {

	return nil, errors.New("not implemented")
}
