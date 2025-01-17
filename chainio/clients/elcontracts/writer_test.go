package elcontracts_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"

	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/testutils/testclients"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterOperator(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	require.NoError(t, err)

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 anvilHttpEndpoint,
		EthWsUrl:                   anvilWsEndpoint,
		RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
		OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
		AvsName:                    "exampleAvs",
		PromMetricsIpPortAddress:   ":9090",
	}

	t.Run("register as an operator", func(t *testing.T) {
		// Fund the new address with 5 ether
		fundedAccount := "0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1"
		fundedPrivateKeyHex := "3339854a8622364bcd5650fa92eac82d5dccf04089f5575a761c9b7d3c405b1c"
		richPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
		code, _, err := anvilC.Exec(
			context.Background(),
			[]string{"cast",
				"send",
				"0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1",
				"--value",
				"5ether",
				"--private-key",
				richPrivateKeyHex,
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 0, code)

		ecdsaPrivateKey, err := crypto.HexToECDSA(fundedPrivateKeyHex)
		require.NoError(t, err)

		clients, err := clients.BuildAll(
			chainioConfig,
			ecdsaPrivateKey,
			logger,
		)
		require.NoError(t, err)

		operator :=
			types.Operator{
				Address:                   fundedAccount,
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			}

		receipt, err := clients.ElChainWriter.RegisterAsOperator(context.Background(), operator, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("register as an operator already registered", func(t *testing.T) {
		operatorAddress := "0x408EfD9C90d59298A9b32F4441aC9Df6A2d8C3E1"
		operatorPrivateKeyHex := "3339854a8622364bcd5650fa92eac82d5dccf04089f5575a761c9b7d3c405b1c"

		ecdsaPrivateKey, err := crypto.HexToECDSA(operatorPrivateKeyHex)
		require.NoError(t, err)

		clients, err := clients.BuildAll(
			chainioConfig,
			ecdsaPrivateKey,
			logger,
		)
		require.NoError(t, err)

		operator :=
			types.Operator{
				Address:                   operatorAddress,
				DelegationApproverAddress: "0xd5e099c71b797516c10ed0f0d895f429c2781142",
				MetadataUrl:               "https://madhur-test-public.s3.us-east-2.amazonaws.com/metadata.json",
			}

		_, err = clients.ElChainWriter.RegisterAsOperator(context.Background(), operator, true)
		assert.Error(t, err)
	})
}

func TestChainWriter(t *testing.T) {
	clients, anvilHttpEndpoint := testclients.BuildTestClients(t)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	t.Run("update operator details", func(t *testing.T) {
		walletModified, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)

		operatorModified := types.Operator{
			Address:                   walletModifiedAddress.Hex(),
			DelegationApproverAddress: walletModifiedAddress.Hex(),
			MetadataUrl:               "eigensdk-go",
		}

		receipt, err := clients.ElChainWriter.UpdateOperatorDetails(context.Background(), operatorModified, true)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("update metadata URI", func(t *testing.T) {
		walletModified, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		assert.NoError(t, err)
		walletModifiedAddress := crypto.PubkeyToAddress(walletModified.PublicKey)
		receipt, err := clients.ElChainWriter.UpdateMetadataURI(
			context.Background(),
			walletModifiedAddress,
			"https://0.0.0.0",
			true,
		)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})

	t.Run("deposit ERC20 into strategy", func(t *testing.T) {
		amount := big.NewInt(1)
		receipt, err := clients.ElChainWriter.DepositERC20IntoStrategy(
			context.Background(),
			contractAddrs.Erc20MockStrategy,
			amount,
			true,
		)
		assert.NoError(t, err)
		assert.True(t, receipt.Status == 1)
	})
}

func TestSetClaimerFor(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	waitForReceipt := true
	claimer := common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")

	// call SetClaimerFor
	receipt, err := chainWriter.SetClaimerFor(context.Background(), claimer, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestSetOperatorPISplit(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	waitForReceipt := true

	activationDelay := uint32(0)
	// Set activation delay to zero so that the new PI split can be retrieved immediately after setting it
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	expectedInitialSplit := uint16(1000)
	initialSplit, err := chainReader.GetOperatorPISplit(context.Background(), operatorAddr)
	require.NoError(t, err)
	require.Equal(t, expectedInitialSplit, initialSplit)

	newSplit := initialSplit + 1
	// Set a new operator PI split
	receipt, err = chainWriter.SetOperatorPISplit(context.Background(), operatorAddr, newSplit, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Retrieve the operator PI split to check it has been set
	updatedSplit, err := chainReader.GetOperatorPISplit(context.Background(), operatorAddr)
	require.NoError(t, err)
	require.Equal(t, newSplit, updatedSplit)
}

func TestSetOperatorAVSSplit(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	avsAddr := contractAddrs.ServiceManager
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	waitForReceipt := true
	activationDelay := uint32(0)

	// Set activation delay to zero so that the new AVS split can be retrieved immediately after setting it
	receipt, err := setTestRewardsCoordinatorActivationDelay(anvilHttpEndpoint, privateKeyHex, activationDelay)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	expectedInitialSplit := uint16(1000)
	initialSplit, err := chainReader.GetOperatorAVSSplit(context.Background(), operatorAddr, avsAddr)
	require.NoError(t, err)
	require.Equal(t, expectedInitialSplit, initialSplit)

	newSplit := initialSplit + 1
	// Set a new operator AVS split
	receipt, err = chainWriter.SetOperatorAVSSplit(
		context.Background(),
		operatorAddr,
		avsAddr,
		newSplit,
		waitForReceipt,
	)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Retrieve the operator AVS split to check it has been set
	updatedSplit, err := chainReader.GetOperatorAVSSplit(context.Background(), operatorAddr, avsAddr)
	require.NoError(t, err)
	require.Equal(t, newSplit, updatedSplit)
}

func TestSetAllocationDelay(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	config := elcontracts.Config{
		DelegationManagerAddress:  contractAddrs.DelegationManager,
		RewardsCoordinatorAddress: rewardsCoordinatorAddr,
	}

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	waitForReceipt := true

	// Create ChainWriter
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	delay := uint32(10)
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func TestSetAndRemovePermission(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)
	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	accountAddress := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	appointeeAddress := common.HexToAddress("009440d62dc85c73dbf889b7ad1f4da8b231d2ef")
	target := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	selector := [4]byte{0, 1, 2, 3}
	waitForReceipt := true

	setPermissionRequest := elcontracts.SetPermissionRequest{
		AccountAddress:   accountAddress,
		AppointeeAddress: appointeeAddress,
		Target:           target,
		Selector:         selector,
		WaitForReceipt:   waitForReceipt,
	}

	removePermissionRequest := elcontracts.RemovePermissionRequest{
		AccountAddress:   accountAddress,
		AppointeeAddress: appointeeAddress,
		Target:           target,
		Selector:         selector,
		WaitForReceipt:   waitForReceipt,
	}
	receipt, err := chainWriter.SetPermission(context.Background(), setPermissionRequest)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	canCall, err := chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
	require.NoError(t, err)
	require.True(t, canCall)

	receipt, err = chainWriter.RemovePermission(context.Background(), removePermissionRequest)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	canCall, err = chainReader.CanCall(context.Background(), accountAddress, appointeeAddress, target, selector)
	require.NoError(t, err)
	require.False(t, canCall)
}

func TestModifyAllocations(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}

	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	strategyAddr := contractAddrs.Erc20MockStrategy
	avsAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	operatorSetId := uint32(1)

	waitForReceipt := true
	delay := uint32(1)
	// The allocation delay must be initialized before modifying the allocations
	receipt, err := chainWriter.SetAllocationDelay(context.Background(), operatorAddr, delay, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	allocationConfigurationDelay := 1200
	// Advance the chain by the required number of blocks
	// (ALLOCATION_CONFIGURATION_DELAY) to apply the allocation delay
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), allocationConfigurationDelay+1, anvilC)

	// Retrieve the allocation delay so that the delay is applied
	_, err = chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)

	err = createOperatorSet(anvilHttpEndpoint, privateKeyHex, avsAddr, operatorSetId, strategyAddr)
	require.NoError(t, err)

	operatorSet := allocationmanager.OperatorSet{
		Avs: avsAddr,
		Id:  operatorSetId,
	}
	newAllocation := uint64(100)
	allocateParams := []allocationmanager.IAllocationManagerTypesAllocateParams{
		{
			OperatorSet:   operatorSet,
			Strategies:    []common.Address{strategyAddr},
			NewMagnitudes: []uint64{newAllocation},
		},
	}

	receipt, err = chainWriter.ModifyAllocations(context.Background(), operatorAddr, allocateParams, waitForReceipt)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Check that the new allocation is pending and the current magnitude is zero
	allocationInfo, err := chainReader.GetAllocationInfo(context.Background(), operatorAddr, strategyAddr)
	require.NoError(t, err)
	pendingDiff := allocationInfo[0].PendingDiff
	require.Equal(t, big.NewInt(int64(newAllocation)), pendingDiff)
	require.Equal(t, allocationInfo[0].CurrentMagnitude, big.NewInt(0))

	// Retrieve the allocation delay and advance the chain
	allocationDelay, err := chainReader.GetAllocationDelay(context.Background(), operatorAddr)
	require.NoError(t, err)
	testutils.AdvanceChainByNBlocksExecInContainer(context.Background(), int(allocationDelay), anvilC)

	// Check the new allocation has been updated after the delay
	allocationInfo, err = chainReader.GetAllocationInfo(context.Background(), operatorAddr, strategyAddr)
	require.NoError(t, err)

	currentMagnitude := allocationInfo[0].CurrentMagnitude
	require.Equal(t, big.NewInt(int64(newAllocation)), currentMagnitude)
}

func TestAddAndRemovePendingAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	// TODO: unhardcode permissionControllerAddr
	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	operatorAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	privateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	chainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, privateKeyHex, config)
	require.NoError(t, err)
	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	pendingAdmin := common.HexToAddress("009440d62dc85c73dbf889b7ad1f4da8b231d2ef")
	request := elcontracts.AddPendingAdminRequest{
		AccountAddress: operatorAddr,
		AdminAddress:   pendingAdmin,
		WaitForReceipt: true,
	}

	removePendingAdminRequest := elcontracts.RemovePendingAdminRequest{
		AccountAddress: operatorAddr,
		AdminAddress:   pendingAdmin,
		WaitForReceipt: true,
	}
	t.Run("add pending admin", func(t *testing.T) {
		receipt, err := chainWriter.AddPendingAdmin(context.Background(), request)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdmin)
		require.NoError(t, err)
		require.True(t, isPendingAdmin)
	})
	t.Run("remove pending admin", func(t *testing.T) {
		receipt, err := chainWriter.RemovePendingAdmin(context.Background(), removePendingAdminRequest)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isPendingAdmin, err := chainReader.IsPendingAdmin(context.Background(), operatorAddr, pendingAdmin)
		require.NoError(t, err)
		require.False(t, isPendingAdmin)
	})
}

func TestAcceptAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	// TODO: unhardcode permissionControllerAddr
	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	accountAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	accountPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	accountChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, accountPrivateKeyHex, config)
	require.NoError(t, err)

	pendingAdminPrivateKeyHex := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	adminChainWriter, err := testclients.NewTestChainWriterFromConfig(
		anvilHttpEndpoint,
		pendingAdminPrivateKeyHex,
		config,
	)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	pendingAdminAddr := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	request := elcontracts.AddPendingAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   pendingAdminAddr,
		WaitForReceipt: true,
	}
	receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), request)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	acceptAdminRequest := elcontracts.AcceptAdminRequest{
		AccountAddress: accountAddr,
		WaitForReceipt: true,
	}
	t.Run("accept admin", func(t *testing.T) {
		receipt, err = adminChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isAdmin, err := chainReader.IsAdmin(context.Background(), accountAddr, pendingAdminAddr)
		require.NoError(t, err)
		require.True(t, isAdmin)
	})
}

func TestRemoveAdmin(t *testing.T) {
	testConfig := testutils.GetDefaultTestConfig()
	anvilC, err := testutils.StartAnvilContainer(testConfig.AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	// TODO: unhardcode permissionControllerAddr
	permissionControllerAddr := common.HexToAddress(testutils.PERMISSION_CONTROLLER_ADDRESS)

	accountAddr := common.HexToAddress(testutils.ANVIL_FIRST_ADDRESS)
	accountPrivateKeyHex := testutils.ANVIL_FIRST_PRIVATE_KEY
	config := elcontracts.Config{
		DelegationManagerAddress:     contractAddrs.DelegationManager,
		PermissionsControllerAddress: permissionControllerAddr,
	}
	accountChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, accountPrivateKeyHex, config)
	require.NoError(t, err)

	// Adding two admins and removing one. Cannot remove the last admin, so one must remain
	admin1 := common.HexToAddress("14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	admin1PrivateKeyHex := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"

	admin2 := common.HexToAddress("23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f")
	admin2PrivateKeyHex := "dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97"

	admin1ChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, admin1PrivateKeyHex, config)
	require.NoError(t, err)

	admin2ChainWriter, err := testclients.NewTestChainWriterFromConfig(anvilHttpEndpoint, admin2PrivateKeyHex, config)
	require.NoError(t, err)

	chainReader, err := testclients.NewTestChainReaderFromConfig(anvilHttpEndpoint, config)
	require.NoError(t, err)

	addAdmin1Request := elcontracts.AddPendingAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin1,
		WaitForReceipt: true,
	}
	addAdmin2Request := elcontracts.AddPendingAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin2,
		WaitForReceipt: true,
	}
	acceptAdminRequest := elcontracts.AcceptAdminRequest{
		AccountAddress: accountAddr,
		WaitForReceipt: true,
	}
	// Add and accept admin 1
	receipt, err := accountChainWriter.AddPendingAdmin(context.Background(), addAdmin1Request)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	receipt, err = admin1ChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Add and accept admin 2
	receipt, err = admin1ChainWriter.AddPendingAdmin(context.Background(), addAdmin2Request)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	receipt, err = admin2ChainWriter.AcceptAdmin(context.Background(), acceptAdminRequest)
	require.NoError(t, err)
	require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

	removeAdminRequest := elcontracts.RemoveAdminRequest{
		AccountAddress: accountAddr,
		AdminAddress:   admin2,
		WaitForReceipt: true,
	}
	t.Run("remove admin 2", func(t *testing.T) {
		receipt, err = admin1ChainWriter.RemoveAdmin(context.Background(), removeAdminRequest)
		require.NoError(t, err)
		require.Equal(t, gethtypes.ReceiptStatusSuccessful, receipt.Status)

		isAdmin, err := chainReader.IsAdmin(context.Background(), accountAddr, admin2)
		require.NoError(t, err)
		require.False(t, isAdmin)
	})
}

// Creates an operator set with `avsAddress`, `operatorSetId` and `erc20MockStrategyAddr`.
func createOperatorSet(
	anvilHttpEndpoint string,
	privateKeyHex string,
	avsAddress common.Address,
	operatorSetId uint32,
	erc20MockStrategyAddr common.Address,
) error {
	testConfig := testutils.GetDefaultTestConfig()
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	config := elcontracts.Config{
		DelegationManagerAddress: contractAddrs.DelegationManager,
	}
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: testConfig.LogLevel})
	ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
	if err != nil {
		return err
	}

	elBindings, err := elcontracts.NewBindingsFromConfig(config, ethHttpClient, logger)
	if err != nil {
		return err
	}

	allocationManager := elBindings.AllocationManager
	registryCoordinatorAddress := contractAddrs.RegistryCoordinator
	registryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		registryCoordinatorAddress,
		ethHttpClient,
	)
	if err != nil {
		return err
	}
	txManager, err := testclients.NewTestTxManager(anvilHttpEndpoint, privateKeyHex)
	if err != nil {
		return err
	}
	noSendTxOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return err
	}

	tx, err := allocationManager.SetAVSRegistrar(noSendTxOpts, avsAddress, registryCoordinatorAddress)
	if err != nil {
		return err
	}

	waitForReceipt := true

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	tx, err = registryCoordinator.EnableOperatorSets(noSendTxOpts)
	if err != nil {
		return err
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	operatorSetParam := regcoord.IRegistryCoordinatorOperatorSetParam{
		MaxOperatorCount:        10,
		KickBIPsOfOperatorStake: 100,
		KickBIPsOfTotalStake:    1000,
	}
	minimumStake := big.NewInt(0)

	strategyParams := regcoord.IStakeRegistryStrategyParams{
		Strategy:   erc20MockStrategyAddr,
		Multiplier: big.NewInt(1),
	}
	strategyParamsArray := []regcoord.IStakeRegistryStrategyParams{strategyParams}
	lookAheadPeriod := uint32(0)
	tx, err = registryCoordinator.CreateSlashableStakeQuorum(
		noSendTxOpts,
		operatorSetParam,
		minimumStake,
		strategyParamsArray,
		lookAheadPeriod,
	)
	if err != nil {
		return err
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return err
	}

	strategies := []common.Address{erc20MockStrategyAddr}
	operatorSetParams := allocationmanager.IAllocationManagerTypesCreateSetParams{
		OperatorSetId: operatorSetId,
		Strategies:    strategies,
	}
	operatorSetParamsArray := []allocationmanager.IAllocationManagerTypesCreateSetParams{operatorSetParams}
	tx, err = allocationManager.CreateOperatorSets(noSendTxOpts, avsAddress, operatorSetParamsArray)
	if err != nil {
		return err
	}

	_, err = txManager.Send(context.Background(), tx, waitForReceipt)
	return err
}

// Sets the testing RewardsCoordinator's activationDelay.
// This is useful to test ChainWriter setter functions that depend on activationDelay.
func setTestRewardsCoordinatorActivationDelay(
	httpEndpoint string,
	privateKeyHex string,
	activationDelay uint32,
) (*gethtypes.Receipt, error) {
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(httpEndpoint)
	rewardsCoordinatorAddr := contractAddrs.RewardsCoordinator
	ethHttpClient, err := ethclient.Dial(httpEndpoint)
	if err != nil {
		return nil, utils.WrapError("Failed to create eth client", err)
	}

	rewardsCoordinator, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardsCoordinatorAddr, ethHttpClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create rewards coordinator", err)
	}

	txManager, err := testclients.NewTestTxManager(httpEndpoint, privateKeyHex)
	if err != nil {
		return nil, utils.WrapError("Failed to create tx manager", err)
	}

	noSendOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		return nil, utils.WrapError("Failed to get NoSend tx opts", err)
	}

	tx, err := rewardsCoordinator.SetActivationDelay(noSendOpts, activationDelay)
	if err != nil {
		return nil, utils.WrapError("Failed to create SetActivationDelay tx", err)
	}

	receipt, err := txManager.Send(context.Background(), tx, true)
	if err != nil {
		return nil, utils.WrapError("Failed to send SetActivationDelay tx", err)
	}
	return receipt, err
}
