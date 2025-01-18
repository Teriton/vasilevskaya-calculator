package resistorcontrollerIml

import (
    "math"
	"github.com/shpakunya/pkg/components"
	"strconv"
)

func trimWideInit(resistor *components.Resistor) {

	resistor.TrimWide.N = components.Parameter{
		Name:   "Количество подгоночных секций",
		Symbol: "m",
		Unit:   "",
		Value:  math.Ceil(calculateGammaR(resistor.Material.Senescence.Value, resistor.Env.GammaRokv.Value,
            resistor.Env.GammaRcontact.Value, calculateGammaRdeltaTrim(resistor.FormFactor.Value, resistor.Rectangle.Width.Value, resistor.Env.Deltab.Value), resistor.GammaRt.Value) / resistor.Tolerance.Value),
	}

	resistor.TrimWide.Bo = components.Parameter{
		Name:   "Максимальная возможная ширина резистора",
		Symbol: "b₀",
		Unit:   "мм",
		Value:  tehnRound(calculateBo(resistor.Precision.Rokvmax.Value, resistor.Precision.Lmax.Value, resistor.Precision.Rmax.Value), 0.1),
	}

	resistor.TrimWide.Rdashmin = components.Parameter{
		Name:   "Минимально возможное сопротивление резистора",
		Symbol: "R'min",
		Unit:   "Ом",
		Value:  calculateRdashmin(resistor.Precision.Rokvmin.Value, resistor.Precision.Lmin.Value, resistor.TrimWide.Bo.Value),
	}

	resistor.TrimWide.Deltar = components.Parameter{
		Name:   "Величина сопротивления, которую необходимо скомпенсировать",
		Symbol: "Δr",
		Unit:   "Ом",
		Value:  calculateDeltarTrim(resistor.Precision.Rmin.Value, resistor.TrimWide.Rdashmin.Value),
	}

	resistor.TrimWide.Bn = components.Parameter{
		Name:   "Ширина нерегулируеыой части подгоняемого резистора",
		Symbol: "bₙ",
		Unit:   "мм",
		Value:  calculateBn(resistor.Precision.Rokvmin.Value, resistor.Precision.Lmin.Value, resistor.Precision.Rmin.Value),
	}

	resistor.TrimWide.DeltaR = components.Parameter{
		Name:   "Величина сопротивления, которую необходимо скомпенсировать на одну секцию",
		Symbol: "ΔR",
		Unit:   "Ом",
		Value:  calculateDeltaRTrimWide(resistor.TrimWide.Deltar.Value, resistor.TrimWide.N.Value),
	}

	resistor.TrimWide.Ir = []components.Parameter{}
	resistor.TrimWide.IR = []components.Parameter{}
	resistor.TrimWide.Deltabi = []components.Parameter{}

	for i := 0; i < int(resistor.TrimWide.N.Value); i++ {
		ri := calculateRi(resistor.TrimWide.Rdashmin.Value, float64(i), resistor.TrimWide.DeltaR.Value)
		ril := calculateri(ri, resistor.TrimWide.DeltaR.Value)
		deltabi := calculateDeltaBi(resistor.Precision.Rokvmin.Value, resistor.Precision.Lmin.Value, ril)

		resistor.TrimWide.IR = append(resistor.TrimWide.IR, components.Parameter{
			Name:   "Сопротивление резистора перед удалением " + strconv.Itoa(i+1) + " секции",
			Symbol: "IR",
			Unit:   "Ом",
			Value:  ri,
		})
		resistor.TrimWide.Ir = append(resistor.TrimWide.Ir, components.Parameter{
			Name:   "Сопротивление " + strconv.Itoa(i+1) + " подгоночной секции",
			Symbol: "Ir",
			Unit:   "Ом",
			Value:  ril,
		})
		resistor.TrimWide.Deltabi = append(resistor.TrimWide.Deltabi, components.Parameter{
			Name:   "Ширина " + strconv.Itoa(i+1) + " секции подгонки",
			Symbol: "Δbi",
			Unit:   "мм",
			Value:  deltabi,
		})
	}

	var bregValue float64
	for i := range resistor.TrimWide.Deltabi {
		bregValue += resistor.TrimWide.Deltabi[i].Value
	}

	resistor.TrimWide.Breg = components.Parameter{
		Name:   "Проверка рассчетов для подгонки",
		Symbol: "Breg(b₀ - bₙ)",
		Unit:   "мм",
		Value:  bregValue,
	}
}

func calculateBo(rokvmax, lmax, rmax float64) float64 {
	if rmax == 0 {
		return 0
	}
	return (rokvmax * lmax) / rmax
}

func calculateBn(rokvmin, lmin, rmin float64) float64 {
	if rmin == 0 {
		return 0
	}
	return (rokvmin * lmin) / rmin
}

func calculateDeltaRTrimWide(deltar, m float64) float64 {
	if m == 0 {
		return 0
	}
	return deltar / m
}

func calculateRi(rdashmin float64, n float64, deltaR float64) float64 {
	return rdashmin + n*deltaR
}

func calculateri(Ri float64, deltaR float64) float64 {
	if deltaR == 0 {
		return 0
	}
	return Ri * (1 + (Ri / deltaR))
}

func calculateDeltaBi(rokvmin, lmin, ri float64) float64 {
	if ri == 0 {
		return 0
	}
	return (rokvmin * lmin) / ri
}
