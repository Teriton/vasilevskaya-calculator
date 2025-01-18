package resistorcontrollerIml

import (
    "math"

    "github.com/shpakunya/pkg/components"
)

func initTrimLength(resistor *components.Resistor) {
    
    resistor.TrimLength.GammaRdeltaTrim = components.Parameter{
        Name:   "Погрешность воспроизведения геометрических размеров резистора",
        Symbol: "γRΔ",
        Unit:   "%",
        Value:  calculateGammaRdeltaTrim(resistor.FormFactor.Value, resistor.Rectangle.Width.Value, resistor.Env.Deltab.Value),
    }

    resistor.TrimLength.GammaR = components.Parameter{
        Name:   "Относительная погрешность сопротивления резистора",
        Symbol: "γR",
        Unit:   "%",
        Value:  calculateGammaR(resistor.Material.Senescence.Value, resistor.Env.GammaRokv.Value,
            resistor.Env.GammaRcontact.Value, resistor.TrimLength.GammaRdeltaTrim.Value, resistor.GammaRt.Value),
    }

    resistor.TrimLength.MOfTrim = components.Parameter{
        Name:   "Количество подгоночных секций",
        Symbol: "m",
        Unit:   "",
        Value:  math.Ceil(resistor.TrimLength.GammaR.Value / resistor.Tolerance.Value),
    }

    resistor.TrimLength.LnTrim = components.Parameter{
        Name:   "Длина нерегулируемого участка резистора",
        Symbol: "lₙ",
        Unit:   "мм",
        Value:  calculateLn(resistor.Precision.Rmax.Value, resistor.Precision.Bmin.Value,
            resistor.Precision.Rokvmax.Value, resistor.Env.Deltal.Value),
    }

    resistor.TrimLength.RdashminTrim = components.Parameter{
        Name:   "Минимально возможное сопротивление резистора",
        Symbol: "R'min",
        Unit:   "Ом",
        Value:  calculateRdashmin(resistor.Precision.Rokvmin.Value, resistor.TrimLength.LnTrim.Value,
            resistor.Precision.Bmax.Value),
    }

    resistor.TrimLength.DeltarTrim = components.Parameter{
        Name:   "Величина сопротивления, которую необходимо скомпенсировать",
        Symbol: "Δr",
        Unit:   "Ом",
        Value:  calculateDeltarTrim(resistor.Precision.Rmin.Value, resistor.TrimLength.RdashminTrim.Value),
    }

    resistor.TrimLength.LoTrim = components.Parameter{
        Name:   "Общая длина резистора без проводящих перемычек",
        Symbol: "l'₀",
        Unit:   "мм",
        Value:  calculateLo(resistor.Precision.Rmin.Value, resistor.Precision.Bmax.Value,
            resistor.Precision.Rokvmin.Value),
    }

    resistor.TrimLength.Ltune = components.Parameter{
        Name:   "Длина регулируемой части",
        Symbol: "ltune(lp)",
        Unit:   "мм",
        Value:  calculateLtune(resistor.TrimLength.LoTrim.Value, resistor.TrimLength.LnTrim.Value),
    }

    resistor.TrimLength.DeltaLrTrim = components.Parameter{
        Name:   "Длина одной подгоночной секции",
        Symbol: "Δlp",
        Unit:   "мм",
        Value:  calculateDeltaLr(resistor.TrimLength.Ltune.Value, resistor.TrimLength.MOfTrim.Value),
    }

    resistor.TrimLength.DeltaLdashTrim = components.Parameter{
        Name:   "Проверка для подгонки",
        Symbol: "Δ'l(Δ'r)",
        Unit:   "мм",
        Value:  calculateDeltaLdash(resistor.Precision.Rokvmin.Value, resistor.TrimLength.DeltaLrTrim.Value,
            resistor.Precision.Bmax.Value),
    }

    resistor.TrimLength.DeltaRTrim = components.Parameter{
        Name:   "Величина сопротивления, которую необходимо скомпенсировать на одну секцию",
        Symbol: "ΔR",
        Unit:   "Ом",
        Value:  calculateDeltaR(resistor.TrimLength.DeltarTrim.Value, resistor.TrimLength.MOfTrim.Value),
    }

    resistor.TrimLength.Lpodg = components.Parameter{
        Name:   "Длина одной подгоночной секции",
        Symbol: "lподг",
        Unit:   "мм",
        Value:  tehnRound(calculateLpodg(resistor.TrimLength.Ltune.Value, resistor.TrimLength.MOfTrim.Value), 0.1),
    }

//совпадение не найдено
    resistor.TrimLength.Lsum = components.Parameter{
        Name:   "Общая длина резистора",
        Symbol: "l₀",
        Unit:   "мм",
        Value:  calculateLsum(resistor.TrimLength.LnTrim.Value, resistor.TrimLength.Lpodg.Value,
            resistor.TrimLength.MOfTrim.Value),
    }

//совпадение не найдено
    resistor.TrimLength.Gammakf = components.Parameter{
        Name:   "Gamma KF",
        Symbol: "γKF",
        Unit:   "",
        Value:  calculateGammaKf(resistor.Rectangle.Height.Value, resistor.Rectangle.Width.Value,
            resistor.Env.Deltal.Value, resistor.Env.Deltab.Value),
    }

//совпадение не найдено
    resistor.TrimLength.LooTrim = components.Parameter{
        Name:   "Loo Trim",
        Symbol: "looTrim",
        Unit:   "мм",
        Value:  calculateLooTrim(resistor.TrimLength.LnTrim.Value, resistor.TrimLength.Ltune.Value,
            0.1, resistor.TrimLength.MOfTrim.Value),
    }
}

func calculateGammaR(senescence, gammaRokv, gammaRcontact, gammaRdeltaTrim, gammaRt float64) float64 {
    return senescence + gammaRokv + gammaRcontact + gammaRdeltaTrim + gammaRt
}

func calculateGammaRdeltaTrim(formFactor, b, deltab float64) float64 {
    return 2 * ((1 + (deltab / (formFactor * deltab))) / ((b / deltab) - (deltab / b))) * 100
}

func calculateLn(rmax, bmin, rokvmax, deltal float64) float64 {
    return ((rmax * bmin) / rokvmax) - deltal
}

func calculateRdashmin(rokvmin, lmin, bmax float64) float64 {
    return (rokvmin * lmin) / bmax
}

func calculateLo(rmin, bmax, rokvmin float64) float64 {
    return (rmin * bmax) / rokvmin
}

func calculateLtune(lo, ln float64) float64 {
    return lo - ln
}

func calculateDeltarTrim(rmin, rdashmin float64) float64 {
    return rmin - rdashmin
}
//!!! одинаковые функции
func calculateDeltaLr(ltune, m float64) float64 {
    return ltune / m
}
//!!! одинаковые функции
func calculateLpodg(ltune, m float64) float64 {
    return ltune / m
}

func calculateDeltaLdash(rokvmin, deltaLr, bmax float64) float64 {
    return (rokvmin * deltaLr) / bmax
}

func calculateDeltaR(deltaR, m float64) float64 {
    return deltaR / m
}

func calculateLsum(ln, lpodg, m float64) float64 {
    return ln + lpodg*m
}

func calculateGammaKf(l, b, deltal, deltab float64) float64 {
    return (deltal/l + deltab/b) * 100
}

func calculateLooTrim(ln, lreg, deltaltehn, m float64) float64 {
    return ln + lreg + deltaltehn*m
}