//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package system

type DAO interface {
}
