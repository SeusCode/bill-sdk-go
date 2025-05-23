package pos

type (
	PointOfSale struct {
		Nro         int32  `xml:"Nro,omitempty" json:"Nro,omitempty"`
		EmisionTipo string `xml:"EmisionTipo,omitempty" json:"EmisionTipo,omitempty"`
		Bloqueado   string `xml:"Bloqueado,omitempty" json:"Bloqueado,omitempty"`
		FchBaja     string `xml:"FchBaja,omitempty" json:"FchBaja,omitempty"`
	}
)
