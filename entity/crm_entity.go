package entity

type CRMEntity string

const (
	Account CRMEntity = "Account"

	Contact CRMEntity = "Contact"
)

func (d CRMEntity) Get() []CRMEntity {
	return []CRMEntity{
		Account,
		Contact,
	}
}
