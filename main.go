package main

import (
	"fmt"
	"math"
)

const (
	CORRECTION_FACTOR  = 1.08
	LANGUAGE_FACTOR    = 1
	APPS_FACTOR        = 0.7
	DEVELOPMENT_FACTOR = 1.4

	YEAR_FUND = 262.0

	ELECTRICITY_PRICE = 2.64
)

func main() {
	var timeStandart float64
	var projectDuration int
	var wage float64
	var costOfPC float64
	var operatingCosts float64

	for {
		// Getting time
		fmt.Printf("Enter a time standart: ")
		fmt.Scan(&timeStandart)

		// Getting duration of the project
		fmt.Printf("Enter a duration of the project: ")
		fmt.Scan(&projectDuration)

		// Getting basic wage for a dev
		fmt.Printf("Enter a wage for a developer: ")
		fmt.Scan(&wage)

		// Getting price of PCs used on project
		fmt.Printf("Enter a price of a PC: ")
		fmt.Scan(&costOfPC)

		// Getting operating costs
		fmt.Printf("Enter operating costs: ")
		fmt.Scan(&operatingCosts)

		// Calculating labor intensity
		laborIntensity := CalculateLaborIntensity(&timeStandart)

		// Calculating amount of developers for a project
		developers := CalculateAmountOfDevelopers(&laborIntensity, &projectDuration)

		// Calculating salary and bonuses
		baseSalary, additionalSalary := CalculateSalary(&wage, &laborIntensity)
		salary := baseSalary + additionalSalary

		// Calculating taxes
		tax := CalculateTax(&salary)

		// Calculating machine time costs
		machineTimeCosts := CalculateMachineTimeCosts(&laborIntensity, &costOfPC, &operatingCosts)

		// General production costs
		generalProductionCosts := CalculateGeneralProductionCosts(&baseSalary)

		// Calculating salary fund
		salaryFund := math.Ceil(salary * developers)

		// Calculating materials costs
		materialsCosts := CalculateMaterialsCosts(&salaryFund)

		// PRODUCTION COST
		productionCost := CalculateProductionCosts(&baseSalary, &additionalSalary, &tax,
			&machineTimeCosts, &generalProductionCosts, &materialsCosts)

		// Administration costs
		administrationCosts := CalculateAdministrationCosts(&baseSalary)

		// Sales expences
		salesExpences := CalculateSalesExpences(&productionCost)

		// TOTAL
		totalCost := CalculateTotalCost(&productionCost, &administrationCosts, &salesExpences)

		// Outputs
		fmt.Println("============RESULTS============")
		fmt.Printf("Labor intensity: %.1f \n", laborIntensity)
		fmt.Printf("Amount of developers: %.1f \n", developers)
		fmt.Printf("Basic salary: %.1f | Bonuses: %.1f \n", baseSalary, additionalSalary)
		fmt.Printf("Tax: %.1f \n", tax)
		fmt.Printf("Machine time costs: %.1f \n", machineTimeCosts)
		fmt.Printf("General production costs: %.1f \n", generalProductionCosts)
		fmt.Printf("Materials costs: %.1f \n", materialsCosts)
		fmt.Printf("PRODUCTION COSTS: %.1f \n", productionCost)
		fmt.Println("===============================")
		fmt.Printf("Administration costs: %.1f \n", administrationCosts)
		fmt.Printf("Sales expences: %.1f \n", salesExpences)
		fmt.Printf("TOTAL COST: %.1f \n", totalCost)
		fmt.Println("===============================")

		var q string
		fmt.Print("Calculate more? (y/n): ")
		fmt.Scan(&q)
		if q == "n" {
			break
		}
	}
}

func CalculateLaborIntensity(ts *float64) float64 {
	return *ts * CORRECTION_FACTOR * LANGUAGE_FACTOR * APPS_FACTOR * DEVELOPMENT_FACTOR
}
func CalculateAmountOfDevelopers(ts *float64, d *int) float64 {
	return math.Ceil(*ts / (YEAR_FUND * float64(float64(*d)/12)))
}
func CalculateSalary(wage *float64, ts *float64) (float64, float64) {
	baseSalary := *wage * 12 * *ts / YEAR_FUND
	return baseSalary, baseSalary * 0.3
}
func CalculateTax(salary *float64) float64 {
	return *salary * 0.3676
}
func CalculateMachineTimeCosts(ts *float64, pc *float64, oc *float64) float64 {
	return MachineTimePrice(oc) * MachineTimeDuration(ts, pc)
}
func MachineTimePrice(oc *float64) float64 {
	return *oc / (YEAR_FUND * 6.8)
}
func MachineTimeDuration(ts *float64, pc *float64) float64 {
	return *pc*0.15 + (0.12 * 8 * ELECTRICITY_PRICE)
}

func CalculateGeneralProductionCosts(baseSalary *float64) float64 {
	return *baseSalary * 1.1
}
func CalculateMaterialsCosts(fund *float64) float64 {
	return *fund * 0.04
}
func CalculateProductionCosts(bs *float64, b *float64, tax *float64, mtc *float64, gpc *float64, mc *float64) float64 {
	return *bs + *b + *tax + *mtc + *gpc + *mc
}
func CalculateAdministrationCosts(baseSalary *float64) float64 {
	return *baseSalary * 1.5
}
func CalculateSalesExpences(productionCost *float64) float64 {
	return *productionCost * 0.03
}
func CalculateTotalCost(productionCost *float64, admCost *float64, salesExpences *float64) float64 {
	return *productionCost + *admCost + *salesExpences
}
