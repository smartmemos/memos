//go:generate mockgen -source $GOFILE -destination ./service_mock.go -package $GOPACKAGE
package memo

type Service interface {
}
