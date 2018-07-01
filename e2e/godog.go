package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/e2e/impl"
)

var (
	chainID = params.TestnetChainConfig.ChainID
)

type FeatureContextOpts struct {
	suite        *godog.Suite
	logsToStdout bool
}

func FeatureContext(opts *FeatureContextOpts) {
	context := impl.NewTestContext(chainID)
	validationCtx := impl.NewValidationContext(context)
	walletBackendCtx := impl.NewWalletBackendContext(context)

	opts.suite.BeforeFeature(func(ft *gherkin.Feature) {
		context.Name = getFeatureName(ft.Name)

		if err := context.InitCluster(opts.logsToStdout); err != nil {
			log.Fatal(err)
		}
	})

	opts.suite.BeforeScenario(func(scenario interface{}) {
		if err := context.RunCluster(); err != nil {
			log.Fatal(err)
		}
	})

	opts.suite.AfterFeature(func(ft *gherkin.Feature) {
		if err := context.DeleteCluster(); err != nil {
			log.Fatal(err)
		}
	})

	opts.suite.AfterScenario(func(scenario interface{}, err error) {
		context.Reset()
		validationCtx.Reset()
		walletBackendCtx.Reset()
	})

	// Creating accounts
	opts.suite.Step(`^I have the following accounts:$`, validationCtx.IHaveTheFollowingAccounts)
	opts.suite.Step(`^I created an account with password '(\w+)'$`, context.ICreatedAnAccountWithPassword)

	// Unlocking accounts
	opts.suite.Step(`^I unlock the account (\w+) with password '(\w+)'$`, context.IUnlockAccountWithPassword)
	opts.suite.Step(`^I try to unlock the account (\w+) with password '(\w+)'$`, context.ITryUnlockAccountWithPassword)
	opts.suite.Step(`^I try to unlock my account with password '(\w+)'$`, context.ITryUnlockMyAccountWithPassword)

	opts.suite.Step(`^I should get my account unlocked$`, context.IGotAccountUnlocked)
	opts.suite.Step(`^I should get an error unlocking the account$`, context.IGotErrorUnlocking)

	// Transactions
	opts.suite.Step(`^I transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITransferKUSD)
	opts.suite.Step(`^I try to transfer (\d+) kcoins? from (\w+) to (\w+)$`, context.ITryTransferKUSD)
	opts.suite.Step(`^the transaction should fail$`, context.LastTransactionFailed)
	opts.suite.Step(`^only one transaction should be done$`, context.OnlyOneTransactionIsDone)
	opts.suite.Step(`^the transaction hash the same$`, context.TransactionHashTheSame)

	// Balances
	opts.suite.Step(`^the balance of (\w+) should be (\d+) kcoins?$`, context.TheBalanceIsExactly)
	opts.suite.Step(`^the balance of (\w+) should be around (\d+) kcoins?$`, context.TheBalanceIsAround)
	opts.suite.Step(`^the balance of (\w+) should be greater (\d+) kcoins?$`, context.TheBalanceIsGreater)
	opts.suite.Step(`^the transaction should fail$`, context.LastTransactionFailed)

	// validation
	opts.suite.Step(`^I start validator with (\d+) mTokens? deposit$`, validationCtx.IStartTheValidator)
	opts.suite.Step(`^I have my node running using account (\w+)$`, validationCtx.IHaveMyNodeRunning)
	opts.suite.Step(`^I stop validation$`, validationCtx.IStopValidation)
	opts.suite.Step(`^I wait for the unbonding period to be over$`, validationCtx.IWaitForTheUnbondingPeriodToBeOver)
	opts.suite.Step(`^I withdraw my node from validation$`, validationCtx.IWithdrawMyNodeFromValidation)
	opts.suite.Step(`^there should be (\d+) mTokens? available to me after (\d+) days$`, validationCtx.ThereShouldBeTokensAvailableToMeAfterDays)
	opts.suite.Step(`^My node should be not be a validator$`, validationCtx.MyNodeShouldBeNotBeAValidator)
	opts.suite.Step(`^I wait for my node to be synced$`, validationCtx.IWaitForMyNodeToBeSynced)

	// mTokens
	opts.suite.Step(`^the deposit of (\w+) should be (\d+) mTokens?$`, validationCtx.IsMTokensBalanceExact)
	opts.suite.Step(`^I transfer (\d+) mTokens? from (\w+) to (\w+)$`, validationCtx.ITransferMTokens)

	// Nodes
	opts.suite.Step(`^I start a new node$`, context.IStartANewNode)
	opts.suite.Step(`^my node should sync with the network$`, context.MyNodeShouldSyncWithTheNetwork)
	opts.suite.Step(`^my node is already synchronised$`, validationCtx.MyNodeIsAlreadySynchronised)
	opts.suite.Step(`^I disconnect my node for (\d+) blocks and reconnect it$`, context.IDisconnectMyNodeForBlocksAndReconnectIt)
	opts.suite.Step(`^I start a new node with a different network ID$`, context.IStartANewNodeWithADifferentNetworkID)
	opts.suite.Step(`^my node should not sync with the network$`, context.MyNodeShouldNotSyncWithTheNetwork)
	opts.suite.Step(`^I start a new node with a different chain ID$`, context.IStartANewNodeWithADifferentChainID)
	opts.suite.Step(`^I start validator with (\d+) deposit and coinbase A$`, context.IStartValidatorWithDepositAndCoinbaseA)

	// Wallet backend
	opts.suite.Step(`^the wallet backend node is running$`, walletBackendCtx.TheWalletBackendNodeIsRunning)
	opts.suite.Step(`^I check the current block height in the wallet backend API$`, walletBackendCtx.ICheckTheCurrentBlockHeightInTheWalletBackendAPI)
	opts.suite.Step(`^I wait for (\d+) blocks$`, walletBackendCtx.IWaitForBlocks)
	opts.suite.Step(`^the new block height in the wallet backend API has increased by at least (\d+)$`, walletBackendCtx.TheNewBlockHeightInTheWalletBackendAPIHasIncreasedByAtLeast)
	opts.suite.Step(`^I transfer (\d+) kcoin from A to B using the wallet API$`, walletBackendCtx.ITransferKcoinFromAToBUsingTheWalletAPI)
	opts.suite.Step(`^the transaction is listed in the wallet backend API$`, walletBackendCtx.TheTransactionIsListedInTheWalletBackendAPI)
	opts.suite.Step(`^the balance of A using the wallet backend should be around (\d+) kcoins$`, walletBackendCtx.TheBalanceOfAUsingTheWalletBackendShouldBeAroundKcoins)
	opts.suite.Step(`^the balance of B using the wallet backend should be (\d+) kcoins$`, walletBackendCtx.TheBalanceOfBUsingTheWalletBackendShouldBeKcoins)
	opts.suite.Step(`^I wait for (\d+) seconds$`, walletBackendCtx.IWaitForSeconds)

}

func getFeatureName(feature string) string {
	feature = strings.ToLower(feature)
	reg, _ := regexp.Compile("[^a-z0-9 ]+")
	feature = reg.ReplaceAllString(feature, "")
	feature = strings.Replace(feature, " ", "_", -1)

	return feature
}
