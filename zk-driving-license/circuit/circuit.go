// circuit/circuit.go
package circuit

import "github.com/consensys/gnark/frontend"

// DrivingLicenseCircuit enforces Age ≥ 18, with Age as a public input.
type DrivingLicenseCircuit struct {
    Age frontend.Variable `gnark:",public"`
}

func (c *DrivingLicenseCircuit) Define(api frontend.API) error {
    // Age - 18 ≥ 0  →  Age ≥ 18
    diff := api.Sub(c.Age, 18)
    api.AssertIsLessOrEqual(0, diff)
    return nil
}
