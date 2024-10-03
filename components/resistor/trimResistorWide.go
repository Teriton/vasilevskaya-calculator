package resistor

import "fmt"

type TrimWide struct {
	n        float64
	bo       float64
	bn       float64
	rdashmin float64
	deltar   float64
	deltaR   float64
	ir       []float64
	iR       []float64
	deltabi  []float64
	breg     float64
}

func (r *Resistor) InitTrimWide() {
	r.trimWide.n = r.GetMOfTrimL()
	r.trimWide.bo = TehnRound(CalculateBo(r.Getrokvmax(), r.Getlmax(), r.Getrmax()))
	r.trimWide.rdashmin = CalculateRdashmin(r.Getrokvmin(), r.Getlmin(), r.trimWide.bo)
	r.trimWide.deltar = CalculateDeltarTrimWide(r.Getrmin(), r.GetrdashminTrimW())
	r.trimWide.bn = CalculateBn(r.Getrokvmin(), r.Getlmin(), r.Getrmin())
	r.trimWide.deltaR = CalculateDeltaRTrimWide(r.GetdeltarTrimW(), r.GetnTrimW())
	for i := range int(r.GetnTrimW()) {
		Ril := CalculateRi(r.GetrdashminTrimW(), i, r.GetdeltaRTrimW())
		ril := Calculateri(Ril, r.GetdeltaRTrimW())
		deltabil := CalculateDeltaBi(r.Getrokvmin(), r.Getlmin(), ril)
		*r.GetiRTrimW() = append(*r.GetiRTrimW(), Ril)
		*r.GetirTrimW() = append(*r.GetirTrimW(), ril)
		*r.GetdeltabiTrimW() = append(*r.GetdeltabiTrimW(), deltabil)
	}
	for i := range int(r.GetnTrimW()) {
		fmt.Println("ri ", i+1, ": ", (*r.GetirTrimW())[i])
		fmt.Println("Ri ", i+1, ": ", (*r.GetiRTrimW())[i])
		fmt.Println("deltaBi  ", i+1, ": ", (*r.GetdeltabiTrimW())[i])
		r.trimWide.breg += (*r.GetdeltabiTrimW())[i]
	}
	// fmt.Println("bmax: ", bmax)
	// fmt.Println("bmin: ", bmin)
	// fmt.Println("lmin: ", lmin)
	// fmt.Println("lmax: ", lmax)
	// fmt.Println("rokvmax: ", rokvmax)
	// fmt.Println("rokvmin: ", rokvmin)
	// fmt.Println("rmax: ", rmax)
	// fmt.Println("rmin: ", rmin)
	// fmt.Println("R'min: ", rdashmin)
	// fmt.Println("deltar: ", deltartrim)
	// fmt.Println("deltaR: ", deltaRtrim)
	// fmt.Println("bn ", bn)
	// fmt.Println("bo ", bo)
	// fmt.Println("sum ", sum)
	// fmt.Println("bo - bn ", bo-bn)
	// fmt.Println("Длинна ", r.GetHeight())
	// fmt.Println("Ширина ", r.GetWidth())
}

func CalculateBo(rokvmax float64, lmax float64, rmax float64) float64 {
	return (rokvmax * lmax) / rmax
}

func CalculateDeltarTrimWide(rmin float64, rdashmin float64) float64 {
	return rmin - rdashmin
}

func CalculateBn(rokvmin float64, lmin float64, rmin float64) float64 {
	return (rokvmin * lmin) / rmin
}

func CalculateDeltaRTrimWide(deltar float64, m float64) float64 {
	return deltar / m
}

func CalculateRi(rdashmin float64, n int, deltaR float64) float64 {
	return rdashmin + float64(n)*deltaR
}

func Calculateri(Ri float64, deltaR float64) float64 {
	return Ri * (1 + (Ri / deltaR))
}

func CalculateDeltaBi(rokvmin float64, lmin float64, ri float64) float64 {
	return (rokvmin * lmin) / ri
}

// Getters

func (r *Resistor) GetnTrimW() float64 {
	return r.trimWide.n
}
func (r *Resistor) GetboTrimW() float64 {
	return r.trimWide.bo
}

func (r *Resistor) GetbnTrimW() float64 {
	return r.trimWide.bn
}

func (r *Resistor) GetrdashminTrimW() float64 {
	return r.trimWide.rdashmin
}

func (r *Resistor) GetdeltarTrimW() float64 {
	return r.trimWide.deltar
}

func (r *Resistor) GetdeltaRTrimW() float64 {
	return r.trimWide.deltaR
}

func (r *Resistor) GetirTrimW() *[]float64 {
	return &r.trimWide.ir
}

func (r *Resistor) GetiRTrimW() *[]float64 {
	return &r.trimWide.iR
}

func (r *Resistor) GetdeltabiTrimW() *[]float64 {
	return &r.trimWide.deltabi
}

func (r *Resistor) GetbregTrimW() float64 {
	return r.trimWide.breg
}
