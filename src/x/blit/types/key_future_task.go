package types

import (
	"encoding/binary"
	fmt "fmt"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
// FutureTaskKeyPrefix is the prefix to retrieve all FutureTask
//
//	FutureTaskKeyPrefix = "FutureTask/value/"
)

func FutureTaskKeyPool() []byte {
	var key []byte
	key = append(key, []byte(fmt.Sprintf("%s/", FutureTaskStatus_POOL))...)
	return key
}

func FutureTaskKeyPending() []byte {
	var key []byte
	key = append(key, []byte(fmt.Sprintf("%s/", FutureTaskStatus_PENDING))...)
	return key
}

// FutureTaskKey returns the store key to retrieve a FutureTask from the timestamp and optional taskId
func FutureTaskKey(
	status FutureTaskStatus,
	scheduledOn time.Time,
	taskId uint64,
	gasPrice *sdk.DecCoin,
) []byte {
	var key []byte

	statusBytes := []byte(fmt.Sprint(status))
	key = append(key, statusBytes...)
	key = append(key, []byte("/")...)

	if status == FutureTaskStatus_PENDING {

		scheduledOnBytes := []byte(scheduledOn.Format(time.RFC3339))
		key = append(key, scheduledOnBytes...)
		key = append(key, []byte("/")...)

		gasPriceBytes := []byte(fmt.Sprint(gasPrice))
		key = append(key, gasPriceBytes...)
		key = append(key, []byte("/")...)

	}

	if status == FutureTaskStatus_POOL {

		gasPriceBytes := []byte(fmt.Sprint(gasPrice))
		key = append(key, gasPriceBytes...)
		key = append(key, []byte("/")...)

		scheduledOnBytes := []byte(scheduledOn.Format(time.RFC3339))
		key = append(key, scheduledOnBytes...)
		key = append(key, []byte("/")...)

	}
	taskIdBytes := []byte(fmt.Sprint(taskId))
	key = append(key, taskIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
