package citizen

import (
	"encoding/json"
	"fmt"
)

type (
	PersonAddress struct {
		Street          string  `json:"calle"`
		PostalCode      string  `json:"codigoPostal"`
		ProvinceDesc    string  `json:"descripcionProvincia"`
		AddressLine     string  `json:"direccion"`
		AddressState    string  `json:"estadoDomicilio"`
		ProvinceID      int     `json:"idProvincia"`
		Locality        *string `json:"localidad,omitempty"`
		Number          int     `json:"numero"`
		OfficeDeptLocal *string `json:"oficinaDptoLocal,omitempty"`
		Floor           *string `json:"piso,omitempty"`
		AdditionalData  *string `json:"datoAdicional,omitempty"`
		AdditionalType  *string `json:"tipoDatoAdicional,omitempty"`
		AddressType     string  `json:"tipoDomicilio"`
	}

	FlexibleAddress struct {
		Addresses []*PersonAddress
	}

	// Tipo personalizado para fechas
	FlexibleDate string

	// Estructura de respuesta
	GetPersonInformationResponse struct {
		Surname                 *string         `json:"apellido,omitempty"`
		DescriptionMainActivity *string         `json:"descripcionActividadPrincipal,omitempty"`
		Address                 FlexibleAddress `json:"domicilio,omitempty"`
		StateKey                *string         `json:"estadoClave,omitempty"`
		BirthDate               FlexibleDate    `json:"fechaNacimiento,omitempty"`
		SocialContractDate      FlexibleDate    `json:"fechaContratoSocial,omitempty"`
		DateOfDeath             FlexibleDate    `json:"fechaFallecimiento,omitempty"`
		LegalForm               *string         `json:"formaJuridica,omitempty"`
		MainActivityID          *int            `json:"idActividadPrincipal,omitempty"`
		PersonID                int             `json:"idPersona"`
		ClosingMonth            *int            `json:"mesCierre,omitempty"`
		MainActivityPeriod      *int            `json:"periodoActividadPrincipal,omitempty"`
		BusinessName            *string         `json:"razonSocial,omitempty"`
		Name                    *string         `json:"nombre,omitempty"`
		DocumentNumber          *string         `json:"numeroDocumento,omitempty"`
		KeyType                 string          `json:"tipoClave"`
		DocumentType            *string         `json:"tipoDocumento,omitempty"`
		PersonType              string          `json:"tipoPersona"`
	}
)

// To solve afip response irregularies set a custom json unmarshal function
func (fa *FlexibleAddress) UnmarshalJSON(data []byte) error {
	var singleAddress *PersonAddress
	if err := json.Unmarshal(data, &singleAddress); err == nil {
		fa.Addresses = []*PersonAddress{singleAddress}
		return nil
	}

	var multipleAddresses []*PersonAddress
	if err := json.Unmarshal(data, &multipleAddresses); err == nil {
		fa.Addresses = multipleAddresses
		return nil
	}

	return fmt.Errorf("data does not match expected structure for domicilio")
}

func (fd *FlexibleDate) UnmarshalJSON(data []byte) error {
	if string(data) == "{}" || string(data) == "null" {
		*fd = ""
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*fd = FlexibleDate(str)
	return nil
}
