package main

import (
    "fmt"
    "log"

    "github.com/consensys/gnark-crypto/ecc"
    "github.com/consensys/gnark/backend/groth16"
    "github.com/consensys/gnark/frontend"
    "github.com/consensys/gnark/frontend/cs/r1cs"
    "github.com/consensys/gnark/test"
    "github.com/nickthelegend/zk-driving-license/circuit"
)

func main() {
    const ageValue = 17

    var ckt circuit.DrivingLicenseCircuit
    assignment := circuit.DrivingLicenseCircuit{Age: ageValue}

    // 1. compile
    r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &ckt)
    if err != nil {
        log.Fatalf("compile error: %v", err)
    }

    // 2. enforce circuit constraints in Go
    if err := test.IsSolved(&ckt, &assignment, ecc.BN254.ScalarField()); err != nil {
        log.Fatalf("circuit constraint failed (age < 18): %v", err)
    }

    // 3. setup, witness, prove, verify (as before)…
    pk, vk, err := groth16.Setup(r1cs)
    if err != nil {
        log.Fatal(err)
    }
    fullWit, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
    publicWit, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField(), frontend.PublicOnly())
    proof, err := groth16.Prove(r1cs, pk, fullWit)
    if err != nil {
        log.Fatalf("proof generation failed: %v", err)
    }
    if err := groth16.Verify(proof, vk, publicWit); err != nil {
        log.Fatalf("proof verification failed: %v", err)
    }

    fmt.Println("✅ Proof succeeded for age", ageValue)
}
