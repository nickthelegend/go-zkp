# go-zkp


Here’s a step-by-step walkthrough of what main.go is doing

# 1. Circuit Instantiation & Compilation

```go
var ckt circuit.DrivingLicenseCircuit

r1cs, err := frontend.Compile(
    ecc.BN254.ScalarField(),
    r1cs.NewBuilder,
    &ckt,
)


```