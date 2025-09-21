# Factory Pattern
![](https://mermaid.ink/img/pako:eNrNVt2O20QUfhVrpEppSaJNtnGzVlmppNuqEiDUrbiAcDGxx1mr9kw0HpeGZaXugkDQCyQuueKCB1hWrBqWbXgF-404M57YY8dJWwkkcpF45vzM933nzHGOkcs8ghzkhjiO7wd4ynE0phZ81I41CgNChXWc78nPexEOaOumsXGEqReSEeYjTrAIGG3B8wPsCsbnbQuen8xnxAzwSMRoLDgWMqr1CL7aFqwDOtVuJ2NqopAeJoa7dwMqCPexS_b3jcQPifiAA5rWTZ2uavsIuIabbPj54YwQGQq5q7YDOg0okSyagg8F5qLJcM91SUgky-YwNqvt11jXSHcmktpaok4kWTVsa0JVOh1ScGmU4X8kX4NA22WtyfcYu2CqiygSPmEHFE9CkGbCWGjYMOHsE-w-xdMGcT4mXxYZ4WDZku9Q6w0cQZ0nEhA4iCQGjyoiab9XgtrO95B4mNbpuizyGRcfkmdNPeInJDzw_cCFS-7OLT9kWNi3DYcY-0TMH8OtptNaYUGP1Yn_phwjA_AGlwcV1OC0hlt6HRrQza5cv2R6VL3NgFETTg2t1Viz8ulFOGe8BiCZzYAJ8aSfrO3nX-igZiCMupyo5A2ApNylBZKVi_8InvJ7FH-Kw8DThjJr2abrcqrZUtGS0CSSzQCvBlPN4jpVD111VYNSN25Y6S_pX-kivUxfZ2fZaXoln630InuZvkqX8tdKF7CG1TmYrmHvD3D-MzvNzmB3kf2UZ9LvtU5n3-wAx_L10_tWTfDtYa4WGeJ0hm5ZC7fyBjRSqNeaCu6Wo7et1vqKQFLhdpvOVio7VhKT2FTnVyB7rhQCAbLvJGEQRIqVXmYvsm9zLaRSOikg6Ha_LrAE0SwkEZwTa_t6Vyr3CnszyCzUEvT_G75fAYwcSgsqBJXIXqjqXMD3a11A6yHTApVz-1anY-WwSDQhnoZUzLl1c3n4b5D6dzhmkX0P5F_KI2RHyA44zftCdcRiC8tCE7AknG4WpKkiW7KW_KD0qkm2-xd8Dfd3Lvh5TgbuyxIkvwJlLsFDKnMt78sVZDgHRcBRbtQLsa1DCnibnEqsP6teuJDVz36Aky7T67w0S6jVj7C8gtaAC7zMzixAriso7zggzEGubv43Rfu-qWAmi7qHhddJNPigNprywEOO4AlpI5hl8B8YlkhNuTESR8B1jBx49DB_OkZjegIxM0w_YyxahXGWTI-Q4-MwhlUy86CU-h93scsJ9QgfsYQK5PT7vYHKgpxj9Bw5nb1h9_adwcDeG9j9_rC_s9dGc-QM7a7d6w-Gg907Pdve2T1po6_Usb2uvTvo7fR2pWu_bw_tk38ALXQu1A)
Creates objects without knowing their exact type. You ask for what you want, factory gives it to you.

## Files

```
factory-pattern/
├── car.go              # Car interface
├── car_factory.go      # Creates cars
├── racing_car.go       # Racing car
├── sedan_car.go        # Sedan car
└── main.go             # Example
```

## How it works

### Car interface
```go
type ICar interface {
    GetBrand() string
    Start() string
    Accelerate() string
}
```

### Car types
- **Racing**: Fast, has turbo
- **Sedan**: Comfortable, safe

### Factory
```go
factory := NewCarFactory()
car, err := factory.CreateCar(RacingCarType)
car.Start()
```

## Why use it

- Add new car types easily
- One place handles car creation
- Type-safe
- Clean code


