package capacitor

type Material struct {
	name       string
	cud        float64
	elStrength float64
	e          float64
	tgdelta    float64
	tke        float64
}

var Materials = []Material{
	{
		name:       "Моноокись кремния", // 0
		cud:        100,
		elStrength: 3,
		e:          6,
		tgdelta:    0.02,
		tke:        2,
	},
	{
		name:       "Моноокись германия", // 1
		cud:        150,
		elStrength: 1,
		e:          12,
		tgdelta:    0.007,
		tke:        3,
	},
	{
		name:       "Трехсернистая сурьма", // 2
		cud:        150,
		elStrength: 0.5,
		e:          21,
		tgdelta:    0.01,
		tke:        5,
	},
	{
		name:       "Халькогенидное стекло", // 3
		cud:        150,
		elStrength: 2.8,
		e:          10,
		tgdelta:    0.02,
		tke:        28,
	},
	{
		name:       "Халькогенидное стекло2", // 4
		cud:        100,
		elStrength: 0.4,
		e:          6,
		tgdelta:    0.02,
		tke:        9,
	},
	{
		name:       "Боросиликатное стекло(140)", // 5
		cud:        140,
		elStrength: 4,
		e:          4,
		tgdelta:    0.0015,
		tke:        0.35,
	},
	{
		name:       "Стекло электровакуумное", // 6
		cud:        400,
		elStrength: 4,
		e:          5.2,
		tgdelta:    0.003,
		tke:        1.8,
	},
	{
		name:       "Пятиокись тантала(1000) WARNING", // 7
		cud:        1000,
		elStrength: 2,
		e:          23,
		tgdelta:    0.02,
		tke:        4,
	},
	{
		name:       "Окись алюминия", // 8
		cud:        850,
		elStrength: 9,
		e:          10,
		tgdelta:    0.015,
		tke:        5,
	},
}

func GetHollowMaterial() Material {
	return Material{
		name:       "Hollow",
		cud:        0,
		elStrength: 0,
		e:          0,
		tgdelta:    0,
		tke:        0,
	}
}

func SetMaterialsForCapacitors(arrOfRes []Capacitor, mat Material) {
	for i := range arrOfRes {
		(arrOfRes)[i].SetMaterial(mat)
	}
}

// Getters

func (m *Material) GetName() string {
	return m.name
}

func (m *Material) GetCud() float64 {
	return m.cud
}

func (m *Material) GetElStrength() float64 {
	return m.elStrength
}

func (m *Material) Gete() float64 {
	return m.e
}

func (m *Material) Getthdelta() float64 {
	return m.tgdelta
}

func (m *Material) GetTKE() float64 {
	return m.tke
}
