package model

type HelloIn struct {
	Name string `form:"name" json:"name" binding:"required,oneof=PointLight SpotLight RectLight"`
}
type HelloOut struct {
	Message string `json:"message"`
}
