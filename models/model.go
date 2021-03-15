package models

type BeforeSave interface {
	BeforeSave() error
}
