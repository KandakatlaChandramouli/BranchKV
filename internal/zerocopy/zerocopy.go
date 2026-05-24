package zerocopy

func Slice(
	data []byte,
	start int,
	end int,
) []byte {

	return data[start:end]
}
