package components

type Resistor struct {
	Resistance Parameter
	Tolerance  Parameter
	Power      Parameter
	Material   ResistorMaterial
	Env        Environment

	//Main parametrs
	FormFactor     Parameter
	GammaRt        Parameter
	GammaRdelta    Parameter
	FormOfResistor Form

	Rectangle Rectangle

	Precision Precision

	Meander Meander

	Ccp CCP
	// Trim
	trimLength TrimLength
	trimWide   TrimWide
}

type Form string

const (
	RectangleForm Form = "Прямоугольник"
	MeanderForm   Form = "Меандр"
	CCPForm       Form = "ЦКП"
)

type ResistorMaterial struct {
	Name                                ParameterMain[string]
	SquareResistance                    Parameter
	PermissibleSpecificPowerDissipation Parameter
	TemperatureCoefficientOfResistance  Parameter
	Senescence                          Parameter
}

type Precision struct {
	Bmin, Bmax       Parameter
	Lmin, Lmax       Parameter
	Rokvmin, Rokvmax Parameter
	Rmin, Rmax       Parameter
}

type Rectangle struct {
	Bp     Parameter
	Bdelta Parameter
	Width  Parameter
	Lp     Parameter
	Ldelta Parameter
	Height Parameter
}

type Meander struct {
	numberOfLinks  Parameter
	meanderXLength Parameter
	meanderYLength Parameter
	meanderArea    Parameter
}

type CCP struct {
	bpCCP     Parameter
	bdeltaCCP Parameter
}

type TrimWide struct {
	N        Parameter
	Bo       Parameter
	Bn       Parameter
	Rdashmin Parameter
	Deltar   Parameter
	DeltaR   Parameter
	Ir       []Parameter
	IR       []Parameter
	Deltabi  []Parameter
	Breg     Parameter
}

type TrimLength struct {
	GammaRdeltaTrim Parameter
	GammaR          Parameter
	MOfTrim         Parameter
	LnTrim          Parameter
	LoTrim          Parameter
	RdashminTrim    Parameter
	Ltune           Parameter
	DeltarTrim      Parameter
	DeltaLrTrim     Parameter
	DeltaLdashTrim  Parameter
	DeltaRTrim      Parameter
	Lpodg           Parameter
	Lsum            Parameter
	Gammakf         Parameter
	LooTrim         Parameter
}
