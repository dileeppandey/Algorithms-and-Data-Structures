// Calculate wage of a worker in a company based on following:
// basePay, hoursWorked, overTimePay
// 0 < hoursWorked <= 60, overTime = no of hours worked after regular 40 hours

package main

func CalculateWage(basePay float32, hoursWorked float32, overTimePay float32) float32 {
	if hoursWorked <= 0 {
		return 0.0
	} else if hoursWorked <= 40 {
		return basePay * hoursWorked
	} else if hoursWorked > 40 && hoursWorked <= 60 {
		return basePay*40 + (60-hoursWorked)*overTimePay
	} else {
		return 0.0
	}
}
