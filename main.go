// package main

// import (
// 	"fmt"
// 	"math/big"

// 	"github.com/mohithchintu/final_year_project_support/helpers"
// 	"github.com/mohithchintu/final_year_project_support/models"
// )

// func main() {
// 	// Create devices and assign them private keys
// 	device1 := models.NewDevice("device1", big.NewInt(12345), 2)
// 	device2 := models.NewDevice("device2", big.NewInt(67890), 2)

// 	// Add peers
// 	device1.AddPeer(device2)
// 	device2.AddPeer(device1)

//		// Generate group key using ECC-based approach
//		err := helpers.GenerateGroupKeyWithECDH(device1)
//		if err != nil {
//			fmt.Println("Error generating group key for device1:", err)
//		}
//		fmt.Printf("Device ID: %s\n", device1.ID)
//		fmt.Printf("  Private Key: %s\n", device1.PrivateKey.String())
//		fmt.Printf("  Threshold: %d\n", device1.Threshold)
//		fmt.Printf("  Group Key: %s\n", device1.GroupKey.String())
//		// fmt.Printf("  Peers: ")
//		// for _, peer := range device1.Peers {
//		// 	fmt.Printf("%s ", peer.ID)
//		// }
//		// for i, coeff := range device1.Coefficients {
//		// 	fmt.Printf("\n  Coefficient %d: %s", i, coeff.String())
//		// }
//		// for _, share := range device1.Shares {
//		// 	fmt.Printf("\n  Share X: %s, Y: %s", share.X.String(), share.Y.String())
//		// }
//		fmt.Println("\n---------------------")
//		//----------------------------------------------------------------------------------------
//		fmt.Printf("Device ID: %s\n", device2.ID)
//		fmt.Printf("  Private Key: %s\n", device2.PrivateKey.String())
//		fmt.Printf("  Threshold: %d\n", device2.Threshold)
//		fmt.Printf("  Group Key: %s\n", device2.GroupKey.String())
//		// fmt.Printf("  Peers: ")
//		// for _, peer := range device2.Peers {
//		// 	fmt.Printf("%s ", peer.ID)
//		// }
//		// for i, coeff := range device2.Coefficients {
//		// 	fmt.Printf("\n  Coefficient %d: %s", i, coeff.String())
//		// }
//		// for _, share := range device2.Shares {
//		// 	fmt.Printf("\n  Share X: %s, Y: %s", share.X.String(), share.Y.String())
//		// }
//		fmt.Println("\n---------------------")
//	}
package main

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project_support/helpers"
	"github.com/mohithchintu/final_year_project_support/models"
)

func main() {
	// Create devices and assign them private keys
	device1 := models.NewDevice("device1", big.NewInt(12345), 2)
	device2 := models.NewDevice("device2", big.NewInt(67890), 2)

	// Add peers
	device1.AddPeer(device2)
	device2.AddPeer(device1)

	// Generate group key using ECC-based approach for Device 1
	err := helpers.GenerateGroupKeyWithECDH(device1)
	if err != nil {
		fmt.Println("Error generating group key for device1:", err)
	}

	// Generate group key using ECC-based approach for Device 2
	err = helpers.GenerateGroupKeyWithECDH(device2)
	if err != nil {
		fmt.Println("Error generating group key for device2:", err)
	}

	// Now both devices should have the same group key
	fmt.Printf("Device ID: %s\n", device1.ID)
	fmt.Printf("  Private Key: %s\n", device1.PrivateKey.String())
	fmt.Printf("  Threshold: %d\n", device1.Threshold)
	fmt.Printf("  Group Key: %s\n", device1.GroupKey.String())

	// Display Device 2's information
	fmt.Println("---------------------")
	fmt.Printf("Device ID: %s\n", device2.ID)
	fmt.Printf("  Private Key: %s\n", device2.PrivateKey.String())
	fmt.Printf("  Threshold: %d\n", device2.Threshold)
	fmt.Printf("  Group Key: %s\n", device2.GroupKey.String())
}
