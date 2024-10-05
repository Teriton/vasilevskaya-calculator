# vasilevskaya-calculator

### Пример запроса по /arrOfRes
```json
{
    "temperature":50.0,
    "material": 18,
    "res":
    [
        {
            "resistance": "1200",
            "power": "0.150",
            "tolerance": "20"
        },
        {
            "resistance": "2200",
            "power": "0.050",
            "tolerance": "10"
        },
        {
            "resistance": "6000",
            "power": "0.040",
            "tolerance": "10"
        },
            {
            "resistance": "1000",
            "power": "0.010",
            "tolerance": "20"
        },
            {
            "resistance": "120000",
            "power": "0.051",
            "tolerance": "10"
        },
            {
            "resistance": "130000",
            "power": "0.090",
            "tolerance": "2"
        },
            {
            "resistance": "1000000",
            "power": "0.012",
            "tolerance": "10"
        },
            {
            "resistance": "500000",
            "power": "0.032",
            "tolerance": "20"
        },
            {
            "resistance": "12000",
            "power": "0.045",
            "tolerance": "20"
        },
            {
            "resistance": "667000",
            "power": "0.123",
            "tolerance": "10"
        }
    ]
}

```

### Получение материалов резисторов по /resistorMaterials

### Пример запроса по /arrOfCaps
```json
{
    "temperature": 70.0,
    "material": 0,
    "cap":
    [
        {
            "capacity": "4500",
            "urab": "6.0",
            "tolerance": "20"
        },
        {
            "capacity": "400",
            "urab": "6.0",
            "tolerance": "20"
        },
        {
            "capacity": "1700",
            "urab": "6.0",
            "tolerance": "20"
        }
    ]
}
```

### Получение материалов конденсаторов по /capacitorMaterials