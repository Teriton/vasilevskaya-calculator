package resistorcontrollerIml

import (
    "github.com/shpakunya/pkg/components"
)

func trimWideInit(resistor *components.Resistor) {
    // N
    resistor.TrimWide.N = components.Parameter{
        Name:   "M of Trim Wide",
        Symbol: "n",
        Unit:   "",
        Value:  resistor.TrimLength.MOfTrim.Value, // или свой расчёт
    }

    // Bo
    resistor.TrimWide.Bo = components.Parameter{
        Name:   "Bo",
        Symbol: "bo",
        Unit:   "мм",
        Value:  tehnRound(calculateBo(resistor.Precision.Rokvmax.Value, resistor.Precision.Lmax.Value, resistor.Precision.Rmax.Value), 0.01),
    }

    // Rdashmin
    resistor.TrimWide.Rdashmin = components.Parameter{
        Name:   "R' min Trim Wide",
        Symbol: "r'dashmin",
        Unit:   "Ом",
        Value:  calculateRdashmin(resistor.Precision.Rokvmin.Value, resistor.Precision.Lmin.Value, resistor.TrimWide.Bo.Value),
    }

    // Deltar
    resistor.TrimWide.Deltar = components.Parameter{
        Name:   "Δr Trim Wide",
        Symbol: "deltar",
        Unit:   "Ом",
        Value:  calculateDeltarTrimWide(resistor.Precision.Rmin.Value, resistor.TrimWide.Rdashmin.Value),
    }

    // Bn
    resistor.TrimWide.Bn = components.Parameter{
        Name:   "Bn",
        Symbol: "bn",
        Unit:   "мм",
        Value:  calculateBn(resistor.Precision.Rokvmin.Value, resistor.Precision.Lmin.Value, resistor.Precision.Rmin.Value),
    }

    // DeltaR
    resistor.TrimWide.DeltaR = components.Parameter{
        Name:   "ΔR Trim Wide",
        Symbol: "deltaR",
        Unit:   "Ом",
        Value:  calculateDeltaRTrimWide(resistor.TrimWide.Deltar.Value, resistor.TrimWide.N.Value),
    }

    // Очистка массивов
    resistor.TrimWide.Ir = []components.Parameter{}
    resistor.TrimWide.IR = []components.Parameter{}
    resistor.TrimWide.Deltabi = []components.Parameter{}

    // Заполнение массивов
    for i := 0; i < int(resistor.TrimWide.N.Value); i++ {
        ri := calculateRi(resistor.TrimWide.Rdashmin.Value, float64(i), resistor.TrimWide.DeltaR.Value)
        ril := calculateri(ri, resistor.TrimWide.DeltaR.Value)
        deltabi := calculateDeltaBi(resistor.Precision.Rokvmin.Value, resistor.Precision.Lmin.Value, ril)

        resistor.TrimWide.IR = append(resistor.TrimWide.IR, components.Parameter{
            Name:   "Ri",
            Symbol: "Ri",
            Unit:   "Ом",
            Value:  ri,
        })
        resistor.TrimWide.Ir = append(resistor.TrimWide.Ir, components.Parameter{
            Name:   "ri",
            Symbol: "ri",
            Unit:   "Ом",
            Value:  ril,
        })
        resistor.TrimWide.Deltabi = append(resistor.TrimWide.Deltabi, components.Parameter{
            Name:   "Δbi",
            Symbol: "Δbi",
            Unit:   "мм",
            Value:  deltabi,
        })
    }

    // Суммируем breg
    var bregValue float64
    for i := range resistor.TrimWide.Deltabi {
        bregValue += resistor.TrimWide.Deltabi[i].Value
    }

    resistor.TrimWide.Breg = components.Parameter{
        Name:   "Breg",
        Symbol: "breg",
        Unit:   "мм",
        Value:  bregValue,
    }
}

// ...existing code...

func calculateBo(rokvmax, lmax, rmax float64) float64 {
    if rmax == 0 {
        return 0
    }
    return (rokvmax * lmax) / rmax
}

func calculateWideRdashmin(rokvmin, lmin, bo float64) float64 {
    if bo == 0 {
        return 0
    }
    return (rokvmin * lmin) / bo
}

func calculateDeltarTrimWide(rmin, rdashmin float64) float64 {
    return rmin - rdashmin
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

// func tehnRound(value, step float64) float64 {
//     if step == 0 {
//         return value
//     }
//     return math.Round(value/step) * step
// }