package citizen

import "time"

type (
	CitizenP4 struct {
		Actividad                 []*Actividad         `xml:"actividad,omitempty" json:"actividad,omitempty"`
		Apellido                  string               `xml:"apellido,omitempty" json:"apellido,omitempty"`
		CantidadSociosEmpresaMono int32                `xml:"cantidadSociosEmpresaMono,omitempty" json:"cantidadSociosEmpresaMono,omitempty"`
		Categoria                 []*Categoria         `xml:"categoria,omitempty" json:"categoria,omitempty"`
		ClaveInactivaAsociada     []*int64             `xml:"claveInactivaAsociada,omitempty" json:"claveInactivaAsociada,omitempty"`
		Dependencia               *Dependencia         `xml:"dependencia,omitempty" json:"dependencia,omitempty"`
		Domicilio                 []*Domicilio         `xml:"domicilio,omitempty" json:"domicilio,omitempty"`
		Email                     []*Email             `xml:"email,omitempty" json:"email,omitempty"`
		EstadoClave               string               `xml:"estadoClave,omitempty" json:"estadoClave,omitempty"`
		FechaContratoSocial       time.Time            `xml:"fechaContratoSocial,omitempty" json:"fechaContratoSocial,omitempty"`
		FechaFallecimiento        time.Time            `xml:"fechaFallecimiento,omitempty" json:"fechaFallecimiento,omitempty"`
		FechaInscripcion          time.Time            `xml:"fechaInscripcion,omitempty" json:"fechaInscripcion,omitempty"`
		FechaJubilado             time.Time            `xml:"fechaJubilado,omitempty" json:"fechaJubilado,omitempty"`
		FechaNacimiento           time.Time            `xml:"fechaNacimiento,omitempty" json:"fechaNacimiento,omitempty"`
		FechaVencimientoMigracion time.Time            `xml:"fechaVencimientoMigracion,omitempty" json:"fechaVencimientoMigracion,omitempty"`
		FormaJuridica             string               `xml:"formaJuridica,omitempty" json:"formaJuridica,omitempty"`
		IdPersona                 int64                `xml:"idPersona,omitempty" json:"idPersona,omitempty"`
		Impuesto                  []*ImpuestoDetallado `xml:"impuesto,omitempty" json:"impuesto,omitempty"`
		LeyJubilacion             int32                `xml:"leyJubilacion,omitempty" json:"leyJubilacion,omitempty"`
		LocalidadInscripcion      string               `xml:"localidadInscripcion,omitempty" json:"localidadInscripcion,omitempty"`
		MesCierre                 int32                `xml:"mesCierre,omitempty" json:"mesCierre,omitempty"`
		Nombre                    string               `xml:"nombre,omitempty" json:"nombre,omitempty"`
		NumeroDocumento           string               `xml:"numeroDocumento,omitempty" json:"numeroDocumento,omitempty"`
		NumeroInscripcion         int64                `xml:"numeroInscripcion,omitempty" json:"numeroInscripcion,omitempty"`
		OrganismoInscripcion      string               `xml:"organismoInscripcion,omitempty" json:"organismoInscripcion,omitempty"`
		OrganismoOriginante       string               `xml:"organismoOriginante,omitempty" json:"organismoOriginante,omitempty"`
		PorcentajeCapitalNacional float64              `xml:"porcentajeCapitalNacional,omitempty" json:"porcentajeCapitalNacional,omitempty"`
		ProvinciaInscripcion      string               `xml:"provinciaInscripcion,omitempty" json:"provinciaInscripcion,omitempty"`
		RazonSocial               string               `xml:"razonSocial,omitempty" json:"razonSocial,omitempty"`
		Regimen                   []*RegimenExtendido  `xml:"regimen,omitempty" json:"regimen,omitempty"`
		Relacion                  []*Relacion          `xml:"relacion,omitempty" json:"relacion,omitempty"`
		Sexo                      string               `xml:"sexo,omitempty" json:"sexo,omitempty"`
		Telefono                  []*Telefono          `xml:"telefono,omitempty" json:"telefono,omitempty"`
		TipoClave                 string               `xml:"tipoClave,omitempty" json:"tipoClave,omitempty"`
		TipoDocumento             string               `xml:"tipoDocumento,omitempty" json:"tipoDocumento,omitempty"`
		TipoOrganismoOriginante   string               `xml:"tipoOrganismoOriginante,omitempty" json:"tipoOrganismoOriginante,omitempty"`
		TipoPersona               string               `xml:"tipoPersona,omitempty" json:"tipoPersona,omitempty"`
		TipoResidencia            string               `xml:"tipoResidencia,omitempty" json:"tipoResidencia,omitempty"`
	}

	CitizenP5 struct {
		DatosGenerales      *DatosGenerales      `xml:"datosGenerales,omitempty" json:"datosGenerales,omitempty"`
		DatosMonotributo    *DatosMonotributo    `xml:"datosMonotributo,omitempty" json:"datosMonotributo,omitempty"`
		DatosRegimenGeneral *DatosRegimenGeneral `xml:"datosRegimenGeneral,omitempty" json:"datosRegimenGeneral,omitempty"`
		ErrorConstancia     *ErrorConstancia     `xml:"errorConstancia,omitempty" json:"errorConstancia,omitempty"`
		ErrorMonotributo    *ErrorMonotributo    `xml:"errorMonotributo,omitempty" json:"errorMonotributo,omitempty"`
		ErrorRegimenGeneral *ErrorRegimenGeneral `xml:"errorRegimenGeneral,omitempty" json:"errorRegimenGeneral,omitempty"`
	}

	CitizenP10 struct {
		Apellido                      string       `xml:"apellido,omitempty" json:"apellido,omitempty"`
		ClaveInactivaAsociada         []*int64     `xml:"claveInactivaAsociada,omitempty" json:"claveInactivaAsociada,omitempty"`
		Dependencia                   *Dependencia `xml:"dependencia,omitempty" json:"dependencia,omitempty"`
		DescripcionActividadPrincipal string       `xml:"descripcionActividadPrincipal,omitempty" json:"descripcionActividadPrincipal,omitempty"`
		Domicilio                     []*Domicilio `xml:"domicilio,omitempty" json:"domicilio,omitempty"`
		EstadoClave                   string       `xml:"estadoClave,omitempty" json:"estadoClave,omitempty"`
		IdActividadPrincipal          int64        `xml:"idActividadPrincipal,omitempty" json:"idActividadPrincipal,omitempty"`
		IdPersona                     int64        `xml:"idPersona,omitempty" json:"idPersona,omitempty"`
		Nombre                        string       `xml:"nombre,omitempty" json:"nombre,omitempty"`
		NumeroDocumento               string       `xml:"numeroDocumento,omitempty" json:"numeroDocumento,omitempty"`
		RazonSocial                   string       `xml:"razonSocial,omitempty" json:"razonSocial,omitempty"`
		TipoClave                     string       `xml:"tipoClave,omitempty" json:"tipoClave,omitempty"`
		TipoDocumento                 string       `xml:"tipoDocumento,omitempty" json:"tipoDocumento,omitempty"`
		TipoPersona                   string       `xml:"tipoPersona,omitempty" json:"tipoPersona,omitempty"`
	}

	CitizenP13 struct {
		Apellido                      string                `xml:"apellido,omitempty" json:"apellido,omitempty"`
		ClaveInactivaAsociada         []*int64              `xml:"claveInactivaAsociada,omitempty" json:"claveInactivaAsociada,omitempty"`
		DescripcionActividadPrincipal string                `xml:"descripcionActividadPrincipal,omitempty" json:"descripcionActividadPrincipal,omitempty"`
		Domicilio                     []*DomicilioExtendido `xml:"domicilio,omitempty" json:"domicilio,omitempty"`
		EstadoClave                   string                `xml:"estadoClave,omitempty" json:"estadoClave,omitempty"`
		FechaContratoSocial           time.Time             `xml:"fechaContratoSocial,omitempty" json:"fechaContratoSocial,omitempty"`
		FechaFallecimiento            time.Time             `xml:"fechaFallecimiento,omitempty" json:"fechaFallecimiento,omitempty"`
		FechaNacimiento               time.Time             `xml:"fechaNacimiento,omitempty" json:"fechaNacimiento,omitempty"`
		FormaJuridica                 string                `xml:"formaJuridica,omitempty" json:"formaJuridica,omitempty"`
		IdActividadPrincipal          int64                 `xml:"idActividadPrincipal,omitempty" json:"idActividadPrincipal,omitempty"`
		IdPersona                     int64                 `xml:"idPersona,omitempty" json:"idPersona,omitempty"`
		MesCierre                     int32                 `xml:"mesCierre,omitempty" json:"mesCierre,omitempty"`
		Nombre                        string                `xml:"nombre,omitempty" json:"nombre,omitempty"`
		NumeroDocumento               string                `xml:"numeroDocumento,omitempty" json:"numeroDocumento,omitempty"`
		PeriodoActividadPrincipal     int32                 `xml:"periodoActividadPrincipal,omitempty" json:"periodoActividadPrincipal,omitempty"`
		RazonSocial                   string                `xml:"razonSocial,omitempty" json:"razonSocial,omitempty"`
		TipoClave                     string                `xml:"tipoClave,omitempty" json:"tipoClave,omitempty"`
		TipoDocumento                 string                `xml:"tipoDocumento,omitempty" json:"tipoDocumento,omitempty"`
		TipoPersona                   string                `xml:"tipoPersona,omitempty" json:"tipoPersona,omitempty"`
	}

	ErrorConstancia struct {
		Apellido  string    `xml:"apellido,omitempty" json:"apellido,omitempty"`
		Error     []*string `xml:"error,omitempty" json:"error,omitempty"`
		IdPersona int64     `xml:"idPersona,omitempty" json:"idPersona,omitempty"`
		Nombre    string    `xml:"nombre,omitempty" json:"nombre,omitempty"`
	}

	ErrorMonotributo struct {
		Error   []*string `xml:"error,omitempty" json:"error,omitempty"`
		Mensaje string    `xml:"mensaje,omitempty" json:"mensaje,omitempty"`
	}

	ErrorRegimenGeneral struct {
		Error   []*string `xml:"error,omitempty" json:"error,omitempty"`
		Mensaje string    `xml:"mensaje,omitempty" json:"mensaje,omitempty"`
	}

	DatosGenerales struct {
		Apellido            string                `xml:"apellido,omitempty" json:"apellido,omitempty"`
		Caracterizacion     []*Caracterizacion    `xml:"caracterizacion,omitempty" json:"caracterizacion,omitempty"`
		Dependencia         *DependenciaDetallada `xml:"dependencia,omitempty" json:"dependencia,omitempty"`
		DomicilioFiscal     *Domicilio            `xml:"domicilioFiscal,omitempty" json:"domicilioFiscal,omitempty"`
		EsSucesion          string                `xml:"esSucesion,omitempty" json:"esSucesion,omitempty"`
		EstadoClave         string                `xml:"estadoClave,omitempty" json:"estadoClave,omitempty"`
		FechaContratoSocial time.Time             `xml:"fechaContratoSocial,omitempty" json:"fechaContratoSocial,omitempty"`
		FechaFallecimiento  time.Time             `xml:"fechaFallecimiento,omitempty" json:"fechaFallecimiento,omitempty"`
		IdPersona           int64                 `xml:"idPersona,omitempty" json:"idPersona,omitempty"`
		MesCierre           int32                 `xml:"mesCierre,omitempty" json:"mesCierre,omitempty"`
		Nombre              string                `xml:"nombre,omitempty" json:"nombre,omitempty"`
		RazonSocial         string                `xml:"razonSocial,omitempty" json:"razonSocial,omitempty"`
		TipoClave           string                `xml:"tipoClave,omitempty" json:"tipoClave,omitempty"`
		TipoPersona         string                `xml:"tipoPersona,omitempty" json:"tipoPersona,omitempty"`
	}

	DatosMonotributo struct {
		Actividad               []*Actividad      `xml:"actividad,omitempty" json:"actividad,omitempty"`
		ActividadMonotributista *Actividad        `xml:"actividadMonotributista,omitempty" json:"actividadMonotributista,omitempty"`
		CategoriaMonotributo    *CategoriaFiscal  `xml:"categoriaMonotributo,omitempty" json:"categoriaMonotributo,omitempty"`
		ComponenteDeSociedad    []*RelacionFiscal `xml:"componenteDeSociedad,omitempty" json:"componenteDeSociedad,omitempty"`
		Impuesto                []*Impuesto       `xml:"impuesto,omitempty" json:"impuesto,omitempty"`
	}

	DatosRegimenGeneral struct {
		Actividad         []*Actividad     `xml:"actividad,omitempty" json:"actividad,omitempty"`
		CategoriaAutonomo *CategoriaFiscal `xml:"categoriaAutonomo,omitempty" json:"categoriaAutonomo,omitempty"`
		Impuesto          []*Impuesto      `xml:"impuesto,omitempty" json:"impuesto,omitempty"`
		Regimen           []*Regimen       `xml:"regimen,omitempty" json:"regimen,omitempty"`
	}

	Caracterizacion struct {
		DescripcionCaracterizacion string `xml:"descripcionCaracterizacion,omitempty" json:"descripcionCaracterizacion,omitempty"`
		IdCaracterizacion          int32  `xml:"idCaracterizacion,omitempty" json:"idCaracterizacion,omitempty"`
		Periodo                    int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
	}

	Telefono struct {
		Numero       int64  `xml:"numero,omitempty" json:"numero,omitempty"`
		TipoLinea    string `xml:"tipoLinea,omitempty" json:"tipoLinea,omitempty"`
		TipoTelefono string `xml:"tipoTelefono,omitempty" json:"tipoTelefono,omitempty"`
	}

	Impuesto struct {
		DescripcionImpuesto string `xml:"descripcionImpuesto,omitempty" json:"descripcionImpuesto,omitempty"`
		EstadoImpuesto      string `xml:"estadoImpuesto,omitempty" json:"estadoImpuesto,omitempty"`
		IdImpuesto          int32  `xml:"idImpuesto,omitempty" json:"idImpuesto,omitempty"`
		Motivo              string `xml:"motivo,omitempty" json:"motivo,omitempty"`
		Periodo             int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
	}

	ImpuestoDetallado struct {
		DescripcionImpuesto string    `xml:"descripcionImpuesto,omitempty" json:"descripcionImpuesto,omitempty"`
		DiaPeriodo          int32     `xml:"diaPeriodo,omitempty" json:"diaPeriodo,omitempty"`
		Estado              string    `xml:"estado,omitempty" json:"estado,omitempty"`
		FfInscripcion       time.Time `xml:"ffInscripcion,omitempty" json:"ffInscripcion,omitempty"`
		IdImpuesto          int32     `xml:"idImpuesto,omitempty" json:"idImpuesto,omitempty"`
		Periodo             int32     `xml:"periodo,omitempty" json:"periodo,omitempty"`
	}

	Regimen struct {
		DescripcionRegimen string `xml:"descripcionRegimen,omitempty" json:"descripcionRegimen,omitempty"`
		IdImpuesto         int32  `xml:"idImpuesto,omitempty" json:"idImpuesto,omitempty"`
		IdRegimen          int32  `xml:"idRegimen,omitempty" json:"idRegimen,omitempty"`
		Periodo            int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
		TipoRegimen        string `xml:"tipoRegimen,omitempty" json:"tipoRegimen,omitempty"`
	}

	RegimenExtendido struct {
		DescripcionRegimen string `xml:"descripcionRegimen,omitempty" json:"descripcionRegimen,omitempty"`
		DiaPeriodo         int32  `xml:"diaPeriodo,omitempty" json:"diaPeriodo,omitempty"`
		Estado             string `xml:"estado,omitempty" json:"estado,omitempty"`
		IdImpuesto         int32  `xml:"idImpuesto,omitempty" json:"idImpuesto,omitempty"`
		IdRegimen          int32  `xml:"idRegimen,omitempty" json:"idRegimen,omitempty"`
		Periodo            int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
		TipoRegimen        string `xml:"tipoRegimen,omitempty" json:"tipoRegimen,omitempty"`
	}

	RelacionFiscal struct {
		ApellidoPersonaAsociada    string    `xml:"apellidoPersonaAsociada,omitempty" json:"apellidoPersonaAsociada,omitempty"`
		FfRelacion                 time.Time `xml:"ffRelacion,omitempty" json:"ffRelacion,omitempty"`
		FfVencimiento              time.Time `xml:"ffVencimiento,omitempty" json:"ffVencimiento,omitempty"`
		IdPersonaAsociada          int64     `xml:"idPersonaAsociada,omitempty" json:"idPersonaAsociada,omitempty"`
		NombrePersonaAsociada      string    `xml:"nombrePersonaAsociada,omitempty" json:"nombrePersonaAsociada,omitempty"`
		RazonSocialPersonaAsociada string    `xml:"razonSocialPersonaAsociada,omitempty" json:"razonSocialPersonaAsociada,omitempty"`
		TipoComponente             string    `xml:"tipoComponente,omitempty" json:"tipoComponente,omitempty"`
	}

	Relacion struct {
		FfRelacion        time.Time `xml:"ffRelacion,omitempty" json:"ffRelacion,omitempty"`
		FfVencimiento     time.Time `xml:"ffVencimiento,omitempty" json:"ffVencimiento,omitempty"`
		IdPersona         int64     `xml:"idPersona,omitempty" json:"idPersona,omitempty"`
		IdPersonaAsociada int64     `xml:"idPersonaAsociada,omitempty" json:"idPersonaAsociada,omitempty"`
		SubtipoRelacion   string    `xml:"subtipoRelacion,omitempty" json:"subtipoRelacion,omitempty"`
		TipoRelacion      string    `xml:"tipoRelacion,omitempty" json:"tipoRelacion,omitempty"`
	}

	CategoriaFiscal struct {
		DescripcionCategoria string `xml:"descripcionCategoria,omitempty" json:"descripcionCategoria,omitempty"`
		IdCategoria          int32  `xml:"idCategoria,omitempty" json:"idCategoria,omitempty"`
		IdImpuesto           int32  `xml:"idImpuesto,omitempty" json:"idImpuesto,omitempty"`
		Periodo              int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
	}

	Categoria struct {
		DescripcionCategoria string `xml:"descripcionCategoria,omitempty" json:"descripcionCategoria,omitempty"`
		Estado               string `xml:"estado,omitempty" json:"estado,omitempty"`
		IdCategoria          int32  `xml:"idCategoria,omitempty" json:"idCategoria,omitempty"`
		IdImpuesto           int32  `xml:"idImpuesto,omitempty" json:"idImpuesto,omitempty"`
		Periodo              int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
	}

	Email struct {
		Direccion string `xml:"direccion,omitempty" json:"direccion,omitempty"`
		Estado    string `xml:"estado,omitempty" json:"estado,omitempty"`
		TipoEmail string `xml:"tipoEmail,omitempty" json:"tipoEmail,omitempty"`
	}

	Actividad struct {
		DescripcionActividad string `xml:"descripcionActividad,omitempty" json:"descripcionActividad,omitempty"`
		IdActividad          int64  `xml:"idActividad,omitempty" json:"idActividad,omitempty"`
		Nomenclador          int32  `xml:"nomenclador,omitempty" json:"nomenclador,omitempty"`
		Orden                int32  `xml:"orden,omitempty" json:"orden,omitempty"`
		Periodo              int32  `xml:"periodo,omitempty" json:"periodo,omitempty"`
	}

	Dependencia struct {
		DescripcionDependencia string `xml:"descripcionDependencia,omitempty" json:"descripcionDependencia,omitempty"`
		IdDependencia          int32  `xml:"idDependencia,omitempty" json:"idDependencia,omitempty"`
	}

	DependenciaDetallada struct {
		CodPostal              string `xml:"codPostal,omitempty" json:"codPostal,omitempty"`
		DescripcionDependencia string `xml:"descripcionDependencia,omitempty" json:"descripcionDependencia,omitempty"`
		DescripcionProvincia   string `xml:"descripcionProvincia,omitempty" json:"descripcionProvincia,omitempty"`
		Direccion              string `xml:"direccion,omitempty" json:"direccion,omitempty"`
		IdDependencia          int32  `xml:"idDependencia,omitempty" json:"idDependencia,omitempty"`
		IdProvincia            int32  `xml:"idProvincia,omitempty" json:"idProvincia,omitempty"`
		Localidad              string `xml:"localidad,omitempty" json:"localidad,omitempty"`
	}

	Domicilio struct {
		CodPostal            string `xml:"codPostal,omitempty" json:"codPostal,omitempty"`
		DatoAdicional        string `xml:"datoAdicional,omitempty" json:"datoAdicional,omitempty"`
		DescripcionProvincia string `xml:"descripcionProvincia,omitempty" json:"descripcionProvincia,omitempty"`
		Direccion            string `xml:"direccion,omitempty" json:"direccion,omitempty"`
		IdProvincia          int32  `xml:"idProvincia,omitempty" json:"idProvincia,omitempty"`
		Localidad            string `xml:"localidad,omitempty" json:"localidad,omitempty"`
		TipoDatoAdicional    string `xml:"tipoDatoAdicional,omitempty" json:"tipoDatoAdicional,omitempty"`
		TipoDomicilio        string `xml:"tipoDomicilio,omitempty" json:"tipoDomicilio,omitempty"`
	}

	DomicilioExtendido struct {
		Calle                string `xml:"calle,omitempty" json:"calle,omitempty"`
		CodigoPostal         string `xml:"codigoPostal,omitempty" json:"codigoPostal,omitempty"`
		DatoAdicional        string `xml:"datoAdicional,omitempty" json:"datoAdicional,omitempty"`
		DescripcionProvincia string `xml:"descripcionProvincia,omitempty" json:"descripcionProvincia,omitempty"`
		Direccion            string `xml:"direccion,omitempty" json:"direccion,omitempty"`
		EstadoDomicilio      string `xml:"estadoDomicilio,omitempty" json:"estadoDomicilio,omitempty"`
		IdProvincia          int32  `xml:"idProvincia,omitempty" json:"idProvincia,omitempty"`
		Localidad            string `xml:"localidad,omitempty" json:"localidad,omitempty"`
		Manzana              string `xml:"manzana,omitempty" json:"manzana,omitempty"`
		Numero               int32  `xml:"numero,omitempty" json:"numero,omitempty"`
		OficinaDptoLocal     string `xml:"oficinaDptoLocal,omitempty" json:"oficinaDptoLocal,omitempty"`
		Piso                 string `xml:"piso,omitempty" json:"piso,omitempty"`
		Sector               string `xml:"sector,omitempty" json:"sector,omitempty"`
		TipoDatoAdicional    string `xml:"tipoDatoAdicional,omitempty" json:"tipoDatoAdicional,omitempty"`
		TipoDomicilio        string `xml:"tipoDomicilio,omitempty" json:"tipoDomicilio,omitempty"`
		Torre                string `xml:"torre,omitempty" json:"torre,omitempty"`
	}
)
