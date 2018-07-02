package structsCustom

var quoteBytes = []byte(`"`)

const jsonStartLength = len(`{"Field01":"`)
const jsonEndLength = len(`"}`)

const shortFieldLength = len(`","FieldX":"`)
const longFieldLength = len(`","FieldXX":"`)
