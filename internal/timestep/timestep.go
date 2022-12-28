package timestep

import (
	"encoding/binary"
	"time"
)

const Period = 30 * time.Second

func Get() []byte {
	step := time.Now().UTC().Truncate(Period)

	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(step.Unix()))
	return result
}
