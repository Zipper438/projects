package flipbeats

func plippingBits(n int64) uint32 {
	return uint32(^uint32(n))
}
