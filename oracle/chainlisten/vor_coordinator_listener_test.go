package chainlisten_test

import (
	"oracle/chainlisten"
	"testing"
)

const (
	VORCoordinatorAddress = "0xBcD9D6aa17802953b455c1225e0Bcf86fe4Ec57A"
)

func TestNewVORCoordinatorListener(t *testing.T) {
	VORCoordinator, err := chainlisten.NewVORCoordinatorListener(VORCoordinatorAddress)

}
