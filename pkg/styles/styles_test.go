package styles

func TestInvalidStyle(t *testing.T) {

	_, err := GetPrettyBool(true, "")
	require.NotNil(err)
}

func TestValidStyle(t *testing.T) {
	_, err := GetPrettyBool(true, "ok")
	require.Nil(err)
}
